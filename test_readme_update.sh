#!/bin/bash
set -e

echo "Testing specs/README.md updates..."

# Test 1: documentation-lifecycle.md is in Topics of Concern table
if ! grep -q "documentation-lifecycle" specs/README.md; then
    echo "FAIL: documentation-lifecycle.md not found in README"
    exit 1
fi

# Test 2: documentation-lifecycle.md is in Specification Status table
if ! grep -q "\[documentation-lifecycle\](documentation-lifecycle.md)" specs/README.md; then
    echo "FAIL: documentation-lifecycle.md link not found in Specification Status table"
    exit 1
fi

# Test 3: JTBD for documentation-lifecycle.md is present
if ! grep -q "Specify how existing procedures interact with docs through read-verify-update cycles" specs/README.md; then
    echo "FAIL: documentation-lifecycle.md JTBD not found"
    exit 1
fi

# Test 4: documentation-lifecycle.md is in the Configuration & Interface section
if ! grep -A 10 "### Configuration & Interface" specs/README.md | grep -q "documentation-lifecycle"; then
    echo "FAIL: documentation-lifecycle.md not in Configuration & Interface section"
    exit 1
fi

echo "All tests passed!"
