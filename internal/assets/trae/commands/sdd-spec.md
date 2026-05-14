---
Name: sdd-spec
Description: Create a specification with scenarios and success criteria
Instructions: |
  Use this command to create a formal specification for an SDD change.
  Execution flow:
  1. Read the proposal artifact
  2. Define all requirements with scenarios
  3. Specify success criteria for each scenario
  4. Identify modified, added, and removed requirements
  5. Create the spec artifact
  6. Present the spec summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`