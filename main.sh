export GITHUB_BRANCH=${GITHUB_REF##*heads/}
export SLACK_USERNAME=${SLACK_USERNAME:-"GitHub Actions"}
export COMMIT_MESSAGE=$(cat "$GITHUB_EVENT_PATH" | jq -r '.commits[-1].message')

slack-notify "$@"
