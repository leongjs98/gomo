# gomo

A terminal pomodoro app for the minimalists. Developed and tested on Linux.

## Features

- [x] chain into other commands (e.g. `notify-send`)
- [ ] log (append) the task details
    - start and end time
- [ ] detect Ctrl-c to cancel the task and stdout the stop time

### Roadmap
- tags
- customise loging format using golang template
- toggle pausing by pressing `p`
- logout pause and unpause time in logging
- minimalistic way to cycle between task and rest
- progress bar

## Usage
```
gomo --time "2s" --name "Write Gomo CLI" --description "Writing an minimalistic CLI pomodoro" && notify-send "Task done" "Take a rest"
```

### Rest timer
```
gomo --name "Rest" --time "10m"
```

### Disable logging
```
gomo --log /dev/null ...
```

