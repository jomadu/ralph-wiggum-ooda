# OODA Component System

## Job to be Done
Enable agents to execute procedures through composable, reusable prompt components that maintain consistency across iterations while allowing flexible procedure definitions.

## Activities
- **Component Composition** - Combine four OODA phase components (observe, orient, decide, act) into complete procedures
- **Common Step Reference** - Use step codes (O1-O15, R1-R22, D1-D15, A1-A9) to reference standardized instructions
- **AGENTS.md Update Guidance** - Define what qualifies as operational learning worthy of AGENTS.md updates

## Acceptance Criteria
- [ ] Components follow four-phase OODA structure (observe/orient/decide/act)
- [ ] Common steps defined and referenceable by code (O1-O15, R1-R22, D1-D15, A1-A9)
- [ ] Key principles for writing components documented
- [ ] A6 step explicitly defines operational vs non-operational learning
- [ ] Clear examples distinguish proper vs improper AGENTS.md updates
- [ ] References specs/agents-md-format.md for content boundaries

## Data Structures

### Component File Format
```markdown
# [Phase Name]: [Purpose]

## [Step Code]: [Step Name]

[Detailed instructions for this step]

## [Step Code]: [Step Name]

[Detailed instructions for this step]
```

**Fields:**
- `Phase Name` - One of: Observe, Orient, Decide, Act
- `Purpose` - Brief description of what this component does
- `Step Code` - Unique identifier (O1-O15, R1-R22, D1-D15, A1-A9)
- `Step Name` - Human-readable step description
- `Detailed instructions` - Specific guidance for executing this step

### Procedure Configuration
```yaml
procedures:
  procedure-name:
    display: "Human Readable Name"
    summary: "Brief description"
    observe: src/components/observe_*.md
    orient: src/components/orient_*.md
    decide: src/components/decide_*.md
    act: src/components/act_*.md
    default_iterations: N
```

## Algorithm

1. **Load Procedure Configuration** - Read rooda-config.yml to identify which four components to use
2. **Load Component Files** - Read each of the four markdown files (observe, orient, decide, act)
3. **Combine Components** - Concatenate components into single prompt in OODA order
4. **Execute Iteration** - Pass combined prompt to AI CLI
5. **Agent Interprets Step Codes** - Agent reads step codes and executes corresponding instructions
6. **Exit and Clear Context** - Iteration completes, script exits, context cleared

**Component Reference Resolution:**
```
Component references step code (e.g., "O1, O3, O7")
  ↓
Agent looks up step definition in common steps list
  ↓
Agent executes step instructions
  ↓
Agent proceeds to next step code
```

## Edge Cases

| Condition | Expected Behavior |
|-----------|-------------------|
| Component references undefined step code | Agent should skip or report error (implementation-dependent) |
| Multiple components define same step code | Last definition wins (avoid this through design) |
| Component omits step codes entirely | Agent executes inline instructions without common step reference |
| Conflicting guidance between components | Later phases override earlier phases (Act > Decide > Orient > Observe) |

## Dependencies

- `specs/agents-md-format.md` - Defines AGENTS.md structure and content boundaries
- `src/rooda-config.yml` - Procedure configuration schema
- `src/rooda.sh` - Script that loads and combines components

## Implementation Mapping

**Source files:**
- `src/components/observe_*.md` - Observation phase components
- `src/components/orient_*.md` - Analysis phase components
- `src/components/decide_*.md` - Decision phase components
- `src/components/act_*.md` - Execution phase components

**Related specs:**
- `specs/agents-md-format.md` - AGENTS.md content definition
- `specs/prompt-composition.md` - How prompts are assembled
- `specs/configuration-schema.md` - rooda-config.yml structure

## Common Steps Reference

### Observe Steps (O1-O15)

- **O1:** Study AGENTS.md as a whole (operational guide, conventions, learnings)
- **O2:** Study AGENTS.md for build/test commands
- **O3:** Study AGENTS.md for specification/implementation definitions
- **O4:** Study AGENTS.md for quality criteria definitions
- **O5:** Study AGENTS.md for task file location
- **O6:** Study AGENTS.md for draft plan location
- **O7:** Study AGENTS.md for work tracking system
- **O8:** Study work tracking system per AGENTS.md (query ready work)
- **O9:** Study task file per AGENTS.md (story/bug description)
- **O10:** Study draft plan file per AGENTS.md (current plan state, may not exist)
- **O11:** Study specifications per AGENTS.md definition
- **O12:** Study implementation per AGENTS.md definition (file tree, symbols)
- **O13:** Study repository structure (file tree, languages, build files)
- **O14:** Study existing documentation (README, specs if present)
- **O15:** Study implementation patterns

### Orient Steps (R1-R22)

- **R1:** Identify project type and tech stack
- **R2:** Determine what constitutes "specification" vs "implementation"
- **R3:** Identify build/test/run commands empirically
- **R4:** Synthesize operational understanding
- **R5:** Understand task requirements
- **R6:** Search codebase (don't assume not implemented)
- **R7:** Identify what needs to be built/modified
- **R8:** Determine test strategy
- **R9:** Analyze story from task file (scope, requirements, integration points)
- **R10:** Analyze bug from task file (symptoms, root cause, affected functionality)
- **R11:** Understand existing spec structure and patterns
- **R12:** Determine how story should be incorporated (create new specs, update existing, refactor)
- **R13:** Determine how spec should be adjusted to drive bug fix (acceptance criteria, edge cases, clarifications)
- **R14:** If draft plan exists: critique it (completeness, accuracy, priorities, clarity)
- **R15:** Identify tasks needed
- **R16:** Gap analysis: compare specs vs implementation
- **R17:** Assess completeness and accuracy
- **R18:** Apply boolean criteria per AGENTS.md
- **R19:** Identify human markers (TODOs, code smells, unclear language)
- **R20:** Score each criterion PASS/FAIL
- **R21:** Parse draft plan structure
- **R22:** Understand task breakdown and dependencies

### Decide Steps (D1-D15)

- **D1:** Determine AGENTS.md structure
- **D2:** Define specification and implementation locations
- **D3:** Identify quality criteria for this project
- **D4:** Pick the most important task from work tracking
- **D5:** Determine implementation approach using parallel subagents
- **D6:** Identify which files to modify
- **D7:** Generate complete plan for story incorporation into specs
- **D8:** Generate complete plan for spec adjustments to drive bug fix
- **D9:** Structure plan by priority (most important first)
- **D10:** Break into tight, actionable tasks
- **D11:** Determine task dependencies
- **D12:** If criteria fail threshold: propose refactoring
- **D13:** Prioritize by impact
- **D14:** Map plan tasks to work tracking issues (title, description, dependencies)
- **D15:** Identify order of issue creation

### Act Steps (A1-A9)

- **A1:** Create AGENTS.md with operational guide
- **A2:** Commit changes
- **A3:** Implement using parallel subagents (only 1 subagent for build/tests)
- **A4:** Run tests per AGENTS.md (backpressure)
- **A5:** Update work tracking per AGENTS.md (mark complete/update status)
- **A6:** Update AGENTS.md if learned something new (capture the why, keep it up to date)
- **A7:** Commit when tests pass
- **A8:** Write draft plan file per AGENTS.md with prioritized bullet-point task list
- **A9:** Execute work tracking commands per AGENTS.md to create issues from draft plan

## A6 Operational Learning Criteria

**A6 Purpose:** Update AGENTS.md when operational learnings occur during procedure execution.

**What Qualifies as Operational Learning:**

✅ **Commands that failed or succeeded differently than documented**
- Example: AGENTS.md said `npm test` but actual command is `npm run test:unit`
- Example: Build requires `make deps` before `make build`

✅ **File paths discovered or corrected**
- Example: Specs are in `docs/specs/*.md` not `specs/*.md`
- Example: Implementation includes `lib/**/*.rb` not just `src/**/*.rb`

✅ **Quality criteria refined based on project needs**
- Example: Added "API documentation completeness" as boolean criterion
- Example: Changed test coverage threshold from 80% to 90%

✅ **Workflow patterns learned**
- Example: Tests require Docker containers to be running first
- Example: Work tracking uses labels `ready` and `in-progress` not status field

**What Does NOT Qualify as Operational Learning:**

❌ **Test artifacts or validation files**
- Example: Created `VALIDATION-issue-123.md` with test cases
- Example: Generated `test-results.json` from test run

❌ **Validation patterns or verification steps**
- Example: Discovered how to verify acceptance criteria for this task
- Example: Created checklist for manual testing

❌ **Historical notes or progress tracking**
- Example: Completed 3 of 5 subtasks
- Example: Previous iteration failed due to X

❌ **Temporary debugging files or logs**
- Example: Created `debug.log` to troubleshoot issue
- Example: Generated `trace.txt` for analysis

**Key Distinction:** Operational learning changes how future agents interact with the repository. Test artifacts and validation patterns are task-specific and don't generalize.

## Examples

### Example 1: Proper A6 Update (Operational Learning)

**Scenario:** Agent runs tests per AGENTS.md

**Input:**
```bash
# AGENTS.md says:
Test: npm test

# Agent runs:
npm test
# Error: Script "test" not found
```

**Agent Action:**
```bash
# Agent discovers correct command:
npm run test:unit
# Tests pass

# Agent updates AGENTS.md:
Test: npm run test:unit
```

**Commit Message:**
```
Update test command in AGENTS.md

Why: `npm test` script doesn't exist. Correct command is `npm run test:unit`.
Verified by running command successfully.
```

**Verification:**
- AGENTS.md updated with correct command
- Rationale captured in commit message
- Future agents will use correct command

### Example 2: Improper A6 Update (Non-Operational)

**Scenario:** Agent implements feature and creates validation file

**Input:**
```bash
# Task: Implement user authentication
# Agent creates test cases in VALIDATION-issue-123.md
```

**Agent Action (INCORRECT):**
```bash
# Agent should NOT update AGENTS.md with:
"Validation files stored in VALIDATION-*.md at project root"

# This is NOT operational learning - it's a test artifact
```

**Correct Behavior:**
- Create `VALIDATION-issue-123.md` if needed for task
- Do NOT update AGENTS.md
- Validation file is task-specific, not operational guidance

**Verification:**
- AGENTS.md unchanged
- Validation file exists but not documented in AGENTS.md
- Future agents don't need to know about this specific validation file

### Example 3: Proper A6 Update (File Path Discovery)

**Scenario:** Agent searches for specifications

**Input:**
```bash
# AGENTS.md says:
Specification Definition: specs/*.md

# Agent searches and finds specs in different location:
ls docs/specifications/*.md
# Files found: api.md, database.md, auth.md
```

**Agent Action:**
```bash
# Agent updates AGENTS.md:
Specification Definition: docs/specifications/*.md

# Commit with rationale
```

**Commit Message:**
```
Correct specification location in AGENTS.md

Why: Specs are in docs/specifications/*.md not specs/*.md.
Discovered through file tree search. Verified 3 spec files exist at this location.
```

**Verification:**
- AGENTS.md reflects actual repository structure
- Future agents will search correct location
- Rationale explains discovery process

### Example 4: Improper A6 Update (Validation Pattern)

**Scenario:** Agent verifies acceptance criteria

**Input:**
```bash
# Task acceptance criteria:
- [ ] User can log in with email/password
- [ ] Invalid credentials show error message
- [ ] Successful login redirects to dashboard

# Agent creates verification checklist
```

**Agent Action (INCORRECT):**
```bash
# Agent should NOT update AGENTS.md with:
"Acceptance criteria verification: manual testing checklist in VALIDATION-*.md"

# This is task-specific validation, not operational guidance
```

**Correct Behavior:**
- Verify acceptance criteria as needed
- Do NOT update AGENTS.md with verification approach
- Verification pattern is task-specific

**Verification:**
- AGENTS.md unchanged
- Acceptance criteria verified through appropriate means
- No task-specific validation patterns added to AGENTS.md

## Key Principles for Writing Components

**Reference Common Steps by Code** - Use step codes (O1, R5, D3, A2) rather than rewriting instructions. This ensures consistency and makes updates propagate automatically.

**One Component, One Concern** - Each component handles exactly one OODA phase. Observe gathers, Orient analyzes, Decide chooses, Act executes. No overlap.

**AGENTS.md is Always the Source of Truth** - Components must defer to AGENTS.md for all project-specific definitions: what constitutes specs/implementation, where files live, what commands to run, what quality criteria apply.

**Explicit Over Implicit** - State exactly what to read, analyze, decide, or do. "Study specifications per AGENTS.md definition" is better than "look at the specs."

**Use Precise Language** - Follow terminology from the Ralph Loop methodology:
- "study" (not "read" or "look at")
- "don't assume not implemented" (critical - the Achilles' heel)
- "using parallel subagents" / "only 1 subagent for build/tests"
- "capture the why" when updating AGENTS.md
- "keep it up to date" for maintaining accuracy
- "resolve them or document them" for issues found

**Search Before Assuming** - Orient components must emphasize searching the codebase before concluding something doesn't exist. This is the critical failure mode.

**Backpressure is Mandatory** - Act components that modify code must run tests and only commit when passing. No exceptions.

**Capture the Why** - When updating AGENTS.md, components must instruct agents to document rationale, not just changes. Why this command? Why this location? What was learned?

**Parallel Subagents for Scale** - Act components should use parallel subagents for independent work, but only 1 subagent for build/test operations to avoid conflicts.

**Plans are Disposable** - Planning components should generate complete plans each iteration, not incrementally patch. Cheap to regenerate beats expensive to maintain.

**Tight Tasks Win** - Decide components should break work into the smallest implementable units. One task per build iteration maximizes smart zone utilization.

**Commit After Complete** - Act components must complete all file modifications before committing. No partial work commits.

**Boolean Criteria Only** - Quality assessment components use PASS/FAIL criteria, not subjective scores. Clear thresholds trigger refactoring.

## Notes

The component system enables significant reuse across procedures. For example, `observe_plan_specs_impl.md` is shared by `build`, `draft-plan-spec-to-impl`, and `draft-plan-impl-to-spec` procedures. They differ only in their orient, decide, and act components.

The A6 operational learning criteria are critical to prevent AGENTS.md from becoming cluttered with task-specific artifacts. AGENTS.md should contain only information that helps future agents interact with the repository, not historical records of what was done.

## Known Issues

None identified during specification creation.

## Areas for Improvement

- Consider adding validation for step code references (detect undefined codes)
- Could add tooling to generate component documentation from step definitions
- May benefit from examples of custom procedure creation using existing components
