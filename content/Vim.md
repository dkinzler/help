 To learn, you have to slow down a bit, think about what you want to accomplish and then think of a way to accomplish it.

| Move in line | |
| --- | --- |
`ea` | insert at end of word
`<esc>Ea` | when inserting `()` move after `)` and keep typing


| Move between lines | |
| --- | --- |
`5j` | relative jumps
`/, ?` | find where possible
`]] [[` | go to next/prev occurrence of word under cursor
`*` |  to search for word under cursor, `#` for prev
`o / O` | in visual mode move between ends
`-` | first char in line above
`C-o / C-i` | navigate jumplist back/forward, e.g. after goto def or `gg`, not all movements are recorded
`m'` | add a pos to jumplist
`gi` | last position you were in insert mode in
`m[a-z]` | set mark, lowercase specific to buffer, uppercase between buffers
`'[a-z]`| jump to mark
`{ }` | jump between paragraphs

### Navigating files
* fuzzy finder/telescope to open
* `gs` `gws` to fuzzy find LSP symbols
    * can type e.g. `function` after keyword to match only a certain type of symbol
* quickfix list
	* populate
		* with search results from telescope `C-q`
		* run a tool like make/compiler/linter/tests that output a list of files + lines
	* how to use
		* `:cope(n)` show list or use telescope
		* `:cnext :cprev` or keybinds to jump to next/prev
		* `:cdo` to apply an operation to every item in quickfix list, can be used with macros
		* can have multiple quickfix lists and you can move between them `:colder  :cnewer`
	

| Editing | |
| --- | --- |
`C+w` | delete word before cursor
`C+j` | insert line break
`C+t/d` | indent/deindent
`sa / sd / sr` | surround
`ciw fn() <esc> p`	| surround word with function call
`J / gJ` | join line below (without spaces)
`viwp` | pase and replace a word
`vi({[p` | replace inside bracket
`gp / gP` | paste and move cursor after pasted text
`gpa` | paste and keep typing
`cc / S / C` | change line, to end of line
`gcc` | toggle line comment
`gu / gU` |  switch to lower/upper up to motion
` = ap` | align paragraph
`C-a C-a` | increment
`C-x` | decrement
`:earlier 1f` | undo to last write, back with `later 1f`


| Text objects | |
| --- | --- |
`ax / ix` | outer/inner textobject
`anx / inx / alx / ilx` | next/last instance of outer/inner textobject
`b` | bracket `( { [`
`B` | curly bracket `{`
`q` | quote
`a` | argument
`F` | function
`f` | function call
`i` | block (if, for, ...), support depends on language


| Ex Commands | |
| --- | --- |
`C-n / C-p` | next / prev command in history
`q: / C-f` | open command line window from normal/command line mode
`:quit :close :exit` | close command line window
`:norm` | to run a key sequence in normal mode
`:15,20norm @q` | replay macro on lines 15 to 20
`:s` | search and replace, see below
`:5,$x` | last line in file
`:%x` | entire file
`:1,/that line/x` | first line that matches "that line"


| Search and Replace | |
| --- | --- |
`:s/old/new` | replace first match in line
`:s/old/new/g` | replace all in line
`:%s/old/new/g` | replace all in file
`:%s/old/new/g` | ask for confirmation
`:%s/old/new/gi` | ignore case
`:g/pattern/norm A` |  run command on every line with a matches


| Macros | |
| --- | --- |
`qx` |  record into register x
`q` |  stop recording
`@x` |  replay x

For better macros
* `^` in the beginning to go to start of line, so that you can later apply the macro anywhere in the line
* avoid `hjkl` motions if possible, if you e.g. use `f` or `/` to find it makes it more repeatable
* can be nested and recursive, i.e. call another macro or itself


| LSP, Diagnostics and co. | |
| --- | --- |
`K K` |  to move into hover window
`L` | signature help
`ee` | show diagnostics
`en / ep` | next/prev diagnostic


| Windows | |
| --- | --- |
`C+ws/C+wv` | split window horizontally/vertically
`:q / C+wq` | close window
`C+w +/-/>/</=` | to change height/width of a window, make equal with ` = `
