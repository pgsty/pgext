## Usage

Sources:

- [README.md](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/README.md)
- [pgquarry.control](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/pgquarry.control)
- [sql/pgquarry--1.0.sql](https://github.com/bugraaktug/pgquarry/blob/2690fdb537bb325f84e4c2d3293738b05780e88b/sql/pgquarry--1.0.sql)

`pgquarry` 1.0 installs a SQL queue and APIs for local embedding and text generation. It requires `vector` and a separate `pgquarry_worker` OS process using local GGUF models. Installing the SQL extension alone does not run inference.

### Core workflow

```sql
CREATE EXTENSION pgquarry CASCADE;
SELECT pgquarry.embed_async('A document to embed');
SELECT id, status FROM pgquarry.jobs ORDER BY id DESC LIMIT 5;
```

Configure the worker database and model paths in its configuration file, then start the worker. No PostgreSQL shared preload or server restart is required.

### Queue, triggers and privileges

`pgquarry.watch` installs an insert/update trigger to enqueue embeddings and write them back to the source or a separate target table. `pgquarry.watch_generate` adds generation jobs; `pgquarry.generate_async` queues text directly. Synchronous procedures require the documented output argument and timeout, and still depend on a running worker. A generation model must be configured separately.

Creation requires a superuser. The script grants `PUBLIC` access to its schema, tables and functions; review that boundary before using a shared database. Watch registration requires the relevant table ownership. Deleting a source row does not cancel queued jobs: cross-table write-back can create orphan results. Worker retention eventually purges completed jobs; plan auditing and model dimensions accordingly.
