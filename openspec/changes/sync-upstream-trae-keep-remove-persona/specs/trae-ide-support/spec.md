# Delta for Trae IDE Support

## ADDED Requirements

### Requirement: Trae IDE Agent Detection

The system SHALL detect the presence of Trae IDE by checking for the `.trae` configuration directory in the user's home directory.

The system MUST check platform-specific paths:
- Windows: `%APPDATA%\trae`
- Linux: `$XDG_CONFIG_HOME/trae` or `$HOME/.config/trae`
- macOS: `~/Library/Application Support/trae`

#### Scenario: Trae detected on Windows

- GIVEN the user has Trae IDE installed on Windows
- WHEN the system scans for agent configurations
- THEN it MUST find `~/.trae` (resolved via `os.UserConfigDir`)
- AND report Trae as a available agent

#### Scenario: Trae not installed

- GIVEN the user does not have Trae IDE installed
- WHEN the system scans for agent configurations
- THEN it MUST return `found: false` for Trae
- AND no error SHALL be returned

#### Scenario: Trae detected on Linux

- GIVEN the user has Trae IDE installed on Linux with XDG support
- WHEN the system scans for agent configurations
- THEN it MUST find `XDG_CONFIG_HOME/trae` if `XDG_CONFIG_HOME` is set
- AND report Trae as available

### Requirement: Trae Agent Registration

The system SHALL register Trae as a first-class agent in the agent catalog with `TierFull` support.

#### Scenario: Trae in agent catalog

- GIVEN the system loads the agent catalog
- WHEN `AllAgents()` is called
- THEN Trae MUST appear in the returned list
- AND its config path MUST be `~/.trae`

### Requirement: Trae Adapter Interface Compliance

The Trae adapter MUST implement the `AgentAdapter` interface with `Agent()`, `Tier()`, and `Detect()` methods.

#### Scenario: Trae adapter returns correct agent ID

- GIVEN a Trae adapter instance
- WHEN `Agent()` is called
- THEN it MUST return `model.AgentTrae`

#### Scenario: Trae adapter returns correct tier

- GIVEN a Trae adapter instance
- WHEN `Tier()` is called
- THEN it MUST return `model.TierFull`
