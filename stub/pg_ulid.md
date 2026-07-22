## Usage

Sources:

- [Official README](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/README.md)
- [Version 0.0.4 SQL API](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/pg_ulid--0.0.4.sql)
- [Generator implementation](https://github.com/edoceo/pg-ulid/blob/eda5ace98b0f1668297e95476cde5136ff06c97b/pg_ulid.c)

`pg_ulid` 0.0.4 exposes one C function that generates a 26-character ULID string. Despite broader claims and examples in the README, the installable SQL does not define a PostgreSQL `ulid` type, comparison operators, validation helpers, formatting helpers, or a timestamp-accepting overload.

### Actual SQL API

```sql
CREATE EXTENSION pg_ulid;

CREATE TABLE events (
    id text PRIMARY KEY DEFAULT ulid_create(),
    payload jsonb NOT NULL
);

INSERT INTO events (payload) VALUES ('{"kind":"created"}');
SELECT id, payload FROM events ORDER BY id;
```

The only callable object is `ulid_create() RETURNS text`, declared volatile and strict. Calls such as `ulid_create(now())` shown in the README are not defined by version `0.0.4`, and values receive ordinary text comparison and indexing semantics rather than a custom binary type.

### Identifier Boundaries

The implementation initializes an embedded `ulid-c` generator on every call and returns its text output. It exposes no way to provide a timestamp, select an entropy policy, validate canonical encoding, decode the embedded time, or request monotonic generation within the same millisecond. Validate format and uniqueness in the application if those properties are required, and do not assume cluster-wide monotonic ordering.

ULIDs expose approximate creation time and lexical order only when all producers follow the same canonical format. The project has no release artifacts or upgrade chain and is abandoned at a small prototype. For durable identifiers, prefer a maintained ULID extension with explicit type and randomness guarantees, or store application-generated values under a checked contract.
