## Usage

Sources:

- [Official README](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/README.md)
- [Extension SQL](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/sql/jansson-json.sql)
- [Regression examples](https://github.com/claesjac/pg-json/blob/551e0067fe7731778362aa3eb265b810002ba9de/test/sql/json.sql)

`jansson-json` 0.0.2 is a historical, Jansson-backed JSON type and operator extension written before PostgreSQL acquired its built-in `json` and `jsonb` types. Its SQL attempts to create a type literally named `json`, so it conflicts with modern PostgreSQL and should be studied as legacy code rather than installed as a current JSON solution.

### Historical API

On a compatible pre-built-in-JSON server, the intended workflow was:

```sql
CREATE EXTENSION "jansson-json";

SELECT '{"foo":{"answer":42}}'::json ~ 'foo.answer';
SELECT json_set_value('{"a":1}'::json, 'a', '2'::json);
SELECT '{"a":1}'::json || '{"b":2}'::json;
```

`json_get_value(json, text)` and the `~` operator extract a path as text. Paths can traverse object keys and arrays, for example `foo.quax`, `foo[0].quax`, and `[0]`. `json_set_value(json, text, json)` replaces a value. `json_concat` and `||` combine objects or arrays, while `json_equals`, `json_not_equals`, `=`, and `!=` compare parsed JSON values rather than raw textual formatting.

### Compatibility Boundary

The extension requires the native Jansson library and exposes its own C I/O functions for the custom `json` type. PostgreSQL has shipped a built-in type with that same name since 9.2, making the extension's installation SQL collide on supported modern releases. Its API also lacks the indexing, containment, path, validation, and maintenance behavior expected from modern `jsonb`.

Use PostgreSQL's built-in `jsonb`, extraction operators, subscripting, SQL/JSON path functions, and GIN indexes for new applications. Do not rename the legacy type or SQL objects casually: its compiled function signatures and operator definitions were designed around the original type identity and old server internals.
