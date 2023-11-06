# gomo

A minimalistic terminal pomodoro app written with Golang.

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

## Installation
```sh
git clone https://github.com/leongjs98/gomo
cd gomo && go build main.go

# Copy to your $PATH
cp main ~/.local/bin/gomo
```

## usage
```sh
gomo --task "Your Task" --duration "30m" && notify-send "Task done" "Take a rest"
```

### Alias for notification

You are most likely to repeat the notifcation part on every task.
I suggest using alias for it since they are lengthy.
In your `.bashrc` or `.zshrc`, set your alias.

```sh
custom-notification='notify-send "Task done" "Take a rest" && mpv /path/to/notification_sound.mpv'
```

### Notification sounds

Use `&&` to play your notification sound in the terminal
E.g. On Linux Mint, the notification sound is on `/usr/share/mint-artwork/sounds/notification.oga`

```sh
gomo --task "Work" --duration "60m" && mpv /usr/share/mint-artwork/sounds/notification.oga
```

### Logging your task
The default log file is `./gomo.log`, you can change it

```sh
gomo --task "Work" --duration "60m" --log "./work.log"
```

### Rest timer
```sh
gomo --task "Rest" --duration "10m" && notify-send "Let's go!" "Work on your stuff again"
```

### Disable logging
```sh
gomo --log /dev/null ...
```
