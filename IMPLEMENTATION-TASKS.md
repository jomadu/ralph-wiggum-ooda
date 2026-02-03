# Repository Restructuring Implementation Tasks

Created from PLAN.md - 17 beads issues for implementing the repository reorganization.

## Execution Order

### Pre-Execution (Optional)
- **ralph-wiggum-ooda-bpa** (P2) - Test rollback procedure before starting

### Phase 1-3: Core Restructuring (P0 - Blocking)
1. **ralph-wiggum-ooda-460** (P0) - Create directory structure (src/, specs/)
2. **ralph-wiggum-ooda-f7u** (P0) - Move implementation files to src/ with path resolution fix
3. **ralph-wiggum-ooda-8pj** (P0) - Update rooda-config.yml path references

### Phase 4-5: Documentation Organization (P1)
4. **ralph-wiggum-ooda-edn** (P1) - Create specification files in specs/
5. **ralph-wiggum-ooda-x5y** (P1) - Reorganize user documentation in docs/

### Phase 6a-d: README Updates (P1)
6. **ralph-wiggum-ooda-clj** (P1) - Update README installation instructions
7. **ralph-wiggum-ooda-2hg** (P1) - Update README example commands and sample structure
8. **ralph-wiggum-ooda-26r** (P1) - Update README internal documentation links
9. **ralph-wiggum-ooda-j83** (P1) - Add README framework developer vs consumer usage section

### Phase 7-8: Final Touches (P1-P2)
10. **ralph-wiggum-ooda-izm** (P1) - Update AGENTS.md definitions
11. **ralph-wiggum-ooda-r0e** (P2) - Create wrapper script for framework developers

### Phase 9: Verification (P1)
12. **ralph-wiggum-ooda-6bt** (P1) - Final verification and testing

### Dogfooding Validation (P1)
13. **ralph-wiggum-ooda-2jp** (P1) - Run draft-plan-impl-to-spec
14. **ralph-wiggum-ooda-367** (P1) - Publish plan to beads
15. **ralph-wiggum-ooda-ja4** (P1) - Build specs from work tracking
16. **ralph-wiggum-ooda-dwp** (P1) - Verify spec quality

## Priority Breakdown

- **P0 (Blocking):** 3 tasks - Core restructuring that must happen first
- **P1 (High):** 12 tasks - Documentation, verification, and dogfooding
- **P2 (Medium):** 2 tasks - Optional rollback testing and wrapper script

## Estimated Timeline

- **Phases 1-3:** 40 minutes (core restructuring)
- **Phases 4-5:** 45 minutes (documentation organization)
- **Phases 6a-d:** 55 minutes (README updates)
- **Phases 7-8:** 30 minutes (AGENTS.md and wrapper)
- **Phase 9:** 30 minutes (verification)
- **Dogfooding:** 40 minutes (4 steps)
- **Rollback testing:** 15 minutes (optional, before execution)

**Total:** 4-6 hours (including debugging and iteration)

## Dependencies

Each task lists its dependencies in the description. Key dependency chains:

1. Phase 1 → Phase 2 → Phase 3 (sequential core restructuring)
2. Phase 1 → Phase 4 → Phase 5 (documentation organization)
3. Phases 2,3 → Phase 6a → Phase 6b → Phase 6d (README updates)
4. Phases 4,5 → Phase 6c (documentation links)
5. Phases 2,3 → Phase 7 (AGENTS.md)
6. Phases 2,3 → Phase 8 (wrapper script)
7. All phases → Phase 9 (verification)
8. Phase 9 → Dogfooding steps 1-4 (sequential validation)

## Success Criteria

The restructuring is successful when:

1. All 12 main phases complete (1-9 including 6a-d)
2. Verification checklist passes (Phase 9)
3. All 4 dogfooding steps complete successfully
4. Framework generates its own specifications using its own methodology

## Notes

- Each phase commits changes incrementally (no single large commit)
- Git history preserved through `git mv` commands
- Rollback testing validates recovery procedure before execution
- Dogfooding validates the entire restructuring by using the framework on itself
