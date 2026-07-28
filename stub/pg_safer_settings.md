## Usage

Sources:

- [Official upstream README](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/README.md)
- [Official extension control file (pg_safer_settings.control)](https://github.com/bigsmoke/pg_safer_settings/blob/6ef74b74e6ad2799cae9d29c87990d353786e821/pg_safer_settings.control)

`pg_safer_settings` — pg_safer_settings provides a handful of functions and mechanisms to make dealing with settings in Postgres a bit … safer. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_safer_settings;

-- To change for the duration of the session:
SET app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', false);

-- To change for the duration of the transaction:
SET LOCAL app.settings.bla = 'blegherrerbypass';  -- or:
SELECT set_config('app.settings.bla', 'blegherrerbypass', true);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
