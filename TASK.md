# Refactor Task: Unify Component System and Prompt Composition Specifications

## Problem

Two specifications cover overlapping material but from different angles:

1. **specs/component-system.md** - Describes components as referencing common step codes (O1-O15, R1-R22, D1-D15, A1-A9)
2. **specs/prompt-composition.md** - Describes how four prompt files are assembled into a single executable prompt

Reading these specs might lead someone to believe components are just lists of step codes. However, examining actual component files (e.g., `src/components/act_bootstrap.md`, `src/components/observe_bootstrap.md`) reveals they contain **full prose instructions**, not just step code references.

The specs don't accurately describe what's in `src/components/*.md` or how to write new components.

## Root Cause

**specs/component-system.md** was generated from **src/README.md**, which is a quick reference guide showing which components use which step codes. This created a spec that describes the reference system but not the actual component content structure.

**specs/prompt-composition.md** describes the mechanical assembly (bash heredoc concatenation) but doesn't describe what the component files contain or how they're structured.

Neither spec would enable someone to write a new component that matches the existing patterns in `src/components/*.md`.

## Expected Outcome

A single unified specification that:

1. **Describes actual component file structure** - Components contain markdown with section headers and full prose instructions, not just step code lists
2. **Explains the dual nature** - Components can reference common steps by code OR provide inline instructions OR mix both approaches
3. **Shows real examples** - Uses actual component files to illustrate structure
4. **Covers prompt assembly** - Includes the mechanical assembly process (currently in prompt-composition.md)
5. **Provides authoring guidance** - Explains how to write new components that match existing patterns
6. **Documents the common steps system** - Explains what step codes mean and when to use them vs inline instructions

## Proposed Solution

Create a single specification that replaces both:

**File:** `specs/component-authoring.md`

**Structure:**
- **Job to be Done:** Enable developers to create and modify OODA component prompt files that can be composed into executable procedures
- **Component File Structure:** Markdown format with phase header, step sections, prose instructions
- **Component Patterns:** Three patterns observed in actual components:
  1. Step code references only (rare, mainly in src/README.md documentation)
  2. Full prose instructions with step code headers (most common, e.g., act_bootstrap.md)
  3. Inline instructions without step codes (possible but uncommon)
- **Common Steps Reference:** Complete list of O1-O15, R1-R22, D1-D15, A1-A9 with descriptions
- **Prompt Assembly Algorithm:** How create_prompt() combines four files with OODA headers
- **Authoring Guidelines:** Key principles for writing components (from component-system.md)
- **Examples:** Real component files showing actual structure

**Actions:**
1. Create `specs/component-authoring.md` with unified content
2. Mark `specs/component-system.md` as deprecated (add note at top referencing new spec)
3. Mark `specs/prompt-composition.md` as deprecated (add note at top referencing new spec)
4. Update `specs/README.md` to list new spec and mark old ones as deprecated
5. Update `src/README.md` to reference `specs/component-authoring.md` as authoritative source

## Acceptance Criteria

- [ ] New spec accurately describes structure of files in `src/components/*.md`
- [ ] New spec includes prompt assembly algorithm from prompt-composition.md
- [ ] New spec includes common steps reference from component-system.md
- [ ] New spec includes authoring guidelines from component-system.md
- [ ] New spec shows real examples from actual component files
- [ ] Old specs marked as deprecated with clear references to new spec
- [ ] specs/README.md updated to reflect deprecation
- [ ] src/README.md references new spec as authoritative

## Why This Matters

Without accurate specifications, developers can't:
- Understand how existing components work
- Write new components that match existing patterns
- Modify components confidently
- Understand the relationship between step codes and prose instructions

The current specs describe the reference system and assembly mechanism but not the actual artifact structure, creating a documentation gap.
