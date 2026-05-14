---
Name: sdd-onboard
Description: Guided end-to-end walkthrough of SDD using your real codebase
Instructions: |
  Use this command for a guided walkthrough of the SDD workflow.
  Execution flow:
  1. Present an overview of the SDD workflow
  2. Walk through each phase with the user's actual codebase
  3. Explain decisions and tradeoffs as they arise
  4. Answer questions about patterns and conventions
  5. Complete a full SDD cycle if time permits
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`