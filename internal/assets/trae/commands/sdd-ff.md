---
description: Fast-forward de fases SDD de planificación — proposal → spec → design → tasks
agent: sdd-orchestrator
---

FORM FIELDS (Trae Command):
- Name: SDD: Fast-Forward
- Description: Genera proposal → spec → design → tasks en secuencia, con resumen final
- Instructions (mínimo): `/sdd-ff <change-name>`

Instrucciones (completas para pegar en el campo Instructions):

```
/sdd-ff <change-name>

GUARD:
- Verificá sdd-init; si falta, ejecutalo primero.

WORKFLOW (en cadena):
1. sdd-propose — crear proposal
2. sdd-spec — escribir especificaciones
3. sdd-design — diseño técnico
4. sdd-tasks — checklist de implementación

Presentá un resumen combinado al final (no entre fases).

CONTEXT:
- Working dir: !`echo -n "$(pwd)"`
- Proyecto: !`echo -n "$(basename $(pwd))"`
- Cambio: <change-name>
- Artifact store: engram
- Modo: auto recomendado para FF (sin pausas intermedias)

ENGRAM:
- Persistí cada fase en topic_key "sdd/<change-name>/{type}"
- type: proposal, spec, design, tasks

NOTAS:
- No ejecutes trabajo de fase inline; delegá a sub-agentes
```

