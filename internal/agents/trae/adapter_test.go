package trae

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/gentleman-programming/gentle-ai/internal/model"
	"github.com/gentleman-programming/gentle-ai/internal/system"
)

func TestDetect(t *testing.T) {
	tests := []struct {
		name            string
		stat            statResult
		wantInstalled   bool
		wantConfigPath  string
		wantConfigFound bool
		wantErr         bool
	}{
		{
			name:            "config directory found",
			stat:            statResult{isDir: true},
			wantInstalled:   true,
			wantConfigPath:  filepath.Join("/tmp/home", ".trae"),
			wantConfigFound: true,
		},
		{
			name:            "config missing",
			stat:            statResult{err: os.ErrNotExist},
			wantInstalled:   false,
			wantConfigPath:  filepath.Join("/tmp/home", ".trae"),
			wantConfigFound: false,
		},
		{
			name:    "stat error bubbles up",
			stat:    statResult{err: errors.New("permission denied")},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Adapter{
				statPath: func(string) statResult {
					return tt.stat
				},
			}

			installed, binaryPath, configPath, configFound, err := a.Detect(context.Background(), "/tmp/home")
			if (err != nil) != tt.wantErr {
				t.Fatalf("Detect() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if binaryPath != "" {
				t.Fatalf("Detect() binaryPath = %q, want %q (Trae is desktop app)", binaryPath, "")
			}

			if installed != tt.wantInstalled {
				t.Fatalf("Detect() installed = %v, want %v", installed, tt.wantInstalled)
			}

			if configPath != tt.wantConfigPath {
				t.Fatalf("Detect() configPath = %q, want %q", configPath, tt.wantConfigPath)
			}

			if configFound != tt.wantConfigFound {
				t.Fatalf("Detect() configFound = %v, want %v", configFound, tt.wantConfigFound)
			}
		})
	}
}

func TestConfigPathsCrossPlatform(t *testing.T) {
	a := NewAdapter()
	home := "/tmp/home"

	if got := a.GlobalConfigDir(home); got != filepath.Join(home, ".trae") {
		t.Fatalf("GlobalConfigDir() = %q, want %q", got, filepath.Join(home, ".trae"))
	}

	if got := a.SystemPromptDir(home); got != filepath.Join(home, ".trae", "user_rules") {
		t.Fatalf("SystemPromptDir() = %q, want %q", got, filepath.Join(home, ".trae", "user_rules"))
	}

	if got := a.SkillsDir(home); got != filepath.Join(home, ".trae", "skills") {
		t.Fatalf("SkillsDir() = %q, want %q", got, filepath.Join(home, ".trae", "skills"))
	}

	if got := a.MCPConfigPath(home, "ctx7"); got != filepath.Join(home, ".trae", "mcp.json") {
		t.Fatalf("MCPConfigPath() = %q, want %q", got, filepath.Join(home, ".trae", "mcp.json"))
	}

	if got := a.SettingsPath(home); got != filepath.Join(home, ".trae", "settings.json") {
		t.Fatalf("SettingsPath() = %q, want %q", got, filepath.Join(home, ".trae", "settings.json"))
	}

	if got := a.SystemPromptFile(home); got != filepath.Join(home, ".trae", "user_rules", "gentle-ai.md") {
		t.Fatalf("SystemPromptFile() = %q, want %q", got, filepath.Join(home, ".trae", "user_rules", "gentle-ai.md"))
	}
}

func TestStrategies(t *testing.T) {
	a := NewAdapter()

	if got := a.SystemPromptStrategy(); got != model.StrategyMarkdownSections {
		t.Fatalf("SystemPromptStrategy() = %v, want %v", got, model.StrategyMarkdownSections)
	}

	if got := a.MCPStrategy(); got != model.StrategyMCPConfigFile {
		t.Fatalf("MCPStrategy() = %v, want %v", got, model.StrategyMCPConfigFile)
	}
}

func TestDesktopAppNotAutoInstallable(t *testing.T) {
	a := NewAdapter()

	if a.SupportsAutoInstall() {
		t.Fatalf("Trae should not support auto-install (desktop app)")
	}

	_, err := a.InstallCommand(system.PlatformProfile{})
	if err == nil {
		t.Fatalf("InstallCommand() should return error for desktop app")
	}
}

func TestAgentIdentity(t *testing.T) {
	a := NewAdapter()

	if got := a.Agent(); got != model.AgentTrae {
		t.Fatalf("Agent() = %v, want %v", got, model.AgentTrae)
	}

	if got := a.Tier(); got != model.TierFull {
		t.Fatalf("Tier() = %v, want %v", got, model.TierFull)
	}
}

func TestAgentNotInstallableError(t *testing.T) {
	err := AgentNotInstallableError{Agent: model.AgentTrae}
	expected := "agent trae is a desktop app and cannot be installed via CLI"
	if err.Error() != expected {
		t.Errorf("expected %s, got %s", expected, err.Error())
	}
}

func TestSupportsSubAgents(t *testing.T) {
	a := NewAdapter()

	if a.SupportsSubAgents() {
		t.Fatalf("SupportsSubAgents() = true, want false (Trae does not support sub-agents)")
	}

	if got := a.SubAgentsDir("/tmp/home"); got != "" {
		t.Fatalf("SubAgentsDir() = %q, want %q", got, "")
	}

	if got := a.EmbeddedSubAgentsDir(); got != "" {
		t.Fatalf("EmbeddedSubAgentsDir() = %q, want %q", got, "")
	}
}

func TestCapabilities(t *testing.T) {
	a := NewAdapter()

	caps := []struct {
		name string
		got  bool
		want bool
	}{
		{"SupportsOutputStyles", a.SupportsOutputStyles(), false},
		{"SupportsSlashCommands", a.SupportsSlashCommands(), false},
		{"SupportsSkills", a.SupportsSkills(), true},
		{"SupportsSystemPrompt", a.SupportsSystemPrompt(), true},
		{"SupportsMCP", a.SupportsMCP(), true},
		{"SupportsSubAgents", a.SupportsSubAgents(), false},
	}

	for _, c := range caps {
		if c.got != c.want {
			t.Errorf("%s() = %v, want %v", c.name, c.got, c.want)
		}
	}
}
