---
Name: sdd-explore
Description: Investigate an idea — reads codebase, compares approaches, creates no files
Instructions: |
  Use this command to investigate a feature idea or problem before planning.
  Execution flow:
  1. Read the codebase to understand the current structure
  2. Identify relevant files and patterns
  3. Compare alternative approaches
  4. Create an exploration summary with findings
  5. Present the summary to the user — no files created
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`