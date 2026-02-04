# Prompt Composition

This document explains how rooda.sh assembles OODA phase prompts into a single executable prompt for the AI CLI.

## Overview

The `create_prompt()` function combines four separate markdown files (one per OODA phase) into a single prompt that gets piped to the AI CLI. This composition approach enables:

- **Reusability** - Same observe/orient/decide/act components can be mixed and matched
- **Modularity** - Each phase is independently maintainable
- **Configuration-driven** - Procedures specify which components to combine

## The create_prompt() Function

**Location:** `src/rooda.sh` lines 370-389

**Implementation:**
```bash
create_prompt() {
    # Assemble four OODA phase prompt files into single executable prompt
    # Uses heredoc (<<EOF) to create template with embedded command substitution
    # Each $(cat "$VAR") is evaluated when heredoc executes, inserting file contents
    cat <<EOF
# OODA Loop Iteration

## OBSERVE
$(cat "$OBSERVE")

## ORIENT
$(cat "$ORIENT")

## DECIDE
$(cat "$DECIDE")

## ACT
$(cat "$ACT")
EOF
}
```

**How it works:**

1. Uses bash heredoc (`<<EOF`) to create a multi-line template
2. Embeds command substitution `$(cat "$VAR")` within the heredoc
3. When heredoc executes, each `$(cat)` command runs and inserts file contents
4. Variables ($OBSERVE, $ORIENT, $DECIDE, $ACT) contain paths to prompt files
5. Output is valid markdown with clear section headers

## Assembled Prompt Structure

The output follows this structure:

```markdown
# OODA Loop Iteration

## OBSERVE
[Contents of observe prompt file]

## ORIENT
[Contents of orient prompt file]

## DECIDE
[Contents of decide prompt file]

## ACT
[Contents of act prompt file]
```

## Prompt Component Files

**Location:** `src/prompts/*.md`

**Naming convention:** `{phase}_{variant}.md`

Examples:
- `observe_plan_specs_impl.md` - Observes AGENTS.md, work tracking, specs, and implementation
- `orient_build.md` - Orients around understanding tasks and searching codebase
- `decide_build.md` - Decides which task to implement and how
- `act_build.md` - Acts by implementing code and running tests

## Configuration

Procedures are defined in `src/rooda-config.yml` by specifying which four prompt files to use:

```yaml
procedures:
  build:
    display: "Build from Plan"
    observe: prompts/observe_plan_specs_impl.md
    orient: prompts/orient_build.md
    decide: prompts/decide_build.md
    act: prompts/act_build.md
    default_iterations: 5
```

See `specs/configuration-schema.md` for complete configuration specification.

## Related Specifications

- `specs/iteration-loop.md` - Defines loop execution behavior
- `specs/configuration-schema.md` - Defines YAML structure
- `specs/ai-cli-integration.md` - Defines how assembled prompt is piped to AI CLI
- `specs/component-authoring.md` - Guidelines for writing prompt components
