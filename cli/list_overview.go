package cli

// ext/_index.md must not declare a print output. When Hugo renders the print
// variant for the full extension catalog, the generated `_print/ext/index.html`
// can exceed Cloudflare Pages' 25 MiB asset limit.
const extOverviewOutputsFrontMatter = `outputs:
  - HTML
cascade:
  type: docs
  outputs:
    - HTML`
