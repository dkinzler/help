| General | |
| --- | --- |
`man ascii` | ASCII table
`take x`  | create dir and cd
`rd` | remove dir
`d / cd -3`  | print last pwds, goto 3 in that list
`copypath`  | copy pwd to clipboard
`copyfile`  | copy file contents
`ls **/*.txt`  | all dirs and subdirs
`ls **/*(.Lm+250)` | files larger than 250mb
`... \| xargs -i cmd {}` | repeat for every input, placeholder `{}`
`tr` | replace characters in input
`cut -d , -f 5- file` | extract segments of each line of the input
`head -n 5 file` | output first 5 lines, `-n -5` for last, can also work on bytes
`tail -n 5 file` | output last 5 lines, `-n +5` to start at line 5
`join -t ',' file1 file2` | join lines of file1 and file2 that have equal first field
`sort file` | sort file, `-n` numeric, `--reverse`,  `--unique`, `--ignore-case` and more

| fd | |
| --- | --- |
`fd pattern [path...]` | basic usage
`-t f/d/x/...` | type file/dir/executable/...
`-d --maxdepth` | max dir traversal depth
`--min-depth` | min dir traversal depth
`-S, -size +5m` | only files with size greater than 5 megabytes
`--changed-before 2weeks` | filter by modification time

| ripgrep (rg) | |
| --- | --- |
`rg pattern [path]` | recursively search every file in the given directory (or cwd if none) for given pattern
`rg -tpy exp` | only search .py files
`rg exp --glob "*abc.*"` | search files matching glob
`rg --hidden --no-ignore exp` | also search hidden files and those in .gitignore
`rg --files-with-matches exp` | only list matched files in output

| sed | |
| --- | --- |
`sed [options] command [file]` | basic usage, by default reads stdin, operates line by line
`sed '1,10 d'` | delete lines 1 to 10, relative range `5,+2`
`sed '/pattern/ d'` | delete lines matching pattern
`sed 's/a/b/'` | replace first occurence of pattern a with b in all lines, can use `/2` to replace only second instance
`sed 's/a/b/g'` | replace all occurences of pattern a with b in all lines
`sed 's/a/b/i'` | ignore case
`sed 's/a/&/'` | can refer to matched text with `&`, can also refer to regex capture groups `\1, \2, ...`
`a i c` | commands to append/insert/change a line

| jq | |
| --- | --- |
`... \| jq '.'` | output pretty json
`... \| jq '.key1, .key2'` | select specific keys
`... \| jq '.[] \| {k1: .a.b, k2: .a.c}'` | chain filters, for each element of array will create a new object with keys k1, k2

| awk | |
| --- | --- |
`awk '/foo/ {print $5}' file` | for every line containing `foo` print fifth column, in space-separated file
`awk -F ',' ...` | use , as separator
`awk 'NR%3==1' file` | print every third line
`awk '($10 == value)' file` | print all lines where 10th column value matches

### Regex
* different flavors: BRE (basic), ERE (extended), PCRE (perl-compatible)
    * BRE default in many older tools, have to escape some special characters
    * no escaping in ERE, PCRE has many advanced features
    * some tools let you choose the style e.g. `grep -E/P`
    * not super-fixed standards, many tools have their own quirks and subset of supported features
    * for basic expressions all will be mostly the same
* generally use `'expr'`, for `"expr"` shell will expand variables, escape certain chars and more
* basic syntax
	* `.` any char
	* character classes
		* `[aeh]`
		* ranges `[0-9a-zA-Z]`
		* negation `[^0-9]`
	* quantifiers
		* `* +` at least 0/1 times
		* `?` 0/1 times
		* `{m}` m times
		* `{m,n}` range
	* grouping `(...)`
		* "capture groups", can refer to them via `\1, \2, ...` not supported everywhere
	* alternatives `x | y`
	* anchors `^ $` start/beginning
* advanced
	* character class shortcuts
		* `\d` digit, `\D` for negation
		* `\w` (usually) equivalent to `[0-9a-zA-Z_]`, `\W` negation
		* `\s` white space (space, tab, line feed etc.), `\S` negation
	+ in ERE and PCRE you escape metacharacters with `\`
		* but e.g. not inside `[...]`
	* in PCRE quote mode `\Q...\E`, everything in between treated as literal text
