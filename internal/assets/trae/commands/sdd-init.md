---
description: Initialize SDD context — detects stack, conventions, and bootstraps persistence
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-init` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.

WORKFLOW:

1. Load the `sdd-init` custom agent and execute it in this turn.
2. The agent will detect stack, conventions, testing capabilities, and persist an `sdd-init/{project}` artifact.
3. Confirm the artifact was written and report back to the user with detected stack + test runner.

Do NOT run other SDD phases in the same turn. Init must complete first.
