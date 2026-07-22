## Usage

Sources:

- [Official README](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/README.md)
- [Official version 0.0.1 SQL](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views--0.0.1.sql)
- [Official control file](https://github.com/Sednai/xl_global_views/blob/804ab3fbd91ca84441af86d520717ce22583cd98/xl_global_views.control)
- [Required Postgres-XL cursor patch](https://git.postgresql.org/gitweb/?p=postgres-xl.git;a=commit;h=1b6ada225da45c82529d56f71e3c6a62fabcfd55)

`xl_global_views` creates cluster-wide `pgxl_` monitoring views for Postgres-XL by running each node's local PostgreSQL system view through `EXECUTE DIRECT`. Every generated row is prefixed with `node_name` and `node_type`, allowing coordinators and data nodes to be compared in one query. It is specific to Postgres-XL, not ordinary PostgreSQL.

### Core Workflow

Install version `0.0.1` on every required Postgres-XL node, create the non-relocatable extension in `public`, then explicitly generate the views:

```sql
CREATE EXTENSION xl_global_views;
SELECT public.pgxl_create_views();

SELECT node_name, node_type, relname, heap_blks_read, heap_blks_hit
FROM public.pgxl_statio_all_tables
WHERE node_type = 'D' AND relname = 'pg_class';
```

Rerun `pgxl_create_views()` after an underlying system-view shape changes; it drops and recreates matching `pgxl_` views from `information_schema.columns`.

### Objects

- `pgxl_global_view(localtable text, fields text [, cond text, limitGiven int])`: executes a selected local system view on each coordinator/data node and returns prefixed records.
- `pgxl_create_views()`: builds public views for matching `pg_stat*`/`pg_statio*` sources plus selected catalog views.
- Documented outputs include `pgxl_locks`, `pgxl_stat_activity`, `pgxl_stat_database`, `pgxl_stat_replication`, `pgxl_stat_ssl`, `pgxl_stats`, and the table/index/function I/O and transaction-statistics families.

### Caveats

The Postgres-XL build must contain upstream cursor patch `1b6ada225da45c82529d56f71e3c6a62fabcfd55`; without it, most views fail with SPI errors. In the targeted Postgres-XL implementation, `EXECUTE DIRECT` requires superuser privileges, so merely granting `SELECT` on generated views does not make them safely usable by ordinary roles.

The generic function constructs dynamic SQL from text and defaults to a high per-node limit. Expose it only to trusted administrators. Columns of `anyarray` and `pg_node_tree` types are omitted, so a global view may not exactly match the corresponding local system view. The project targets an old Postgres-XL fork and should not be installed or tested as if it were a current PostgreSQL monitoring extension.
