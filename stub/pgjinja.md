## Usage

Sources:

- [Upstream source 1](https://github.com/thejens/pgjinja/blob/8298a381f88437a70ae9d22d6b35fb6db64fd73c/README.md)
- [Upstream source 2](https://github.com/thejens/pgjinja/blob/8298a381f88437a70ae9d22d6b35fb6db64fd73c/pgjinja.control)

`pgjinja` 0.1.0 renders MiniJinja templates directly from JSONB data in PostgreSQL.

### Core Workflow

`jinja_render` renders inline templates. `jinja_render_named` reads `jinja_templates`, including inheritance and includes; `jinja_syntax_error` returns parse diagnostics. Stored-template changes invalidate session caches through PostgreSQL's transactional invalidation mechanism.

```sql
CREATE EXTENSION pgjinja;
SELECT jinja_render('Hello {{ name }}', '{"name":"Ada"}'::jsonb, strict => true);
SELECT jinja_is_valid('{{ name }}');
```

### Operational Boundaries

The README supports PostgreSQL 14–18 on Linux/macOS; installation is superuser-only. Inline escaping defaults to off, so enable autoescaping for untrusted HTML. Keep `pgjinja.fuel` bounded: statement timeouts cannot interrupt rendering until control returns. JSON numbers become 64-bit integers/floats; pass exact large numbers as text. Templates do not receive filesystem or network access.
