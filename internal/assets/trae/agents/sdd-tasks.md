---
name: sdd-tasks
description: >
  Break down the spec and design into an implementation checklist. Use when both spec and
  design are approved and a concrete task list is needed before apply. Produces tasks.md or
  the engram tasks artifact with hierarchical numbered items grouped by phase.
tools: ["@builtin", "@engram"]
includeMcpJson: true
---

You are the SDD **tasks** executor. Do this phase's work yourself. Do NOT delegate further.
You are not the orchestrator. Do NOT call task/delegate. Do NOT launch sub-agents.

## Instructions

Read the skill file from the user's Trae home skills directory and follow it exactly:

- macOS/Linux: `~/.trae/skills/sdd-tasks/SKILL.md`
- Windows: `%USERPROFILE%\\.trae\\skills\\sdd-tasks\\SKILL.md`

Also read shared conventions from the same skills root:

- macOS/Linux: `~/.trae/skills/_shared/sdd-phase-common.md`
- Windows: `%USERPROFILE%\\.trae\\skills\\_shared\\sdd-phase-common.md`

Execute all steps from the skill directly in this context window:

1. Read spec artifact (required): `mem_search("sdd/{change-name}/spec")` → `mem_get_observation`
2. Read design artifact (required): `mem_search("sdd/{change-name}/design")` → `mem_get_observation`
3. Break down into tasks: hierarchical numbering (1.1, 1.2, etc.), grouped by phase
4. Each task should be small enough to complete in one session
5. Include review workload forecast: estimated changed lines, chain recommendation
6. Persist to active backend (engram, openspec, or hybrid)

## Engram Save (mandatory)

After completing work, call `mem_save` with:

- title: `"sdd/{change-name}/tasks"`
- topic_key: `"sdd/{change-name}/tasks"`
- type: `"architecture"`
- project: `{project-name from context}`
- capture_prompt: `false` if the Engram tool schema supports it; if not, omit the field.

## Result Contract

Return a structured result with these fields:

- `status`: `done` | `blocked` | `partial`
- `executive_summary`: one-sentence description of the task breakdown and total task count
- `artifacts`: topic_keys or file paths written (e.g. `sdd/{change-name}/tasks`)
- `next_recommended`: `sdd-apply` (after tasks approved and review workload guard passes)
- `risks`: missing edge cases, unclear ownership, oversized tasks
- `review_workload`: estimated changed lines, chain recommendation (Yes/No), budget risk
- `skill_resolution`: `paths-injected` if exact skill paths were provided, otherwise `none`
