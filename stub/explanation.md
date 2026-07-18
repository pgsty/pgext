## Usage

Sources:

- [Official extension control file](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/explanation.control)
- [Official PGXN distribution page](https://pgxn.org/dist/explanation/)
- [Official upstream README](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/README.md)

`explanation` — EXPLAIN plan parser that returns plan nodes as relational rows organized as a proximity tree.

The reviewed catalog snapshot records version `0.3.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.
The curated compatibility set is `10,11,12,13,14`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "explanation";
SELECT extversion
FROM pg_extension
WHERE extname = 'explanation';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
