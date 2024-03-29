# toggl

(This is a port of my [toggl tool](https://github.com/mmcclimon/toggl) to go,
without the Linear integration.)

`toggl` is a little tool for tracking your time with [Toggl](https://toggl.com/).
Here's the help, at time of writing:

```
Usage:
  toggl [command]

Available Commands:
  abort       actually, you weren't doing that thing after all
  projects    list the buckets things can go in
  shortcuts   list the things you can start easily
  start       start doing a new thing
  stop        stop doing the thing you're doing
  timer       what are you doing right now?

Use "toggl [command] --help" for more information about a command.
```

## Config file

`toggl` is driven by a config file, which by default is at `~/.togglrc`, but
ocnfigurable if you set `TOGGL_CONFIG_FILE` in your environment. It's a
[TOML](https://toml.io/en/) file, pretty simple.

```
api_token = "your-token"    # your api token
workspace_id = 1234         # your workspace id

# This is a map of shortcut name to toggl project id.
[project_shortcuts]
evergreen = 123456
meetings  = 234567

# This is a map of shortcuts to description/projects.
# You can say `toggl start @email` to start a timer with the description
# "read email" in the "evergreen project", for example.
[task_shortcuts]
email = { desc = "read email",       project = "evergreen" }
11s   = { desc = "1:1 with manager", project = "meetings"  }
```

For everything else, the built-in help should do you.

## Jira

There's also Jira integration; it's pretty specific to my own needs, but it
might be useful.

```
[jira]
url = "https://jira.example.org"
access_token = "my token"
access_secret = "my secret"
consumer_key = "my key"
key_file = "/Users/michael/.ssh/my-jira.crt"

[jira.projects]
REP-1256   = 190548682  # optimization
BF_DEFAULT = 190665570  # bfs
DEFAULT    = 190160292  # misc
```
