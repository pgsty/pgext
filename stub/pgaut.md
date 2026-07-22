## Usage

Sources:

- [Archived official README at the reviewed commit](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/README.md)
- [Version 1.0.0 SQL implementation](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut--1.0.0.sql)
- [Extension control file](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut.control)

`pgaut` emulates automatic update timestamps with an `auto_update_timestamp` domain, database-wide event triggers, and generated per-table row triggers. The domain is based on `timestamp` without time zone. The archived, early-version implementation requires superuser-level event-trigger installation and careful review before use in any shared database.

### Core Workflow

```sql
CREATE EXTENSION pgaut;

CREATE TABLE app_user (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  updated_at auto_update_timestamp DEFAULT current_timestamp
);

INSERT INTO app_user (name) VALUES ('alice');
UPDATE app_user SET name = 'Alice' WHERE name = 'alice';
```

Event triggers react to table creation, alteration, and deletion. For each table they create a BEFORE UPDATE trigger function that assigns `clock_timestamp()` to every column whose domain is recognized as `auto_update_timestamp`.

### Actual Semantics and Caveats

Despite the README's description, the generated row trigger does not compare OLD and NEW values. Every UPDATE resets all recognized timestamp-domain columns, even when no other business column changed or the timestamp itself was the only target.

The generator creates per-table functions and triggers even for tables without such a column. It constructs schema, table, column, trigger, and function identifiers without robust identifier quoting; unusual or crafted names can break generation or change the interpreted SQL. Domain lookup also does not reliably disambiguate the domain schema, and generated names can collide with existing objects.

Install only in a controlled schema namespace, inspect every generated object, and test create/alter/drop migrations. The repository is archived and dates from the PostgreSQL 9-era; exact current-major compatibility is unproven.
