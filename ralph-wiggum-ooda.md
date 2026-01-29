# Ralph Wiggum OODA

An autonomous AI coding methodology using composable OODA-based prompts to maintain fresh context across iterations.

## Core Concept

Each loop iteration: **observe → orient → decide → act → clear context → repeat**

Fresh context each iteration keeps the AI in its "smart zone" (40-60% context utilization). File-based memory (AGENTS.md, PLAN.md) persists learnings between iterations.

## The Loop Mechanism

```bash
./ooda --observe prompts/observe_X.md \
       --orient prompts/orient_Y.md \
       --decide prompts/decide_Z.md \
       --act prompts/act_W.md \
       [--max-iterations N]
```

The script interpolates the 4 prompt components into a template and feeds it to an LLM agent. Each iteration:
1. Loads prompt template with 4 OODA phase components
2. Agent executes observe → orient → decide → act
3. Updates PLAN.md on disk
4. Exits (context cleared)
5. Loop restarts with fresh context

Exits when max iterations reached.

## OODA Phase Responsibilities

### Observe - Gather information from environment
- Read specs, implementation, PLAN.md, AGENTS.md, test results
- Surface raw data without interpretation
- Variants read different sources (specs, implementation, gaps, quality metrics)

### Orient - Analyze and synthesize based on mental models
- Discover/verify AGENTS.md (create if missing based on codebase inspection)
- Discover/verify PLAN.md accuracy
- Apply refactoring criteria (cohesion, coupling, complexity, completeness)
- Gap analysis between specs and implementation
- Synthesize observations into understanding
- **This is the "heavy" phase with most logic**

### Decide - Determine course of action
- Pick highest priority task from analysis
- Decide what to write to PLAN.md
- For building: decide which task to implement
- For planning: decide plan structure and priorities

### Act - Execute the decision
- Write to PLAN.md (for planning tasks)
- Implement code + run tests (for building tasks)
- Commit changes
- Update PLAN.md with progress
- **Backpressure lives here** (tests, lints, type checks)

## Task Types

The methodology supports multiple task types through prompt composition:

1. **Building from plan** - Implement tasks from PLAN.md
   - Only task type that modifies implementation code
   - Backpressure from tests ensures correctness

2. **Plan spec→impl** - Create plan to make implementation match specifications
   - Gap analysis: what's in specs but not in code

3. **Plan impl→spec** - Create plan to make specifications match implementation
   - Gap analysis: what's in code but not in specs

4. **Plan spec refactoring** - Create plan to refactor specs out of local optimums
   - Orient applies boolean criteria (cohesion, coupling, completeness, complexity)
   - Triggers on threshold failures or human markers (TODOs, comments, "REFACTORME")

5. **Plan impl refactoring** - Create plan to refactor implementation out of local optimums
   - Same criteria as spec refactoring
   - Proposes refactoring in PLAN.md, doesn't execute

## Key Principles

### Composable Prompts
- Minimal yet complete set of prompt variants per phase
- Most variants in observe (different data sources) and orient (different analysis types)
- Decide/act more stable across task types
- Same orient variant can be reused with different observe inputs

### File-Based State

**AGENTS.md** - Operational guide for the repository
- How to build/run/test the project
- Definition of what constitutes "specification" vs "implementation"
- Created by orient phase if missing
- Assumed inaccurate/incomplete until verified empirically

**PLAN.md** - Prioritized task list and progress tracking
- Generated and updated by act phase
- Can contain refactoring proposals with criteria scores
- Assumed inaccurate until verified

**prompts/** - OODA phase component library
- `prompts/observe_*.md` - Different observation sources
- `prompts/orient_*.md` - Different analysis types
- `prompts/decide_*.md` - Different decision strategies
- `prompts/act_*.md` - Different execution modes

### Context Management
- 200K tokens advertised ≈ 176K usable
- 40-60% utilization = "smart zone"
- Fresh context each iteration prevents degradation
- Use main agent as scheduler, spawn subagents for parallel work

### Steering via Backpressure
- Tests, lints, type checks reject invalid work (in act phase)
- Refactoring criteria provide quality gates (in orient phase)
- Existing code patterns guide generation
- Eventual consistency through iteration

### Refactoring Triggers
Boolean criteria scored as PASS/FAIL in orient phase:
- **Cohesion** - Do related things belong together?
- **Coupling** - Are dependencies minimized?
- **Completeness** - Are specs/implementation complete?
- **Complexity** - Is it unnecessarily complex?
- **Human markers** - TODOs, comments, "REFACTORME", spec phrases

When criteria fail threshold, decide/act write refactoring proposal to PLAN.md. Future iteration with building task executes it.

## Why It Works

1. **Fresh context** - No degradation from context pollution
2. **Composable prompts** - Reusable components for different task types
3. **OODA framework** - Clear separation of concerns across phases
4. **File-based state** - PLAN.md and AGENTS.md persist learnings
5. **Backpressure** - Tests and criteria force correctness
6. **Eventual consistency** - Iteration converges to solution
7. **Simplicity** - Bash loop, prompt interpolation, file I/O

## Safety

Requires `--dangerously-skip-permissions` to run autonomously. Run in isolated sandbox environments:
- Docker containers (local)
- Fly Sprites / E2B (remote)
- Minimum viable access (only needed API keys)
- No access to private data beyond requirements

Philosophy: "It's not if it gets popped, it's when. And what is the blast radius?"

## Escape Hatches

- Max iterations prevents infinite loops
- Ctrl+C stops the loop
- `git reset --hard` reverts uncommitted changes
- Regenerate PLAN.md if trajectory goes wrong

---

*Methodology evolved from [Ralph Loop](https://ghuntley.com/ralph/) by Geoff Huntley*
