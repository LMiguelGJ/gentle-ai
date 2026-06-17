---
description: Write the spec — requirements and Given/When/Then scenarios
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-spec` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

HARD GATES:

1. Proposal must exist in the active artifact store. If missing, run `sdd-propose` first and STOP.
2. Resolve the active change: if `$ARGUMENTS` is missing or ambiguous, ask the user to choose and STOP.

WORKFLOW:

1. Load the `sdd-spec` custom agent.
2. The agent writes the spec with Given/When/Then scenarios, RFC 2119 keywords, and acceptance criteria.
3. Persist to the active artifact store.
