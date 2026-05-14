---
Name: sdd-archive
Description: Close a change and persist final state in the active artifact store
Instructions: |
  Use this command to archive an SDD change after verification passes.
  Execution flow:
  1. Read all artifacts for the change
  2. Create an archive report with final state
  3. Persist the archive to the artifact store
  4. Clean up any temporary files
  5. Present the archive summary to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`