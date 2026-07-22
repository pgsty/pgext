## Usage

Sources:

- [Official extension control file](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/jsonpath.control)
- [Official upstream README](https://github.com/choplin/pg-jsonpath/blob/968eff51fcb958b2641fd81cb5739b3c033a7979/README.md)

`jsonpath` 0.1 is a PL/v8 wrapper around Stefan Goessner's historical JavaScript JSONPath implementation. It predates PostgreSQL's built-in SQL/JSON path support and implements that JavaScript library's syntax, not the SQL-standard `jsonpath` type.

### Core Workflow

Install `plv8` first, then create the extension and call `jsonPath(json, text)`.

```sql
CREATE EXTENSION plv8;
CREATE EXTENSION jsonpath;

SELECT jsonPath('{"x":{"a":1,"b":2}}'::json, '$.x.[a,b]');
```

The function returns a PostgreSQL `json[]` containing every match. Use expressions accepted by the bundled JavaScript implementation; they are not interchangeable with operators and functions built around PostgreSQL's native `jsonpath` type.

### Caveats

The project is small and old, and execution occurs inside PL/v8. Prefer native SQL/JSON path features on supported PostgreSQL releases unless compatibility with this exact API is required. Validate expression semantics, bound input size, and restrict execution where untrusted path expressions or documents could consume excessive backend CPU or memory.
