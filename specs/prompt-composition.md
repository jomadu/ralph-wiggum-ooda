# Prompt Composition

## Job to be Done

Assemble four OODA phase files (observe, orient, decide, act) and optional user-provided context into a single prompt that can be piped to an AI CLI tool, supporting both embedded defaults and user-provided custom prompts with clear path resolution.

## Activities

1. **Resolve prompt file paths** — For each OODA phase, determine whether to use embedded default prompts (builtin: prefix) or filesystem prompts (relative paths)
2. **Read prompt content** — Load content from embedded resources or filesystem
3. **Inject user context** — If --context flag provided, insert user-supplied text into the assembled prompt
4. **Format with section markers** — Wrap each phase with clear delimiters for readability and debugging
5. **Validate completeness** — Ensure all four phases are present and non-empty before returning

## Acceptance Criteria

- [ ] Assembles prompts from four OODA phase files (observe, orient, decide, act)
- [ ] Supports embedded default prompts via builtin: prefix (e.g., builtin:observe_build.md)
- [ ] Supports filesystem prompts via relative paths (e.g., ./custom-prompts/observe.md)
- [ ] Filesystem paths resolved relative to config file directory, not current working directory
- [ ] Injects user-provided context when --context flag is supplied
- [ ] Wraps user context with "# CONTEXT" section marker
- [ ] Wraps each phase with section markers (e.g., "# OBSERVE", "# ORIENT")
- [ ] Validates all prompt files at config load time (fail fast)
- [ ] Returns error if any phase file is missing or empty (after trimming whitespace)
- [ ] Preserves markdown formatting from source files
- [ ] Normalizes trailing newlines (trim then add consistent spacing)
- [ ] Handles multi-line context injection without breaking prompt structure

## Data Structures

### Procedure Definition (from configuration.md)

```yaml
procedures:
  build:
    observe: builtin:observe_plan_specs_impl.md
    orient: builtin:orient_build.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
    max_iterations: 5
```

### Assembled Prompt Structure

```
# CONTEXT
[User Context - if provided via --context]

# OBSERVE
[Content from observe phase file]

# ORIENT
[Content from orient phase file]

# DECIDE
[Content from decide phase file]

# ACT
[Content from act phase file]
```

## Algorithm

```
function AssemblePrompt(procedure, userContext, configDir):
    prompt = ""
    
    // Inject user context first if provided
    if userContext != "":
        prompt += "# CONTEXT\n"
        prompt += userContext + "\n\n"
    
    // Process each OODA phase in order
    for phase in [observe, orient, decide, act]:
        filePath = procedure[phase]
        
        // Resolve path (builtin: vs filesystem)
        if filePath.startsWith("builtin:"):
            content = readEmbeddedPrompt(filePath.removePrefix("builtin:"))
        else:
            // Resolve relative to config file directory
            absolutePath = resolvePath(configDir, filePath)
            content = readFilesystemPrompt(absolutePath)
        
        // Validate content exists (trim whitespace first)
        if strings.TrimSpace(content) == "":
            return error("Missing or empty prompt file: " + filePath)
        
        // Add section marker and content (normalize trailing newlines)
        prompt += "# " + phase.toUpperCase() + "\n"
        prompt += strings.TrimRight(content, "\n") + "\n\n"
    
    return prompt
```

### Validation at Config Load Time

```
function ValidateProcedure(procedure, configDir):
    // Validate all four phases exist and are non-empty
    for phase in [observe, orient, decide, act]:
        filePath = procedure[phase]
        
        if filePath.startsWith("builtin:"):
            embeddedPath = filePath.removePrefix("builtin:")
            if !embeddedPromptExists(embeddedPath):
                return error("Prompt file not found: " + filePath + "\n" +
                           "Available builtin prompts for " + phase + " phase:\n" +
                           listBuiltinPromptsForPhase(phase))
        else:
            absolutePath = resolvePath(configDir, filePath)
            if !fileExists(absolutePath):
                return error("Prompt file not found: " + filePath + "\n" +
                           "Resolved to: " + absolutePath + "\n\n" +
                           "Tip: Check that the file exists and the path is correct.\n" +
                           "     Paths are resolved relative to the config file directory.")
            
            content = readFile(absolutePath)
            if strings.TrimSpace(content) == "":
                return error("Prompt file is empty: " + filePath + "\n" +
                           "Resolved to: " + absolutePath + "\n\n" +
                           "Tip: Each OODA phase must contain prompt instructions.\n" +
                           "     Check that the file is not empty or whitespace-only.")
    
    return nil
```

## Edge Cases

### Missing Prompt Files

**Scenario:** Procedure references a prompt file that doesn't exist

**Behavior:** Return error immediately, don't attempt to assemble partial prompt

**Example:**
```
Error: Prompt file not found: builtin:observe_missing.md
Procedure: build
Phase: observe
```

### Empty Prompt Files

**Scenario:** Prompt file exists but contains only whitespace

**Behavior:** Treat as missing, return error

**Rationale:** Empty phases indicate configuration error, not intentional omission

### Relative Path Resolution

**Scenario:** User provides custom prompt with relative path `./prompts/custom.md`

**Behavior:** Resolve relative to config file directory, not current working directory

**Example:**
```bash
# Config at /project/root/rooda-config.yml references ./prompts/custom.md
# Resolves to /project/root/prompts/custom.md

# Works regardless of where rooda is invoked from:
cd /project/root
rooda build  # Resolves to /project/root/prompts/custom.md

cd /project/root/src
rooda build  # Still resolves to /project/root/prompts/custom.md
```

**Rationale:** Config-relative paths ensure reproducibility regardless of invocation directory

### Context Injection with Special Characters

**Scenario:** User context contains markdown formatting, code blocks, or special characters

**Behavior:** Inject verbatim, preserve all formatting

**Example:**
```bash
rooda build --context "Focus on the auth module:
\`\`\`go
func Authenticate(token string) error
\`\`\`"
```

Result: Context appears at top of prompt with code block intact

### Builtin Prefix Case Sensitivity

**Scenario:** User writes `Builtin:` or `BUILTIN:` instead of `builtin:`

**Behavior:** Return error, require exact `builtin:` prefix

**Example:**
```
Error: Prompt file not found: Builtin:observe_build.md
Procedure: build
Phase: observe

Did you mean: builtin:observe_build.md
Note: The builtin: prefix is case-sensitive and must be lowercase.
```

**Rationale:** YAML is case-sensitive everywhere else; clear error message guides users to correct syntax

## Dependencies

- **configuration.md** — Procedure definitions specify which prompt files to compose
- **Embedded prompt files** — Default prompts shipped with rooda binary (25 files in prompts/ directory)
- **Filesystem access** — For reading custom user-provided prompts

## Implementation Mapping

### Source Files (v2 Go implementation)

- `internal/prompt/composer.go` — Core assembly logic
- `internal/prompt/resolver.go` — Path resolution (builtin: vs filesystem)
- `internal/prompt/embed.go` — Embedded prompt access via go:embed
- `prompts/*.md` — 25 embedded default prompt files

### Related Specs

- [configuration.md](configuration.md) — Defines procedure structure with OODA phase file references
- [iteration-loop.md](iteration-loop.md) — Consumes assembled prompts for each iteration
- [cli-interface.md](cli-interface.md) — Defines --context flag for user context injection
- [ai-cli-integration.md](ai-cli-integration.md) — Receives assembled prompt as stdin

## Examples

### Example 1: Basic Assembly with Builtin Prompts

**Input:**
```yaml
# rooda-config.yml
procedures:
  build:
    observe: builtin:observe_plan_specs_impl.md
    orient: builtin:orient_build.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
```

**Command:**
```bash
rooda build
```

**Output (assembled prompt):**
```markdown
# OBSERVE
[Content from embedded observe_plan_specs_impl.md]

# ORIENT
[Content from embedded orient_build.md]

# DECIDE
[Content from embedded decide_build.md]

# ACT
[Content from embedded act_build.md]
```

**Verification:** All four phases present, no user context, section markers clear

---

### Example 2: Custom Prompts from Filesystem

**Input:**
```yaml
# rooda-config.yml
procedures:
  custom-build:
    observe: ./my-prompts/observe.md
    orient: ./my-prompts/orient.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
```

**Command:**
```bash
rooda custom-build
```

**Output:**
```markdown
# OBSERVE
[Content from ./my-prompts/observe.md]

# ORIENT
[Content from ./my-prompts/orient.md]

# DECIDE
[Content from embedded decide_build.md]

# ACT
[Content from embedded act_build.md]
```

**Verification:** Mix of filesystem and builtin prompts, paths resolved correctly

---

### Example 3: User Context Injection

**Input:**
```yaml
# rooda-config.yml
procedures:
  build:
    observe: builtin:observe_plan_specs_impl.md
    orient: builtin:orient_build.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
```

**Command:**
```bash
rooda build --context "Focus on the authentication module. The new feature should integrate with the existing OAuth2 flow."
```

**Output:**
```markdown
# CONTEXT
Focus on the authentication module. The new feature should integrate with the existing OAuth2 flow.

# OBSERVE
[Content from embedded observe_plan_specs_impl.md]

# ORIENT
[Content from embedded orient_build.md]

# DECIDE
[Content from embedded decide_build.md]

# ACT
[Content from embedded act_build.md]
```

**Verification:** User context appears first, followed by all four phases

---

### Example 4: Missing Prompt File Error

**Input:**
```yaml
# rooda-config.yml
procedures:
  broken:
    observe: builtin:observe_missing.md
    orient: builtin:orient_build.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
```

**Command:**
```bash
rooda broken
```

**Output (error):**
```
Error: Prompt file not found: builtin:observe_missing.md
Procedure: broken
Phase: observe

Available builtin prompts for observe phase:
  observe_plan_specs_impl.md
  observe_specs.md
  observe_impl.md
  observe_bootstrap.md
  observe_bug_task_specs_impl.md
  observe_draft_plan.md
  observe_story_task_specs_impl.md
```

**Verification:** Clear error message, suggests available alternatives

---

### Example 5: Empty Prompt File Error

**Input:**
```bash
# Create empty custom prompt
touch ./my-prompts/empty.md
```

```yaml
# rooda-config.yml
procedures:
  test:
    observe: ./my-prompts/empty.md
    orient: builtin:orient_build.md
    decide: builtin:decide_build.md
    act: builtin:act_build.md
```

**Command:**
```bash
rooda test
```

**Output (error):**
```
Error: Prompt file is empty: ./my-prompts/empty.md
Resolved to: /project/root/my-prompts/empty.md
Procedure: test
Phase: observe

Tip: Each OODA phase must contain prompt instructions.
     Check that the file is not empty or whitespace-only.
```

**Verification:** Detects empty file, prevents partial prompt assembly

## Notes

### Design Rationale

**Why section markers?**
- Debugging: Easy to identify which phase produced output
- Readability: Clear visual separation in assembled prompt
- Parsing: AI can reference specific phases in its response

**Why builtin: prefix instead of @builtin/ or similar?**
- Simplicity: Single character prefix, no special escaping needed
- Familiarity: Similar to URL schemes (http:, file:)
- Clarity: Unambiguous distinction from filesystem paths

**Why inject user context at the top?**
- Precedence: User intent should frame the entire OODA cycle
- Visibility: AI sees context before any phase-specific instructions
- Simplicity: No need to parse or merge context into specific phases

**Why validate all phases at config load time?**
- Fail fast: Catch configuration errors immediately, before any iteration starts
- Clear feedback: User knows about broken config before invoking a procedure
- Atomicity: Either get complete prompt or clear error, no partial states
- Debugging: Error messages point to exact missing file with resolved absolute path

### Embedded Prompts List (v1 Reference)

Current embedded prompts in `prompts/` directory (25 files):

**Observe phase (7):**
- observe_bootstrap.md
- observe_bug_task_specs_impl.md
- observe_draft_plan.md
- observe_impl.md
- observe_plan_specs_impl.md
- observe_specs.md
- observe_story_task_specs_impl.md

**Orient phase (6):**
- orient_bootstrap.md
- orient_bug_task_incorporation.md
- orient_build.md
- orient_gap.md
- orient_publish.md
- orient_quality.md
- orient_story_task_incorporation.md

**Decide phase (6):**
- decide_bootstrap.md
- decide_bug_task_plan.md
- decide_build.md
- decide_gap_plan.md
- decide_publish.md
- decide_refactor_plan.md
- decide_story_task_plan.md

**Act phase (4):**
- act_bootstrap.md
- act_build.md
- act_plan.md
- act_publish.md

**Note:** v2 may consolidate or rename these prompts. This list reflects v1 implementation.

### Future Enhancements (Out of Scope for v2)

- **Prompt templating** — Variable substitution within prompts (e.g., {{project_name}})
- **Conditional phases** — Skip phases based on runtime conditions
- **Phase composition** — Assemble observe phase from multiple sub-prompts
- **Prompt validation** — Lint prompts for common mistakes or missing instructions
- **Prompt versioning** — Track which prompt versions produced which outputs
