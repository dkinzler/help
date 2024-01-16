
| jq | |
| --- | --- |
`... \| jq '.'` | output pretty json
`... \| jq '.key1, .key2'` | select specific keys
`... \| jq '.[] \| {k1: .a.b, k2: .a.c}'` | chain filters, for each element of array will create a new object with keys k1, k2


| xargs |  |
| --- | --- |
use outputs of another command as arguments |
`... \| xargs -I rm -r` | execute once for each input line
`... \| xargs -I rm -r _` | replace each placeholder

| ripgrep (rg) | |
| --- | --- |
`rg exp` | recursively search every file in the current directory for given regular expression
`rg -tpy exp` | only search .py files
`rg exp --glob "*abc.*"` | search files matching glob
`rg --hidden --no-ignore exp` | also search hidden files and those in .gitignore
`rg --files \| rg exp` | search for filenames matching expression
`rg --files-with-matches exp` | only list matched files in output

| Editing text | |
| --- | --- |
`... \| sed 's/a/b/g'` | replace all occurences of a with b in all input lines
`awk '/foo/ {print $5}' file` | for every line containing `foo` print fifth column, in space-separated file
`awk -F ',' ...` | use , as separator
`awk 'NR%3==1' file` | print every third line
`awk '($10 == value)' file` | print all lines where 10th column value matches

| Others | |
| --- | --- |
`find root_path -name '*.ext'` | find files, options `-type d -maxdepth -mindepth -size -mtime` and more
`tr` | replace characters in input
`cut -d , -f 5- file` | extract segments of each line of the input
`head -n 5 file` | output first 5 lines, `-n -5` for last, can also work on bytes
`tail -n 5 file` | output last 5 lines, `-n +5` start at line 5
`join -t ',' file1 file2` | join lines of file1 and file2 that have equal first field
`sort file` | sort file in ascending order, `-n --reverse --unique --ignore-case` and other options
