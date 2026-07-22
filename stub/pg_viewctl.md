## Usage

Sources:

- [Official README](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/README.md)
- [Exported function implementation](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/src/functions.rs)
- [Official schema-evolution exercise](https://github.com/boringsql/pg_viewctl/blob/adf2c5e2c4ec23148ca16d10c4de3ebe1a0c5044/try/exercise.sql)

`pg_viewctl` is an experimental pgrx extension for analyzing column-level view dependencies and generating ordered SQL plans for schema changes. It helps review the impact of replacing views, dropping or changing table columns, and renaming view columns; it does not automatically execute the generated migration.

### Analyze Before Changing

Install version 0.1.0, inspect dependencies, and ask the extension to generate a plan:

```sql
CREATE EXTENSION pg_viewctl;

SELECT *
FROM pg_viewctl.get_column_dependencies('app', 'orders')
ORDER BY source_column, dependent_view;

SELECT *
FROM pg_viewctl.analyze_drop_column('app', 'orders', 'legacy_code');

SELECT step, operation, sql, target
FROM pg_viewctl.generate_drop_column('app', 'orders', 'legacy_code')
ORDER BY step;
```

Treat the returned rows as a draft migration. Review object ownership, security options, grants, comments, materialized-view refreshes, column references, locking, and data loss before copying SQL into a migration file.

### Generators and Dependency Objects

- `generate_replace_view` plans dependent view/materialized-view teardown, target replacement, recreation, grants, and refreshes.
- `generate_drop_column` plans dependencies around a table-column drop.
- `generate_alter_type` plans a column type change and dependent recreation.
- `generate_rename_view_column` plans a view-column rename and annotates dependent definitions that need manual edits.
- `pgvc_column_dependencies` and `get_column_dependencies` expose column-level relationships, while `pgvc_dependency_order` produces a recursive view order.

The generators return `step`, `operation`, `sql`, and `target` columns. A generated comment or manual-edit marker is a request for human editing, not executable proof that the migration is safe.

### Deprecation Workflow

The extension can record a view column as deprecated before removal:

```sql
SELECT pg_viewctl.deprecate_column(
  'app', 'v_orders', 'legacy_code',
  'Use external_code instead',
  '2026-12-01'::date
);

SELECT * FROM pg_viewctl.get_deprecated_columns('app');
SELECT pg_viewctl.check_column_deprecated('app', 'v_orders', 'legacy_code');
```

`undeprecate_column` removes the marker, and `pgvc_deprecated_with_dependents` summarizes remaining dependent views. This metadata is advisory and does not stop queries from using the column.

### Boundaries

The official README labels the project work in progress and says the interface is likely to change. The control file marks it non-relocatable and does not require superuser or preload; source features target PostgreSQL 13 through 18, with PostgreSQL 18 as the default build.

Dependency discovery is based on PostgreSQL catalogs and stored view definitions. Dynamic SQL, application code, external consumers, function-body references, and dependencies not represented in those catalogs can be missed. Test every generated plan on a restored production-like database, execute it through normal change control, and keep a rollback path.
