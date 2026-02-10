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
	}
}
