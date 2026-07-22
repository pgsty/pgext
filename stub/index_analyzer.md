## Usage

Sources:

- [Official README](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/README.md)
- [Official extension control file](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/index_analyzer.control)
- [Official extension SQL](https://github.com/tvondra/index_analyzer/blob/ee1f224c18fac5b9afb6d5cc393ebdbc80ad719e/sql/index_analyzer--0.0.1.sql)

`index_analyzer` is a set of PL/pgSQL diagnostics for examining existing indexes and detecting foreign keys without a matching index on the referencing table. Its findings are leads for investigation, not automatic drop or create decisions.

### Core Workflow

Run the broad checks for a schema, then inspect individual tables or indexes with a selectivity threshold expressed as a fraction.

```sql
CREATE EXTENSION index_analyzer;

SELECT analyze_tables('public', 0.05);
SELECT analyze_table('public.orders'::regclass::oid, 0.05);
SELECT analyze_index('public.orders_created_at_idx'::regclass::oid, 0.05);

SELECT analyze_fks('public');
SELECT analyze_fks('public.orders'::regclass::oid);
```

The functions report findings with notices rather than returning a recommendations table. Capture client messages when running them in automation.

### Function Index

- `analyze_tables`, `analyze_table`, and `analyze_index` coordinate checks over user tables and indexes.
- `analyze_index_selectivity` estimates distinct-value selectivity from catalog statistics.
- `analyze_index_usage` compares index and sequential-scan statistics and skips unique indexes, including primary keys.
- `analyze_index_count_distinct` samples table data for a distinct-count estimate.
- `analyze_fks` and `analyze_fk` inspect foreign-key indexing by schema, table OID, constraint name, or constraint OID.

### Interpretation and Caveats

The upstream explicitly warns that answers are not definitive. Usage counters depend on workload history, catalog estimates can overstate selectivity for correlated columns, and the sampling check can be expensive and is weak for expression indexes. Only user tables are considered.

The repository is archived. Its control file declares version `0.1.0`, while the checked-in install script is named `index_analyzer--0.0.1.sql`; verify that the packaged artifact resolves this upstream mismatch before installation. Review every suggested index against constraints, rare critical queries, and a representative observation window.
