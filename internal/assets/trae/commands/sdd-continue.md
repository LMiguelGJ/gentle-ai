---
description: Continúa la siguiente fase SDD en la cadena de dependencias
agent: sdd-orchestrator
---

FORM FIELDS (Trae Command):
- Name: SDD: Continuar
- Description: Avanza a la siguiente fase SDD disponible para el cambio activo
- Instructions (mínimo): `/sdd-continue <change-name>`

Instrucciones (completas para pegar en el campo Instructions):

```
/sdd-continue <change-name>

GUARD:
- Verificá sdd-init; si falta, ejecutalo primero.

WORKFLOW:
1. Detectá qué artefactos existen para el cambio:
   - proposal, spec, design, tasks, apply-progress, verify-report, archive-report
2. Determiná la próxima fase según el grafo:
   proposal → [spec ∥ design] → tasks → apply → verify → archive
3. Lanzá el sub-agente correcto para esa fase
4. Mostrá resultado y pedí confirmación para continuar

CONTEXT:
- Working dir: !`echo -n "$(pwd)"`
- Proyecto: !`echo -n "$(basename $(pwd))"`
- Cambio: <change-name>
- Artifact store: engram
- Modo: interactive por defecto

ENGRAM:
- Listá artefactos: mem_search(query: "sdd/<change-name>/", project: "{project}")
- Persistencia automática por fase usando topic_key "sdd/<change-name>/{type}"

NOTAS:
- No ejecutes trabajo de fase inline; delegá a sub-agentes
```

