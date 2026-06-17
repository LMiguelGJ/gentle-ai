---
description: Break down the change into an implementation checklist
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-tasks` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

HARD GATES:

1. Both spec AND design must exist in the active artifact store. If either is missing, STOP and tell the user to run that phase first.
2. Resolve the active change: if `$ARGUMENTS` is missing or ambiguous, ask the user to choose and STOP.

WORKFLOW:

1. Load the `sdd-tasks` custom agent.
2. The agent produces a hierarchical numbered task list (1.1, 1.2, etc.), grouped by phase, sized for one session each.
3. Include review workload forecast: estimated changed lines, chain recommendation, budget risk.
4. Persist to the active artifact store.
