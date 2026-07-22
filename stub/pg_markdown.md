## Usage

Sources:

- [Official README](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/README.md)
- [Official extension control file](https://github.com/sycobuny/pg_markdown/blob/af587ec2e20f87642b6d1e96536036b39c0030da/pg_markdown.control)

`pg_markdown` adds a `markdown` data type and converts stored Markdown to HTML or plain text inside PostgreSQL. It can support small database-centered publishing workflows, but the reviewed code and documentation date from the project's early 2011-era implementation and should be compatibility- and security-tested before use.

### Core Workflow

Store Markdown as the extension type and render it only at the presentation boundary:

```sql
CREATE EXTENSION pg_markdown;

CREATE TABLE blog_posts (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  body markdown NOT NULL
);

INSERT INTO blog_posts(body)
VALUES ('## Welcome

This is **Markdown** with a [link](https://www.postgresql.org/).');

SELECT to_html(body) FROM blog_posts WHERE id = 1;
SELECT to_text(body) FROM blog_posts WHERE id = 1;
```

### SQL Surface

- `markdown` is the extension's stored data type.
- `to_html(markdown)` returns rendered HTML as `text`.
- `to_text(markdown)` returns a plain-text representation as `text`.

### Safety and Compatibility

- The upstream README promises Markdown conversion, not HTML sanitization. Treat `to_html(markdown)` output as untrusted HTML and sanitize it for the target rendering context, especially when authors can submit raw HTML, links, or images.
- Pin and test the accepted Markdown dialect. Modern Markdown variants, Unicode edge cases, malformed input, and renderer differences are not documented by this old release.
- Rendering during large scans consumes database CPU. Consider storing source Markdown as the system of record and caching sanitized output outside hot query paths.
- Before an upgrade, test type binary/text I/O, dump/restore, casts, and stored values on the exact PostgreSQL version; the official README does not claim support for current PostgreSQL releases.
