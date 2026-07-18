## Usage

Sources:

- [pg_git README at the reviewed commit](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/README.md)
- [pg_git.control at the reviewed commit](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/pg_git.control)
- [Repository initialization SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/001-init.sql)
- [Staging SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/002-add.sql)
- [Commit SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/003-commit.sql)
- [Remote-operation SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/011-remote.sql)
- [HTTPS and credential SQL](https://github.com/seanwevans/pg_git/blob/ca7b3c48506e53ccf55b934a98d76bfa559555b1/sql/functions/014-https.sql)

`pg_git` stores a Git-like object model entirely in PostgreSQL. Objects live in the `pggit` schema, with implemented functions for repository initialization, staging byte content, commits, history, branches, tags, diffs, fast-forward merges, garbage collection, and integrity checks.

Core entry points include `init_repository`, `stage_file`, `commit_index`, and `get_log`. The extension requires `plpgsql`, `pgcrypto`, `pg_trgm`, and `plpython3u` at install time.

### Basic Local Workflow

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

### Caveats

- Remote transport is incomplete. The current `push` function only emits a notice, and `fetch_remote` materializes rows already present in `remote_refs`; do not treat it as a drop-in Git network client.
- HTTPS helpers execute outbound requests through untrusted `plpython3u`. Review database egress and Python privileges before enabling them.
- Configure a strong `pggit.credential_key` before storing credentials. Current source falls back to `pg_git_default_key` when the setting is absent.
- Catalog, control, and README identify version `0.4.0`. The README clearly distinguishes implemented operations from aspirational ones; verify the exact function before depending on a listed feature.
