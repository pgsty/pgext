## Usage

Sources:

- [Official README](https://github.com/adjust/pg-roleman/blob/fdd31900d8f56fa260cbeefa8568444f3c717420/README.md)
- [Version 0.3.0 SQL](https://github.com/adjust/pg-roleman/blob/fdd31900d8f56fa260cbeefa8568444f3c717420/extension/roleman--0.3.0.sql)
- [Official PGXN distribution](https://pgxn.org/dist/roleman/0.3.0/)

`roleman` 0.3.0 is a pure-SQL API for parameterized role creation, alteration, password changes, grants, role settings, and removal. It quotes identifiers and whitelists role attributes and privilege keywords, helping applications avoid constructing raw utility SQL themselves.

### Create and Configure a Role

The extension installs into the fixed `roleman` schema and requires superuser installation. Its functions are security-invoker functions: the caller still needs the PostgreSQL privileges required by the resulting `CREATE ROLE`, `ALTER ROLE`, `GRANT`, `REVOKE`, or `DROP ROLE` command.

```sql
CREATE EXTENSION roleman;

SELECT roleman.create_role(
    'app_reader',
    ARRAY['LOGIN', 'NOINHERIT']
);

SELECT roleman.grant_schema(
    'app_reader',
    'reporting',
    ARRAY['USAGE']
);

SELECT roleman.grant_table(
    'app_reader',
    'reporting.daily_sales'::regclass,
    ARRAY['SELECT']
);

SELECT roleman.set_connection_limit('app_reader', 20);
SELECT roleman.set_search_path('app_reader', ARRAY['reporting', 'public']);
```

### Main Function Groups

- `roleman.create_role(text, text[])` and `roleman.alter_base(text, text[])` apply whitelisted role attributes.
- `roleman.role_set_password(text, text, timestamp)` sets a password and optional expiry.
- `roleman.grant_database`, `roleman.grant_schema`, `roleman.grant_schema_all`, `roleman.grant_table`, `roleman.grant_seq`, and `roleman.grant_function` apply whitelisted privileges.
- `roleman.role_grant_role(text, text)` grants role membership.
- `roleman.set_guc`, `roleman.set_guc_from_current`, `roleman.set_connection_limit`, and `roleman.set_search_path` alter role defaults.
- `roleman.role_blank_perms(text)` broadly revokes privileges and memberships visible from the current database; `roleman.drop_role(text)` removes the role.

### Safety and Known Defects

Password values travel as SQL function arguments and can be exposed through statement logging, activity views, monitoring, or error context. Use encrypted connections and a credential workflow that prevents SQL text capture.

`roleman.role_blank_perms(text)` is destructive and incomplete as a global reset: it loops over schemas in the current database and revokes explicit database/object privileges and memberships, but does not transfer ownership, remove default privileges, or inspect every other database. Inventory dependencies and test the exact revocation plan first.

The 0.3.0 implementation of `roleman.reset_guc(text, text)` mistakenly issues `ALTER USER ... SET ... FROM CURRENT` instead of resetting the setting. Do not call it; use a reviewed native `ALTER ROLE ... RESET ...` command. The project is archived and was tested historically around PostgreSQL 9.4 through 9.6, with PostgreSQL 10 only expected, not established. Regression-test its catalog queries and generated utility commands on the exact target release.
