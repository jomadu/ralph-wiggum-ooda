## fragments files

```
fragments/
├── observe/
│   ├── read_agents_md.md              # Load and parse AGENTS.md configuration
│   ├── scan_repo_structure.md          # Examine directory structure and files
│   ├── detect_build_system.md          # Identify build tools (go.mod, package.json, etc.)
│   ├── detect_work_tracking.md         # Identify work tracking system (.beads/, .github/)
│   ├── verify_commands.md              # Test that commands from AGENTS.md work
│   ├── query_work_tracking.md          # Fetch ready work items
│   ├── read_specs.md                   # Load specification files
│   ├── read_impl.md                    # Load implementation files
│   ├── read_task_input.md              # Load task description
│   ├── read_draft_plan.md              # Load draft plan file
│   ├── read_task_details.md            # Load specific task from work tracking
│   ├── run_tests.md                    # Execute test commands
│   └── run_lints.md                    # Execute lint commands
├── orient/
│   ├── compare_detected_vs_documented.md        # Find drift between actual and documented
│   ├── compare_documented_vs_actual.md          # Find drift in AGENTS.md
│   ├── identify_drift.md                        # Categorize inconsistencies
│   ├── evaluate_against_quality_criteria.md     # Check PASS/FAIL criteria
│   ├── understand_task_requirements.md          # Parse task into requirements
│   ├── understand_feature_requirements.md       # Parse feature requirements
│   ├── understand_bug_root_cause.md             # Analyze bug cause
│   ├── search_codebase.md                       # Find relevant code sections
│   ├── identify_affected_files.md               # Determine what needs to change
│   ├── identify_affected_specs.md               # Determine which specs need changes
│   ├── identify_affected_code.md                # Determine which code needs changes
│   ├── identify_spec_deficiencies.md            # Find gaps in specs
│   ├── parse_plan_tasks.md                      # Extract tasks from draft plan
│   ├── map_to_work_tracking_format.md           # Convert plan to work tracking format
│   ├── identify_specified_but_not_implemented.md # Gap analysis (specs → impl)
│   ├── identify_implemented_but_not_specified.md # Gap analysis (impl → specs)
│   ├── identify_structural_issues.md            # Find spec/code structure problems
│   ├── identify_duplication.md                  # Find duplicated content
│   ├── identify_code_smells.md                  # Find code quality issues
│   ├── identify_complexity_issues.md            # Find overly complex code
│   └── identify_maintenance_needs.md            # Find maintenance work
├── decide/
│   ├── determine_sections_to_update.md          # What to change in AGENTS.md
│   ├── check_if_blocked.md                      # Can we proceed? (emit FAILURE if not)
│   ├── pick_task.md                             # Select work item from work tracking
│   ├── plan_implementation_approach.md          # How to implement
│   ├── break_down_into_tasks.md                 # Decompose work into tasks
│   ├── prioritize_tasks.md                      # Order by impact/dependency
│   ├── prioritize_findings.md                   # Order audit findings
│   ├── prioritize_gaps_by_impact.md             # Order gap analysis findings
│   ├── identify_issues.md                       # List problems found
│   ├── categorize_drift_severity.md             # Rank drift items
│   └── determine_import_strategy.md             # How to import plan to work tracking
└── act/
    ├── write_agents_md.md                       # Update AGENTS.md file
    ├── write_audit_report.md                    # Create audit report
    ├── write_gap_report.md                      # Create gap analysis report
    ├── write_draft_plan.md                      # Create draft plan
    ├── modify_files.md                          # Edit specs or implementation files
    ├── commit_changes.md                        # Git commit with message
    ├── update_work_tracking.md                  # Mark tasks complete
    ├── update_draft_plan_status.md              # Update draft plan status
    ├── create_work_items.md                     # Import plan to work tracking
    ├── run_tests.md                             # Execute tests (verification)
    ├── emit_success.md                          # Output SUCCESS promise
    └── emit_failure.md                          # Output FAILURE promise
```

## Procedures Configuration Schema

```yaml
procedures:
  <procedure-name>:                    # string: Unique identifier for the procedure
    display: string                    # Human-readable name for the procedure
    summary: string                    # Brief description of what the procedure does
    description: string                # Detailed explanation of the procedure's purpose
    observe:                           # Array of actions (concatenated to form full prompt)
      - content: string                # Inline prompt content (optional)
        path: string                   # Path to prompt file (optional)
        parameters:                    # Template parameters (optional)
          <param-name>: <param-value>
    orient:                            # Array of actions (concatenated to form full prompt)
      - content: string                # Inline prompt content (optional)
        path: string                   # Path to prompt file (optional)
        parameters:                    # Template parameters (optional)
          <param-name>: <param-value>
    decide:                            # Array of actions (concatenated to form full prompt)
      - content: string                # Inline prompt content (optional)
        path: string                   # Path to prompt file (optional)
        parameters:                    # Template parameters (optional)
          <param-name>: <param-value>
    act:                               # Array of actions (concatenated to form full prompt)
      - content: string                # Inline prompt content (optional)
        path: string                   # Path to prompt file (optional)
        parameters:                    # Template parameters (optional)
          <param-name>: <param-value>
    iteration_mode: string             # "max-iterations" or "unlimited" (optional)
    default_max_iterations: integer    # Maximum number of OODA loop iterations (optional)
    iteration_timeout: integer         # Timeout per iteration in seconds (optional)
    max_output_buffer: integer         # Maximum output buffer size in bytes (optional)
    ai_cmd: string                     # AI command to use (optional)
    ai_cmd_alias: string               # Model configuration alias (optional)
```

## Built-in Procedures Config

```yml
procedures:
  agents-sync:
    display: "Agents Sync"
    summary: "Synchronize AGENTS.md with actual repository state"
    description: "Detects drift between documented and actual repository configuration, then updates AGENTS.md to match reality"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/scan_repo_structure.md"
      - path: "builtin:fragments/observe/detect_build_system.md"
      - path: "builtin:fragments/observe/detect_work_tracking.md"
    orient:
      - path: "builtin:fragments/orient/compare_detected_vs_documented.md"
      - path: "builtin:fragments/orient/identify_drift.md"
    decide:
      - path: "builtin:fragments/decide/determine_sections_to_update.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_agents_md.md"
      - path: "builtin:fragments/act/commit_changes.md"
      - path: "builtin:fragments/act/emit_success.md"

  build:
    display: "Build"
    summary: "Implement a task from work tracking"
    description: "Picks a ready task, implements it, runs tests, and marks it complete"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/query_work_tracking.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
      - path: "builtin:fragments/observe/read_task_details.md"
    orient:
      - path: "builtin:fragments/orient/understand_task_requirements.md"
      - path: "builtin:fragments/orient/search_codebase.md"
      - path: "builtin:fragments/orient/identify_affected_files.md"
    decide:
      - path: "builtin:fragments/decide/pick_task.md"
      - path: "builtin:fragments/decide/plan_implementation_approach.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/modify_files.md"
      - path: "builtin:fragments/act/run_tests.md"
      - path: "builtin:fragments/act/update_work_tracking.md"
      - path: "builtin:fragments/act/commit_changes.md"
      - path: "builtin:fragments/act/emit_success.md"

  publish-plan:
    display: "Publish Plan"
    summary: "Import draft plan into work tracking system"
    description: "Takes a draft plan and creates work items in the configured work tracking system"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_draft_plan.md"
      - path: "builtin:fragments/observe/query_work_tracking.md"
    orient:
      - path: "builtin:fragments/orient/parse_plan_tasks.md"
      - path: "builtin:fragments/orient/map_to_work_tracking_format.md"
    decide:
      - path: "builtin:fragments/decide/determine_import_strategy.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/create_work_items.md"
      - path: "builtin:fragments/act/update_draft_plan_status.md"
      - path: "builtin:fragments/act/emit_success.md"

  audit-spec:
    display: "Audit Specifications"
    summary: "Audit specification files for quality issues"
    description: "Reviews spec files against quality criteria and generates audit report"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_specs.md"
    orient:
      - path: "builtin:fragments/orient/evaluate_against_quality_criteria.md"
    decide:
      - path: "builtin:fragments/decide/identify_issues.md"
      - path: "builtin:fragments/decide/prioritize_findings.md"
    act:
      - path: "builtin:fragments/act/write_audit_report.md"
      - path: "builtin:fragments/act/emit_success.md"

  audit-impl:
    display: "Audit Implementation"
    summary: "Audit implementation files for quality issues"
    description: "Reviews implementation files, runs tests and lints, generates audit report"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_impl.md"
      - path: "builtin:fragments/observe/run_tests.md"
      - path: "builtin:fragments/observe/run_lints.md"
    orient:
      - path: "builtin:fragments/orient/evaluate_against_quality_criteria.md"
    decide:
      - path: "builtin:fragments/decide/identify_issues.md"
      - path: "builtin:fragments/decide/prioritize_findings.md"
    act:
      - path: "builtin:fragments/act/write_audit_report.md"
      - path: "builtin:fragments/act/emit_success.md"

  audit-agents:
    display: "Audit Agents Configuration"
    summary: "Audit AGENTS.md for accuracy and completeness"
    description: "Verifies AGENTS.md matches repository state and commands work correctly"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/scan_repo_structure.md"
      - path: "builtin:fragments/observe/detect_build_system.md"
      - path: "builtin:fragments/observe/verify_commands.md"
    orient:
      - path: "builtin:fragments/orient/compare_documented_vs_actual.md"
      - path: "builtin:fragments/orient/identify_drift.md"
    decide:
      - path: "builtin:fragments/decide/categorize_drift_severity.md"
    act:
      - path: "builtin:fragments/act/write_audit_report.md"
      - path: "builtin:fragments/act/emit_success.md"

  audit-spec-to-impl:
    display: "Audit Spec to Implementation Gap"
    summary: "Find specifications not implemented in code"
    description: "Identifies features specified but not yet implemented"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/identify_specified_but_not_implemented.md"
    decide:
      - path: "builtin:fragments/decide/prioritize_gaps_by_impact.md"
    act:
      - path: "builtin:fragments/act/write_gap_report.md"
      - path: "builtin:fragments/act/emit_success.md"

  audit-impl-to-spec:
    display: "Audit Implementation to Spec Gap"
    summary: "Find implementation not covered by specifications"
    description: "Identifies code that exists but is not documented in specifications"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_impl.md"
      - path: "builtin:fragments/observe/read_specs.md"
    orient:
      - path: "builtin:fragments/orient/identify_implemented_but_not_specified.md"
    decide:
      - path: "builtin:fragments/decide/prioritize_gaps_by_impact.md"
    act:
      - path: "builtin:fragments/act/write_gap_report.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-spec-feat:
    display: "Draft Plan: Spec Feature"
    summary: "Create plan for new specification feature"
    description: "Analyzes feature requirements and creates implementation plan focused on specifications"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/understand_feature_requirements.md"
      - path: "builtin:fragments/orient/identify_affected_specs.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-spec-fix:
    display: "Draft Plan: Spec Bug Fix"
    summary: "Create plan for specification bug fix"
    description: "Analyzes bug root cause and creates fix plan focused on specifications"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/understand_bug_root_cause.md"
      - path: "builtin:fragments/orient/identify_spec_deficiencies.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-spec-refactor:
    display: "Draft Plan: Spec Refactor"
    summary: "Create plan for specification refactoring"
    description: "Identifies structural issues in specs and creates refactoring plan"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
    orient:
      - path: "builtin:fragments/orient/identify_structural_issues.md"
      - path: "builtin:fragments/orient/identify_duplication.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-spec-chore:
    display: "Draft Plan: Spec Maintenance"
    summary: "Create plan for specification maintenance tasks"
    description: "Identifies maintenance needs in specs and creates chore plan"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
    orient:
      - path: "builtin:fragments/orient/identify_maintenance_needs.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-impl-feat:
    display: "Draft Plan: Implementation Feature"
    summary: "Create plan for new implementation feature"
    description: "Analyzes feature requirements and creates implementation plan focused on code"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/understand_feature_requirements.md"
      - path: "builtin:fragments/orient/identify_affected_code.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-impl-fix:
    display: "Draft Plan: Implementation Bug Fix"
    summary: "Create plan for implementation bug fix"
    description: "Analyzes bug root cause and creates fix plan focused on code"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_specs.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/understand_bug_root_cause.md"
      - path: "builtin:fragments/orient/identify_affected_code.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-impl-refactor:
    display: "Draft Plan: Implementation Refactor"
    summary: "Create plan for code refactoring"
    description: "Identifies code smells and complexity issues, creates refactoring plan"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/identify_code_smells.md"
      - path: "builtin:fragments/orient/identify_complexity_issues.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"

  draft-plan-impl-chore:
    display: "Draft Plan: Implementation Maintenance"
    summary: "Create plan for code maintenance tasks"
    description: "Identifies maintenance needs in code and creates chore plan"
    observe:
      - path: "builtin:fragments/observe/read_agents_md.md"
      - path: "builtin:fragments/observe/read_task_input.md"
      - path: "builtin:fragments/observe/read_impl.md"
    orient:
      - path: "builtin:fragments/orient/identify_maintenance_needs.md"
    decide:
      - path: "builtin:fragments/decide/break_down_into_tasks.md"
      - path: "builtin:fragments/decide/prioritize_tasks.md"
      - path: "builtin:fragments/decide/check_if_blocked.md"
    act:
      - path: "builtin:fragments/act/write_draft_plan.md"
      - path: "builtin:fragments/act/emit_success.md"
```