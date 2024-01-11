### Terminal

* `Ctrl + Space` to accept result of autosuggestions
* `Ctrl + t` to insert result of fzf
* `Ctrl + w/u/a/e` delete word/line/start/end
* globbing
	* `ls **/*.txt`  all dirs and subdirs
	* `ls **/*(.Lm+250)` files larger than 250mb
* placeholders
	* `!!` last command
	* `!*`  last command's parameters
	* `Key=value; command` set variables

Commands
* `ls | xargs -i cmd {}`  repeat for every input, placeholder `{}`
* `man ascii` for ASCII table
* `take x`  to create directory and cd
* `rd` remove dir
* `d / cd -3`  print last pwds, goto 3 in that list
* `copypath`  copy pwd to clipboard
* `copyfile`  copy file contents

Regexp (Perl)
* `.`  any char
* `x?` optional
* `x* / x+ / x{n}`  0+ / 1+ / n times
* `x | y` or
* `()`  grouping
* `[a-z0-9] [aez]` char sets 
* `^x`  x at start 
* `x$`  x at end
* `\d \w`  digit, letter + digit + _ 
* `grep -P "^a[0-9]{12}.txt"` 

### Tools

* `glances`  system monitor
* `httpie`  http client

### Tmux

session persistence

commands
* `prefix + :` to enter command mode
* `new -s xyz`  start a new session

`ctrl+d` kill pane/window/session

sessions
* all of these are with prefix
* `tmux kill-session -t xyz` kill a session
* `$` rename session
* `d` detach from session
* `s` list sessions
* `w`  list sessions + window preview
* `( / )`  previous/next session

windows
* `c` create 
* `,` rename
* `&` close 
* `p / n` previous/next
* `0...9` window by number
* `:swap-window -s 2 -t 1`  swap windows, `-t -1` to move current to the left

panes
* `;`  goto last active
* `% / "`  split vertically / horizontally
* `{ / }` move current pane left / right
* `Space`  toggle pane layouts
* `q`  show pane numbers
* `z`  toggle zoom 
* `!` convert pane to window
* `h j k l` resize

copy mode
* `[` enter
* vim keybinds apply
