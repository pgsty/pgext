## Usage

Sources:

- [Official README](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/README.md)
- [Extension control file](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/ekorre.control)
- [Installation SQL](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/ekorre--1.0.0.sql)
- [Build configuration](https://github.com/danielgustafsson/ekorre/blob/96cc2748bac15a9688e53d29be880fd745a25c64/Makefile)

`ekorre` exposes a server-local Git repository as PostgreSQL foreign tables. Upstream describes it as unreleased software; its query interface is intended for repository inspection, with no qualifier pushdown guarantee.

### Core Workflow

With the extension library and libgit2 installed, create the extension as an administrator and import the repository schema. The repository path is resolved on the database host.

```sql
CREATE EXTENSION ekorre;
CREATE SERVER git_server FOREIGN DATA WRAPPER ekorre;
IMPORT FOREIGN SCHEMA git FROM SERVER git_server INTO public
  OPTIONS (repopath '/srv/git/project');
```

### Objects and Limits

`ekorre_handler()` and `ekorre_validator(text[], oid)` implement the wrapper. Import uses `repopath` to select the Git repository and creates the tables described by the implementation. Control version is `1.0.0`; the extension is relocatable. No preload or supported PostgreSQL-major matrix is documented. Grant server and imported-table access only to roles allowed to read that repository; filtering query output is not a substitute for host-file access control.
