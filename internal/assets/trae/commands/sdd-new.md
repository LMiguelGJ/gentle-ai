---
description: Inicia un cambio SDD — corre exploration y luego crea la proposal
agent: sdd-orchestrator
---

FORM FIELDS (Trae Command):
- Name: SDD: Nuevo Cambio
- Description: Inicia un cambio SDD con exploration + proposal
- Instructions (mínimo): `/sdd-new <change-name>`

Instrucciones (completas para pegar en el campo Instructions):

```
/sdd-new <change-name>

GUARD:
- Antes de ejecutar cualquier fase, verificá si existe sdd-init para este proyecto:
  - mem_search(query: "sdd-init/{project}", project: "{project}")
  - Si NO existe, corré sdd-init primero y recién después seguí.

WORKFLOW:
1. Lanzá sdd-explore para investigar el codebase del cambio
2. Mostrá el resumen de exploración
3. Lanzá sdd-propose para crear la propuesta basada en la exploración
4. Mostrá el resumen de la propuesta y preguntá si seguir con specs/design

CONTEXT:
- Working dir: !`echo -n "$(pwd)"`
- Proyecto: !`echo -n "$(basename $(pwd))"`
- Cambio: <change-name>
- Artifact store: engram
- Modo por defecto: interactive (si el usuario no define)

ENGRAM:
- Cada fase persiste en topic_key: "sdd/<change-name>/{type}"
- type: explore, proposal, spec, design, tasks, apply-progress, verify-report, archive-report

NOTAS:
- No ejecutes trabajo de fase inline; delegá a sub-agentes
- Preguntá modo (auto/interactive) y artifact store (engram/openspec/hybrid) si no están definidos
```

