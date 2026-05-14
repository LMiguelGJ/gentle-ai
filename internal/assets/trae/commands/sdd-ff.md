---
Name: sdd-ff
Description: Fast-forward planning: proposal → specs → design → tasks in one pass
Instructions: |
  Use this command to run all planning phases back-to-back without pausing.
  Execution flow:
  1. Launch `sdd-explore` to investigate the codebase
  2. Launch `sdd-propose` to create the proposal
  3. Launch `sdd-spec` to create the specification
  4. Launch `sdd-design` to create the design
  5. Launch `sdd-tasks` to create the implementation tasks
  6. Present the full planning summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`