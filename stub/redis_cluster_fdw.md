## Usage

Sources:

- [Official README](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/README.md)
- [Extension control file](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw.control)
- [Installation SQL](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw--1.0.sql)
- [Implementation (redis_cluster_fdw.c)](https://github.com/jeffreydwalter/redis_cluster_fdw/blob/6ef969825958ce4aef34417d79cddba0052b8fbf/redis_cluster_fdw.c)

`redis_cluster_fdw` is a distinct Redis-cluster variant of the Redis FDW. Its versioned SQL and C validator define the canonical wrapper and options; some README examples still use the older non-cluster name.

### Core Workflow

Install the extension and hiredis-cluster library, then configure cluster nodes and a per-user password. No database-number option is accepted by this cluster validator.

```sql
CREATE EXTENSION redis_cluster_fdw;
CREATE SERVER redis_cluster FOREIGN DATA WRAPPER redis_cluster_fdw
  OPTIONS (nodes '127.0.0.1:6379');
CREATE USER MAPPING FOR CURRENT_USER SERVER redis_cluster
  OPTIONS (password 'replace-with-cluster-password');
CREATE FOREIGN TABLE redis_values (key text, value text)
  SERVER redis_cluster;
SELECT * FROM redis_values;
```

### Options and Caveats

`nodes` is a comma-separated server list; `password` belongs to user mappings. Table options include `tabletype`, `tablekeyprefix`, `tablekeyset` and `singleton_key`. Singleton list/set shapes and sorted-set scores need the documented column layouts. Reads and writes act on Redis; do not assume PostgreSQL rollback reverses external writes. There is no schema import, truncate API or general expression pushdown. Upstream states PostgreSQL 13+ but supplies no full tested-major matrix; use a compatible build and UTF-8 database encoding.
