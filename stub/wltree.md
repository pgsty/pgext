## Usage

Sources:

- [Upstream README at the reviewed commit](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/README.md)
- [Version 2.0 SQL objects](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/wltree--2.0.sql)
- [Extension control file](https://github.com/adjust/wltree/blob/f8fb773edfc793c6bd323beaa6cb8b1b98ce366d/wltree.control)

`wltree` is a patched copy of PostgreSQL's `ltree` module. It uses `::` instead of `.` as the label separator and permits special characters such as braces, exclamation marks, asterisks, and percent signs when escaped in queries.

```sql
CREATE EXTENSION wltree;
SELECT '!foo::{bar}::baz%'::ltree
       ~ '\!foo::\{bar\}::baz\%'::lquery;
```

The extension intentionally creates the familiar `ltree`, `lquery`, and related object names with non-core syntax and semantics. Do not install it alongside the core `ltree` extension in the same database without a complete object-conflict review, and do not assume values can be exchanged between them.

Upstream states that escaped characters are unsupported in full-text `ltxtquery` queries. Pin the exact implementation for every node and client, test parsing, escaping, operators, GiST indexes, dump/restore, logical replication, and upgrades, and rebuild dependent indexes after semantic changes. Validate every externally supplied path or query independently; escaping rules are not a substitute for parameterized SQL.
