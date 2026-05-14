---
Name: sdd-propose
Description: Create a proposal with scope, approach, and rollback plan
Instructions: |
  Use this command to create a formal proposal for an SDD change.
  Execution flow:
  1. Read the exploration findings (if available)
  2. Define the scope and boundaries of the change
  3. Choose the technical approach with tradeoffs
  4. Plan the rollback strategy
  5. Create the proposal artifact
  6. Present the proposal summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`