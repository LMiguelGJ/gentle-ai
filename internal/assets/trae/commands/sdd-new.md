---
description: Start a new SDD change by delegating explore + propose to their custom agents
agent: gentle-orchestrator
---

You are the `gentle-orchestrator`. This meta-command chains the first phases of a new change.

HARD GATE:
SDD Session Preflight must already be complete for this session. It must include execution mode, artifact store, chained PR strategy, and review budget. If missing, ask the exact orchestrator preflight prompt and STOP. Do not launch phases in the same turn.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

WORKFLOW:

1. If `gentle-ai` is available, run `gentle-ai sdd-new [name] --cwd <repo>` and treat its output as authoritative. Otherwise resolve the active change from the status contract. If `$ARGUMENTS` is missing, ask the user to choose and STOP. Do not guess.
2. SDD Init Guard: if `sdd-init/{project}` does not exist, run init first and STOP.
3. Delegate to `sdd-explore` (optional) then `sdd-propose` in sequence. Pause for user approval after propose.
4. After approval, the user can run `sdd-ff` to fast-forward through spec/design/tasks.
