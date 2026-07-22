## Usage

Sources:

- [Official upstream documentation](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/README.md)
- [Official Rust implementation](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/src/lib.rs)
- [Official Rust package manifest](https://github.com/ari-becker/pgcuid2/blob/0782b48bcb6eedbfc8785eca9c1e5b5f6215a22f/Cargo.toml)

`pgcuid2` adds a single PostgreSQL function that generates CUID2 text identifiers through the Rust `cuid2` crate. Use it when an application explicitly requires CUID2-compatible text keys; the result is not a PostgreSQL `uuid` and should be treated as an opaque, case-sensitive string.

### Core Workflow

Install the pgrx-built extension and use the generator as a column default:

```sql
CREATE EXTENSION pgcuid2;

SELECT cuid2_create_id();

CREATE TABLE account (
  id text PRIMARY KEY DEFAULT cuid2_create_id(),
  display_name text NOT NULL
);
```

The reviewed version returns a 24-character text value. `cuid2_create_id()` is declared `VOLATILE`, parallel safe, and security invoker, so PostgreSQL evaluates it for each inserted row rather than folding it as a constant.

### Data-Model Notes

Store generated values in `text` or a domain with a length check that matches the pinned implementation. Do not cast them to `uuid`, derive creation time from them, or sort them as chronological identifiers. Text primary keys are wider than integer or UUID keys and increase index, foreign-key, and cache footprint; benchmark that cost for large tables.

### Compatibility and Maintenance

The archived upstream release uses pgrx 0.11 and declares PostgreSQL 11 through 16 build features. Pin both the extension and Rust dependency versions, build against the exact PostgreSQL major, and test uniqueness and output format before migrations. The function delegates generation to its dependency; do not make stronger unpredictability or collision guarantees than the pinned CUID2 implementation documents. If external services validate a particular CUID2 revision, include cross-language fixtures in the application test suite.
