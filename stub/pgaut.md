## Usage

Sources:

- [Archived upstream README at the reviewed commit](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/README.md)
- [Version 1.0.0 SQL implementation](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut--1.0.0.sql)
- [Extension control file](https://github.com/redraiment/pgaut/blob/7db20e740ec8167c781b4456d77a2090d7ef2782/pgaut.control)

`pgaut` provides MySQL-style automatic update timestamps through an `auto_update_timestamp` domain and PostgreSQL event triggers. Tables created with a column of that domain receive a row trigger that assigns `clock_timestamp()` when another column changes.

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

### Caveats

Installation creates database-wide event triggers for `CREATE TABLE`, `ALTER TABLE`, and `DROP TABLE`, and dynamically creates per-table PL/pgSQL trigger functions. This requires elevated privileges and deserves review before use in a shared database. The repository is archived, its last source change was in 2018, and upstream's PostgreSQL 9.1-or-newer claim was not tested against current majors.
