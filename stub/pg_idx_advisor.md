## Usage

Sources:

- [Official README and deprecation notice](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/README.md)
- [Extension SQL](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/sql/pg_idx_advisor.sql)
- [Planner-hook implementation](https://github.com/cohenjo/pg_idx_advisor/blob/5c7a51797b65b0fd8b8d4b26e2fbc28e6a8ca780/src/idx_adviser.c)

pg_idx_advisor is a deprecated planner-hook extension that compares an EXPLAIN plan with plans using hypothetical candidate indexes and emits CREATE INDEX recommendations. Upstream says it is no longer updated and directs users to HypoPG. Keep it only for reproducing historical behavior.

### Core Workflow

Create the catalog table, then load the library in the session before running EXPLAIN:

```sql
CREATE EXTENSION pg_idx_advisor;
LOAD 'pg_idx_advisor';

EXPLAIN
SELECT * FROM orders WHERE customer_id = 42;

SELECT recommendation, benefit, index_size
FROM index_advisory
WHERE backend_pid = pg_backend_pid()
ORDER BY benefit DESC;
```

The extension prints original and hypothetical plans and, unless read-only mode is enabled, stores recommendations in `index_advisory`.

### Important Objects

- `index_advisory` stores relation OID, indexed attributes, estimated benefit and size, expression or predicate, source query, and recommendation text.
- `index_adviser.cols` lists columns considered for partial indexes.
- `index_adviser.schema` selects the schema containing the advisory table.
- `index_adviser.read_only` disables table inserts and prints advice only.
- `index_adviser.text_pattern_ops` permits text-pattern operator classes.
- `index_adviser.composit_max_cols` limits columns in a composite candidate.

### Operational Notes

Loading the library replaces planner and EXPLAIN hooks for that backend and uses old PostgreSQL virtual-index internals. Recommendations are cost-model hypotheses, not proof that an index helps under real data, writes, cache state, or concurrency. Review locking, build time, storage, write amplification, duplicate indexes, and partial predicates before applying any emitted SQL. Prefer a maintained HypoPG-based workflow on supported PostgreSQL versions.
