## Usage

Sources:

- [README](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/README.md)
- [Reference](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/docs/REFERENCE.md)
- [Control file](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext.control)
- [Extension SQL](https://github.com/cgkanchi/pg_lion/blob/v0.3.0/pg_lion_citext--0.3.0.sql)

`pg_lion_citext` 0.3.0 registers the `citext_ops` operator class for case-insensitive text keys in `lion` indexes. It is a SQL companion shipped with `pg_lion`, with no separate shared library.

### Core Workflow

Create `citext` and `pg_lion` first, then enable the companion. Its default operator class uses `citext` hashing and comparison functions for equality, inequality and ranges (`=`, `<>`, `<`, `<=`, `>=`, `>`).

```sql
CREATE EXTENSION citext;
CREATE EXTENSION pg_lion;
CREATE EXTENSION pg_lion_citext;
CREATE TABLE lion_contacts (email citext);
INSERT INTO lion_contacts VALUES ('Ada@example.com'), ('ADA@example.com');
CREATE INDEX lion_email ON lion_contacts USING lion (email);
SELECT count(*) FROM lion_contacts WHERE email = 'ada@example.com';
```

The example returns two rows in the count because both spellings compare equal. Applications keep using ordinary `citext` predicates.

### Operational Boundaries

Requires PostgreSQL 16 or later and superuser installation. Case folding and ordering follow `citext` semantics. Use matching files from the parent distribution. The parent indexes retain their visibility, `VACUUM`, and optional custom-WAL preload requirements. This companion does not enable preloading or change existing indexes by itself.
