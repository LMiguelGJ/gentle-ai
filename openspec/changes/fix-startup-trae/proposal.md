# Proposal: fix-startup-slowness-and-sync-upstream

## Intent

El binary de gentle-ai tarda demasiado en iniciar (múltiples segundos) en Windows. El flujo de startup ejecuta `detectSystem()` que a su vez llama `DetectDependencies()` sin timeout por comando individual — particularmente `npm config get prefix` puede colgar y bloquear todo. Además, el upstream avanzó ~30 commits desde nuestra última sync y hay nuevos agentes (Pi, Openclaw) que necesito traer sin perder Trae.

## Scope

### In Scope
- Traer cambios nuevos de `origin/main` (upstream v1.28.2)
- Preservar soporte Trae IDE (adapter + todos los archivos wired)
- Identificar y fixear el bottleneck de startup en `detectSystem()` / `DetectDependencies()`
- Compilar y reemplazar binary en `C:\Users\luism\bin\gentle-ai.exe`
- Subir resultado a fork `myfork/feat/trae-from-upstream`

### Out of Scope
- No cambiar comportamiento de install/sync/upgrade (solo startup)
- No agregar nuevos features

## Capabilities

### Modified Capabilities
- `system-detection`: `DetectDependencies()` agrega timeout por comando para evitar bloqueos en comandos externos (npm, etc.)

## Approach

1. `git fetch origin` + reset a `origin/main` en branch nueva `feat/fix-startup-trae`
2. Cherry-pick Trae support desde commits locales (`internal/agents/trae/`, `internal/model/types.go`, `internal/agents/factory.go`, `internal/catalog/agents.go`, `internal/cli/validate.go`, `internal/system/config_scan.go`, `internal/components/engram/setup.go`, `internal/tui/model.go`)
3. Investigar bottleneck: agregar timeouts individuales a `detectSingleDep()` en `deps.go` para comandos externos
4. Compilar: `go build -ldflags="-s -w" -o gentle-ai.exe ./cmd/gentle-ai`
5. Reemplazar `C:\Users\luism\bin\gentle-ai.exe`
6. Push a `myfork/feat/fix-startup-trae`

## Affected Areas

| Area | Impact | Description |
|------|--------|-------------|
| `internal/system/deps.go` | Modified | Agregar timeout por comando en `detectSingleDep()` |
| `internal/agents/trae/adapter.go` | Keep | Trae adapter (preservar de local) |
| `internal/agents/factory.go` | Keep | Trae factory registration |
| `internal/catalog/agents.go` | Keep | Trae en allAgents |
| `internal/model/types.go` | Keep | AgentTrae constant |
| `internal/cli/validate.go` | Keep | Trae detection switch |
| `internal/system/config_scan.go` | Keep | Trae config path |
| `internal/components/engram/setup.go` | Keep | Trae slug |
| `internal/tui/model.go` | Keep | Trae loadSelection |

## Risks

| Risk | Likelihood | Mitigation |
|------|------------|------------|
| Cherry-pick conflictos con nuevos archivos de upstream | Med | Resolver manualmente, preferir upstream para archivos no-Trae |
| Timeout muy corto causa false negatives en detection | Low | Usar 3-5s por comando, no menos |

## Rollback Plan

- `git reset --hard origin/main` descarta todos los cambios
- Binary anterior no está en git — pero el usuario puede reinstalar con `winget install`

## Dependencies

- Go 1.26+ instalado
- Git remote `origin` accesible

## Success Criteria

- [ ] `gentle-ai --help` responde en <1 segundo
- [ ] `gentle-ai install --dry-run` responde en <3 segundos
- [ ] Trae aparece como agente en `gentle-ai sync`
- [ ] Binary en `C:\Users\luism\bin\gentle-ai.exe` es el nuevo build
- [ ] Todos los tests pasan (`go test ./...`)