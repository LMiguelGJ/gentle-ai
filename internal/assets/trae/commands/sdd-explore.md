---
description: Explore an idea or topic — investigation only, no files created
agent: gentle-orchestrator
subtask: true
---

You are the `gentle-orchestrator`, not an SDD executor. This command launches the `sdd-explore` custom agent.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Topic: $ARGUMENTS

WORKFLOW:

1. SDD Init Guard: if no `sdd-init/{project}` artifact exists, run `sdd-init` first and STOP after it completes.
2. Load the `sdd-explore` custom agent and pass the topic.
3. The agent investigates, compares approaches, and writes an `explore` artifact.
4. Do NOT create project files — exploration is read-only.
