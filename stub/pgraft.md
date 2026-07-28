## Usage

Sources:

- [Official upstream README](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/README.md)
- [Official extension control file (pgraft.control)](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/pgraft.control)
- [Official extension SQL (pgraft--1.0.sql)](https://github.com/pgelephant/ram/blob/be10315f3dd94f5492a26c596787b36d1410c2f6/pgraft/pgraft--1.0.sql)

`pgraft` — **pgraft** is a high-performance PostgreSQL extension that implements Raft consensus protocol for distributed PostgreSQL clusters. It enables automatic leader election, log replication, and fault tolerance across multiple PostgreSQL instances. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgraft;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pgraft_add_node(node_id integer, address text, port integer)` is an extension function and returns `boolean`.
- `pgraft_get_cluster_status()` is an extension function and returns `TABLE`.
- `pgraft_get_leader()` is an extension function and returns `bigint`.
- `pgraft_get_nodes()` is an extension function and returns `TABLE`.
- `pgraft_get_queue_status()` is an extension function and returns `TABLE`.
- `pgraft_get_term()` is an extension function and returns `integer`.
- `pgraft_get_version()` is an extension function and returns `text`.
- `pgraft_get_worker_state()` is an extension function and returns `text`.
- `pgraft_init()` is an extension function and returns `boolean`.
- `pgraft_is_leader()` is an extension function and returns `boolean`.
- `pgraft_log_append(term bigint, data text)` is an extension function and returns `boolean`.
- `pgraft_log_apply(index bigint)` is an extension function and returns `boolean`.
- `pgraft_log_commit(index bigint)` is an extension function and returns `boolean`.
- `pgraft_log_get_entry(index bigint)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
