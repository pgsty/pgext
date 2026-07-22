## Usage

Sources:

- [Official README](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/README.md)
- [Version 1.4 SQL API](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink--1.4.sql)
- [Extension control file](https://github.com/dimv36/sphinxlink/blob/19760a52df0e2d900bee06c7a0c59896ea408bbd/sphinxlink.control)

`sphinxlink` sends SphinxSearch or ManticoreSearch SQL over their MySQL-compatible protocol and exposes the result as a PostgreSQL record set. It can retain named connections in a backend or open a connection for one query, and it exposes metadata from the last remote search.

### Connect and Query

The caller must provide a column definition list matching the remote result because `sphinx_query()` returns records of caller-defined shape.

```sql
CREATE EXTENSION sphinxlink;

SELECT sphinx_connect('search', '127.0.0.1', 9306);

SELECT *
FROM sphinx_query(
  'search',
  'SELECT id, weight() AS rank FROM docs WHERE MATCH(?)',
  'postgresql extension'
) AS result(id bigint, rank integer);

SELECT * FROM sphinx_meta('search');
SELECT sphinx_disconnect('search');
```

`sphinx_connections()` lists open named connections. The `sphinx_query_params(host, port, query)` variants connect automatically. Overloads with `match_clause` substitute the `MATCH(?)` value; other query structure remains part of the remote SQL string.

### Remote-System Boundary

Queries execute synchronously from a PostgreSQL backend and are outside PostgreSQL transaction rollback. Network stalls, remote limits, result-shape changes, protocol errors, and Manticore/Sphinx version differences can fail or occupy the database session. Set external network timeouts and statement timeouts, constrain reachable hosts, and match every returned column type on representative data.

Treat the query string as remote SQL and never concatenate untrusted identifiers, predicates, or commands. The match-clause overload only protects its intended `MATCH(?)` value and is not a general SQL parameterization API. Review default function grants, allow only trusted callers, monitor connection cleanup under pooling and cancellation, and do not assume that last-query metadata survives a different backend connection.
