## Usage

Sources:

- [pg_git README at the reviewed commit](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/README.md)
- [Repository initialization SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/001-init.sql)
- [Staging SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/002-add.sql)
- [Commit SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/003-commit.sql)
- [Remote-operation SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/011-remote.sql)
- [HTTPS and credential SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/014-https.sql)

`pg_git` implements a Git-like object model in PostgreSQL tables. Version `0.4.0` provides local repository initialization, staging, commits, logs, branches, tags, diffs, fast-forward merge primitives, resets, garbage collection, and integrity checks in the fixed `pggit` schema. It is useful for experimenting with transactional, queryable version graphs; its remote transport is incomplete and it is not a drop-in replacement for the Git command-line client.

The extension requires `plpgsql`, `pgcrypto`, `pg_trgm`, and untrusted `plpython3u`. Installing it normally requires a superuser and the matching server-side Python language package.

### Local Workflow

```sql
CREATE EXTENSION pg_git CASCADE;

SELECT pggit.init_repository('demo', '/logical/demo');

SELECT pggit.stage_file(
  (SELECT id FROM pggit.repositories WHERE name = 'demo'),
  'README.md',
  convert_to('hello from PostgreSQL', 'UTF8')
);

SELECT pggit.commit_index(
  (SELECT id FROM pggit.repositories WHERE name = 'demo'),
  'alice',
  'initial content'
);

SELECT * FROM pggit.get_log(
  (SELECT id FROM pggit.repositories WHERE name = 'demo')
);
```

Repository content is byte-oriented. `stage_file` accepts a path and `bytea`; `commit_index` writes the staged tree and advances repository history. Related implemented entry points include `unstage_file`, `get_status`, `get_decorated_log`, `create_branch`, `checkout_branch`, `merge_branches`, `diff_blobs`, `diff_commits`, `reset_soft`, `reset_mixed`, `create_tag`, `gc`, `verify_integrity`, and `optimize_indexes`.

### Remote and Credential Boundary

The schema includes `add_remote`, `fetch_remote`, `push`, `pull`, `clone`, `store_credentials`, and `http_fetch`, but the names overstate current interoperability. At the reviewed commit, `push` emits a notice rather than uploading objects, and `fetch_remote` works from reference rows already present in the database. Verify each remote function's SQL before treating it as network Git behavior.

HTTPS helpers execute outbound requests in `plpython3u`. Credential storage uses a configurable `pggit.credential_key` and falls back to `pg_git_default_key` when absent; set a strong deployment-specific key before storing any secret and protect the setting itself.

### Operational Notes

- Untrusted Python and outbound HTTPS enlarge the database security boundary. Restrict function execution, database egress, extension ownership, and who can write remote URLs.
- Local table updates participate in PostgreSQL transactions, but external HTTP effects do not roll back with them. Design retries and reconciliation explicitly.
- Object content and history consume database storage and WAL. Measure large repositories, backup growth, replication load, garbage collection, and query plans before production use.
- The README mixes implemented functions with aspirational features. Depend only on callable SQL verified at the exact `0.4.0` commit and exercise recovery and integrity checks on representative repositories.
