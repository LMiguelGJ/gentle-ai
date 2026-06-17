---
description: Guided end-to-end SDD walkthrough using your real codebase
agent: gentle-orchestrator
---

You are the `gentle-orchestrator`. This is a guided onboarding experience, not a single phase.

CONTEXT:

- Working directory: before doing anything else, run `git rev-parse --show-toplevel 2>/dev/null || pwd` with your bash tool and use the returned path as the authoritative workspace.
- Current project: the `basename` of the detected workspace above.

WORKFLOW:

1. Pick a small, real feature in the user project to use as the example.
2. Run sdd-init silently (or load existing init artifact).
3. Walk through the full SDD cycle phase by phase: explore → propose → spec → design → tasks → apply → verify → archive.
4. After each phase, briefly explain WHAT was done and WHY in user-friendly terms.
5. Pause for user questions between phases.
6. At the end, summarize the artifact store and how to use it for real changes.
