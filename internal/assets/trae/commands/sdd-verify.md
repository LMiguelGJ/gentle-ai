---
Name: sdd-verify
Description: Validate implementation against specs and report CRITICAL/WARNING/SUGGESTION
Instructions: |
  Use this command to validate that implementation matches the specs.
  Execution flow:
  1. Read the spec, tasks, and apply-progress artifacts
  2. Read the implemented code
  3. Check each spec scenario against the implementation
  4. Report CRITICAL (missing/broken), WARNING (tech debt), SUGGESTION (improvement)
  5. Present the verification report to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`