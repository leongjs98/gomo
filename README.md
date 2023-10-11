# gomo

A terminal pomodoro app for the minimalists.

## Features

- chain into other commands (e.g. `notify-send`)
- stdout the task details (for logging purposes)
    - start and end time
- pause the task by pressing p

### Upcoming Features
- stdout pause and unpause time in logging
- minimalistic way to cycle between task and rest

## Usage
```
gomo --time "60m" --task "Write gomo app" --description "Random description" && notify-send "You have completed the task!"
```
