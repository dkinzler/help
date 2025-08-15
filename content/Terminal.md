| General | |
| --- | --- |
`Ctrl+Space` | accept result of autosuggestions
`Ctrl+t` | insert result of fzf
`Ctrl+r` | fuzzy search command history 
`Alt+d` | fuzzy cd
`Ctrl+w/u/a/e` | delete word/line/start/end
`!!` | last command
`!*`  | last command's parameters
`Key=value; command` | run command with variables

<br>

## Tmux
| General | |
| --- | --- |
`:` | enter command mode
`Ctrl+d` | kill pane/window/session
`[` | enter copy mode, vim keybinds work
`s` | list sessions
`w` | list sessions + windows

| Sessions | |
| --- | --- |
`:new -s xyz` |  start a new session
`:kill-session -t xyz` | kill a session
`$` | rename session
`d` | detach from session
`( / )` |  previous/next session

| Windows | |
| --- | --- |
`c` | create 
`,` | rename
`&` | close 
`p / n` | previous/next
`0...9` | window by number
`:swap-window -s 2 -t 1` |  swap windows, `-t -1` to move current to the left

| Panes | |
| --- | --- |
`Ctrl+h/j/k/l` |  goto pane
`;` |  goto last active
`% / "` |  split vertically / horizontally
`{ / }` | move current pane left / right
`Space` |  toggle pane layouts
`q` |  show pane numbers
`z` |  toggle zoom 
`!` | convert pane to window
`h/j/k/l` | resize
