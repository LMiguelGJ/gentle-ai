---
description: Read-only structured status for the active change
agent: gentle-orchestrator
---

You are the `gentle-orchestrator`. This is a read-only status command — no phase work.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.
- Change name: $ARGUMENTS

WORKFLOW:

1. If `gentle-ai` is available, run `gentle-ai sdd-status [name] --cwd <repo> --json --instructions` and surface the JSON.
2. Otherwise resolve the active change from the status contract and inspect artifact store manually.
3. Report: active change, artifact paths, task progress, dependency states, next recommended action, blocked reasons, actionContext.
4. Do NOT modify anything. Status only.
