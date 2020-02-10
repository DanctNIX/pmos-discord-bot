# postmarketOS bot for Discord

Fork of the original [matrix-bot](https://gitlab.com/postmarketOS/matrix-bot) for postmarketOS Matrix channels.

A bot that listens for keywords (e.g. `pma!27`) and sends the full URL to the GitLab issue / merge request back (e.g. `https://gitlab.com/postmarketOS/pmaports/merge_requests/27`).

## Building
```sh
go build
```

## Usage

First, Invite the bot into your server then:

```sh
export BOT_TOKEN="token"
./pmos-bot
```
