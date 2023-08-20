export SLACK_CHANNEL=${SLACK_CHANNEL:-"#tech-updates"}
export SLACK_USERNAME=${SLACK_USERNAME:-"GitHub Actions"}
export SLACK_ICON=${SLACK_ICON:-":x:"}
export COMMIT_MESSAGE=$(cat "$GITHUB_EVENT_PATH" | jq -r '.commits[-1].message')

slack-notify "$@"
