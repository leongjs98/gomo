# gomo

A customisable terminal pomodoro app with defaults for the minimalists. 

Developed and tested on Linux Mint using Go 1.21.2

## Features

- [x] chain into other commands (e.g. `notify-send`)
- [x] log (append) the task details
    - start and end time
- [x] detect Ctrl-c to cancel the task and log task

### Roadmap
- toggle pausing by pressing `p`
- log pause and unpause time in logging
- config using $XDG_CONFIG_HOME
- minimalistic way to cycle between task and rest

## usage
```sh
gomo --time "2s" --name "Write Gomo CLI" && notify-send "Task done" "Take a rest"
```

### Notification sounds

Play your notification sound in the terminal
E.g. On Linux Mint, the notification sound is on `/usr/share/mint-artwork/sounds/notification.oga`

```sh
gomo --name "Work" --time "60m" && mpv /usr/share/mint-artwork/sounds/notification.oga
```

### Rest timer
```sh
gomo --name "Rest" --time "10m" && notify-send "Let's go!" "Work on your stuff again" && mpv /usr/share/mint-artwork/sounds/notification.oga
```

### Disable logging
```sh
gomo --log /dev/null ...
```
