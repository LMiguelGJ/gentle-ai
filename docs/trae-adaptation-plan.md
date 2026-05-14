# Plan: Adaptación Completa de Gentle AI a Trae IDE

## Status
- **Fecha**: 2026-05-14
- **Estado**: Planificado
- **Tipo**: New Agent Integration

---

## Resumen Ejecutivo

Gentle AI tiene un adapter para Trae IDE que actualmente soporta skills pero no slash commands. Esto significa que `/sdd-propose` funciona (es un skill) pero `/sdd-new`, `/sdd-continue`, `/sdd-ff` no (son meta-comandos del orquestador). Este plan documenta la adaptación completa para que Trae tenga paridad total con Claude Code y los demás agentes.

---

## Inventario Actual vs Necesario

### Qué tiene Trae hoy (adapter.go)

| Capacidad | Estado | Detalle |
|-----------|--------|---------|
| `Agent()` | ✅ Listo | `model.AgentTrae = "trae"` |
| `Tier()` | ✅ Listo | `TierFull` |
| `Detect()` | ✅ Listo | Detecta `~/.trae` |
| `GlobalConfigDir()` | ✅ Listo | `~/.trae` |
| `SystemPromptDir()` | ✅ Listo | `~/.trae/user_rules` |
| `SystemPromptFile()` | ✅ Listo | `~/.trae/user_rules/gentle-ai.md` |
| `SkillsDir()` | ✅ Listo | `~/.trae/skills` |
| `MCPStrategy()` | ✅ Listo | `StrategyMCPConfigFile` → `~/.trae/mcp.json` |
| `SettingsPath()` | ✅ Listo | `~/.trae/settings.json` |
| `SupportsSkills()` | ✅ Listo | `true` |
| `SupportsSystemPrompt()` | ✅ Listo | `true` |
| `SupportsMCP()` | ✅ Listo | `true` |
| `SupportsSubAgents()` | ✅ Listo | `false` |

### Qué falta

| Capacidad | Prioridad | Impacto |
|-----------|-----------|---------|
| `SupportsSlashCommands()` | 🔴 Alta | Habilita `/sdd-new`, `/sdd-continue`, `/sdd-ff` |
| `CommandsDir()` | 🔴 Alta | Destino: `~/.trae/commands/` |
| `SupportsWorkflows()` | 🟡 Media | Soporte para workflows nativos de Trae |
| `WorkflowsDir()` | 🟡 Media | Destino: `.trae/workflows/` (workspace) |
| `EmbeddedWorkflowsDir()` | 🟡 Media | Fuente: `trae/workflows` |
| Documentación de Trae en Gentle AI | 🟢 Baja | `docs/trae.md` |

---

## Canal de Slash Commands (Critical Fix)

### Cómo funciona

Trae soporta slash commands exactamente como Claude Code:
- **Global**: `~/.trae/commands/` (Windows: `%userprofile%/.trae/commands`)
- **Project**: `.trae/commands/` en raíz del proyecto (hasta 3 niveles de anidamiento)

### Formato de archivo

Trae usa Markdown con frontmatter YAML:

```yaml
---
Name: <command-name>         # Identificador único
Description: <desc>           # Descripción breve
Instructions: |
  <instrucciones>            # Pasos específicos para la IA
---

# Contenido adicional (opcional)
```

**Nota**: A diferencia de Claude que usa `agent: gentle-orchestrator` en frontmatter, Trae usa `Name`, `Description`, `Instructions`. Necesitamos adaptar el formato.

### Comandos SDD necesarios

| Comando | Descripción | Prioridad |
|---------|-------------|-----------|
| `sdd-new.md` | Iniciar nuevo cambio SDD | 🔴 Alta |
| `sdd-continue.md` | Continuar siguiente phase | 🔴 Alta |
| `sdd-ff.md` | Fast-forward: proposal→specs→design→tasks | 🔴 Alta |
| `sdd-init.md` | Inicializar contexto SDD | 🟡 Media |
| `sdd-explore.md` | Investigar idea/cambios | 🟡 Media |
| `sdd-propose.md` | Crear proposal | 🟡 Media |
| `sdd-spec.md` | Crear specs | 🟡 Media |
| `sdd-design.md` | Crear design | 🟡 Media |
| `sdd-tasks.md` | Crear tasks | 🟡 Media |
| `sdd-apply.md` | Implementar tasks | 🟡 Media |
| `sdd-verify.md` | Verificar implementación | 🟡 Media |
| `sdd-archive.md` | Archivar cambio | 🟡 Media |
| `sdd-onboard.md` | Onboarding guiado | 🟡 Media |
| `judgment-day.md` | Revisión adversarial dual | 🟡 Baja |

### Comparación: Claude vs Trae slash command format

**Claude** (`internal/assets/claude/commands/sdd-init.md`):
```yaml
---
description: Initialize SDD context...
agent: gentle-orchestrator
subtask: true
---
You are an SDD sub-agent. Read the skill file...
```

**Trae** (formato que debemos usar):
```yaml
---
Name: sdd-new
Description: Start a new SDD change — runs exploration then creates a proposal
Instructions: |
  If the native `sdd-orchestrator` agent is available, delegate this command to it.
  Otherwise, follow the SDD orchestrator workflow inline...
---
```

---

## Canal de Workflows (Enhancement)

### Contexto

Windsurf tiene `SupportsWorkflows() = true` porque Trae tiene su propio sistema de workflows nativos en `.windsurf/workflows/`. Trae también soporta workflows pero el adapter actual no lo implementa.

### Formato Trae Workflows

Según la documentación, Trae usa un formato similar a Windsurf:
- Ubicación: `.trae/workflows/` en workspace
- Formato: Markdown con frontmatter

### Decisión arquitectura

> **Pregunta**: ¿Trae debe usar workflows o slash commands para SDD?

**Respuesta**: Slash commands es suficiente y más simple. Los workflows son redundantes para Gentle AI porque:
1. SDD tiene su propio sistema de orchestrator que no necesita ser inyectado como workflow
2. Skills y slash commands ya cubren todo el flujo SDD
3. No duplicar mecanismo

**Recomendación**: No implementar `SupportsWorkflows()` por ahora — se puede hacer en un cambio futuro si hay necesidad específica.

---

## Ubicaciones de archivos de Trae

### Comparación con otros agentes

| Componente | Claude | OpenCode | Windsurf | Trae (actual) | Trae (necesario) |
|------------|--------|----------|----------|---------------|------------------|
| Config root | `~/.claude` | `~/.config/opencode` | `~/.codeium/windsurf` | `~/.trae` | `~/.trae` ✅ |
| Commands | `~/.claude/commands` | `~/.config/opencode/commands` | N/A | — | `~/.trae/commands` |
| Skills | `~/.claude/skills` | `~/.config/opencode/skills` | `~/.codeium/windsurf/skills` | `~/.trae/skills` ✅ | `~/.trae/skills` ✅ |
| System prompt | `~/.claude/CLAUDE.md` | `~/.config/opencode/AGENTS.md` | `~/.codeium/windsurf/memories/global_rules.md` | `~/.trae/user_rules/gentle-ai.md` ✅ | `~/.trae/user_rules/gentle-ai.md` ✅ |
| MCP config | `~/.claude/mcp/*.json` | `settings.json` (merged) | `mcp_config.json` | `~/.trae/mcp.json` ✅ | `~/.trae/mcp.json` ✅ |
| Settings | `~/.claude/settings.json` | `settings.json` (merged) | Platform-specific | `~/.trae/settings.json` ✅ | `~/.trae/settings.json` ✅ |
| Agents | `~/.claude/agents/*.md` | `settings.json` (merged) | N/A | N/A | N/A |
| Memories | — | — | — | — | `~/.trae/user_rules` ✅ |

---

## Estructura de activos necesarios

```
internal/assets/
├── trae/
│   ├── commands/           # NUEVO
│   │   ├── sdd-new.md
│   │   ├── sdd-continue.md
│   │   ├── sdd-ff.md
│   │   ├── sdd-init.md
│   │   ├── sdd-explore.md
│   │   ├── sdd-propose.md
│   │   ├── sdd-spec.md
│   │   ├── sdd-design.md
│   │   ├── sdd-tasks.md
│   │   ├── sdd-apply.md
│   │   ├── sdd-verify.md
│   │   ├── sdd-archive.md
│   │   ├── sdd-onboard.md
│   │   └── judgment-day.md
│   └── workflows/           # NO por ahora
└── ...
```

---

## Cambios en adapter.go

### Método 1: Slash commands estilo Trae (recomendado)

El adapter necesita retornar un path real para CommandsDir:

```go
func (a *Adapter) SupportsSlashCommands() bool {
    return true
}

func (a *Adapter) CommandsDir(homeDir string) string {
    return filepath.Join(homeDir, ".trae", "commands")
}
```

### Método 2: Mantener Skills + inyectar como Skills

Alternativamente, si Trae no puede cambiarse el formato de slash commands, los meta-comandos podrían injectarse como skills con un wrapper que los expone como slash commands.

**Decisión**: Usar Método 1 — Trae ya tiene soporte nativo para slash commands con frontmatter Name/Description/Instructions.

---

## Plan de implementación (orden específico)

### Fase 1: Core Adapter (crítico)
1. Modificar `internal/agents/trae/adapter.go`:
   - Cambiar `SupportsSlashCommands()` de `false` a `true`
   - Implementar `CommandsDir(homeDir string) string` retornando `~/.trae/commands`
   - Agregar tests para los nuevos métodos

### Fase 2: Assets de comandos
2. Crear directorio `internal/assets/trae/commands/`
3. Copiar y adaptar los archivos de `internal/assets/claude/commands/` a formato Trae
4. Formato Trae: agregar `Name`, `Description`, `Instructions` en frontmatter
5. Remover `agent: gentle-orchestrator` y `subtask: true` (no aplica en Trae)

### Fase 3: Testeo
6. Verificar que `inject.go` detecta correctamente `SupportsSlashCommands() = true`
7. Correr tests del adapter (`go test ./internal/agents/trae/...`)
8. Verificar que los archivos se copian a `~/.trae/commands/`

### Fase 4: Documentación
9. Crear `docs/trae.md` con:
   - Descripción de capacidades
   - Comandos disponibles
   - Configuraciones específicas de Trae
   - Limitaciones conocidas

### Fase 5: Compilación y distribución
10. Compilar binary
11. Reemplazar en `C:\Users\luism\bin\gentle-ai.exe`
12. Push a fork

---

## Formato de comandos Trae (spec detallado)

Cada archivo en `internal/assets/trae/commands/` sigue este formato:

```yaml
---
Name: <kebab-case-name>           # e.g., sdd-new, sdd-propose
Description: <one-line-summary>   # Visible en la UI de Trae
Instructions: |
  <multi-line-instructions>       # Lo que la IA debe hacer
---

# Contenido adicional (opcional)
# Puede incluir contexto, ejemplos, o documentación adicional
```

### Ejemplo: sdd-new.md adaptado

```yaml
---
Name: sdd-new
Description: Start a new SDD change — runs exploration then creates a proposal
Instructions: |
  Use this command to initiate a formal Spec-Driven Development change.

  Execution flow:
  1. Launch `sdd-explore` to investigate the codebase for this change
  2. Present the exploration summary to the user
  3. Launch `sdd-propose` to create a proposal based on the exploration
  4. Present the proposal summary and ask the user if they want to continue

  Context:
  - Working directory: !`echo -n "$(pwd)"`
  - Current project: !`echo -n "$(basename $(pwd))"`
  - Change name: $ARGUMENTS
  - Execution mode: ask/cache per orchestrator
  - Artifact store mode: ask/cache per orchestrator

  If the user requests fast-forward mode, skip directly to sdd-ff.
---

## SDD Orchestrator

You are coordinating a Spec-Driven Development workflow.
```

---

## Dependencias y orden de ejecución

```
adapter.go (SupportsSlashCommands + CommandsDir)
    ↓
trae/commands/*.md (assets)
    ↓
inject.go (automatico — usa CommandsDir())
    ↓
tests
    ↓
docs/trae.md
    ↓
binary + push
```

---

## Archivos a modificar

| Archivo | Cambio | Prioridad |
|---------|--------|-----------|
| `internal/agents/trae/adapter.go` | Agregar `SupportsSlashCommands()` y `CommandsDir()` | 🔴 Alta |
| `internal/agents/trae/adapter_test.go` | Agregar tests para nuevos métodos | 🔴 Alta |
| `internal/assets/trae/commands/` (nuevo) | Crear 14 archivos de comandos SDD | 🔴 Alta |
| `docs/trae.md` (nuevo) | Documentación de Trae en Gentle AI | 🟡 Media |
| `internal/assets/trae/` (nuevo) | Estructura de assets para Trae | 🔴 Alta |

---

## Criterios de éxito

- [ ] `/sdd-new` disponible en Trae chat
- [ ] `/sdd-continue` disponible en Trae chat
- [ ] `/sdd-ff` disponible en Trae chat
- [ ] Todos los skills SDD funcionando (`/sdd-propose`, `/sdd-spec`, etc.)
- [ ] Tests pasando: `go test ./internal/agents/trae/...`
- [ ] Binary compilado y替换 en sistema
- [ ] Push a fork exitoso

---

## Riesgos

| Riesgo | Probabilidad | Mitigación |
|--------|--------------|------------|
| Trae no carga comandos desde `~/.trae/commands/` | Baja | Verificar con instalación real; puede necesitar `.trae/commands/` (con punto) |
| Formato frontmatter incompatible | Baja | Usar formato exacto de la docs: `Name`, `Description`, `Instructions` |
| Conflictos con sistema de skills existente | Baja | Skills y commands son independientes en Trae |

---

## Notas adicionales

1. **$ARGUMENTS**: En Trae, los slash commands pueden recibir argumentos después del nombre. Por ejemplo, `/sdd-new add-dark-mode`. El placeholder `$ARGUMENTS` se reemplaza dinámicamente.

2. **Niveles de anidamiento**: Trae soporta hasta 3 niveles de directorios en commands de proyecto. Podemos usar esto para organizar: `sdd/orchestrator/sdd-new.md`.

3. **SOLO Agent**: Trae tiene un agente especial "SOLO" con commands built-in (`/plan`, `/spec`). Nuestra integración debe funcionar tanto en modo IDE normal como en SOLO.

4. **Memories y Rules**: Trae tiene sistemas de memories y rules además de skills y commands. Nuestro adapter usa `user_rules/` para system prompt, lo que sugiere que Trae mapea memories a `user_rules/` también.