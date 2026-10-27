## Usage

Sources:

- [README](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/README.md)
- [Control](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/pg_lock_tree.control)
- [Install and grants](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/pg_lock_tree--1.5.1.sql)

`pg_lock_tree` 1.5.1 is a pure SQL extension for readable lock-wait trees and root blockers. Upstream tests PostgreSQL 14–19rc1. A superuser installs it; no library, preload or restart is required. This source revision is newer than the repository's latest tagged release.

### Inspect a Queue

```sql
CREATE EXTENSION pg_lock_tree;
SELECT * FROM lock_tree_text();
SELECT * FROM lock_tree_roots();
SELECT command FROM lock_tree_command();
```

The command helper returns SQL for review and executes none of it. Do not feed its complete output into an automatic execution command. Prepared transactions can block sessions without a live backend to terminate. The wide and summary readers help with large queues; sampled history and NOTIFY alerts are optional.

### Grants and Maintenance

Monitoring requires both the extension's object/function grants and pg_read_all_stats; pg_monitor alone is insufficient. History readers need their additional grants. Schedule sampling, alerts, retention and archive maintenance outside PostgreSQL's extension; there is no built-in worker. Query text in history may contain sensitive literals. Review the supplied upgrade SQL and installation grant matrix before replacing an existing installation.
