## Usage

Sources:

- [Official extension control file](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/explanation.control)
- [Official PGXN distribution page](https://pgxn.org/dist/explanation/)
- [Official upstream README](https://github.com/theory/explanation/blob/abae6466a140981565479b6160fe6122eb344248/README.md)

`explanation` 0.3.0 parses PostgreSQL `EXPLAIN` output into relational rows: one row per plan node, with child nodes referring to their parent identifiers. It is useful for SQL-side plan analysis when a proximity-tree representation is easier to query than formatted plan text.

### Core Workflow

```sql
CREATE EXTENSION explanation;

SELECT *
FROM explanation('SELECT * FROM orders WHERE customer_id = 42');
```

Pass a single query string to `explanation(text)`. The function runs plain `EXPLAIN`, parses its XML representation, and returns the plan nodes as a table for filtering, joins, or visualization.

### Requirements and Caveats

The upstream release requires PostgreSQL 9.0 or later built with XML support, and its code predates modern JSON-format plan tooling. It does not run `EXPLAIN ANALYZE`, so the result contains planner estimates rather than measured execution. Accept query text only from trusted callers, and test the parser against the exact PostgreSQL major version because plan fields evolve between releases.
