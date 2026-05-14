---
Name: sdd-continue
Description: Continue the next dependency-ready SDD phase for an active change
Instructions: |
  Use this command to continue an existing SDD change through its next phase.
  Execution flow:
  1. Identify the current state of the active change
  2. Determine the next dependency-ready phase (spec → design → tasks → apply → verify → archive)
  3. Launch the next phase sub-agent
  4. Present the phase result summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`