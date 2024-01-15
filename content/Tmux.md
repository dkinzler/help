
| General | |
| --- | --- |
`i` | fuzzy cheat.sh
`:` | to enter command mode
`new -s xyz` |  start a new session
`Ctrl+d` | kill pane/window/session
`[` | enter copy mode, vim keybinds work

| Sessions | |
| --- | --- |
`tmux kill-session -t xyz` | kill a session
`$` | rename session
`d` | detach from session
`s` | list sessions
`w` |  list sessions + window preview
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
`;` |  goto last active
`% / "` |  split vertically / horizontally
`{ / }` | move current pane left / right
`Space` |  toggle pane layouts
`q` |  show pane numbers
`z` |  toggle zoom 
`!` | convert pane to window
`h j k l` | resize
