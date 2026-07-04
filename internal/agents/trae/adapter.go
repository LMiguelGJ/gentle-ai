package trae

import (
	"context"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

type statResult struct {
	isDir bool
	err   error
}

// Adapter implements agents.Adapter for Trae IDE (by ByteDance).
//
// Config path summary (Trae uses a flat cross-platform home-rooted layout):
//   - Detection, skills, rules, MCP: ~/.trae/ (cross-platform, always under home)
//     → skills/                         Skill files
//     → user_rules/gentle-ai.md         Personal rules (StrategyMarkdownSections)
//     → mcp.json                        MCP server configs (Cursor-compatible mcpServers)
//   - Settings: OS-specific Trae User config dir (VSCode-style split root)
//     macOS:   ~/Library/Application Support/Trae/User/
//     Linux:   ~/.config/Trae/User/   (respects XDG_CONFIG_HOME)
//     Windows: %APPDATA%\Trae\User\
//     → settings.json                   Editor settings (rarely used by gentle-ai)
//
// Detection: Trae is a desktop app. If ~/.trae exists as a directory, it's installed.
// No binary appears on PATH.
type Adapter struct {
	statPath func(string) statResult
}

func NewAdapter() *Adapter {
	return &Adapter{statPath: defaultStat}
}

// --- Identity ---

func (a *Adapter) Agent() model.AgentID    { return model.AgentTrae }
func (a *Adapter) Tier() model.SupportTier { return model.TierFull }

// --- Detection ---

// Detect checks for the ~/.trae directory, which Trae creates on its first launch.
// No binary appears on PATH (desktop app).
func (a *Adapter) Detect(_ context.Context, homeDir string) (bool, string, string, bool, error) {
	configPath := a.GlobalConfigDir(homeDir)
	stat := a.statPath(configPath)
	if stat.err != nil {
		if os.IsNotExist(stat.err) {
			return false, "", configPath, false, nil
		}
		return false, "", "", false, stat.err
	}
	return stat.isDir, "", configPath, stat.isDir, nil
}

// --- Installation ---

func (a *Adapter) SupportsAutoInstall() bool { return false }

func (a *Adapter) InstallCommand(_ system.PlatformProfile) ([][]string, error) {
	return nil, AgentNotInstallableError{Agent: model.AgentTrae}
}

// --- Config paths ---

// GlobalConfigDir returns ~/.trae, the root of Trae's config directory.
// Trae uses a flat cross-platform layout with no OS-specific split.
func (a *Adapter) GlobalConfigDir(homeDir string) string {
	return filepath.Join(homeDir, ".trae")
}

// SystemPromptDir returns the Trae rules directory under ~/.trae/.
// Trae reads personal rules from ~/.trae/user_rules/ (a directory), not from
// the OS-specific app config dir.
func (a *Adapter) SystemPromptDir(homeDir string) string {
	return filepath.Join(a.GlobalConfigDir(homeDir), "user_rules")
}

// SystemPromptFile returns the personal rules file that gentle-ai manages
// inside the Trae user_rules/ directory. Trae loads all .md files in that
// directory as rules, so writing gentle-ai.md here makes our sections visible
// without clobbering other user-managed rules files.
// gentle-ai injects its sections via StrategyMarkdownSections markers.
func (a *Adapter) SystemPromptFile(homeDir string) string {
	return filepath.Join(a.SystemPromptDir(homeDir), "gentle-ai.md")
}

// SkillsDir returns the skills directory for Trae.
func (a *Adapter) SkillsDir(homeDir string) string {
	return filepath.Join(a.GlobalConfigDir(homeDir), "skills")
}

// SettingsPath returns the platform-specific editor settings.json.
func (a *Adapter) SettingsPath(homeDir string) string {
	return filepath.Join(a.traeUserDir(homeDir), "settings.json")
}

// --- Config strategies ---

// SystemPromptStrategy uses MarkdownSections: gentle-ai markers are injected
// into user_rules/gentle-ai.md without clobbering other user content.
func (a *Adapter) SystemPromptStrategy() model.SystemPromptStrategy {
	return model.StrategyMarkdownSections
}

func (a *Adapter) MCPStrategy() model.MCPStrategy {
	return model.StrategyMCPConfigFile
}

// --- MCP ---

// MCPConfigPath returns the MCP servers config file.
// Trae reads MCP config from ~/.trae/mcp.json (Cursor-compatible
// mcpServers object format). The OS-specific Trae User dir is NOT used
// for MCP — that path was historically documented but real Trae installs
// ignore it. See docs/trae-integration-plan.md P0.2.
func (a *Adapter) MCPConfigPath(homeDir string, _ string) string {
	return filepath.Join(a.GlobalConfigDir(homeDir), "mcp.json")
}

// traeUserDir returns the OS-specific Trae User config directory.
// Only used by SettingsPath (which follows VSCode-style split-root layout
// for editor settings.json). Skills, rules, and MCP all live under the
// cross-platform ~/.trae/ root.
func (a *Adapter) traeUserDir(homeDir string) string {
	switch runtime.GOOS {
	case "darwin":
		return filepath.Join(homeDir, "Library", "Application Support", "Trae", "User")
	case "windows":
		appData := os.Getenv("APPDATA")
		if appData == "" {
			appData = filepath.Join(homeDir, "AppData", "Roaming")
		}
		return filepath.Join(appData, "Trae", "User")
	default: // linux and others
		xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
		if xdgConfigHome == "" {
			xdgConfigHome = filepath.Join(homeDir, ".config")
		}
		return filepath.Join(xdgConfigHome, "Trae", "User")
	}
}

// --- Optional capabilities ---

func (a *Adapter) SupportsOutputStyles() bool     { return false }
func (a *Adapter) OutputStyleDir(_ string) string { return "" }

// SupportsSlashCommands is true. Trae loads custom slash commands from
// `~/.trae/commands/<name>.md` and exposes them in chat autocomplete.
func (a *Adapter) SupportsSlashCommands() bool { return true }

// CommandsDir returns the Trae slash commands directory.
func (a *Adapter) CommandsDir(homeDir string) string {
	return filepath.Join(a.GlobalConfigDir(homeDir), "commands")
}

// SupportsSubAgents is true. Trae loads Skill files from `~/.trae/skills/<name>/`
// as callable custom agents. Gentle AI installs SDD phase skills under
// `~/.trae/skills/sdd-<phase>/SKILL.md` and Judgment Day agents under
// `~/.trae/skills/jd-<name>/SKILL.md`.
func (a *Adapter) SupportsSubAgents() bool { return true }

// SubAgentsDir returns the Trae skills directory (Trae uses skills as agents).
func (a *Adapter) SubAgentsDir(homeDir string) string {
	return a.SkillsDir(homeDir)
}

// EmbeddedSubAgentsDir is the embedded assets subdirectory containing
// Trae custom-agent templates, relative to internal/assets/.
func (a *Adapter) EmbeddedSubAgentsDir() string {
	return "trae/agents"
}

func (a *Adapter) SupportsSkills() bool       { return true }
func (a *Adapter) SupportsSystemPrompt() bool { return true }
func (a *Adapter) SupportsMCP() bool          { return true }

// AgentNotInstallableError is returned when InstallCommand is called on a desktop-only agent.
type AgentNotInstallableError struct {
	Agent model.AgentID
}

func (e AgentNotInstallableError) Error() string {
	return "agent " + string(e.Agent) + " is a desktop app and cannot be installed via CLI"
}

func defaultStat(path string) statResult {
	info, err := os.Stat(path)
	if err != nil {
		return statResult{err: err}
	}
	return statResult{isDir: info.IsDir()}
}
