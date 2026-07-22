## Usage

Sources:

- [postgrest_auth README at the reviewed commit](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/README.md)
- [postgrest_auth.control at the reviewed commit](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth.control)
- [Version 0.1 install SQL](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/postgrest_auth--0.1.sql)
- [Upstream Makefile](https://github.com/monacoremo/postgrest-auth/blob/26b70693abc1897d1dc4812e0c40911ee9cb15a2/Makefile)

`postgrest_auth` is an unfinished extension skeleton, not a usable PostgREST authentication component. The reviewed version `0.1` install SQL is empty and creates no functions, tables, roles, policies, hooks, or authentication API.

### Core Workflow

There is no functional workflow to document. If a locally repaired or manually installed copy has been registered, the following query only reports catalog registration; it does not test authentication:

```sql
SELECT extname, extversion
FROM pg_extension
WHERE extname = 'postgrest_auth';
```

### Object Index

The reviewed install script defines no SQL objects. In particular, it provides no login function, JWT verifier, role-switching helper, credential table, or row-level-security policy.

### Operational Boundaries

- The README says “Work in progress” and documents no API, configuration, or enablement sequence.
- The Makefile names `postgrest-auth.control` and `postgrest-auth--0.1.sql`, while the tracked files are `postgrest_auth.control` and `postgrest_auth--0.1.sql`. The normal PGXS control and data-file names therefore do not match the repository contents as reviewed.
- Do not infer a PostgREST security model from the project name. Applications still need an independently reviewed authentication, role, privilege, and row-level-security design.
