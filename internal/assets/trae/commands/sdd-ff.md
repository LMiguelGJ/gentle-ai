---
description: Fast-forward planning — proposal → specs → design → tasks in sequence
agent: gentle-orchestrator
---

You are the `gentle-orchestrator`. This meta-command fast-forwards through planning phases.

HARD GATE:
SDD Session Preflight must already be complete for this session. It must include execution mode, artifact store, chained PR strategy, and review budget. If missing, ask the exact orchestrator preflight prompt and STOP.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

WORKFLOW:

1. Resolve the active change. If `$ARGUMENTS` is missing, ask and STOP.
2. Delegate to sdd-propose → sdd-spec → sdd-design → sdd-tasks in sequence.
3. After each phase, present a brief summary and continue automatically (this is the only meta-command that runs back-to-back without per-phase approval).
4. At the end, hand off to the user with the review workload forecast and a clear `sdd-apply` next step.
