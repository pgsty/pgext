## Usage

Sources:

- [README](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/README.md)
- [Control file](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/pg_prosupport.control)
- [pg_prosupport--1.0.sql](https://github.com/danolivo/pg_prosupport/blob/fb410817eb4c6315572e873f6950ed62570661fd/pg_prosupport--1.0.sql)

`pg_prosupport` rewrites eligible numeric aggregates using planner support functions. The current 1.0 source works on stock PostgreSQL 19; PostgreSQL 18 requires the project’s additional kernel patch and explicit library loading.

### PostgreSQL 19 Workflow

```sql
CREATE EXTENSION pg_prosupport;
CREATE TABLE amounts (amount numeric(12,2));
INSERT INTO amounts VALUES (12.50), (8.25);
EXPLAIN (VERBOSE, COSTS OFF) SELECT sum(amount), avg(amount) FROM amounts;
```

### Rewrites and Settings

Installation attaches support functions to the built-in numeric aggregates and creates objects in the fixed `prosupport` schema. Eligible bounded numeric expressions can use specialized sum/average states; unsupported precision or expressions remain on the ordinary path. This is not a blanket speed guarantee.

`pg_prosupport.bounded_numeric_agg` enables bounded rewrites by default. `pg_prosupport.fold_const_sum` controls constant-sum rewriting and is off by default. PostgreSQL 19 loads the library on demand. The patched PostgreSQL 18 path instead requires `LOAD` or preloading; changing startup preload configuration requires a restart.

### Maintenance Boundary

Superuser privileges are required for installation and support attachment. After a restore, reattach with `prosupport.pps_attach_support()` as directed upstream, and reconnect or discard cached plans when changing planner behavior. Before dropping the extension on PostgreSQL 19, detach support first to avoid dangling catalog references:

```sql
SELECT prosupport.pps_detach_support();
DROP EXTENSION pg_prosupport;
```

The project supersedes the earlier repository name pg_numeric_agg_support. Treat this as a planner experiment and test against the precise server build.
