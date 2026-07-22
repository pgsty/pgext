## Usage

Sources:

- [Official extension control file](https://github.com/komamitsu/foreign_table_exposer/blob/1c57a554d1c50c872fac3b103317b11dec10e95f/foreign_table_exposer.control)
- [Official upstream documentation](https://pgxn.org/dist/foreign_table_exposer/1.0.0/README.html)
- [Official PGXN distribution page](https://pgxn.org/dist/foreign_table_exposer/)

`foreign_table_exposer` 1.0 makes catalog clients that search only for ordinary tables and views discover PostgreSQL foreign tables. It does this by rewriting qualifying `pg_catalog.pg_class` queries, not by changing foreign-table storage or transaction semantics.

### Core Workflow

Load the module at server start, restart PostgreSQL, and enable its SQL objects in each database that needs the behavior.

```conf
shared_preload_libraries = 'foreign_table_exposer'
```

```sql
CREATE EXTENSION foreign_table_exposer;
```

After activation, a catalog predicate equivalent to `relkind IN ('r', 'v')` is rewritten to include foreign tables with relkind `f`. This targets older BI and metadata tools that otherwise hide them.

### Operational Boundaries

The rewrite hook affects matching catalog queries throughout the server and can surprise clients that intentionally exclude foreign tables. It does not make FDW operations local, transactional, or inexpensive. The release dates from 2015; test it on the exact PostgreSQL major, with other hook-using extensions, and against each client query before production use.
