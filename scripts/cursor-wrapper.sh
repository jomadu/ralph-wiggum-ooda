#!/bin/bash
# cursor-wrapper.sh - Wrapper for cursor agent that parses JSON output
# Reads prompt from stdin, pipes to agent, parses JSON, emits text to stdout
# Progress/tool info goes to stderr, accumulated text goes to stdout

# Check dependencies
if ! command -v jq &> /dev/null; then
  echo "Error: jq is required but not installed" >&2
  echo "Install: brew install jq (macOS) or apt-get install jq (Linux)" >&2
  exit 1
fi

if ! command -v agent &> /dev/null; then
  echo "Error: cursor agent is required but not installed" >&2
  exit 1
fi

# Read prompt from stdin
prompt=$(cat)

# Track state
accumulated_text=""
tool_count=0
start_time=$(date +%s)

# Pipe prompt to agent and parse JSON output
echo "$prompt" | agent -p --force --output-format stream-json --stream-partial-output | \
  while IFS= read -r line; do
    type=$(echo "$line" | jq -r '.type // empty')
    subtype=$(echo "$line" | jq -r '.subtype // empty')
    
    case "$type" in
      "system")
        if [ "$subtype" = "init" ]; then
          model=$(echo "$line" | jq -r '.model // "unknown"')
          echo "🤖 Using model: $model" >&2
        fi
        ;;
        
      "assistant")
        # Accumulate text deltas and emit to stdout
        content=$(echo "$line" | jq -r '.message.content[0].text // empty')
        if [ -n "$content" ]; then
          echo -n "$content"
          accumulated_text="$accumulated_text$content"
        fi
        ;;

      "tool_call")
        if [ "$subtype" = "started" ]; then
          tool_count=$((tool_count + 1))

          # Extract tool information (to stderr)
          if echo "$line" | jq -e '.tool_call.writeToolCall' > /dev/null 2>&1; then
            path=$(echo "$line" | jq -r '.tool_call.writeToolCall.args.path // "unknown"')
            echo "🔧 Tool #$tool_count: Creating $path" >&2
          elif echo "$line" | jq -e '.tool_call.readToolCall' > /dev/null 2>&1; then
            path=$(echo "$line" | jq -r '.tool_call.readToolCall.args.path // "unknown"')
            echo "📖 Tool #$tool_count: Reading $path" >&2
          fi

        elif [ "$subtype" = "completed" ]; then
          # Extract and show tool results (to stderr)
          if echo "$line" | jq -e '.tool_call.writeToolCall.result.success' > /dev/null 2>&1; then
            lines=$(echo "$line" | jq -r '.tool_call.writeToolCall.result.success.linesCreated // 0')
            size=$(echo "$line" | jq -r '.tool_call.writeToolCall.result.success.fileSize // 0')
            echo "   ✅ Created $lines lines ($size bytes)" >&2
          elif echo "$line" | jq -e '.tool_call.readToolCall.result.success' > /dev/null 2>&1; then
            lines=$(echo "$line" | jq -r '.tool_call.readToolCall.result.success.totalLines // 0')
            echo "   ✅ Read $lines lines" >&2
          fi
        fi
        ;;

      "result")
        duration=$(echo "$line" | jq -r '.duration_ms // 0')
        end_time=$(date +%s)
        total_time=$((end_time - start_time))

        echo "" # Final newline to stdout
        echo "🎯 Completed in ${duration}ms (${total_time}s total)" >&2
        echo "📊 Final stats: $tool_count tools, ${#accumulated_text} chars generated" >&2
        ;;
    esac
  done
