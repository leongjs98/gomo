# gomo

A terminal pomodoro app for the minimalists.

## Features

- chain into other commands (e.g. `notify-send`)
- record tasks into a log file (toml by default)
- pause the task by pressing p

### Upcoming Features
- customise file template using golang template
- option to output file into stdout

## Usage
```
gomo --minute "60" --task "Write gomo app" --description "Random description" && notify-send "You have completed the task!"
```
