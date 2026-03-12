

## Usage

> [l10n_table_dependent_extension: l10n table dependent extension for pg_xenophile](https://github.com/bigsmoke/pg_xenophile)

The `l10n_table_dependent_extension` is a companion extension to `pg_xenophile` that provides infrastructure for other extensions that depend on its localization (l10n) table mechanism.

```sql
CREATE EXTENSION l10n_table_dependent_extension;
```

This extension works in conjunction with `pg_xenophile`'s `l10n_table` system, allowing dependent extensions to register their tables with the localization framework. When a base table is registered in `xeno.l10n_table`, the system automatically creates companion translation tables and language-specific views.

See the [`pg_xenophile`](https://github.com/bigsmoke/pg_xenophile) documentation for the full l10n table management API.
