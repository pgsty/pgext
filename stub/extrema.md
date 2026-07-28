## Usage

Sources:

- [Official upstream README](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/README.md)
- [Official extension control file (extrema.control)](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/extrema.control)
- [Official extension SQL (extrema--1.0.sql)](https://github.com/ggnmstr/extrema/blob/75fd0c4e50957cc59428e2dbb6933c91fbb82e6c/extrema--1.0.sql)

`extrema` — This extension allows user to limit other extensions usage of resources (currently, only CPU, RAM, VmSwap and cpuset) by adding them to corresponding cgroups. These limitations can be easily configured using PostgreSQL's GUC mechanism. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION extrema;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ema_lib_info()` is an extension function and returns `SETOF`.
- `ema_reload()` is an extension function and returns `VOID`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
