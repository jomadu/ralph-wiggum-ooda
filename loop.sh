#!/bin/bash
# Usage: ./loop.sh <prompt_file> [max_iterations]
# Examples:
#   ./loop.sh PROMPT_build.md                    # Unlimited iterations
#   ./loop.sh PROMPT_build.md 20                 # Max 20 iterations

# Parse arguments
if [ -z "$1" ]; then
    echo "Error: Prompt file required"
    echo "Usage: ./loop.sh <prompt_file> [max_iterations]"
    exit 1
elif [ -f "$1" ]; then
    PROMPT_FILE="$1"
    MAX_ITERATIONS=${2:-0}
else
    echo "Error: Prompt file '$1' not found"
    exit 1
fi

ITERATION=0
CURRENT_BRANCH=$(git branch --show-current)

echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "Prompt: $PROMPT_FILE"
echo "Branch: $CURRENT_BRANCH"
[ $MAX_ITERATIONS -gt 0 ] && echo "Max:    $MAX_ITERATIONS iterations"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

while true; do
    if [ $MAX_ITERATIONS -gt 0 ] && [ $ITERATION -ge $MAX_ITERATIONS ]; then
        echo "Reached max iterations: $MAX_ITERATIONS"
        break
    fi

    cat "$PROMPT_FILE" | kiro-cli chat --no-interactive --trust-all-tools

    git push origin "$CURRENT_BRANCH" || {
        echo "Failed to push. Creating remote branch..."
        git push -u origin "$CURRENT_BRANCH"
    }

    ITERATION=$((ITERATION + 1))
    echo -e "\n\n======================== LOOP $ITERATION ========================\n"
done
