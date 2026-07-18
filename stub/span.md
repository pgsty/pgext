## Usage

Sources:

- [Official extension control file](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/span.control)
- [Official upstream documentation](https://github.com/arunchaganty/pg-span/blob/0ff0b6f1070cf6410299701b7fcfe6368513eaa2/README.md)

`span` — C data type for indexed document text spans, with casts, comparisons, B-tree/hash operator classes, aggregates, and accessors.

The reviewed catalog snapshot records version `0.1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "span";
SELECT extversion
FROM pg_extension
WHERE extname = 'span';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
