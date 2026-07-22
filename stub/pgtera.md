## Usage

Sources:

- [Official README](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/README.md)
- [Rust implementation](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/src/lib.rs)
- [Extension control file](https://github.com/f-prime/pgtera/blob/9ca7b37b4448c6583351777a5530f1582f99867d/pgtera.control)

pgtera renders Tera templates from PostgreSQL. It can load named templates from a server-local glob or render a template string with a JSON context, making database-driven HTML possible for systems such as PostgREST. The extension reads the database server's filesystem and renders application output inside a backend, so its functions require strict privilege and content boundaries.

### Core Workflow

Set one administrator-controlled template glob, then render a named template with an array of name/value context objects:

```sql
CREATE EXTENSION pgtera;

SELECT pgtera_set_render_path('/srv/app/templates/**/*.html');

SELECT pgtera_render(
  'index.html',
  '[{"name":"title","value":"Status"},
    {"name":"items","value":["ready","healthy"]}]'
);
```

Use `pgtera_render_str` for a template supplied directly as text. Both paths parse the same JSON context shape, and the string-render path still initializes the configured template glob.

### Important Objects

- `pgtera.render_path` stores configured template globs; the most recently inserted row is used.
- `pgtera_set_render_path` inserts a new glob.
- `pgtera_render` loads the template set and renders one named template.
- `pgtera_render_str` renders template text with the supplied context.

### Security and Operational Notes

The PostgreSQL operating-system account must be able to read every matched file. Restrict the path table and all three functions so untrusted users cannot select arbitrary server files or consume backend CPU and memory with large template sets. The path setter builds SQL text without identifier-safe parameter binding, so do not pass user-controlled strings. Tera escaping rules and the `safe` filter determine output safety; review templates before serving HTML, and prefer an application tier when rendering needs sandboxing, caching, or filesystem isolation.
