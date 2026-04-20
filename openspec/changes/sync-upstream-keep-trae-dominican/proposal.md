# Proposal: Sync Upstream + Keep Trae Dominican Personality

## Intent

This fork added Trae agent support and DominicanRD personality to the original repo. The upstream (`Gentleman-Programming/gentle-ai`) has advanced significantly (~70+ commits). This change merges upstream while preserving our Trae platform-specific paths and all Dominican personality assets.

## Scope

### In Scope
- Fetch and merge `origin/main` into current branch
- Preserve our Trae adapter multi-platform settings path implementation (vs upstream's simple `.trae/settings.json`)
- Preserve all 4 DominicanRD persona variants (claude, generic, opencode, trae)
- Preserve Trae SDD command assets (sdd-new, sdd-continue, sdd-ff, sdd-onboard)
- Compile Windows binary and replace user's existing `gentle-ai.exe`

### Out of Scope
- Re-testing the entire suite — only verify build succeeds
- Modifying upstream features not related to our changes

## Capabilities

### New Capabilities
- `sync-upstream`: Bringing in upstream v1.11.0 changes including Strict TDD, SDD auto-init, antigravity/kimi/qwen agents, opencode profile sync

### Modified Capabilities
- `trae-agent`: Platform-specific settings path is MORE correct than upstream's simple path approach

## Approach

1. `git fetch origin main`
2. `git merge origin/main --no-edit` — our changes (Trae + Dominican persona) are already rebased on top
3. Resolve any conflicts (expected: minimal, our changes are orthogonal to most upstream work)
4. `go build -ldflags="-s -w" -o gentle-ai.exe ./cmd/gentle-ai/`
5. Replace user's binary at `%LOCALAPPDATA%\gentle-ai\gentle-ai.exe` or wherever it's installed

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/agents/trae/adapter.go` | Modified | Keep ours (multi-platform), upstream has simple path |
| `internal/agents/trae/adapter_test.go` | Modified | Keep ours (multi-platform tests) |
| `internal/assets/claude/persona-gentleman.md` | Modified | DominicanRD voseo — no upstream changes expected |
| `internal/assets/generic/persona-gentleman.md` | Modified | DominicanRD — no upstream changes expected |
| `internal/assets/opencode/persona-gentleman.md` | Modified | DominicanRD — no upstream changes expected |
| `internal/assets/trae/persona-gentleman.md` | Modified | DominicanRD — no upstream changes expected |
| `internal/assets/trae/commands/*.md` | New | Trae SDD commands — preserved |
| `internal/app/app.go` | Modified | `--skip-upgrade` flag — our addition, preserve |
| `internal/cli/run.go` | Modified | `OverrideUserHomeDirForTest` — our addition, preserve |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Merge conflicts in shared files | Low | Our changes are additive + orthogonal; conflicts unlikely |
| Build fails on Windows | Low | Test `go build` before replacing binary |
| Trae settings path breaks on new upstream | Medium | Our adapter.go is kept intact; compare after merge |

## Rollback Plan

If merge produces broken build:
1. `git merge --abort` — revert to pre-merge state
2. User's binary unchanged — no action needed

If binary replacement breaks runtime:
1. User has the old binary backed up implicitly via git state
2. `git reset --hard HEAD~1` restores pre-merge commit

## Dependencies

- Git Bash / terminal with git access to `origin`
- Go 1.24 toolchain
- Write access to user's local install directory

## Success Criteria

- [ ] `git fetch` + `git merge origin/main` completes without manual conflict resolution
- [ ] `go build -o gentle-ai.exe ./cmd/gentle-ai/` succeeds with zero errors
- [ ] Windows `.exe` replaces existing installation and runs without crash on startup
- [ ] Trae adapter works (settings path preserved)
- [ ] Dominican persona still speaks in character