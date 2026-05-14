---
Name: sdd-design
Description: Create a design with architecture decisions and data flow
Instructions: |
  Use this command to create a formal design for an SDD change.
  Execution flow:
  1. Read the proposal and spec artifacts
  2. Make architecture decisions with rationale
  3. Document the data flow and interfaces
  4. Specify testing strategy
  5. Create the design artifact
  6. Present the design summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`