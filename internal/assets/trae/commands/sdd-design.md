---
description: Write the design — architecture, file changes, sequence diagrams
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-design` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

HARD GATES:

1. Proposal must exist in the active artifact store. If missing, run `sdd-propose` first and STOP.
2. Resolve the active change: if `$ARGUMENTS` is missing or ambiguous, ask the user to choose and STOP.

WORKFLOW:

1. Load the `sdd-design` custom agent.
2. The agent writes the design with architecture decisions, file changes, sequence diagrams for complex flows, and rationale.
3. Spec and design can run in parallel; tasks requires both.
