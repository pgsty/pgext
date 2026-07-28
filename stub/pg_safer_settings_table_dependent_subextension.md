## Usage

Sources:

- [Official upstream README](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/README.md)
- [Official extension control file (pg_safer_settings_table_dependent_subextension.control)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings_table_dependent_subextension/pg_safer_settings_table_dependent_subextension.control)
- [Official extension SQL (pg_safer_settings_table_dependent_subextension--monover.sql)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings_table_dependent_subextension/pg_safer_settings_table_dependent_subextension--monover.sql)

`pg_safer_settings_table_dependent_subextension` — A handful of functions and mechanisms to make dealing with settings in Postgres a bit … safer. Use it when administering or automating the database behavior described above. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION pg_safer_settings_table_dependent_subextension;

-- To change for the duration of the session:
SET app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', false);

-- To change for the duration of the transaction:
SET LOCAL app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', true);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `monover`.
- Install the confirmed extension dependencies first: `pg_safer_settings_table_dependent_extension`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
