package procedures

import "github.com/jomadu/rooda/internal/config"

// BuiltInProcedures returns the 16 built-in procedures that ship with rooda.
// Each procedure is composed of reusable fragments organized by OODA phase.
func BuiltInProcedures() map[string]config.Procedure {
	return map[string]config.Procedure{
		"agents-sync": {
			Display:     "Agents Sync",
			Summary:     "Synchronize AGENTS.md with actual repository state",
			Description: "Detects drift between documented and actual repository configuration, then updates AGENTS.md to match reality",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/scan_repo_structure.md"},
				{Path: "builtin:fragments/observe/detect_build_system.md"},
				{Path: "builtin:fragments/observe/detect_work_tracking.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/compare_detected_vs_documented.md"},
				{Path: "builtin:fragments/orient/identify_drift.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/determine_sections_to_update.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_agents_md.md"},
				{Path: "builtin:fragments/act/commit_changes.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"build": {
			Display:     "Build",
			Summary:     "Implement a task from work tracking",
			Description: "Picks a ready task, implements it, runs tests, and marks it complete",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/query_work_tracking.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
				{Path: "builtin:fragments/observe/read_task_details.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/understand_task_requirements.md"},
				{Path: "builtin:fragments/orient/search_codebase.md"},
				{Path: "builtin:fragments/orient/identify_affected_files.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/pick_task.md"},
				{Path: "builtin:fragments/decide/plan_implementation_approach.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/modify_files.md"},
				{Path: "builtin:fragments/act/run_tests.md"},
				{Path: "builtin:fragments/act/update_work_tracking.md"},
				{Path: "builtin:fragments/act/commit_changes.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"publish-plan": {
			Display:     "Publish Plan",
			Summary:     "Import draft plan into work tracking system",
			Description: "Takes a draft plan and creates work items in the configured work tracking system",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_draft_plan.md"},
				{Path: "builtin:fragments/observe/query_work_tracking.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/parse_plan_tasks.md"},
				{Path: "builtin:fragments/orient/map_to_work_tracking_format.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/determine_import_strategy.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/create_work_items.md"},
				{Path: "builtin:fragments/act/update_draft_plan_status.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"audit-spec": {
			Display:     "Audit Specifications",
			Summary:     "Audit specification files for quality issues",
			Description: "Reviews spec files against quality criteria and generates audit report",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/evaluate_against_quality_criteria.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/identify_issues.md"},
				{Path: "builtin:fragments/decide/prioritize_findings.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_audit_report.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"audit-impl": {
			Display:     "Audit Implementation",
			Summary:     "Audit implementation files for quality issues",
			Description: "Reviews implementation files, runs tests and lints, generates audit report",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
				{Path: "builtin:fragments/observe/run_tests.md"},
				{Path: "builtin:fragments/observe/run_lints.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/evaluate_against_quality_criteria.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/identify_issues.md"},
				{Path: "builtin:fragments/decide/prioritize_findings.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_audit_report.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"audit-agents": {
			Display:     "Audit Agents Configuration",
			Summary:     "Audit AGENTS.md for accuracy and completeness",
			Description: "Verifies AGENTS.md matches repository state and commands work correctly",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/scan_repo_structure.md"},
				{Path: "builtin:fragments/observe/detect_build_system.md"},
				{Path: "builtin:fragments/observe/verify_commands.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/compare_documented_vs_actual.md"},
				{Path: "builtin:fragments/orient/identify_drift.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/categorize_drift_severity.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_audit_report.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"audit-spec-to-impl": {
			Display:     "Audit Spec to Implementation Gap",
			Summary:     "Find specifications not implemented in code",
			Description: "Identifies features specified but not yet implemented",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_specified_but_not_implemented.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/prioritize_gaps_by_impact.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_gap_report.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"audit-impl-to-spec": {
			Display:     "Audit Implementation to Spec Gap",
			Summary:     "Find implementation not covered by specifications",
			Description: "Identifies code that exists but is not documented in specifications",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_implemented_but_not_specified.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/prioritize_gaps_by_impact.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_gap_report.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-spec-feat": {
			Display:     "Draft Plan: Spec Feature",
			Summary:     "Create plan for new specification feature",
			Description: "Analyzes feature requirements and creates implementation plan focused on specifications",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/understand_feature_requirements.md"},
				{Path: "builtin:fragments/orient/identify_affected_specs.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-spec-fix": {
			Display:     "Draft Plan: Spec Bug Fix",
			Summary:     "Create plan for specification bug fix",
			Description: "Analyzes bug root cause and creates fix plan focused on specifications",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/understand_bug_root_cause.md"},
				{Path: "builtin:fragments/orient/identify_spec_deficiencies.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-spec-refactor": {
			Display:     "Draft Plan: Spec Refactor",
			Summary:     "Create plan for specification refactoring",
			Description: "Identifies structural issues in specs and creates refactoring plan",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_structural_issues.md"},
				{Path: "builtin:fragments/orient/identify_duplication.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-spec-chore": {
			Display:     "Draft Plan: Spec Maintenance",
			Summary:     "Create plan for specification maintenance tasks",
			Description: "Identifies maintenance needs in specs and creates chore plan",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_maintenance_needs.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-impl-feat": {
			Display:     "Draft Plan: Implementation Feature",
			Summary:     "Create plan for new implementation feature",
			Description: "Analyzes feature requirements and creates implementation plan focused on code",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/understand_feature_requirements.md"},
				{Path: "builtin:fragments/orient/identify_affected_code.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-impl-fix": {
			Display:     "Draft Plan: Implementation Bug Fix",
			Summary:     "Create plan for implementation bug fix",
			Description: "Analyzes bug root cause and creates fix plan focused on code",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_specs.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/understand_bug_root_cause.md"},
				{Path: "builtin:fragments/orient/identify_affected_code.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-impl-refactor": {
			Display:     "Draft Plan: Implementation Refactor",
			Summary:     "Create plan for code refactoring",
			Description: "Identifies code smells and complexity issues, creates refactoring plan",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_code_smells.md"},
				{Path: "builtin:fragments/orient/identify_complexity_issues.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
		"draft-plan-impl-chore": {
			Display:     "Draft Plan: Implementation Maintenance",
			Summary:     "Create plan for code maintenance tasks",
			Description: "Identifies maintenance needs in code and creates chore plan",
			Observe: []config.FragmentAction{
				{Path: "builtin:fragments/observe/read_agents_md.md"},
				{Path: "builtin:fragments/observe/read_task_input.md"},
				{Path: "builtin:fragments/observe/read_impl.md"},
			},
			Orient: []config.FragmentAction{
				{Path: "builtin:fragments/orient/identify_maintenance_needs.md"},
			},
			Decide: []config.FragmentAction{
				{Path: "builtin:fragments/decide/break_down_into_tasks.md"},
				{Path: "builtin:fragments/decide/prioritize_tasks.md"},
				{Path: "builtin:fragments/decide/check_if_blocked.md"},
			},
			Act: []config.FragmentAction{
				{Path: "builtin:fragments/act/write_draft_plan.md"},
				{Path: "builtin:fragments/act/emit_success.md"},
			},
		},
	}
}
