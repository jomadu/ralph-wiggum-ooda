# Minimal Composable Prompt Set

## Task Breakdown by Phase

### Task 1: Building from plan

**OBSERVE**
- AGENTS.md (how to build/test, what is specification/implementation)
- PLAN.md (find next task)
- Specifications (per AGENTS.md definition)
- Implementation (per AGENTS.md definition, file tree, symbols)

**ORIENT**
- Understand task requirements
- Identify what needs to be built/modified
- Determine test strategy

**DECIDE**
- Pick highest priority task from PLAN.md
- Determine implementation approach
- Identify which files to modify

**ACT**
- Implement the task
- Run tests per AGENTS.md
- Update PLAN.md (mark complete/update)
- Update AGENTS.md if learned something new
- Commit if passing

---

### Task 2: Plan spec-to-impl

**OBSERVE**
- AGENTS.md (what is specification/implementation)
- PLAN.md (current state)
- Specifications (per AGENTS.md definition)
- Implementation (per AGENTS.md definition, file tree, symbols)

**ORIENT**
- Gap analysis: what's in specifications but missing from implementation
- Identify unimplemented features
- Assess implementation completeness and accuracy

**DECIDE**
- Structure plan by priority
- Break gaps into implementable tasks
- Determine task dependencies

**ACT**
- Write PLAN.md with prioritized tasks
- Update AGENTS.md if learned something new
- Commit changes

---

### Task 3: Plan impl-to-spec

**OBSERVE**
- AGENTS.md (what is specification/implementation)
- PLAN.md (current state)
- Specifications (per AGENTS.md definition)
- Implementation (per AGENTS.md definition, file tree, symbols)

**ORIENT**
- Gap analysis: what's in implementation but missing from specifications
- Identify undocumented features
- Assess specification completeness and accuracy

**DECIDE**
- Structure plan by priority
- Break gaps into documentation tasks
- Determine which specs need updates

**ACT**
- Write PLAN.md with prioritized tasks
- Update AGENTS.md if learned something new
- Commit changes

---

### Task 4: Plan spec refactoring

**OBSERVE**
- AGENTS.md (what is specification, quality criteria definitions)
- Specifications (per AGENTS.md definition)

**ORIENT**
- Apply boolean criteria: clarity, completeness, consistency, testability
- Identify human markers (TODOs, "REFACTORME", unclear language)
- Score each criterion PASS/FAIL

**DECIDE**
- If criteria fail threshold: propose spec refactoring
- Structure spec refactoring plan
- Prioritize by impact

**ACT**
- Write PLAN.md with prioritized tasks
- Update AGENTS.md if learned something new
- Commit changes

---

### Task 5: Plan impl refactoring

**OBSERVE**
- AGENTS.md (what is implementation, quality criteria definitions)
- Implementation (per AGENTS.md definition, file tree, symbols)

**ORIENT**
- Apply boolean criteria: cohesion, coupling, complexity, maintainability
- Identify human markers (TODOs, long functions, code smells)
- Score each criterion PASS/FAIL

**DECIDE**
- If criteria fail threshold: propose implementation refactoring
- Structure implementation refactoring plan
- Prioritize by impact

**ACT**
- Write PLAN.md with prioritized tasks
- Update AGENTS.md if learned something new
- Commit changes

---

## Prompt Component Set

**OBSERVE (3 variants)**
1. `observe_plan_specs_impl.md` - AGENTS.md + PLAN.md + specifications + implementation
2. `observe_specs.md` - AGENTS.md + specifications
3. `observe_impl.md` - AGENTS.md + implementation

**ORIENT (3 variants)**
1. `orient_build.md` - Understand task, identify what to build
2. `orient_gap.md` - Compare sources, identify gaps, assess completeness/accuracy
3. `orient_quality.md` - Apply criteria, identify markers, score PASS/FAIL

**DECIDE (3 variants)**
1. `decide_build.md` - Pick task, determine approach, identify files
2. `decide_gap_plan.md` - Structure plan, break gaps into tasks, determine dependencies/updates
3. `decide_refactor_plan.md` - If threshold fails: propose refactoring, structure plan, prioritize

**ACT (2 variants)**
1. `act_build.md` - Implement, test, update PLAN.md/AGENTS.md, commit if passing
2. `act_plan.md` - Write PLAN.md, update AGENTS.md, commit

**Total: 11 prompt files** (3+3+3+2)

---

## Task Compositions

| Task | Observe | Orient | Decide | Act |
|------|---------|--------|--------|-----|
| 1. Building from plan | plan_specs_impl | build | build | build |
| 2. Plan spec-to-impl | plan_specs_impl | gap | gap_plan | plan |
| 3. Plan impl-to-spec | plan_specs_impl | gap | gap_plan | plan |
| 4. Plan spec refactoring | specs | quality | refactor_plan | plan |
| 5. Plan impl refactoring | impl | quality | refactor_plan | plan |

---

## Principles

- **AGENTS.md always read first** - Defines what constitutes "specification" and "implementation", plus quality criteria
- **Definitions defer to AGENTS.md** - What files/locations constitute specs and implementation varies by project
- **Commit after updates** - All file modifications complete before commit
- **Composability** - 11 files generate 5 task types through different combinations
