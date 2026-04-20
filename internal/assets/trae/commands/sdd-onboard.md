---
description: Walkthrough guiado del ciclo SDD completo usando el codebase real del usuario
agent: sdd-orchestrator
---

FORM FIELDS (Trae Command):
- Name: SDD: Onboard
- Description: Walkthrough end-to-end del ciclo SDD en tu proyecto real
- Instructions (mínimo): `/sdd-onboard`

Instrucciones (completas para pegar en el campo Instructions):

```
/sdd-onboard

GUARD:
- Verificá sdd-init; si falta, ejecutalo primero.

TASK:
- Guiá al usuario por un ciclo SDD real con su codebase:
  exploration → proposal → spec → design → tasks → apply → verify → archive

CONTEXT:
- Working dir: !`echo -n "$(pwd)"`
- Proyecto: !`echo -n "$(basename $(pwd))"`
- Artifact store: engram

ENGRAM (progreso de onboarding):
- mem_save(title: "sdd-onboard/{project}",
          topic_key: "sdd-onboard/{project}",
          type: "architecture",
          project: "{project}",
          content: "{estado del onboarding}")
- topic_key permite upsert — re-ejecutar actualiza, no duplica.

RESULTADO:
- Devolvé: status, executive_summary, artifacts, next_recommended
```

