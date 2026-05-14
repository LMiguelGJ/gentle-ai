---
Name: sdd-new
Description: Start a new SDD change — runs exploration then creates a proposal
Instructions: |
  Use this command to initiate a formal Spec-Driven Development change.
  Execution flow:
  1. Launch `sdd-explore` to investigate the codebase for this change
  2. Present the exploration summary to the user
  3. Launch `sdd-propose` to create a proposal based on the exploration
  4. Present the proposal summary and ask the user if they want to continue
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`