package alerts

import (
	"strings"
	"testing"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/testutil"
)

// Additional test cases for comprehensive coverage
var additionalTestCases = [...]TestCase{
	// Edge cases and malformed syntax
	{
		desc: "Invalid alert type with numbers",
		md:   `> [!info123] title`,
		html: `<div class="gh-alert gh-alert-info123" data-callout="info123"><div class="gh-alert-title"><p><svg></svg>title</p></div></div>`,
	},
	{
		desc: "Alert type with special characters",
		md:   `> [!warning-special] title`,
		html: `<blockquote>
<p>[!warning-special] title</p>
</blockquote>
`,
	},
	{
		desc: "Empty alert type",
		md:   `> [!] title`,
		html: `<blockquote>
<p>[!] title</p>
</blockquote>
`,
	},
	{
		desc: "Missing closing bracket",
		md:   `> [!info title`,
		html: `<blockquote>
<p>[!info title</p>
</blockquote>
`,
	},
	{
		desc: "Missing opening bracket",
		md:   `> !info] title`,
		html: `<blockquote>
<p>!info] title</p>
</blockquote>
`,
	},
	{
		desc: "Closed alert with dash",
		md:   `> [!warning]- This is a closed alert`,
		html: `<div class="gh-alert gh-alert-warning" data-callout="warning"><div class="gh-alert-title"><p><svg></svg>This is a closed alert</p></div></div>`,
	},
	{
		desc: "Closed alert without title",
		md: `> [!tip]-
> content here`,
		html: `<div class="gh-alert gh-alert-tip" data-callout="tip"><div class="gh-alert-title"><p><svg></svg>Tip</p></div><div class="gh-alert-body"><p>content here</p>
</div></div>`,
	},

	// Case sensitivity tests
	{
		desc: "Mixed case alert type",
		md:   `> [!WaRnInG] Mixed case`,
		html: `<div class="gh-alert gh-alert-warning" data-callout="warning"><div class="gh-alert-title"><p><svg></svg>Mixed case</p></div></div>`,
	},
	{
		desc: "Uppercase alert type",
		md:   `> [!ERROR] Uppercase alert`,
		html: `<div class="gh-alert gh-alert-error" data-callout="error"><div class="gh-alert-title"><p><svg></svg>Uppercase alert</p></div></div>`,
	},

	// Complex content within alerts
	{
		desc: "Alert with code block",
		md: `> [!note]
> Here's some code:
> ` + "```" + `
> function test() {
>   return true;
> }
> ` + "```",
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p><svg></svg>Note</p></div><div class="gh-alert-body"><p>Here's some code:</p>
<pre><code>function test() {
  return true;
}
</code></pre>
</div></div>`,
	},
	{
		desc: "Alert with inline code",
		md: `> [!info]
> Use the ` + "`alert`" + ` function`,
		html: `<div class="gh-alert gh-alert-info" data-callout="info"><div class="gh-alert-title"><p><svg></svg>Info</p></div><div class="gh-alert-body"><p>Use the <code>alert</code> function</p>
</div></div>`,
	},
	{
		desc: "Alert with nested list",
		md: `> [!tip]
> - item 1
>   - nested item
>   - another nested
> - item 2`,
		html: `<div class="gh-alert gh-alert-tip" data-callout="tip"><div class="gh-alert-title"><p><svg></svg>Tip</p></div><div class="gh-alert-body"><ul>
<li>item 1
<ul>
<li>nested item</li>
<li>another nested</li>
</ul>
</li>
<li>item 2</li>
</ul>
</div></div>`,
	},
	{
		desc: "Alert with ordered list",
		md: `> [!important]
> 1. First step
> 2. Second step
> 3. Third step`,
		html: `<div class="gh-alert gh-alert-important" data-callout="important"><div class="gh-alert-title"><p><svg></svg>Important</p></div><div class="gh-alert-body"><ol>
<li>First step</li>
<li>Second step</li>
<li>Third step</li>
</ol>
</div></div>`,
	},
	{
		desc: "Alert with blockquote inside",
		md: `> [!note]
> > This is a nested quote
> > with multiple lines`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p><svg></svg>Note</p></div><div class="gh-alert-body"><blockquote>
<p>This is a nested quote
with multiple lines</p>
</blockquote>
</div></div>`,
	},

	// Multiple line breaks and empty lines
	{
		desc: "Alert with multiple empty lines",
		md: `> [!info]
> First paragraph
>
>
> Second paragraph after empty lines`,
		html: `<div class="gh-alert gh-alert-info" data-callout="info"><div class="gh-alert-title"><p><svg></svg>Info</p></div><div class="gh-alert-body"><p>First paragraph</p>
<p>Second paragraph after empty lines</p>
</div></div>`,
	},
	{
		desc: "Alert with trailing empty lines",
		md: `> [!warning]
> Content here
>
>
`,
		html: `<div class="gh-alert gh-alert-warning" data-callout="warning"><div class="gh-alert-title"><p><svg></svg>Warning</p></div><div class="gh-alert-body"><p>Content here</p>
</div></div>`,
	},

	// HTML escaping and special characters
	{
		desc: "Alert with HTML entities in title",
		md:   `> [!note] Title with &amp; <script> tags`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p><svg></svg>Title with &amp; <!-- raw HTML omitted --> tags</p></div></div>`,
	},
	{
		desc: "Alert with HTML in content",
		md: `> [!warning]
> Be careful with <strong>HTML</strong> & scripts`,
		html: `<div class="gh-alert gh-alert-warning" data-callout="warning"><div class="gh-alert-title"><p><svg></svg>Warning</p></div><div class="gh-alert-body"><p>Be careful with <!-- raw HTML omitted -->HTML<!-- raw HTML omitted --> &amp; scripts</p>
</div></div>`,
	},

	// Unicode and international content
	{
		desc: "Alert with unicode in title",
		md:   `> [!info] Título con acentos é çharacters 中文`,
		html: `<div class="gh-alert gh-alert-info" data-callout="info"><div class="gh-alert-title"><p><svg></svg>Título con acentos é çharacters 中文</p></div></div>`,
	},
	{
		desc: "Alert with emoji",
		md: `> [!tip] 🚀 Rocket tip
> Use emojis sparingly 😊`,
		html: `<div class="gh-alert gh-alert-tip" data-callout="tip"><div class="gh-alert-title"><p><svg></svg>🚀 Rocket tip</p></div><div class="gh-alert-body"><p>Use emojis sparingly 😊</p>
</div></div>`,
	},

	// Long content tests
	{
		desc: "Alert with very long title",
		md:   `> [!note] This is a very long title that goes on and on and on and should still work properly even with lots of text in the title section of the alert`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p><svg></svg>This is a very long title that goes on and on and on and should still work properly even with lots of text in the title section of the alert</p></div></div>`,
	},
	{
		desc: "Alert with very long paragraph",
		md: `> [!info]
> This is a very long paragraph that contains lots of text and should wrap properly in the alert body. It includes multiple sentences and should demonstrate that the alert can handle substantial amounts of content without any issues. The formatting should remain intact and the HTML output should be properly structured.`,
		html: `<div class="gh-alert gh-alert-info" data-callout="info"><div class="gh-alert-title"><p><svg></svg>Info</p></div><div class="gh-alert-body"><p>This is a very long paragraph that contains lots of text and should wrap properly in the alert body. It includes multiple sentences and should demonstrate that the alert can handle substantial amounts of content without any issues. The formatting should remain intact and the HTML output should be properly structured.</p>
</div></div>`,
	},

	// Edge cases with indentation
	{
		desc: "Alert with tabs in content",
		md: `> [!note]
>	Indented with tab
>		Double tab indent`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p><svg></svg>Note</p></div><div class="gh-alert-body"><blockquote>
<blockquote>
<p>Indented with tab
Double tab indent</p>
</blockquote>
</blockquote>
</div></div>`,
	},

	// Various alert types that might not be tested
	{
		desc: "Caution alert type",
		md: `> [!caution]
> Be very careful here`,
		html: `<div class="gh-alert gh-alert-caution" data-callout="caution"><div class="gh-alert-title"><p><svg></svg>Caution</p></div><div class="gh-alert-body"><p>Be very careful here</p>
</div></div>`,
	},
	{
		desc: "Important alert type",
		md: `> [!important] Critical information
> This is very important`,
		html: `<div class="gh-alert gh-alert-important" data-callout="important"><div class="gh-alert-title"><p><svg></svg>Critical information</p></div><div class="gh-alert-body"><p>This is very important</p>
</div></div>`,
	},

	// Multiple alerts in sequence
	{
		desc: "Multiple alerts back to back",
		md: `> [!info]
> First alert

> [!warning]
> Second alert`,
		html: `<div class="gh-alert gh-alert-info" data-callout="info"><div class="gh-alert-title"><p><svg></svg>Info</p></div><div class="gh-alert-body"><p>First alert</p>
</div></div>
<div class="gh-alert gh-alert-warning" data-callout="warning"><div class="gh-alert-title"><p><svg></svg>Warning</p></div><div class="gh-alert-body"><p>Second alert</p>
</div></div>`,
	},
}

// Test with different icon configurations
var markdownWithIcons = goldmark.New(
	goldmark.WithExtensions(
		&GhAlerts{
			Icons: map[string]string{
				"note":      "📝",
				"tip":       "💡",
				"warning":   "⚠️",
				"caution":   "🔥",
				"important": "❗",
			},
		},
	),
)

var iconTestCases = [...]TestCase{
	{
		desc: "Alert with icon",
		md: `> [!note]
> Content with icon`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p>📝Note</p></div><div class="gh-alert-body"><p>Content with icon</p>
</div></div>`,
	},
	{
		desc: "Alert using 'noicon' to suppress icon",
		md: `> [!noicon] Error
> No icon for this type`,
		html: `<div class="gh-alert gh-alert-error" data-callout="error"><div class="gh-alert-title"><p>Error</p></div><div class="gh-alert-body"><p>No icon for this type</p>
</div></div>`,
	},
	{
		desc: "Alert using 'noicon' to suppress icon with complex title formatting",
		md: `> [!noicon] Error **BAD* [text](#link) ` + "`some code`" + `
> No icon for this type`,
		html: `<div class="gh-alert gh-alert-error-bad" data-callout="error-bad"><div class="gh-alert-title"><p>Error *<em>BAD</em> <a href="#link">text</a> <code>some code</code></p></div><div class="gh-alert-body"><p>No icon for this type</p>
</div></div>`,
	},
	{
		desc: "Custom title with icon",
		md:   `> [!tip] Custom tip title`,
		html: `<div class="gh-alert gh-alert-tip" data-callout="tip"><div class="gh-alert-title"><p>💡Custom tip title</p></div></div>`,
	},
	{
		desc: "Unsupported alert type 'cite' should use 'note' icon",
		md:   `> [!cite]`,
		html: `<div class="gh-alert gh-alert-cite" data-callout="cite"><div class="gh-alert-title"><p>📝Cite</p></div></div>`,
	},
	{
		desc: "Unsupported alert type 'cite' with 'Quote' title should use 'note' icon",
		md:   `> [!cite] Quote`,
		html: `<div class="gh-alert gh-alert-cite" data-callout="cite"><div class="gh-alert-title"><p>📝Quote</p></div></div>`,
	},
}

// Test with no icons configuration
var markdownNoIcons = goldmark.New(
	goldmark.WithExtensions(
		&GhAlerts{},
	),
)

var noIconTestCases = [...]TestCase{
	{
		desc: "Alert without any icons configured",
		md: `> [!note]
> Content without icon`,
		html: `<div class="gh-alert gh-alert-note" data-callout="note"><div class="gh-alert-title"><p>Note</p></div><div class="gh-alert-body"><p>Content without icon</p>
</div></div>`,
	},
}

func TestAdditionalAlerts(t *testing.T) {
	for i, c := range additionalTestCases {
		testutil.DoTestCase(markdown, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}

func TestAlertsWithIcons(t *testing.T) {
	for i, c := range iconTestCases {
		testutil.DoTestCase(markdownWithIcons, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}

func TestAlertsWithoutIcons(t *testing.T) {
	for i, c := range noIconTestCases {
		testutil.DoTestCase(markdownNoIcons, testutil.MarkdownTestCase{
			No:          i,
			Description: c.desc,
			Markdown:    c.md,
			Expected:    c.html,
		}, t)
	}
}

// Test AST node functionality
func TestASTNodeCreation(t *testing.T) {
	// These tests verify that the AST nodes are created correctly
	// by attempting to convert to HTML and checking for errors

	testCases := []string{
		`> [!note] Test AST`,
		`> [!warning]
> Body content`,
		`> [!info]- Closed alert`,
	}

	for _, md := range testCases {
		var buf strings.Builder
		err := markdown.Convert([]byte(md), &buf)
		if err != nil {
			t.Errorf("Failed to parse markdown: %s, error: %v", md, err)
		}

		// Verify that some HTML was generated
		if buf.Len() == 0 {
			t.Errorf("No HTML generated for markdown: %s", md)
		}
	}
}

// Test extension registration
func TestExtensionRegistration(t *testing.T) {
	ext := &GhAlerts{
		Icons: map[string]string{"test": "icon"},
	}

	md := goldmark.New(goldmark.WithExtensions(ext))

	// Test that the extension was registered by converting a simple alert
	var buf strings.Builder
	err := md.Convert([]byte(`> [!test] Extension test`), &buf)
	if err != nil {
		t.Errorf("Extension registration failed: %v", err)
	}

	if buf.Len() == 0 {
		t.Error("Extension registration failed: no output")
	}
}

// Benchmark tests for performance
func BenchmarkSimpleAlert(b *testing.B) {
	md := `> [!info] Simple alert
> Content here`

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := markdown.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}

func BenchmarkComplexAlert(b *testing.B) {
	md := `> [!warning] Complex alert with lots of content
> This is a paragraph with *emphasis* and **strong** text.
>
> - List item 1
> - List item 2
>   - Nested item
>
> ` + "```" + `
> code block
> ` + "```" + `
>
> Final paragraph`

	for i := 0; i < b.N; i++ {
		var buf strings.Builder
		err := markdown.Convert([]byte(md), &buf)
		if err != nil {
			b.Error("Failed to parse")
		}
	}
}
