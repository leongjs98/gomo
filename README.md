# gomo

A terminal pomodoro app for the minimalists.

## Features

- chain into other commands (e.g. `notify-send`)
- record tasks into a log file (toml by default)
    - start and end time
- pause the task by pressing p

### Upcoming Features
- customise file template using golang template
- record pause and unpause time in logging
- minimalistic way to do tasks, and rest, and repeat
- option to output file into stdout

## Usage
```
gomo --minute "60" --task "Write gomo app" --description "Random description" && notify-send "You have completed the task!"
```
