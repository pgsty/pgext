## Usage

Sources:

- [Upstream README](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/README.md)
- [Extension control file](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach.control)
- [SQL API](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach--1.0.sql)
- [C implementation](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/pg_mustach.c)
- [Distribution metadata](https://github.com/RekGRpth/pg_mustach/blob/fc36a2203d2007fd93a4d1d6a49e8c24f02fd8e2/META.json)

`pg_mustach` renders Mustache templates from PostgreSQL `json` values through the C `mustach` library. The short two-argument form returns rendered text:

### Render a template

```sql
CREATE EXTENSION pg_mustach;

SELECT mustach(
  '{"name":"PostgreSQL"}'::json,
  'Hello {{name}}!'
);
```

The extension also exposes `mustach_cjson()`, `mustach_jansson()`, and `mustach_json_c()` adapters. Session-local process flags are changed by the `mustach_with_*()` functions; `mustach_with_noextensions()` resets the flags, while the other switches enable individual Mustach extensions. These flag changes are not transaction-scoped.

Every renderer also has a three-argument overload that accepts a server-side filename and opens it for writing as the PostgreSQL operating-system account. The install SQL does not revoke default `EXECUTE`, so revoke or tightly grant all filename overloads before allowing non-administrators to use the extension. Template output and errors also need resource limits when JSON or templates are untrusted.

The catalog/control object version is `1.0`, while the same commit's `META.json` reports distribution version `1.0.16`. Keep those version channels distinct when packaging or upgrading.
