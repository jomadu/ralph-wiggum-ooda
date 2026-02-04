# Agent Instructions

## Work Tracking System

**System:** beads (bd CLI)

**Query ready work:**
```bash
bd ready --json
```

**Update status:**
```bash
bd update <id> --status in_progress
```

**Close issue:**
```bash
bd close <id> --reason "Completed X"
```

**Create issue with dependencies:**
```bash
bd create --title "Title" --description "Desc" --deps blocks:issue-id --priority 2
```

## Story/Bug Input

Stories and bugs are documented in `TASK.md` at project root.

## Planning System

**Draft plan location:** `PLAN.md` at project root

**Publishing mechanism:** Agent reads `PLAN.md` and runs `bd create` commands to file issues

## Build/Test/Lint Commands

**Test:** Manual verification (no automated tests)

**Build:** Not required (bash scripts are interpreted)

**Lint:**
```bash
shellcheck src/rooda.sh
```

**Verification:**
```bash
./src/rooda.sh bootstrap --max-iterations 1
bd ready --json
```

## Specification Definition

**Location:** `specs/*.md`

**Format:** Markdown specifications following JTBD structure

**Exclude:** `specs/README.md`, `specs/TEMPLATE.md`, `specs/specification-system.md`

## Implementation Definition

**Location:** `src/rooda.sh`, `src/components/*.md`, and `docs/*.md`

**Patterns:**
- `src/rooda.sh` - Main loop script
- `src/rooda-config.yml` - Procedure configuration
- `src/components/*.md` - OODA prompt components
- `docs/*.md` - User-facing documentation

**Exclude:**
- `.beads/*` (work tracking database)
- `specs/*` (specifications)
- `README.md`, `AGENTS.md`, `PLAN.md`, `TASK.md`, `LICENSE.md` (root files)

## Quality Criteria

**For specifications:**
- Clarity: Can a new user understand the framework?
- Completeness: Are all 9 procedures documented?
- Consistency: Do docs match script behavior?
- Accuracy: Do command examples work?

**For implementation:**
- Correctness: Does script execute procedures as documented?
- Robustness: Does error handling work?
- Maintainability: Is bash code readable?
- Compatibility: Works on macOS and Linux?

**Refactoring triggers:**
- Documentation contradicts script behavior
- Script fails on documented use cases
- Error messages unclear or misleading

