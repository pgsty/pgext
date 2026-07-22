## Usage

Sources:

- [Official version 1.1.0 README](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/README.md)
- [Official control file](https://github.com/franckverrot/git_fdw/blob/7f0426d392854909677dd7f2367ff3b09e3a3a99/git_fdw.control)
- [Official repository](https://github.com/franckverrot/git_fdw)
- [Official PGXN distribution](https://pgxn.org/dist/git_fdw/)

`git_fdw` exposes commits from a local Git repository as a read-only PostgreSQL foreign table. It uses libgit2 and returns commit identity/message/timestamp plus insertion, deletion, and changed-file counts for a selected branch.

### Core Workflow

Create the wrapper and an option-free server, then import a repository schema on PostgreSQL 9.5 or later:

```sql
CREATE EXTENSION git_fdw;
CREATE SERVER git_repo FOREIGN DATA WRAPPER git_fdw;

IMPORT FOREIGN SCHEMA git_data
FROM SERVER git_repo
INTO public
OPTIONS (
  path '/srv/git/project.git',
  branch 'refs/heads/main',
  prefix 'project_'
);

SELECT sha1, name, email, commit_date, message
FROM project_repository
ORDER BY commit_date DESC
LIMIT 10;
```

On PostgreSQL 9.4, define the documented foreign-table columns manually. One foreign table addresses one repository/branch; combine multiple tables with views.

### Options and Columns

- `path` (required): repository path on the PostgreSQL server's filesystem.
- `branch` (required): full Git reference to scan.
- `git_search_path` (optional): helps libgit2 locate its configuration.
- `prefix` applies when importing the generated table name.
- Commit columns include `sha1`, `message`, author `name`/`email`, `commit_date`, `insertions`, `deletions`, and `files_changed`.

### Caveats

The PostgreSQL server account needs filesystem access to the repository; the FDW does not clone a remote URL or use a user mapping for credentials. Treat repository contents and commit metadata as server-side data, and expose the foreign table only to roles allowed to read them.

`IMPORT FOREIGN SCHEMA` does not support `LIMIT TO` or `EXCEPT`. Version `1.1.0` targets libgit2 through 0.27 and old PostgreSQL versions, so current libgit2/PostgreSQL compatibility must be validated from the actual package. Queries walk repository history and compute diff statistics; use limits and measure cost on large histories.
