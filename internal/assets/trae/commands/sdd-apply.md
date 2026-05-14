---
Name: sdd-apply
Description: Implement tasks from specs and design
Instructions: |
  Use this command to implement SDD tasks from specs and design.
  Execution flow:
  1. Read the tasks, spec, and design artifacts
  2. Read existing code to understand current patterns
  3. Implement each task following the design decisions
  4. Run tests and verify compilation
  5. Mark tasks as complete in the tasks artifact
  6. Persist apply progress to the artifact store
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`