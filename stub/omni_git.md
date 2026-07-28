## Usage

Sources:

- [Official upstream README](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/README.md)
- [Official extension control file (omni_git.control)](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/omni_git.control)
- [Official implementation source](https://github.com/andrew/omni_git/blob/9bc1d4e34f11d49a10071e4c7bafe15ef5ac9fcc/omni_git.c)

`omni_git` — A PostgreSQL extension that stores git repositories in database tables and serves the git smart HTTP protocol, turning Postgres into a git remote. Use it when porting or emulating the corresponding database API. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION omni_git;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- Install the confirmed extension dependencies first: `gitgres`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
