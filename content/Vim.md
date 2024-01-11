 To learn, you have to slow down a bit, think about what you want to accomplish and then think of a way to accomplish it.

Start
* open project dir not specific file, will work nicely with live_grep etc. if cwd is set correctly

Movement
* in line
	* `ea`  insert at end of word
	* when inserting `()` how to move after `)` and keep typing
		* could use `<esc> Ea`
* between lines
	* relative jumps `5j`
	* `/, ?` where possible
	* `]] [[` go to next/prev occurrence of word under cursor
	* `*`  to search for word under cursor, `#` for prev
	* in visual mode `o / O` to move between ends
	* `-` first char in line above
	* `C-o / C-i` navigate jumplist back/forward
		* moving back from a goto def or after `gg`, not all movements are recorded
		* add a pos to jumplist with `m'`
	* `gi` last position you were in insert mode in
	* marks
		* `m[a-z]` to set, `'[a-z]`  to jump 
		* lowercase specific to buffer, uppercase between buffers
	* `{ }` jump between paragraphs
* navigating files
	* fuzzy finder/telescope to open
	* quickfix list
		* populate
			* with search results from telescope `C-q`
			* run a tool like make/compiler/linter/tests that output a list of files + lines
		* how to use
			* `:cope(n)` show list or use trouble/telescope
			* `:cnext :cprev` or keybinds to jump to next/prev
			* `:cdo` to apply an operation to every item in quickfix list, can be used with macros
			* can have multiple quickfix lists and you can move between them `:colder  :cnewer`
	
Editing
* insert mode
	* `C+w` delete word before cursor
	* `C+j` insert line break
	* `C+t/d` indent/deindent
* surround 
	* `sa / sd / sr`
	* word with function call: `ciw` type `fn()`, exit normal mode, paste with `p`
* `J / gJ` join line below (without spaces)
* paste
	* and replace
		* a word, do `viw` to select, `p` to paste over
		* inside `(, {, [`, select with `vi(`, paste
	* `gp / gP` paste and move cursor after pasted text, `gpa` paste and keep typing
* `cc / S / C` change line, to end of line
* `gcc` toggle line comment
* `gu / gU`  switch to lower/upper up to motion
* Text objects
	* `b = (), B = {}`
	* `aa / ia` argument
	* `af / if` function
	* `ai / ii` blocks
		* `ii` works for languages like python that don't have {}
		* `ai` works for both python like and those with explicit blocks, can e.g. change an entire if, for, ...
* aligning, e.g. ` = ap` to align paragraph
* increment `C-a C-a`, decrement `C-x`
* undoing
	* to last write `:earlier 1f` and back with `later 1f`
 * search and replace
	* `:s/old/new` replace first match in line
	* `:s/old/new/g` replace all in line
	* `:%s/old/new/g` replace all in file
	* `:%s/old/new/g` ask for confirmation
	* `:%s/old/new/gi` ignore case
	 * `:g/pattern/norm A`  run command on every line with a matches
* macros
	* `qx`  record into register x
	* `q`  stop recording
	* `@x`  replay x
	* some notes
		* `^` in the beginning to go to start of line, so that you can later apply the macro anywhere in the line
		* avoid `hjkl` motions if possible, if you e.g. use `f` or `/` to find it makes it more repeatable
		* can be recursive

LSP and co.
* `K K`  to move into hover window
* `L` signature help
* diagnostics
	* `tt` toggle trouble
	* `ee` show diagnostics
	* `en / ep` next/prev diagnostic


Windows
* `C+ws/C+wv` split window horizontally/vertically
* `:q / C+wq` close window
* `C+w +/-/>/</=` to change height/width of a window, make equal with ` = `