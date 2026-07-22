## Usage

Sources:

- [Official examples](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Examples.md)
- [Official control file](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/pg_nanoid.control)
- [Official build manifest](https://github.com/lameferret/pg_nanoid/blob/d216e0b481d8ffb8a2c93b1dc3420bdafe31b486/Cargo.toml)

`pg_nanoid` adds a `nanoid` identifier type plus generators for individual, custom-length, and batched Nano IDs. It is useful for compact application identifiers, but uniqueness remains probabilistic and should still be enforced with a primary key or unique constraint.

### Core Workflow

```sql
CREATE EXTENSION pg_nanoid;

CREATE TABLE users (
    id nanoid PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO users (id, name)
SELECT id, 'User #' || row_number() OVER ()
FROM nanoid_batch(5);

CREATE TABLE sessions (
    user_id nanoid REFERENCES users(id),
    token text UNIQUE NOT NULL
);

INSERT INTO sessions (user_id, token)
SELECT id, nanoid_len(64)
FROM users
LIMIT 1;
```

Use `nanoid_batch` to amortize generation for a set of rows and `nanoid_len` when an application requires an explicitly chosen length.

### Important Objects

- `nanoid` is the extension's identifier type.
- `nanoid_batch(count)` returns rows containing generated IDs.
- `nanoid_len(length)` returns a generated text ID of the requested length.

Inspect the installed function signatures before using APIs not shown in the pinned examples; the reviewed package declares version `0.0.0` rather than a stable semantic release.

### Caveats

The reviewed manifest enables only the PostgreSQL 18 pgrx feature and requires superuser installation; other server majors need a separately built and tested artifact. The control file marks the extension untrusted and non-relocatable.

Size the identifier length and alphabet for the expected population and risk tolerance, retain database uniqueness constraints, and retry on a rare collision. Do not automatically treat generated IDs as authentication secrets: separately validate entropy source, alphabet, exposure lifetime, comparison, logging, and revocation requirements for security tokens.
