## Usage

Sources:

- [Official README](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/README.md)
- [Extension control file](https://github.com/glynastill/pg_jsonb_opx/blob/adb985a8c3192e2577bf7ef863560baf4e566c5e/jsonb_opx.control)

`jsonb_opx` supplies hstore-style mutation operators that were missing from the original PostgreSQL 9.4 `jsonb` implementation. Version `1.0` provides shallow deletion, concatenation, and replacement; use it mainly when maintaining an application built against this historical operator behavior.

### Core Workflow

```sql
CREATE EXTENSION jsonb_opx;

SELECT '{"a":1,"b":2}'::jsonb - 'a';
SELECT '{"a":1}'::jsonb || '{"b":2}'::jsonb;
SELECT '{"a":1,"b":2}'::jsonb #= '{"a":9}'::jsonb;
```

The `-` overloads remove matching scalar values or arrays of values without traversing nested objects. `||` concatenates two `jsonb` values, and `#=` replaces top-level values using another `jsonb` document.

### User-Facing Objects

- `jsonb_delete(jsonb, text|numeric|boolean)`: shallow deletion by scalar value.
- `jsonb_delete(jsonb, text[]|numeric[]|boolean[])`: shallow deletion by an array of values.
- `jsonb_delete(jsonb, jsonb)`: shallow deletion driven by another JSON value.
- `jsonb_concat(jsonb, jsonb)`: implementation behind `||`.
- `jsonb_replace(jsonb, jsonb)`: implementation behind `#=`.
- `jsonb_delete_path(jsonb, text[])` and `jsonb_replace_path(jsonb, text[], jsonb)`: nested-path functions described for upstream version `1.1`.

### Version and Compatibility Notes

The control file defaults to `1.0`. The README also describes an experimental `1.1` API with `#-` and nested replacement, but that version is not implied by a `1.0` package; check `pg_available_extension_versions()` before using it. Modern PostgreSQL releases already provide several similarly spelled built-in `jsonb` operators, so test resolution and semantics on the exact server version. Upstream explicitly labels the project experimental and asks users to validate it before production use. No preload or fixed schema is declared.
