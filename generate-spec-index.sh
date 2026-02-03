#!/usr/bin/env bash
set -euo pipefail

# Generate specs/README.md index from existing spec files

SPECS_DIR="specs"
OUTPUT_FILE="$SPECS_DIR/README.md"

# Extract JTBD from a spec file
extract_jtbd() {
    local file="$1"
    # Look for "## Job to be Done" or "## Jobs to be Done" section
    # Extract the next non-empty line after the header
    awk '
        /^## Jobs? to be Done/ { found=1; next }
        found && /^[^#]/ && NF > 0 { print; exit }
    ' "$file"
}

# Generate README content
generate_readme() {
    cat <<'EOF'
# Specifications

This directory contains specifications for the ralph-wiggum-ooda framework.

See [TEMPLATE.md](TEMPLATE.md) for the specification structure.
See [specification-system.md](specification-system.md) for the spec system design.

## Specifications

EOF

    # List all spec files (excluding README.md and TEMPLATE.md)
    for spec in "$SPECS_DIR"/*.md; do
        basename_spec=$(basename "$spec")
        
        # Skip README.md, TEMPLATE.md, and specification-system.md
        if [[ "$basename_spec" == "README.md" ]] || \
           [[ "$basename_spec" == "TEMPLATE.md" ]] || \
           [[ "$basename_spec" == "specification-system.md" ]]; then
            continue
        fi
        
        # Extract JTBD
        jtbd=$(extract_jtbd "$spec")
        
        # If no JTBD found, use filename as fallback
        if [[ -z "$jtbd" ]]; then
            jtbd="(No JTBD specified)"
        fi
        
        # Output entry
        echo "### [$basename_spec]($basename_spec)"
        echo "$jtbd"
        echo ""
    done
}

# Main execution
echo "Generating $OUTPUT_FILE..."
generate_readme > "$OUTPUT_FILE"
echo "Done. Generated index for $(find "$SPECS_DIR" -name "*.md" ! -name "README.md" ! -name "TEMPLATE.md" ! -name "specification-system.md" | wc -l | tr -d ' ') specifications."
