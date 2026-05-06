# Tasks: sync-upstream-trae-keep-remove-persona

## Phase 1: Git Sync — Rebase with Upstream

- [x] 1.1 Run `git fetch origin` to fetch latest upstream commits
- [x] 1.2 Run `git rebase origin/main` to rebase current branch on upstream main
- [x] 1.3 Resolve any merge conflicts (prioritize upstream for non-persona files)
- [x] 1.4 Verify `git log HEAD --oneline` shows alignment with `origin/main` history

## Phase 2: Re-add Trae Adapter (Removed in Upstream)

- [x] 2.1 Create `internal/agents/trae/adapter.go` with multi-platform `.trae` path detection (APPDATA on Windows, XDG_CONFIG_HOME on Linux, ~/Library on macOS)
- [x] 2.2 Create `internal/agents/trae/adapter_test.go` with platform-aware path tests
- [x] 2.3 Verify `internal/catalog/agents.go` still includes `{ID: model.AgentTrae, ...}` after rebase
- [x] 2.4 Run `go build ./cmd/gentle-ai` to confirm Trae adapter compiles without errors

## Phase 3: Clean Dominican Persona — Neutral Style

- [x] 3.1 Overwrite `internal/assets/trae/persona-gentleman.md` (currently 0 bytes) with neutral official persona content
- [x] 3.2 Rewrite `internal/assets/generic/persona-gentleman.md` removing all DominicanRD expressions: remove Lines 17-24 (Language section with criolloRD), keep neutral Spanish/English tone without vernacular
- [x] 3.3 Verify `internal/assets/generic/persona-gentleman.md` contains NO DominicanRD terms: "klk pp", "lider", "jevi", "hay un maco", "to' frio", "sicopata", "duro", "a bueno", "dale tato"
- [x] 3.4 (Optional) Check `internal/assets/claude/persona-gentleman.md` and `internal/assets/opencode/persona-gentleman.md` — if DominicanRD present and user wants neutral there too, clean those as well

## Phase 4: Testing & Verification

- [x] 4.1 Run `go test ./...` — all tests MUST pass
- [x] 4.2 Run `go vet ./...` — zero warnings
- [ ] 4.3 Run `go build -o gentle-ai.exe ./cmd/gentle-ai` — compiles without errors
- [ ] 4.4 Run `./gentle-ai.exe --help` — verify binary works

## Phase 5: Push to Fork

- [ ] 5.1 Run `git push myfork main` — push all changes to `LMiguelGJ/gentle-ai`
- [ ] 5.2 Confirm push completes without rejection
