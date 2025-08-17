# Goldmark GitHub Alerts

> [!WARNING]
>
> **DEPRECATED:** This extension is deprecated and should not be used. A newer extension named
> [`gm-alert-callouts`](https://github.com/ZMT-Creative/gm-alert-callouts) (*based on the same
> underlying code*) should be used instead. This new package is already available on
> [`pkg.dev.go`](https://pkg.go.dev/github.com/ZMT-Creative/gm-alert-callouts).

## Legacy Information

This `goldmark-gh-alerts` package is an extension for the
[Goldmark](http://github.com/yuin/goldmark) Markdown Rendering Package that allows you to use
[GitHub alerts](https://docs.github.com/en/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax#alerts).
It also supports Obsidian-style callouts, although the Open/Close feature of Obsidian Callouts is
not yet implemented.

> [!NOTE]
> This extension does **not** directly include any icons -- it just provides the parsing and
> rendering functionality to create the alerts/callouts. The user must provide the list of valid
> alert/callout names with a string map containing the alert/callout identifier as the key (*e.g.,
> `note`, `important`, etc.*) and the icon as the string value. The icon is usually an SVG in HTML
> format (*e.g., `<svg>...svg-definiton...</svg>`*), but can be any string that a browser or
> application can render (*e.g., a Unicode glyph or an HTML entity code*).

## State of the Project

> [!WARNING]
> This project, under this name (*`github.com/ZMT-Creative/goldmark-gh-alerts`*), will not be
> supported beyond the `0.2.2` version. This project will remain available on GitHub, but will be
> removed from `pkg.go.dev` (*eventually*). We are working on a revision of this package with more
> features and will be creating a new project (*likely named
> `github.com/ZMT-Creative/gm-alert-callouts`*).

This ZMT-Creative project is a hard fork of:
[thiagokokada/goldmark-gh-alerts](https://github.com/thiagokokada/goldmark-gh-alerts).

As stated on Thiago Okada's original project's page, his `goldmark-gh-alerts` extension package has
been created primarily to support another of his projects and is not meant for general usage.

ZMT-Creative's package is also being used for a specific (*currently private*) project to create a
standalone Markdown Reader application.

If you want to use this package in your own project feel free, but it is recommended that you
should either pin a commit or fork since the API is not guarantee to be stable at this time.

## Changes from the Original Extension by Thiago Okada

This modified version of the GitHub Alerts extension adds `<div>` wrappers around the alert
**Title** text and the alert **Body** text. This allows more detailed styling with CSS. A new
`examples` folder containing a more detailed usage example has also been added (see [More Detailed
Example](#more-detailed-example) below).

## Examples

### Basic example

#### **Go**

```go
var markdown = goldmark.New(
  goldmark.WithExtensions(
    &GhAlerts{
      Icons: map[string]string{"note": `<svg class="octicon octicon-info mr-2" viewBox="0 0 16 16"
      version="1.1" width="16" height="16" aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0
      1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5 6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0
      1 .75.75v2.75h.25a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8
      6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z"></path></svg>`},
    },
  ),
)
```

#### **Markdown**

```markdown
> [!NOTE]
> Useful information that users should know, even when skimming content.
```

#### **HTML**

```html
<div class="gh-alert gh-alert-note">
  <div class="gh-alert-title">
    <p>
      <svg class="octicon octicon-info mr-2" viewBox="0 0 16 16" version="1.1" width="16" height="16"
      aria-hidden="true"><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8Zm8-6.5a6.5 6.5 0 1 0 0 13 6.5
      6.5 0 0 0 0-13ZM6.5 7.75A.75.75 0 0 1 7.25 7h1a.75.75 0 0 1 .75.75v2.75h.25a.75.75 0 0 1 0
      1.5h-2a.75.75 0 0 1 0-1.5h.25v-2h-.25a.75.75 0 0 1-.75-.75ZM8 6a1 1 0 1 1 0-2 1 1 0 0 1 0 2Z">
      </path></svg>
      Note
    </p>
  </div>
  <div class="gh-alert-body">
    <p>Useful information that users should know, even when skimming content.</p>
  </div>
</div>
```

### More Detailed Example

A more detailed code example is located in the `examples` folder. If you are on Windows you can run
the `run-ghalerts.ps1` script which will generate the HTML output of sample GitHub Alerts markdown
text. This will write the output to `example.html` and then start the default web browser to view
it.

If you are on MacOS or Linux, just do the following in the `examples` folder (*this assumes the
`open` command is available on your MacOS or Linux system*):

```sh
go run ./ghalerts.go > example.html
open example.html
```

The example shows one possible way to implement a set of alert types and their icons.

## License

This project is licensed under the [MIT License](LICENSE.md)

## Project Lineage

Portions of this software are based on the work of others, used under their
respective MIT Licenses. Details can be found in the following files:

- [Adam Chovanec](LICENSE-chovanec.md)
- [Thiago Okada](LICENSE-thiagokokada)
