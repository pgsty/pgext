## Usage

Sources:

- [Official README](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/README.md)
- [Extension SQL for version 1.0](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/replace_empty_string--1.0.sql)
- [Trigger implementation](https://github.com/okbob/replace_empty_string/blob/ee225b0b52a97ade45e52f6e71d2281c45c01b11/replace_empty_string.c)

`replace_empty_string` provides a generic row trigger that changes zero-length string values to SQL nulls before an insert or update. It is useful for Oracle-oriented migrations where empty strings are expected to behave like null, but it changes data semantics for every string-category column on the attached table.

### Core Workflow

Install the extension and attach its trigger function as a `BEFORE` row trigger to each table that needs the conversion:

```sql
CREATE EXTENSION replace_empty_string;

CREATE TABLE contacts (
  contact_id bigint PRIMARY KEY,
  display_name text,
  external_code varchar(32)
);

CREATE TRIGGER contacts_replace_empty_string
BEFORE INSERT OR UPDATE ON contacts
FOR EACH ROW
EXECUTE FUNCTION replace_empty_string();

INSERT INTO contacts VALUES (1, '', 'A-001');
SELECT display_name IS NULL FROM contacts WHERE contact_id = 1;
```

The trigger scans all attributes whose PostgreSQL type category is string, unwraps domains to their base type, and replaces only values whose stored length is zero. Whitespace-only strings are left unchanged.

### Warning Mode

Pass the literal trigger argument `on` to emit a warning for each column changed:

```sql
DROP TRIGGER contacts_replace_empty_string ON contacts;

CREATE TRIGGER contacts_replace_empty_string
BEFORE INSERT OR UPDATE ON contacts
FOR EACH ROW
EXECUTE FUNCTION replace_empty_string('on');
```

Any absent or different first argument leaves warnings disabled. Warning mode can produce substantial log and client-message volume during bulk loads, so use it for migration discovery or testing rather than as a default observability mechanism.

### Operational Notes

The function must be invoked by a row-level `BEFORE INSERT OR UPDATE` trigger; direct calls, statement triggers, after triggers, and delete events are rejected by the C implementation. Existing stored rows are not rewritten when the trigger is created. Backfill them explicitly if uniform historical data is required.

Version 1.0 is relocatable and declares no preload or extension dependency. The upstream README states that this code was merged into Orafce, so evaluate the maintained Orafce implementation for new Oracle-compatibility deployments. Converting empty strings to null can affect `NOT NULL` constraints, unique indexes, predicates, application serialization, and length checks; test those semantics before attaching the trigger broadly.
