#!/usr/bin/env bash
set -euo pipefail

# Test Makefile on clean checkout
# This script verifies that the Makefile works correctly on a fresh repository clone

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMP_DIR=$(mktemp -d)
REPO_URL="file://${SCRIPT_DIR}"

cleanup() {
    echo "Cleaning up temporary directory: ${TEMP_DIR}"
    rm -rf "${TEMP_DIR}"
}
trap cleanup EXIT

echo "=== Testing Makefile on Clean Checkout ==="
echo "Temporary directory: ${TEMP_DIR}"
echo "Repository: ${REPO_URL}"
echo

# Clone repository to temporary directory
echo "Step 1: Cloning repository..."
git clone "${REPO_URL}" "${TEMP_DIR}/ralph-wiggum-ooda"
cd "${TEMP_DIR}/ralph-wiggum-ooda"
echo "✓ Repository cloned"
echo

# Run make all
echo "Step 2: Running 'make all'..."
if make all; then
    echo "✓ make all succeeded"
else
    echo "✗ make all failed"
    exit 1
fi
echo

# Verify binary created
echo "Step 3: Verifying bin/rooda binary created..."
if [ -f "bin/rooda" ]; then
    echo "✓ bin/rooda exists"
    echo "Binary info:"
    ls -lh bin/rooda
    file bin/rooda
else
    echo "✗ bin/rooda not found"
    exit 1
fi
echo

# Verify binary is executable
echo "Step 4: Verifying binary is executable..."
if [ -x "bin/rooda" ]; then
    echo "✓ bin/rooda is executable"
else
    echo "✗ bin/rooda is not executable"
    exit 1
fi
echo

# Test binary runs
echo "Step 5: Testing binary runs..."
if ./bin/rooda --version; then
    echo "✓ Binary runs successfully"
else
    echo "✗ Binary failed to run"
    exit 1
fi
echo

# Verify tests passed (already run by make all, but check explicitly)
echo "Step 6: Running tests explicitly..."
if make test; then
    echo "✓ Tests pass"
else
    echo "✗ Tests failed"
    exit 1
fi
echo

# Verify lint runs
echo "Step 7: Running lint explicitly..."
if make lint; then
    echo "✓ Lint runs (warnings are acceptable)"
else
    echo "✗ Lint failed"
    exit 1
fi
echo

# Test clean target
echo "Step 8: Testing 'make clean'..."
if make clean; then
    echo "✓ make clean succeeded"
else
    echo "✗ make clean failed"
    exit 1
fi
echo

# Verify binary removed
echo "Step 9: Verifying binary removed after clean..."
if [ ! -f "bin/rooda" ]; then
    echo "✓ bin/rooda removed by clean"
else
    echo "✗ bin/rooda still exists after clean"
    exit 1
fi
echo

echo "=== All Tests Passed ==="
echo "✓ Repository clones successfully"
echo "✓ make all works"
echo "✓ Binary is created and executable"
echo "✓ Tests pass"
echo "✓ Lint runs"
echo "✓ make clean works"
