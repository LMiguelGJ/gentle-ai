---
title: Trae IDE Integration — Analysis and Improvement Plan
source_url: https://docs.trae.ai/ide/trae-solo-is-now-available?_lang=en
local_install_evidence: C:\Users\luism\.trae
status: analysis-and-plan
author: gentle-ai scout (2026-06-17)
related: docs/trae-solo-is-now-available.md, docs/kiro.md, docs/agents.md
---

# Trae IDE Integration — Analysis and Improvement Plan

## TL;DR

The current Trae integration in gentle-ai is **functional but superficial**.
The adapter exists, the catalog entry is correct, but:

- **3 path bugs** point to outdated or wrong locations.
- **2 capability flags are wrong** (Trae has them, adapter says no).
- **Zero dedicated assets** — no `sdd-orchestrator.md`, no persona, no
  commands, no embedded sub-agents. The Go embed pattern explicitly omits
  `trae`.
- **No support for modern Trae features**: `builtin_skills/`, `commands/`,
  `extensions/`, `memory/`, `memory-config.json`, `skill-config.json`,
  marketplace integration, hooks, custom agents, SOLO mode.

Trae is on the same level as Kiro, Antigravity, and Cursor (IDE-fork agents),
yet its integration is closer to OpenClaw in depth. Kiro has 14 embedded
agents and multi-mode model assignments; Trae has nothing.

**This plan turns Trae into a first-class IDE-fork integration** comparable
to Kiro, with extras where Trae is stronger (marketplace, built-in skills,
extensions, memory subsystem).

---

## Part 1 — Evidence

### 1.1 What `docs/agents.md` and the Trae adapter currently claim

From `internal/agents/trae/adapter.go` and `docs/agents.md`:

| Capability                | Adapter says | docs.trae.ai says | Real `~/.trae` shows |
|---------------------------|--------------|-------------------|----------------------|
| Skills                    | ✅ Yes       | ✅ Yes            | ✅ `skills/` exists  |
| System prompt / Rules     | ✅ Yes       | ✅ Rules          | ⚠ `user_rules/` is a **directory**, not `user_rules.md` |
| MCP                       | ✅ Yes       | ✅ MCP            | ✅ `mcp.json` exists, **plus** `mcps/` directory |
| Output styles             | ❌ No        | ❌ Not mentioned  | ❌ Not present       |
| Slash commands            | ❌ No        | ✅ Commands       | ✅ `commands/` exists |
| Sub-agents (native)       | ❌ No        | ✅ Custom agents  | ⚠ No `agents/` but `builtin_skills/` + Skill mechanism |
| Hooks                     | ❌ No        | ✅ Hooks          | ❌ No `hooks/` dir seen |
| Memory subsystem          | ❌ No        | ✅ Memories       | ✅ `memory/` + `memory-config.json` |
| Built-in skills           | ❌ No        | ✅ Built-in       | ✅ 4 skills (`TRAE-code-review`, `TRAE-debugger`, `TRAE-generate-mini-app`, `TRAE-dynamic-ui`) |
| Marketplace               | ❌ No        | ✅ Skills         | ✅ `marketplace` referenced in `skill-config.json` |
| VSCode-style extensions   | ❌ No        | ❌ Not mentioned  | ✅ 14+ extensions in `extensions/` (claude-code, gitlens, ruff, prettier, etc.) |
| Tools / worktrees         | ❌ No        | ❌ Not mentioned  | ✅ `tools/`, `worktrees/` exist |
| Multi-mode SDD            | ❌ No        | ❌ Not mentioned  | ❌ Not present       |

### 1.2 Real `~/.trae` layout (Windows, from this user)

```text
~/.trae/
├── .nomedia
├── argv.json                      ← locale + crash reporter
├── mcp.json                       ← ACTIVE MCP servers (Excel, MiniMax, context7, engram)
├── memory-config.json             ← memory subsystem toggle per workspace
├── skill-config.json              ← disabledSkills + builtinSkillStatus + managedSkills
├── trae-jwt-token
├── binaries/node/24.14.1/         ← embedded Node 24 + Claude Code CLI + firecrawl
├── builtin/                       ← built-in resources (code/, global/, trae/, work/)
├── builtin_skills/                ← 4 built-in skills (TRAE-* prefix)
├── commands/                      ← slash commands
├── extensions/                    ← VSCode-compatible extensions marketplace
├── mcps/                          ← additional MCP servers (directory, not just mcp.json)
├── memory/                        ← memory subsystem data
├── skills/                        ← user-installed skills
├── tools/                         ← tool definitions
├── user_rules/                    ← rules are a DIRECTORY (not a file)
└── worktrees/                     ← git worktrees for isolation
```

### 1.3 Real `skill-config.json` (from this user)

```json
{
  "disabledSkills": [],
  "builtinSkillStatus": {
    "TRAE-dynamic-ui": false,
    "TRAE-debugger": true,
    "TRAE-generate-mini-app": true,
    "TRAE-code-review": true
  },
  "managedSkills": {
    "gh-cli": "marketplace",
    "web-design-guidelines": "marketplace"
  }
}
```

This proves:

- Trae has a **granular enable/disable** model for built-in skills.
- Trae has a **marketplace** with managed install provenance.

Neither is exposed by gentle-ai today.

### 1.4 Real `mcp.json` (from this user)

User already has Engram MCP wired (`command: engram, args: [mcp, --tools=agent]`)
and context7. gentle-ai currently writes MCP config to
`%APPDATA%\Trae\User\mcp.json` (per `MCPStrategy: StrategyMCPConfigFile`),
which is the **OS-specific app config dir**, but the user's active MCP
config is at `~/.trae/mcp.json`. **Path bug** — the user's MCP config is in
a different root than the adapter writes to.

### 1.5 Adapter code review (key issues)

`internal/agents/trae/adapter.go`:

```go
// BUG 1: points to a file, but real layout is a directory
func (a *Adapter) SystemPromptFile(homeDir string) string {
    return filepath.Join(a.traeUserDir(homeDir), "user_rules.md")
}
// → should be ~/.trae/user_rules/<something>.md or a StrategyMarkdownSections-friendly entrypoint

// BUG 2: writes to OS-specific app dir, but Trae also reads from ~/.trae/mcp.json
func (a *Adapter) MCPConfigPath(homeDir string, _ string) string {
    return filepath.Join(a.traeUserDir(homeDir), "mcp.json")
}
// → most users (and the doc site) expect ~/.trae/mcp.json
// → on macOS: ~/Library/Application Support/Trae/User/mcp.json (probably unused)

// BUG 3: claims no slash commands, but ~/.trae/commands/ exists in every install
func (a *Adapter) SupportsSlashCommands() bool { return false }
func (a *Adapter) CommandsDir(_ string) string  { return "" }
// → should be true; CommandsDir should return ~/.trae/commands/
```

`internal/assets/assets.go`:

```go
// BUG 4: explicit embed pattern omits trae entirely
//go:embed all:claude all:opencode all:generic all:skills all:gga all:gemini all:codex all:antigravity all:windsurf all:cursor all:kimi all:qwen all:kiro all:hermes
var FS embed.FS
// → no all:trae, so internal/assets/trae/ is never created (and would be empty today anyway)
```

`internal/catalog/agents.go`:

```go
// Trae is registered correctly here, with proper TierFull.
{ID: model.AgentTrae, Name: "Trae IDE", Tier: model.TierFull, ConfigPath: "~/.trae"},
```

This part is fine.

`internal/installcmd/resolver.go` (and `install_deps.go`):

- No Trae-specific branch. Trae falls through to the default path which uses
  the adapter's `SkillsDir`, `MCPConfigPath`, and `SystemPromptFile`. Because
  of bugs 1–3, this default path produces **stale or incorrect installs**.

### 1.6 Comparison vs Kiro (the gold standard for IDE-fork support)

| Aspect                         | Kiro                                    | Trae (current)                          |
|--------------------------------|-----------------------------------------|-----------------------------------------|
| Embedded sub-agent assets      | 14 (`sdd-*.md`, `jd-*.md`, `review-*`)  | 0                                       |
| Persona asset                  | `kiro/persona-gentleman.md`             | None                                    |
| Orchestrator asset             | `kiro/sdd-orchestrator.md`              | None                                    |
| Slash commands                 | ❌ No                                   | ❌ No (should be ✅ Yes)                |
| Multi-mode SDD                 | ✅ `KiroModelID()` alias resolver       | ❌ None                                 |
| Split-root layout handling     | ✅ Documented in adapter comments       | ⚠ Half-documented (only skills/MCP split, not rules) |
| Custom agents per-phase        | ✅ Native sub-agents at `~/.kiro/agents/`| ❌ Disabled                             |
| Hooks support                  | ❌ No                                   | ❌ No (Trae supports them — opportunity) |
| Built-in skill enable/disable  | N/A (Kiro has none)                     | ❌ Not managed by gentle-ai             |
| Marketplace integration        | N/A                                     | ❌ Not managed by gentle-ai             |
| Memory subsystem               | N/A                                     | ❌ Not managed by gentle-ai             |

Kiro is the right reference: same product category (VSCode-derived IDE),
gentle-ai invested deeply in it. Trae deserves the same investment.

---

## Part 2 — Plan

Prioritized by impact vs. effort. Each item lists: **what**, **why**, **where**, **effort**, **risk**.

### P0 — Fix the existing adapter (no new features, just correctness)

These are bugs, not improvements. Without them, every other item is built on
a broken foundation.

#### P0.1 — Fix `SystemPromptFile` to handle the `user_rules/` directory

- **What**: change the adapter so it can write to `~/.trae/user_rules/`
  (a markdown file inside the directory, e.g. `~/.trae/user_rules/gentle-ai.md`)
  using `StrategyMarkdownSections`. Keep `StrategyMarkdownSections` so existing
  user content is preserved with markers.
- **Why**: real Trae layout uses a directory. Writing to a non-existent file
  in `%APPDATA%\Trae\User\user_rules.md` is silently ignored by Trae.
- **Where**: `internal/agents/trae/adapter.go` (`SystemPromptFile`,
  `SystemPromptDir`).
- **Effort**: XS (1 file, ~10 lines).
- **Risk**: low; `StrategyMarkdownSections` already handles marker-based
  merge.

#### P0.2 — Fix `MCPConfigPath` to use `~/.trae/mcp.json`

- **What**: change `MCPConfigPath` to return `~/.trae/mcp.json`. Optionally
  keep the `%APPDATA%\Trae\User\mcp.json` path as a secondary write target
  for backwards compatibility (only if Trae actually reads both — verify
  first).
- **Why**: real user MCP configs live at `~/.trae/mcp.json`. gentle-ai's
  writes to `%APPDATA%` are invisible to Trae.
- **Where**: `internal/agents/trae/adapter.go` (`MCPConfigPath`).
- **Effort**: XS.
- **Risk**: low. Worst case: existing users with old `%APPDATA%` configs
  see them orphaned (acceptable; they were never read anyway).

#### P0.3 — Verify on macOS / Linux paths

- **What**: confirm `~/.trae/` is the canonical cross-platform location
  (not OS-specific). The doc site suggests it is; this user's Windows
  install confirms it.
- **Why**: if Trae uses `~/.trae/` on all platforms (not `%APPDATA%`),
  then `MCPConfigPath` should never use `traeUserDir` at all.
- **Where**: `internal/agents/trae/adapter.go`. Likely simplify by
  removing `traeUserDir` for MCP and using `GlobalConfigDir` directly.
- **Effort**: XS (after P0.2).
- **Risk**: low; verify with web research before merging.

### P1 — Enable capabilities Trae actually has

These unlock features with low effort.

#### P1.1 — Enable slash commands

- **What**: flip `SupportsSlashCommands: true`, return
  `~/.trae/commands/` from `CommandsDir`. Add `internal/assets/trae/commands/`
  with the SDD commands (`sdd-init.md`, `sdd-explore.md`, etc., and any
  `/gentle-ai:*` aliases). Add `all:trae` to the embed pattern.
- **Why**: Trae has `commands/`, the adapter says false, docs.trae.ai
  confirms it. We miss an easy win.
- **Where**: `internal/agents/trae/adapter.go`,
  `internal/assets/assets.go`, `internal/assets/trae/commands/`.
- **Effort**: S–M (~6 command files, copy from existing kiro or opencode
  slash commands and reformat if needed).
- **Risk**: low; we already ship slash commands for OpenCode and Qwen.

#### P1.2 — Add a Trae-dedicated `sdd-orchestrator.md`

- **What**: ship `internal/assets/trae/sdd-orchestrator.md` with Trae-aware
  paths (`~/.trae/skills/`, `~/.trae/commands/`, `~/.trae/extensions/`),
  similar to how Kiro has its own. Also add
  `internal/assets/trae/persona-gentleman.md`.
- **Why**: Kiro has both; Trae has neither. The orchestrator is the
  identity of the integration.
- **Where**: `internal/assets/trae/`, embed pattern in `assets.go`.
- **Effort**: S (copy from Kiro, edit paths).
- **Risk**: low.

#### P1.3 — Add `EmbeddedSubAgentsDir` and per-phase custom agents

- **What**: ship a small set of Trae custom agents (analogous to Kiro's
  14). Start with the SDD phases (`sdd-explore.md`, `sdd-apply.md`,
  `sdd-verify.md`) and `jd-judge-a.md`, `jd-judge-b.md`, `jd-fix-agent.md`.
- **Why**: Trae supports custom agents via the Skill mechanism. Even if
  Trae doesn't call them "sub-agents" in the strict sense, the orchestrator
  can reference them and Trae will load them.
- **Where**: `internal/assets/trae/agents/` (new), adapter flags.
- **Effort**: M (8–14 files, content adapted from Kiro).
- **Risk**: medium; needs testing on real Trae.

### P2 — Manage Trae-specific config files

These let gentle-ai actively shape the Trae installation, not just install
files into directories it forgets about.

#### P2.1 — Manage `~/.trae/skill-config.json`

- **What**: add a `strategy` that can read/merge `skill-config.json` to
  enable useful built-in skills (`TRAE-debugger`, `TRAE-code-review`) and
  register gentle-ai-managed skills under `managedSkills`. Read-modify-write
  with marker protection, similar to `StrategyMarkdownSections`.
- **Why**: this is the only way to actually turn on Trae's built-in skills.
- **Where**: new `StrategyTraeSkillConfig` in `internal/model/`, wired
  from the Trae adapter.
- **Effort**: M (new strategy, tests, wiring).
- **Risk**: medium; must not clobber user customizations.

#### P2.2 — Manage `~/.trae/memory-config.json`

- **What**: optionally enable Trae's memory subsystem per workspace, or
  document how to make it coexist with Engram. Most users would prefer
  Engram as the single source of memory, so the default should be
  "disable Trae native memory, keep Engram".
- **Why**: avoids two memory systems fighting each other.
- **Where**: same strategy pattern as P2.1.
- **Effort**: S.
- **Risk**: low; default is conservative (do nothing).

#### P2.3 — Surface built-in skills in the TUI / installer

- **What**: in the Trae install/sync flow, show a checklist of Trae
  built-in skills (`TRAE-debugger`, `TRAE-code-review`,
  `TRAE-generate-mini-app`, `TRAE-dynamic-ui`) and let the user toggle
  them before gentle-ai writes `skill-config.json`.
- **Why**: this is a unique differentiator for Trae — no other IDE has
  built-in skills we can manage.
- **Where**: TUI menu (likely in `internal/tui/`), backed by catalog data
  in `internal/catalog/skills.go` with a new `Builtin: true` flag.
- **Effort**: M.
- **Risk**: medium; new UX, needs design review.

### P3 — Differentiate from Kiro (where Trae is actually stronger)

#### P3.1 — Marketplace integration

- **What**: when a marketplace skill is detected in `managedSkills`,
  surface it in gentle-ai's installer as "installed via marketplace" and
  offer to install complementary skills from the same marketplace source.
- **Why**: Trae has a marketplace, Kiro does not. This is a Trae-only
  feature we can claim.
- **Where**: catalog + TUI.
- **Effort**: M.
- **Risk**: low.

#### P3.2 — Extensions awareness

- **What**: read `~/.trae/extensions/` and warn / recommend installing
  relevant extensions (`anthropic.claude-code-*`, `bradlc.vscode-tailwindcss-*`,
  `charliermarsh.ruff-*`, etc.). gentle-ai does not need to install them,
  but it can document a recommended set.
- **Why**: users with Trae install extensions manually; gentle-ai could
  be the curator.
- **Where**: new `docs/trae-recommended-extensions.md` +
  `gentle-ai install --agent trae --suggest-extensions` flag.
- **Effort**: S (docs + a list of recommendations).
- **Risk**: low.

### P4 — Align with Trae Work / SOLO mode (forward-looking)

This is speculative — the Trae Work standalone client is brand new and we
don't yet know its config layout. Defer until more is known, but keep
hooks in the adapter for future expansion.

#### P4.1 — Detect TRAE Work standalone client

- **What**: detect `work.trae.ai` install (different from IDE), or detect
  the "Work" mode flag. Probably out of scope for the install adapter,
  but document the difference in `docs/trae-solo-is-now-available.md` or
  a follow-up.
- **Where**: TBD; no code change yet.
- **Effort**: TBD.

#### P4.2 — Configure Work vs Code mode

- **What**: same as above. The user picks which mode to install
  (gentle-ai for Code mode; future support for Work mode).
- **Where**: TBD.
- **Effort**: TBD.

---

## Suggested Execution Order

1. **P0.1, P0.2, P0.3** — adapter path fixes (XS, can ship together).
2. **P1.2** — `sdd-orchestrator.md` + persona asset (S, immediate win).
3. **P1.1** — slash commands + embed (S–M).
4. **P1.3** — per-phase custom agents (M).
5. **P2.1, P2.2** — `skill-config.json` + `memory-config.json` strategies (M).
6. **P2.3** — TUI checklist for built-in skills (M, UX-heavy).
7. **P3.1, P3.2** — marketplace + extensions awareness (M).
8. **P4.* — defer until Trae Work config is documented.**

Total P0+P1+P2 estimate: ~2–3 days of focused work, parallelizable.

---

## Files to Touch (summary)

| File                                                      | Change                              |
|-----------------------------------------------------------|-------------------------------------|
| `internal/agents/trae/adapter.go`                         | P0.1, P0.2, P0.3, P1.1, P1.3        |
| `internal/assets/assets.go`                               | Add `all:trae` to embed pattern     |
| `internal/assets/trae/sdd-orchestrator.md`                | New (P1.2)                          |
| `internal/assets/trae/persona-gentleman.md`               | New (P1.2)                          |
| `internal/assets/trae/commands/sdd-*.md`                  | New (P1.1)                          |
| `internal/assets/trae/agents/sdd-*.md`, `jd-*.md`         | New (P1.3)                          |
| `internal/model/strategy.go` (or equivalent)              | New strategies (P2.1, P2.2)         |
| `internal/installcmd/resolver.go`                         | Wire new strategies (P2.x)          |
| `internal/catalog/skills.go`                              | Add `Builtin` flag (P2.3)           |
| `internal/tui/...`                                        | Built-in skills checklist (P2.3)    |
| `docs/agents.md`                                          | Update Trae row (capabilities, paths) |
| `docs/trae-solo-is-now-available.md`                      | Cross-link from this plan           |
| `docs/trae-recommended-extensions.md`                     | New (P3.2)                          |
| `internal/agents/trae/adapter_test.go`                    | Update tests for P0.x, P1.1         |

---

## Open Questions (need user / upstream input)

1. **Does Trae actually read `%APPDATA%\Trae\User\mcp.json` on Windows?**
   This user's MCP lives at `~/.trae/mcp.json`. If Trae only reads the
   home-root, we can drop the OS-specific path entirely.
2. **Does Trae support custom agents as files in `~/.trae/skills/`?**
   The adapter says no sub-agents, but the Skill mechanism could host
   custom-agent content. Real install shows no `~/.trae/agents/` dir.
3. **What does `~/.trae/worktrees/` actually do?** Git-worktree-based
   isolation could be a foundation for a "soft sub-agent" model for Trae.
4. **Trae Work standalone clients — what's the config layout?**
   Source: <https://docs.trae.ai/solo/what-is-trae-solo?_lang=en>. Defer.

---

## References

- `docs/trae-solo-is-now-available.md` — announcement + site structure
- `docs/kiro.md` — gold-standard IDE-fork integration to mirror
- `docs/agents.md` — agent matrix (Trae row needs updating)
- `docs/antigravity-sdd-workaround.md` — solo-agent workaround pattern
- `docs/engram.md` — Engram integration pattern
- `internal/agents/trae/adapter.go` — current adapter
- `internal/agents/kiro/adapter.go` — reference adapter
- `internal/agents/factory.go`, `registry.go` — adapter wiring
- `internal/assets/assets.go` — embed pattern
- `internal/catalog/agents.go` — agent registry
