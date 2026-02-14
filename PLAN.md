# Plan: Fix Test Failures

## Problem

CI tests failing due to:
1. Missing documentation sections in `docs/cli-reference.md` and `docs/configuration.md`
2. Fragment `act/write_agents_md.md` exceeds 100-word limit (178 words)

## Solution

### Issue 1: Add Missing Documentation Sections

**File: docs/cli-reference.md**
- Add "Global Flags" section with `--ai-cmd-alias`, `--max-iterations`, `--context`, `--config`, `--verbose`, `--dry-run`

**File: docs/configuration.md**
- Add "Configuration Tiers" section explaining: built-in defaults, global config (`~/.config/rooda/rooda-config.yml`), workspace config (`./rooda-config.yml`), CLI flags

### Issue 2: Reduce Fragment Word Count

**File: internal/prompt/fragments/act/write_agents_md.md**
- Current: 178 words
- Target: ≤100 words
- Strategy: Remove verbose explanations, condense actions list, remove example format block

## Implementation Order

1. Fix `docs/cli-reference.md` - add Global Flags section
2. Fix `docs/configuration.md` - add Configuration Tiers section
3. Refactor `act/write_agents_md.md` - reduce to ≤100 words
4. Run `make test` to verify fixes
5. Commit and push

## Acceptance Criteria

- `make test` passes without errors
- All documentation structure tests pass
- Fragment word count test passes
- No functionality changes, only documentation/prompt updates
