---
description: Create a change proposal — intent, scope, approach, rollback plan
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-propose` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

HARD GATES:

1. SDD Init Guard: `sdd-init/{project}` must exist or be created first.
2. In interactive mode, offer a proposal question round BEFORE finalizing the proposal. Prefer 3–5 product questions that uncover business rules, implications, edge cases, and product tradeoffs.
3. Resolve the active change: if `$ARGUMENTS` is missing or ambiguous, ask the user to choose and STOP. Do not guess.

WORKFLOW:

1. Load the `sdd-propose` custom agent.
2. The agent drafts the proposal, persists it, and returns a structured result.
3. Present the proposal summary and ASK for user approval before any spec/design work.
