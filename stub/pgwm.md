## Usage

Sources:

- [README.md](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/README.md)
- [pgwm.control](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/pgwm.control)
- [sql/pgwm--0.1.0.sql](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/sql/pgwm--0.1.0.sql)
- [compose.yaml](https://github.com/FranckPachot/pgwm/blob/03d43e81136fceecd210823cfc171345f591275e/compose.yaml)

`pgwm` 0.1.0 is an experimental SQL workspace manager for PostgreSQL table data. The upstream test environment uses PostgreSQL 17. Use disposable evaluation data: the project explicitly excludes production use and tenant isolation.

### Core workflow

```sql
CREATE EXTENSION pgwm;
CREATE TABLE public.workspace_demo (id integer PRIMARY KEY, value text);
INSERT INTO public.workspace_demo VALUES (1, 'original');
SELECT pgwm.enable_versioning('public.workspace_demo');
SELECT pgwm.create_workspace('trial');
SELECT pgwm.goto_workspace('trial');
UPDATE public.workspace_demo SET value = 'proposal' WHERE id = 1;
SELECT pgwm.goto_workspace('LIVE');
SELECT * FROM pgwm.conflicts('trial');
```

### Versioning and maintenance

`pgwm.enable_versioning` requires a primary key, renames the physical table with an `_lt` suffix and exposes a typed view under the original name. Workspace writes become history rows. `pgwm.refresh_workspace`, `pgwm.conflicts` and `pgwm.merge_workspace` support parent/child review and merging; inspect conflicts before merging.

Workspace selection is session-local. Reset pooled connections to `LIVE`, and do not expose the underlying history tables as an isolation boundary. Key changes, destructive schema changes, user-trigger replay and several cleanup operations have limitations. Installation requires a superuser and `plpgsql`; no native library or shared preload is needed. Upstream does not declare a license.
