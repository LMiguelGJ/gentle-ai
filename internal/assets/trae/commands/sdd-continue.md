---
description: Continue the next SDD phase in the dependency chain
agent: gentle-orchestrator
---

Follow the SDD orchestrator workflow to continue the active change.

HARD GATE:
SDD Session Preflight must already be complete for this session. It must include execution mode, artifact store, chained PR strategy, and review budget. If missing, ask the exact orchestrator preflight prompt and STOP. Do not launch the next phase in the same turn.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

WORKFLOW:

1. If `gentle-ai` is available, run `gentle-ai sdd-continue [name] --cwd <repo>` and treat its output as authoritative. Otherwise resolve the active change from the status contract.
2. Determine the next phase needed based on the dependency graph: proposal → [specs ∥ design] → tasks → apply → verify → archive.
3. Launch the appropriate custom agent for the next phase. Pause for user approval after each non-trivial phase.
4. If the next phase is apply, also enforce the review workload guard before launching.
