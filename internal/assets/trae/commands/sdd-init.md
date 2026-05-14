---
Name: sdd-init
Description: Initialize SDD context for the project — detects stack and bootstraps persistence
Instructions: |
  Use this command to initialize the SDD context for the current project.
  Execution flow:
  1. Detect the project stack (Go, Node, Python, etc.)
  2. Check for existing SDD configuration
  3. Bootstrap the artifact persistence store (engram or openspec)
  4. Cache testing capabilities and project conventions
  5. Report initialization status to the user
---

WORKING DIRECTORY: !`echo -n "$(pwd)"`
PROJECT: !`echo -n "$(basename $(pwd))"`