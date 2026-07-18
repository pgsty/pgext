## Usage

Sources:

- [postgrest_auth README at the reviewed commit](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/README.md)
- [postgrest_auth.control at the reviewed commit](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth.control)
- [Version 0.1 install SQL](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth--0.1.sql)
- [Upstream Makefile](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/Makefile)

`postgrest_auth` is an unfinished extension skeleton, not a usable PostgREST authentication implementation. The reviewed version `0.1` install SQL is empty: it creates no functions, tables, roles, policies, hooks, or authentication API.

There is therefore no functional usage example. If a locally repaired or manually installed copy is present, this query only confirms extension registration; it does not confirm authentication behavior:

```sql
SELECT extname, extversion
FROM pg_extension
WHERE extname = 'postgrest_auth';
```

### Caveats

- The README says “Work in progress” and documents no API or setup workflow.
- The Makefile refers to hyphenated `postgrest-auth.control` and `postgrest-auth--0.1.sql` names, while the repository contains underscored files. As reviewed, the normal PGXS `DATA` and control-file lookup do not match the tracked files.
- Do not infer PostgREST JWT, role switching, user management, or row-level-security integration from the project name. None of those features exists in the reviewed source.
