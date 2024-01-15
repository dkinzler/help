
| Terminal | |
| --- | --- |
`Ctrl+Space` | accept result of autosuggestions
`Ctrl+t` | insert result of fzf
`Ctrl+r` | fuzzy search command history 
`Alt+d` | fuzzy cd
`Ctrl+w/u/a/e` | delete word/line/start/end
`!!` | last command
`!*`  | last command's parameters
`Key=value; command` | set variables
`ls **/*.txt`  | all dirs and subdirs
`ls **/*(.Lm+250)` | files larger than 250mb
`ls \| xargs -i cmd {}` | repeat for every input, placeholder `{}`
`man ascii` | for ASCII table
`take x`  | to create directory and cd
`rd` | remove dir
`d / cd -3`  | print last pwds, goto 3 in that list
`copypath`  | copy pwd to clipboard
`copyfile`  | copy file contents
`glances` |  system monitor
`httpie` |  http client

| Regex (Perl) | |
| --- | --- |
`.`  | any char
`x?` | optional
`x* / x+ / x{n}` | 0+ / 1+ / n times
`x \| y` | or
`()` |  grouping
`[a-z0-9] [aez]` | char sets 
`^x` |  x at start 
`x$` |  x at end
`\d \w` |  digit, letter + digit + _ 
`grep -P "^a[0-9]{12}.txt"` | 
