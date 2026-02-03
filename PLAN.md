# Repository Restructuring Plan

## Goal

Reorganize ralph-wiggum-ooda into a clean three-directory structure:
- `src/` - Implementation (rooda.sh, rooda-config.yml, prompt components)
- `specs/` - Specifications (JTBD-based specs following spec-template.md)
- `docs/` - User-facing documentation (guides, tutorials, reference)

**Purpose:** This restructuring enables the framework to use its own methodology. Once implementation is properly located in `src/` and the spec system is established in `specs/`, we can run `draft-plan-impl-to-spec` to generate specifications from the implementation—dogfooding the ralph-wiggum-ooda process itself.

## Design Decisions

### 1. Installation Philosophy: Flat Structure for Consumers

**Decision:** Consumers install files flat at their project root (current approach). The framework repository has internal organization (src/, specs/, docs/), but the installed artifact stays flat.

**Rationale:** 
- Most tools work this way (Makefile, .eslintrc, package.json, etc.)
- Existing installations continue to work without changes
- Simpler for consumers - no nested directory navigation
- Framework repo organization is for framework developers, not consumers

**Impact:** Installation instructions copy from `src/` to consumer's project root. Config file paths remain `prompts/*.md` (not `src/components/*.md`) for consumer installations.

### 2. Backward Compatibility: Non-Breaking

**Decision:** Existing installations continue working without changes. This is not a breaking release.

**Rationale:**
- Consumers have already copied files to their project root
- Those files remain valid and functional
- Restructuring only affects the framework repository, not installed artifacts
- Users can update by re-copying files when ready

**Impact:** No migration required for existing users. Optional update path provided in README.

### 3. Wrapper Script Purpose: Framework Developers Only

**Decision:** The wrapper script is for framework developers working on ralph-wiggum-ooda itself. Consumers do not use or need the wrapper.

**Rationale:**
- Consumers copy files to their project root and run `./rooda.sh` directly
- Wrapper only makes sense in the framework repository where files are in `src/`
- Keeps consumer installation simple and straightforward

**Impact:** Wrapper script created at framework repo root for developer convenience. Not included in installation instructions.

### 4. Dogfooding Timeline: Immediate After Restructuring

**Decision:** Run `draft-plan-impl-to-spec` immediately after restructuring is complete and verified.

**Rationale:**
- Validates the restructuring worked correctly
- Demonstrates the framework's capability on itself
- Generates actual specifications to populate `specs/` directory
- Provides concrete example of the framework in action

**Impact:** Dogfooding verification steps included in success criteria. Separate beads issues will be created for spec generation tasks.

## Current State Analysis

### Current Structure
```
ralph-wiggum-ooda/
├── rooda.sh                    # Main loop script (174 LOC)
├── rooda-config.yml            # Procedure definitions
├── AGENTS.md                   # Operational guide (generated)
├── README.md                   # User-facing overview + guide
├── LICENSE.md                  # License
├── .gitignore, .gitattributes  # Git config
├── prompts/                    # 26 OODA component files
│   ├── README.md               # Component documentation
│   ├── observe_*.md            # 6 observe components
│   ├── orient_*.md             # 6 orient components
│   ├── decide_*.md             # 6 decide components
│   └── act_*.md                # 4 act components
├── docs/                       # Mixed documentation
│   ├── agents-md-specification.md  # AGENTS.md format spec
│   ├── beads.md                    # Beads integration guide
│   ├── ooda-loop.md                # OODA framework explanation
│   ├── ralph-loop.md               # Original methodology
│   ├── specs.md                    # Spec system design
│   └── spec-template.md            # Template for specs
└── .beads/                     # Work tracking database
```

### Current Documentation Classification

**Implementation Documentation (should move to src/):**
- `prompts/README.md` - Documents the prompt component system

**Specification Documentation (should move to specs/):**
- `docs/specs.md` - Defines the specification system
- `docs/spec-template.md` - Template for creating specs
- `docs/agents-md-specification.md` - AGENTS.md format specification

**User-Facing Documentation (stays in docs/):**
- `README.md` - Main entry point (stays at root)
- `docs/ooda-loop.md` - OODA framework explanation
- `docs/ralph-loop.md` - Original methodology background
- `docs/beads.md` - Beads integration guide

### Key Insights

1. **This is a framework repository** - The "implementation" is the bash script and prompt components that execute the framework. The "specifications" are the methodology definitions (JTBD, spec system, AGENTS.md format).

2. **Prompt components are implementation** - The 26 markdown files in `prompts/` are not documentation—they're executable components that get composed into procedures. They belong in `src/`.

3. **No actual specs exist yet** - The `specs/` directory doesn't exist. We have documentation *about* the spec system (`docs/specs.md`, `docs/spec-template.md`) but no actual specifications following that system. **This restructuring enables us to generate them using the framework itself.**

4. **AGENTS.md is generated** - It's created by the bootstrap procedure and maintained by agents. It should stay at project root as the agent-project interface.

5. **README.md is the entry point** - It should remain at project root as the primary user-facing documentation.

6. **Dogfooding opportunity** - Once restructured, we can run `draft-plan-impl-to-spec` to analyze the implementation in `src/` and generate proper specifications in `specs/`—demonstrating the framework by using it on itself.

## Target Structure

```
ralph-wiggum-ooda/
├── README.md                   # Main entry point (unchanged)
├── AGENTS.md                   # Agent-project interface (unchanged)
├── LICENSE.md                  # License (unchanged)
├── .gitignore, .gitattributes  # Git config (unchanged)
├── .beads/                     # Work tracking (unchanged)
│
├── src/                        # Implementation
│   ├── rooda.sh                # Main loop script
│   ├── rooda-config.yml        # Procedure definitions
│   ├── README.md               # Implementation guide (from prompts/README.md)
│   └── components/             # OODA prompt components (from prompts/)
│       ├── observe_*.md        # 6 observe components
│       ├── orient_*.md         # 6 orient components
│       ├── decide_*.md         # 6 decide components
│       └── act_*.md            # 4 act components
│
├── specs/                      # Specifications (NEW)
│   ├── README.md               # Index of JTBDs, topics, specs
│   ├── TEMPLATE.md             # Template for new specs (from docs/spec-template.md)
│   ├── agents-md-format.md     # AGENTS.md specification (from docs/agents-md-specification.md)
│   └── specification-system.md # Spec system design (from docs/specs.md)
│
└── docs/                       # User-facing documentation
    ├── README.md               # Documentation index (NEW)
    ├── ooda-loop.md            # OODA framework explanation
    ├── ralph-loop.md           # Original methodology
    └── beads.md                # Beads integration guide
```

## Migration Tasks

### Phase 1: Create New Directory Structure

**Priority: Critical**
**Dependencies: None**

- [ ] Create `src/` directory
- [ ] Create `src/components/` directory
- [ ] Create `specs/` directory
- [ ] Verify directories created successfully
- [ ] Commit changes: `git commit -m "Create src/, specs/ directory structure"`

**Acceptance Criteria:**
- All three directories exist
- Directory structure matches target layout
- Changes committed to git

---

### Phase 2: Move Implementation Files to src/

**Priority: Critical**
**Dependencies: Phase 1**

- [ ] Move `rooda.sh` to `src/rooda.sh`
- [ ] Update `CONFIG_FILE` path resolution in `src/rooda.sh`
- [ ] Move `rooda-config.yml` to `src/rooda-config.yml`
- [ ] Move `prompts/README.md` to `src/README.md`
- [ ] Move all `prompts/*.md` files (except README.md) to `src/components/`
- [ ] Remove empty `prompts/` directory
- [ ] Update file permissions: `chmod +x src/rooda.sh`
- [ ] Commit changes with descriptive message

**Acceptance Criteria:**
- `src/rooda.sh` exists and is executable
- `src/rooda.sh` correctly resolves config file path relative to script location
- `src/rooda-config.yml` exists
- `src/README.md` exists (implementation guide)
- `src/components/` contains 26 component files
- Old `prompts/` directory removed
- Changes committed to git

**Path Resolution Fix:**

In `src/rooda.sh`, change:
```bash
CONFIG_FILE="rooda-config.yml"
```

To:
```bash
# Resolve config file relative to script location
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CONFIG_FILE="${SCRIPT_DIR}/rooda-config.yml"
```

This ensures the script finds its config when run from any directory, whether as `./src/rooda.sh` or via the wrapper script `./rooda`.

**Files to Move:**
```
rooda.sh                              → src/rooda.sh
rooda-config.yml                      → src/rooda-config.yml
prompts/README.md                     → src/README.md
prompts/observe_bootstrap.md          → src/components/observe_bootstrap.md
prompts/observe_plan_specs_impl.md    → src/components/observe_plan_specs_impl.md
prompts/observe_story_task_specs_impl.md → src/components/observe_story_task_specs_impl.md
prompts/observe_bug_task_specs_impl.md → src/components/observe_bug_task_specs_impl.md
prompts/observe_specs.md              → src/components/observe_specs.md
prompts/observe_impl.md               → src/components/observe_impl.md
prompts/observe_draft_plan.md         → src/components/observe_draft_plan.md
prompts/orient_bootstrap.md           → src/components/orient_bootstrap.md
prompts/orient_build.md               → src/components/orient_build.md
prompts/orient_story_task_incorporation.md → src/components/orient_story_task_incorporation.md
prompts/orient_bug_task_incorporation.md → src/components/orient_bug_task_incorporation.md
prompts/orient_gap.md                 → src/components/orient_gap.md
prompts/orient_quality.md             → src/components/orient_quality.md
prompts/orient_publish.md             → src/components/orient_publish.md
prompts/decide_bootstrap.md           → src/components/decide_bootstrap.md
prompts/decide_build.md               → src/components/decide_build.md
prompts/decide_story_task_plan.md     → src/components/decide_story_task_plan.md
prompts/decide_bug_task_plan.md       → src/components/decide_bug_task_plan.md
prompts/decide_gap_plan.md            → src/components/decide_gap_plan.md
prompts/decide_refactor_plan.md       → src/components/decide_refactor_plan.md
prompts/decide_publish.md             → src/components/decide_publish.md
prompts/act_bootstrap.md              → src/components/act_bootstrap.md
prompts/act_build.md                  → src/components/act_build.md
prompts/act_plan.md                   → src/components/act_plan.md
prompts/act_publish.md                → src/components/act_publish.md
```

---

### Phase 3: Update rooda-config.yml Path References

**Priority: Critical**
**Dependencies: Phase 2**

- [ ] Update all `observe:` paths from `prompts/` to `src/components/`
- [ ] Update all `orient:` paths from `prompts/` to `src/components/`
- [ ] Update all `decide:` paths from `prompts/` to `src/components/`
- [ ] Update all `act:` paths from `prompts/` to `src/components/`
- [ ] Verify YAML syntax is valid with `yq eval . src/rooda-config.yml`
- [ ] Commit changes: `git commit -m "Update rooda-config.yml paths to src/components/"`

**Acceptance Criteria:**
- All procedure definitions reference `src/components/*.md`
- YAML file parses correctly with `yq`
- No references to old `prompts/` directory remain
- Changes committed to git

**Changes Required:**
```yaml
# Before:
observe: prompts/observe_bootstrap.md

# After:
observe: src/components/observe_bootstrap.md
```

Apply to all 9 procedures × 4 phases = 36 path updates.

---

### Phase 4: Create Specification Files in specs/

**Priority: High**
**Dependencies: Phase 1**

- [ ] Create `specs/README.md` (minimal index structure)
- [ ] Move `docs/spec-template.md` to `specs/TEMPLATE.md`
- [ ] Move `docs/agents-md-specification.md` to `specs/agents-md-format.md`
- [ ] Move `docs/specs.md` to `specs/specification-system.md`
- [ ] Update internal cross-references in moved files
- [ ] Commit changes: `git commit -m "Create specs/ directory with specification documents"`

**Acceptance Criteria:**
- `specs/README.md` exists with minimal index structure
- `specs/TEMPLATE.md` exists (spec template)
- `specs/agents-md-format.md` exists (AGENTS.md specification)
- `specs/specification-system.md` exists (spec system design)
- All cross-references updated to new paths
- Changes committed to git

**specs/README.md Minimal Structure:**
```markdown
# Specifications

This directory will contain JTBD-based specifications for the ralph-wiggum-ooda framework.

See [TEMPLATE.md](TEMPLATE.md) for the specification structure.
See [specification-system.md](specification-system.md) for the spec system design.

## Existing Specifications
- [agents-md-format.md](agents-md-format.md) - AGENTS.md format specification

## Generating Specifications

After restructuring is complete, run `draft-plan-impl-to-spec` to analyze the implementation and generate JTBD-based specifications.
```

---

### Phase 5: Reorganize User Documentation in docs/

**Priority: High**
**Dependencies: Phase 4**

- [ ] Create `docs/README.md` (documentation index)
- [ ] Keep `docs/ooda-loop.md` (unchanged)
- [ ] Keep `docs/ralph-loop.md` (unchanged)
- [ ] Keep `docs/beads.md` (unchanged)
- [ ] Remove `docs/specs.md` (moved to specs/)
- [ ] Remove `docs/spec-template.md` (moved to specs/)
- [ ] Remove `docs/agents-md-specification.md` (moved to specs/)
- [ ] Verify source attribution in documentation files
- [ ] Commit changes: `git commit -m "Reorganize docs/ for user-facing documentation only"`

**Acceptance Criteria:**
- `docs/README.md` exists with proper navigation
- Only user-facing documentation remains in `docs/`
- No specification documents remain in `docs/`
- Documentation files properly attribute their sources of truth
- Changes committed to git

**Source Attribution:**

These documentation files are built artifacts derived from external sources of truth:

- `docs/ralph-loop.md` - Source: https://github.com/ghuntley/how-to-ralph-wiggum
- `docs/beads.md` - Source: https://github.com/steveyegge/beads
- `docs/ooda-loop.md` - Source: https://en.wikipedia.org/wiki/OODA_loop

When updating these files, consult the upstream sources to ensure accuracy and consistency. These are explanatory documents for users, not authoritative definitions.

**docs/README.md Structure:**
```markdown
# Documentation

## Getting Started
- [Main README](../README.md) - Overview and installation
- [OODA Loop](ooda-loop.md) - Understanding the framework
- [Ralph Loop](ralph-loop.md) - Original methodology

## Integration
- [Beads](beads.md) - Work tracking integration

## Reference
- [Specifications](../specs/README.md) - Framework specifications
- [AGENTS.md Specification](../specs/agents-md-format.md) - AGENTS.md format
```

**Note:** Does not link to `src/README.md` as that's implementation documentation for framework developers, not user-facing content.

---

### Phase 6a: Update Installation Instructions

**Priority: Critical**
**Dependencies: Phase 2, Phase 3**

- [ ] Update installation commands to copy from `src/` directory
- [ ] Verify commands are correct and complete
- [ ] Commit changes with descriptive message

**Acceptance Criteria:**
- Installation section shows copying from `src/` to project root
- Commands copy rooda.sh, rooda-config.yml, and components
- chmod command included for executable permission
- Changes committed to git

**Changes Required:**

```bash
# Installation
cp ralph-wiggum-ooda/src/rooda.sh .
cp ralph-wiggum-ooda/src/rooda-config.yml .
cp -r ralph-wiggum-ooda/src/components ./prompts
chmod +x rooda.sh
```

---

### Phase 6b: Update Example Commands and Sample Structure

**Priority: Critical**
**Dependencies: Phase 6a**

- [ ] Update example commands to show consumer usage pattern
- [ ] Update "Sample Repository Structure" section
- [ ] Commit changes with descriptive message

**Acceptance Criteria:**
- Example commands show `./rooda.sh` (consumer pattern)
- Sample structure shows consumer's flat installation
- Structure clearly labeled as "consumer-project/"
- Changes committed to git

**Sample Structure:**
```
consumer-project/
├── rooda.sh                   # Copied from ralph-wiggum-ooda/src/
├── rooda-config.yml           # Copied from ralph-wiggum-ooda/src/
├── AGENTS.md                  # Generated by bootstrap
├── prompts/                   # Copied from ralph-wiggum-ooda/src/components/
│   ├── observe_*.md
│   ├── orient_*.md
│   ├── decide_*.md
│   └── act_*.md
├── specs/                     # Consumer's specifications
│   └── ...
└── src/                       # Consumer's implementation
    └── ...
```

---

### Phase 6c: Update Internal Documentation Links

**Priority: High**
**Dependencies: Phase 4, Phase 5**

- [ ] Update "Learn More" section links
- [ ] Verify all links resolve correctly
- [ ] Test links in rendered markdown
- [ ] Commit changes with descriptive message

**Acceptance Criteria:**
- All documentation links point to correct new locations
- Links to specs/ directory work
- Links to docs/ directory work
- Links to src/README.md work
- All links verified by clicking in GitHub/rendered view
- Changes committed to git

**Link Updates:**
```markdown
- [OODA Loop](docs/ooda-loop.md)
- [Ralph Loop](docs/ralph-loop.md)
- [Specs System](specs/specification-system.md)
- [Spec Template](specs/TEMPLATE.md)
- [AGENTS.md Specification](specs/agents-md-format.md)
- [Prompts README](src/README.md)
```

---

### Phase 6d: Add Framework Developer vs Consumer Usage Section

**Priority: High**
**Dependencies: Phase 6a, Phase 6b, Phase 8**

- [ ] Add "For Framework Developers" section
- [ ] Add "For Framework Consumers" section
- [ ] Explain wrapper script purpose
- [ ] Clarify repository structure vs installation structure
- [ ] Commit changes with descriptive message

**Acceptance Criteria:**
- Clear distinction between framework development and consumer usage
- Wrapper script purpose explained (developers only)
- Repository structure vs installation structure clarified
- Changes committed to git

**Content:**
```markdown
## For Framework Developers

When working on ralph-wiggum-ooda itself, you can use the wrapper script:
```bash
./rooda bootstrap  # Runs ./src/rooda.sh bootstrap
```

This wrapper is for framework development convenience only and is not part of the consumer installation.

## For Framework Consumers

Copy files from `src/` to your project root per installation instructions above. The framework repository structure (with `src/`, `specs/`, `docs/`) is for organizing the framework itself—consumers copy the implementation files to their project root and use them directly.

### For Existing Installations

If you previously installed ralph-wiggum-ooda, your installation is unaffected. The files you copied to your project root remain valid and functional.

**To update to the latest version:**

1. **Pull latest changes:**
   ```bash
   cd /path/to/ralph-wiggum-ooda
   git pull origin main
   ```

2. **Copy updated files to your project:**
   ```bash
   cd /path/to/your-project
   cp /path/to/ralph-wiggum-ooda/src/rooda.sh .
   cp /path/to/ralph-wiggum-ooda/src/rooda-config.yml .
   cp -r /path/to/ralph-wiggum-ooda/src/components ./prompts
   chmod +x rooda.sh
   ```

3. **Update AGENTS.md:**
   ```bash
   ./rooda.sh bootstrap --max-iterations 1
   ```

Your existing work tracking, specs, and implementation remain unchanged. Only the framework files are updated.
```

**Note:** Config file keeps `prompts/*.md` paths (not `src/components/*.md`). This works because consumers copy `src/components/` to `./prompts/` during installation.

---

### Phase 7: Update AGENTS.md Definitions

**Priority: High**
**Dependencies: Phase 2, Phase 3**

- [ ] Update "Specification Definition" to reference `specs/*.md`
- [ ] Update "Implementation Definition" to reference `src/rooda.sh` and `src/components/*.md`
- [ ] Add operational learning about directory restructuring
- [ ] Document rationale for new structure
- [ ] Commit changes: `git commit -m "Update AGENTS.md for new directory structure"`

**Acceptance Criteria:**
- AGENTS.md correctly identifies specs location
- AGENTS.md correctly identifies implementation location
- Operational learning captured with rationale
- Changes committed to git

**Changes Required:**

```markdown
## Specification Definition

**Location:** `specs/*.md`

**Format:** Markdown specifications following JTBD structure

**Patterns:**
- `specs/agents-md-format.md` - AGENTS.md format specification
- `specs/specification-system.md` - Spec system design
- `specs/TEMPLATE.md` - Template for new specs

**Rationale:** Specifications define the framework methodology (JTBD, spec system, AGENTS.md format). These are distinct from user-facing documentation.

## Implementation Definition

**Location:** `src/rooda.sh` and `src/components/*.md`

**Patterns:** 
- `src/rooda.sh` - Main loop script
- `src/rooda-config.yml` - Procedure configuration
- `src/components/*.md` - OODA prompt components

**Exclude:** 
- `.beads/*` (work tracking database)
- `docs/*` (user-facing documentation)
- `specs/*` (specifications)
- `README.md`, `AGENTS.md`, `LICENSE.md` (root-level files)

**Rationale:** Implementation is the bash script and composable prompt components that execute the framework. Configuration and components are co-located with the script.

## Operational Learnings

**2026-02-03:** Restructured repository into src/, specs/, docs/ to enable dogfooding:
- Separating implementation (src/) from specifications (specs/) allows running draft-plan-impl-to-spec
- Framework can now use its own methodology to generate specs from implementation
- Clear separation makes it obvious what agents should analyze vs what they should read for guidance
- Consumers copy from src/ to their project root (flat structure), while framework repo has internal organization
- Config file paths remain prompts/*.md for consumer compatibility
```

---

### Phase 8: Create Wrapper Script at Project Root (Developer Convenience)

**Priority: Low**
**Dependencies: Phase 2, Phase 3**

- [ ] Create `rooda` wrapper script at project root
- [ ] Make wrapper executable: `chmod +x rooda`
- [ ] Test wrapper: `./rooda bootstrap --max-iterations 1`
- [ ] Commit changes: `git commit -m "Add wrapper script for framework developers"`

**Acceptance Criteria:**
- `./rooda <procedure>` works from project root (for framework development)
- `./src/rooda.sh <procedure>` still works
- Wrapper is executable
- Changes committed to git

**Wrapper Script (`rooda`):**
```bash
#!/bin/bash
# Wrapper script for framework developers to invoke src/rooda.sh from project root
# NOTE: This wrapper is NOT part of the consumer installation.
# Users copy files from src/ to their project per installation instructions.
exec "$(dirname "$0")/src/rooda.sh" "$@"
```

**Rationale:** Provides convenience for framework developers working on ralph-wiggum-ooda itself. Consumers of the framework copy files from `src/` to their project root per installation instructions, so they don't need or use this wrapper. This is a developer tool, not a consumer artifact.

---

### Phase 9: Update Git Tracking and Commit

**Priority: Critical**
**Dependencies: All previous phases**

- [ ] Stage all moved files with `git mv` (preserves history)
- [ ] Stage all new files
- [ ] Stage all updated files
- [ ] Verify no broken references with `grep -r "prompts/" --exclude-dir=.git`
- [ ] Verify no broken references with `grep -r "docs/specs.md" --exclude-dir=.git`
- [ ] Verify no broken references with `grep -r "docs/spec-template.md" --exclude-dir=.git`
- [ ] Verify no broken references with `grep -r "docs/agents-md-specification.md" --exclude-dir=.git`
- [ ] Run `shellcheck src/rooda.sh` to verify script integrity
- [ ] Test bootstrap procedure: `./src/rooda.sh bootstrap`
- [ ] Commit with descriptive message

**Acceptance Criteria:**
- All file moves tracked with `git mv` (preserves history)
- No broken internal references
- shellcheck passes
- Bootstrap procedure runs successfully
- Commit message documents restructuring rationale

**Commit Message:**
```
Restructure repository into src/, specs/, docs/

Reorganize repository for clarity and separation of concerns:

- src/ - Implementation (rooda.sh, config, prompt components)
  - Moved rooda.sh and rooda-config.yml to src/
  - Moved prompts/ to src/components/ (they are implementation, not docs)
  - Updated all path references in rooda-config.yml

- specs/ - Specifications (JTBD-based methodology definitions)
  - Created specs/ directory for framework specifications
  - Moved docs/spec-template.md to specs/TEMPLATE.md
  - Moved docs/agents-md-specification.md to specs/agents-md-format.md
  - Moved docs/specs.md to specs/specification-system.md
  - Created specs/README.md as index

- docs/ - User-facing documentation (guides, tutorials)
  - Kept ooda-loop.md, ralph-loop.md, beads.md
  - Created docs/README.md as navigation index
  - Removed specification documents (moved to specs/)

Updated README.md installation instructions and internal links.
Updated AGENTS.md to reflect new directory structure.

Rationale: Separates implementation, specifications, and user documentation
for better organization and clarity about what each file represents.
```

---

## Verification Checklist

After completing all phases, verify:

### Basic Functionality
- [ ] `./src/rooda.sh bootstrap --max-iterations 1` runs successfully
- [ ] `./src/rooda.sh build --max-iterations 1` runs successfully (with ready work)
- [ ] Wrapper script works: `./rooda bootstrap --max-iterations 1`

### All 9 Procedures
Test each procedure with `--max-iterations 1`:
- [ ] `./src/rooda.sh bootstrap` - Creates/updates AGENTS.md
- [ ] `./src/rooda.sh build` - Implements from work tracking
- [ ] `./src/rooda.sh draft-plan-story-to-spec` - Story incorporation
- [ ] `./src/rooda.sh draft-plan-bug-to-spec` - Bug incorporation
- [ ] `./src/rooda.sh draft-plan-spec-to-impl` - Gap analysis (specs → code)
- [ ] `./src/rooda.sh draft-plan-impl-to-spec` - Gap analysis (code → specs)
- [ ] `./src/rooda.sh draft-plan-spec-refactor` - Spec quality assessment
- [ ] `./src/rooda.sh draft-plan-impl-refactor` - Code quality assessment
- [ ] `./src/rooda.sh publish-plan` - Publish to work tracking

### Prompt Composition
- [ ] Verify `create_prompt()` function reads files from `src/components/`
- [ ] Verify composed prompt includes all 4 OODA phases
- [ ] Verify `kiro-cli chat` receives complete prompt

### Beads Integration
- [ ] `bd ready --json` works
- [ ] `bd show <id> --json` works
- [ ] Work tracking commands function correctly

### Documentation Links
- [ ] All links in README.md resolve correctly
- [ ] All links in docs/README.md resolve correctly
- [ ] All links in specs/README.md resolve correctly
- [ ] All links in src/README.md resolve correctly

### Configuration
- [ ] `yq eval '.procedures.bootstrap.observe' src/rooda-config.yml` returns `src/components/observe_bootstrap.md`
- [ ] All 36 path references in config point to `src/components/`
- [ ] Config file parses without errors

### Code Quality
- [ ] `shellcheck src/rooda.sh` passes with no errors
- [ ] Script has correct shebang and permissions

### Path References
- [ ] `grep -r "prompts/" --exclude-dir=.git --exclude-dir=.beads` returns only README installation instructions
- [ ] No broken references to old `docs/specs.md`
- [ ] No broken references to old `docs/spec-template.md`
- [ ] No broken references to old `docs/agents-md-specification.md`

### Git History
- [ ] Git history preserved for moved files: `git log --follow src/rooda.sh`
- [ ] All commits have descriptive messages
- [ ] Each phase committed separately

## Rollback Plan

If issues arise during migration:

1. **Before committing:** `git reset --hard` to discard all changes
2. **After committing:** `git revert <commit-hash>` to undo specific phase commit
3. **Partial rollback:** Cherry-pick specific files from previous commit
4. **Full rollback:** `git revert <first-commit>..<last-commit>` to undo all phases

### Rollback Testing Procedure

**Before executing the full plan, test the rollback process:**

1. **Create test branch:**
   ```bash
   git checkout -b test-restructure
   ```

2. **Execute first 3 phases:**
   - Phase 1: Create directories
   - Phase 2: Move files to src/
   - Phase 3: Update config paths

3. **Test rollback:**
   ```bash
   git reset --hard HEAD~3  # Undo last 3 commits
   ```

4. **Verify repository state:**
   - [ ] All files back in original locations
   - [ ] `rooda.sh` at project root (not in src/)
   - [ ] `prompts/` directory exists with all components
   - [ ] `rooda-config.yml` has original `prompts/` paths
   - [ ] Script runs: `./rooda.sh bootstrap --max-iterations 1`

5. **Clean up test branch:**
   ```bash
   git checkout abstract-planning
   git branch -D test-restructure
   ```

**Success:** Rollback procedure verified before executing full plan. If issues arise during actual execution, we know the rollback process works.

## Impact Analysis

### Breaking Changes

**For users who have already installed:**
- Installation instructions change (copy from `src/` instead of root)
- Existing installations continue to work (files copied to their project root)
- No impact on existing installations

**For users who clone the repo:**
- Must use new paths: `./src/rooda.sh` or copy files per updated instructions
- rooda-config.yml already updated with new paths
- No manual path updates required

**For contributors:**
- Prompt components now in `src/components/` instead of `prompts/`
- Specifications now in `specs/` instead of `docs/`
- Must update any local branches to match new structure

### Benefits

1. **Clear separation of concerns** - Implementation, specifications, and documentation are distinct
2. **Easier navigation** - Users know where to find what they need
3. **Better scalability** - As specs grow, they have dedicated space
4. **Consistent with conventions** - src/, specs/, docs/ is a common pattern
5. **Clearer intent** - Prompt components are obviously implementation, not documentation

### Risks

1. **Broken links** - Internal documentation links may break (mitigated by Phase 6, Phase 9 verification)
2. **User confusion** - Existing users may be confused by new structure (mitigated by clear README updates)
3. **Git history** - File moves may complicate history (mitigated by using `git mv`)
4. **Path references** - rooda-config.yml paths must be updated (mitigated by Phase 3)

## Timeline Estimate

- **Phase 1:** 5 minutes (create directories)
- **Phase 2:** 20 minutes (move files + path resolution fix)
- **Phase 3:** 15 minutes (update config paths + verification)
- **Phase 4:** 25 minutes (create specs files + cross-references)
- **Phase 5:** 20 minutes (reorganize docs + source attribution)
- **Phase 6a:** 10 minutes (installation instructions)
- **Phase 6b:** 15 minutes (example commands and structure)
- **Phase 6c:** 15 minutes (documentation links)
- **Phase 6d:** 15 minutes (developer vs consumer section)
- **Phase 7:** 20 minutes (update AGENTS.md + rationale)
- **Phase 8:** 10 minutes (wrapper script)
- **Phase 9:** 30 minutes (final verification + testing all 9 procedures)
- **Rollback testing:** 15 minutes (before execution)
- **Dogfooding verification:** 30 minutes (after completion)

**Total:** ~4-6 hours (including debugging, comprehensive testing, and dogfooding verification)

**Note:** Original estimate of 2.5 hours was optimistic. Realistic estimate accounts for:
- Path resolution complexity
- Testing all 9 procedures
- Comprehensive verification checklist
- Dogfooding validation
- Potential debugging and iteration

## Success Criteria

The restructuring is successful when:

1. All files are in their correct locations per target structure
2. All procedures run successfully with new paths
3. All documentation links resolve correctly
4. Git history is preserved for moved files
5. README.md accurately reflects new structure
6. AGENTS.md correctly identifies specs and implementation
7. No broken references remain in any files
8. shellcheck passes on src/rooda.sh
9. Verification checklist is 100% complete
10. **Framework can dogfood itself** - Verified through concrete steps below

### Dogfooding Verification Steps

After restructuring is complete, verify the framework can use its own methodology:

1. **Run gap analysis:**
   ```bash
   ./src/rooda.sh draft-plan-impl-to-spec --max-iterations 1
   ```
   - [ ] Command completes successfully
   - [ ] PLAN.md created/updated at project root
   - [ ] PLAN.md contains tasks for creating JTBD-based specs
   - [ ] Tasks identify rooda.sh and component system as implementation to document

2. **Publish plan:**
   ```bash
   ./src/rooda.sh publish-plan
   ```
   - [ ] Command completes successfully
   - [ ] Beads issues created for spec generation tasks
   - [ ] Issues have appropriate priorities and dependencies
   - [ ] `bd ready --json` shows spec generation tasks

3. **Build specs:**
   ```bash
   ./src/rooda.sh build --max-iterations 3
   ```
   - [ ] Command completes successfully
   - [ ] Specs created in `specs/` directory
   - [ ] Specs follow TEMPLATE.md structure (JTBD, Activities, Acceptance Criteria, etc.)
   - [ ] Specs document rooda.sh implementation
   - [ ] Specs document component system
   - [ ] Work tracking updated (tasks marked complete)

4. **Verify spec quality:**
   - [ ] Specs are readable and accurate
   - [ ] Specs follow JTBD methodology
   - [ ] Implementation mapping section references `src/` files
   - [ ] Acceptance criteria are testable

**Success:** The framework successfully analyzed its own implementation and generated proper specifications using its own methodology. This validates both the restructuring and the framework's capability.

## Next Steps After Restructuring

Once this plan is complete and dogfooding verification passes, the framework will have:
- Clean separation of implementation (src/), specifications (specs/), and documentation (docs/)
- Self-generated specifications demonstrating the methodology
- Proven capability to use its own process on itself

This demonstrates the framework's capability by using it to document itself—the ultimate validation of the methodology.

## Notes

- This is a pure reorganization—no functionality changes
- All existing procedures continue to work identically
- File contents remain unchanged (except path references)
- Git history is preserved through `git mv`
- The restructuring is reversible through git revert if needed

### Documentation Source Attribution

The following documentation files in `docs/` are built artifacts derived from external sources of truth:

- **docs/ralph-loop.md** - Derived from https://github.com/ghuntley/how-to-ralph-wiggum (Geoff Huntley's original Ralph Loop methodology)
- **docs/beads.md** - Derived from https://github.com/steveyegge/beads (Steve Yegge's beads work tracking system)
- **docs/ooda-loop.md** - Derived from https://en.wikipedia.org/wiki/OODA_loop (John Boyd's OODA framework)

These are explanatory documents for users of the ralph-wiggum-ooda framework. When updating these files, consult the upstream sources to ensure accuracy and consistency with the original concepts. They are not authoritative definitions but rather contextual guides to help users understand the foundations of this framework.
