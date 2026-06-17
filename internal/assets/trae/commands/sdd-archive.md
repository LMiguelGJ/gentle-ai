---
description: Archive a completed SDD change — syncs artifacts and closes the cycle
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-archive` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

HARD GATES:

1. tasks AND verify-report must exist. If either is missing, STOP and report.
2. Resolve the active change: if `$ARGUMENTS` is missing or ambiguous, ask the user to choose and STOP.

WORKFLOW:

1. Load the `sdd-archive` custom agent.
2. The agent confirms verify passed, syncs final state, updates DAG state, and writes `archive-report`.
3. Read-only on project files.
