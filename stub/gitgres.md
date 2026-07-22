## Usage

Sources:

- [Official README](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/README.md)
- [Official extension SQL](https://github.com/andrew/gitgres/blob/eaf8743f2c137d61a75432e44e13467cad7eceaa/ext/sql/gitgres--0.1.sql)
- [Official libgit2 backend source](https://github.com/andrew/gitgres/tree/eaf8743f2c137d61a75432e44e13467cad7eceaa/backend)

`gitgres` stores Git objects, repositories, refs, and reflog entries in PostgreSQL and provides SQL parsers plus a separate libgit2/libpq backend for push and clone operations. Version `0.1` is a young implementation centered on 20-byte SHA-1 Git object IDs; evaluate it with real repository sizes and recovery requirements before replacing filesystem Git storage.

### Core Workflow

Install the extension and its `pgcrypto` dependency, build the companion backend, and initialize a named repository:

```sql
CREATE EXTENSION gitgres CASCADE;
```

```bash
./backend/gitgres-backend init "dbname=gitgres" myrepo
./backend/gitgres-backend push "dbname=gitgres" myrepo /path/to/repo
./backend/gitgres-backend clone "dbname=gitgres" myrepo /path/to/dest
./backend/gitgres-backend ls-refs "dbname=gitgres" myrepo
```

After importing or pushing objects, refresh the query-oriented materialized views before reading them:

```sql
REFRESH MATERIALIZED VIEW commits_view;
REFRESH MATERIALIZED VIEW tree_entries_view;

SELECT sha, author_name, authored_at, message
FROM commits_view
ORDER BY authored_at DESC;
```

### Important Objects

- `repositories`, `objects`, `refs`, and `reflog` are the primary storage tables.
- `git_oid` is a fixed 20-byte type with comparison, B-tree, and hash support.
- `git_object_write(...)`, `git_object_read(...)`, and `git_object_read_prefix(...)` manage raw Git objects.
- `git_tree_entries(bytea)`, `git_ls_tree_r(...)`, and `git_commit_parse(bytea)` expose tree and commit structure.
- `git_ref_update(...)` performs compare-and-swap ref updates; `git_ref_set_symbolic(...)` manages symbolic refs.
- `commits_view` and `tree_entries_view` are materialized reporting views and are not refreshed automatically.

### Operational Caveats

The extension requires `pgcrypto`; the push/clone command also requires libgit2, libpq, OpenSSL, filesystem access to the source or destination repository, and a database role with appropriate table/function privileges. Protect connection strings and restrict direct writes to the storage tables so ref invariants are not bypassed. Git object data can be large and write-heavy, so test WAL volume, vacuum, backups, replicas, materialized-view refresh cost, and restore round trips. The `git_oid` format supports SHA-1 repositories only; do not infer support for Git's SHA-256 object format. A database backup is not proven Git recovery until a cloned repository passes Git integrity checks.
