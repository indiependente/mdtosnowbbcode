# mdtosnowbbcode
Markdown to ServiceNow BBCode

## What it does
This tool converts Markdown code to ServiceNow compliant BBCode. 

## Example
Markdown input
```markdown
## Markdown content
This is example markdown
```code block```
**strong text** ~~strikethrough~~.
```
BBCode output
```html
[code]<h2>Markdown content</h2>

<p>This is example markdown
<code>code block</code>
<strong>strong text</strong> <del>strikethrough</del>.</p>
[/code]
```

## How to install
`go get -u github.com/indiependente/mdtosnowbbcode/cmd/mdtosnow`

## Run
The tool reads content from STDIN and prints the converted text to STDOUT.

`cat markdownfile.md | mdtosnow`

E.g. On macOS

`pbpaste | mdtosnow | pbcopy`

## Manual build
```bash
git clone git@github.com:indiependente/mdtosnowbbcode.git
cd mdtosnowbbcode
make build
cat markdownfile.md | ./mdtosnow
```