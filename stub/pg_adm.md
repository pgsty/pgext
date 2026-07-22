## Usage

Sources:

- [Official README](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/README.md)
- [Extension control file](https://github.com/benizar/pg_adm/blob/78314723ddfeb2128847018d5198d4156b527fb4/pg_adm.control)
- [Official SQL tool collection](https://github.com/benizar/pg_adm/tree/78314723ddfeb2128847018d5198d4156b527fb4/sql)

`pg_adm` version 0.0.1 packages a broad collection of PostgreSQL administration queries as views and functions in the fixed `adm` schema. Use it for interactive inspection of size, bloat, indexes, buffers, dependencies, grants, and extension objects; review every mutating helper before allowing routine users to execute it.

### Core Workflow

The extension requires `pg_buffercache`, which PostgreSQL installs automatically when the dependency is available.

```sql
CREATE EXTENSION pg_adm;

SELECT * FROM adm.size_tables_biggest LIMIT 20;
SELECT * FROM adm.index_size_usage WHERE idx_scan = 0;
SELECT * FROM adm.table_bloat ORDER BY wastedbytes DESC LIMIT 20;
```

Planner statistics and cumulative statistics drive several reports, so their results are estimates or observations rather than proof that an object is safe to remove.

### Important Objects

- Size views include `size_databases`, `size_schemas`, `size_tables`, and `size_relations_biggest`.
- Bloat and statistics views include `table_bloat`, `btree_bloat`, `no_stats`, and `buffers_breakdown`.
- Index diagnostics include `index_summary`, `index_duplicates`, `index_size_usage`, and `unindexed_foreign_keys`.
- `dependency_tree(...)` and `dependency` inspect object dependencies.
- `extension_objects` lists objects owned by installed extensions.
- `explain_count_estimate(text)` estimates a query's row count from its plan.
- `clone_schema(...)` and `grant_on_tables(...)` change database state and privileges.

The extension is non-relocatable and always uses `adm`. Access should be granted object by object: some views expose cluster metadata, while helpers such as schema cloning and bulk grants can create objects or change privileges. Bloat, duplicate-index, and unused-index reports require administrator judgment and should not be wired directly to automated drops.
