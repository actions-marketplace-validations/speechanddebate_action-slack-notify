# Slack Notify - GitHub Action

This is a fork of [rtCamp/action-slack-notify](rtCamp/action/slack-notify) with the following changes:
* removes support for Vault as unnecesary
* Makes the default minimal message even more minimal for max minimalism

## Usage

You can use this action after any other action. Here is an example setup of this action:

1. Create a `.github/workflows/slack-notify.yml` file in your GitHub repo.
2. Add the following code to the `slack-notify.yml` file.

```yml
on: push
name: Slack Notification Demo
jobs:
  slackNotification:
    name: Slack Notification
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
    - name: Slack Notification
      uses: speechanddebate/action-slack-notify@master
      env:
        SLACK_WEBHOOK: ${{ secrets.SLACK_WEBHOOK }}
        SLACK_CHANNEL: "#tech-updates"
        SLACK_USERNAME: "Github Actions"
        SLACK_ICON: ":white_check_mark:"
```

3. Create `SLACK_WEBHOOK` secret using [GitHub Action's Secret](https://help.github.com/en/actions/configuring-and-managing-workflows/creating-and-storing-encrypted-secrets#creating-encrypted-secrets-for-a-repository). You can [generate a Slack incoming webhook token from here](https://slack.com/apps/A0F7XDUAZ-incoming-webhooks).

## Credits
Based on: [rtCamp/action-slack-notify](https://github.com/rtCamp/action-slack-notify)

## License
[MIT](LICENSE) © 2022 rtCamp
