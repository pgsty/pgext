## Usage

Sources:

- [Official upstream README](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/README.rst)
- [Official extension control file (spqrhash.control)](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/spqrhash.control)
- [Official extension SQL (spqrhash--1.1--1.2.sql)](https://github.com/pg-sharding/spqrhash/blob/3859461eb18f0c669d5ed016a38d15768da84ace/sql/spqrhash--1.1--1.2.sql)

`spqrhash` — The extension provides hash funcitons for SPQR to work with PG. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION spqrhash;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `spqrhash_city32(bytea)` is an extension function and returns `int8`.
- `spqrhash_city32(id uuid)` is an extension function and returns `int8`.
- `spqrhash_city32(int8)` is an extension function and returns `int8`.
- `spqrhash_city32(int8[])` is an extension function and returns `int8`.
- `spqrhash_city32(text)` is an extension function and returns `int8`.
- `spqrhash_murmur3(bytea)` is an extension function and returns `int8`.
- `spqrhash_murmur3(id uuid)` is an extension function and returns `int8`.
- `spqrhash_murmur3(int8)` is an extension function and returns `int8`.
- `spqrhash_murmur3(int8[])` is an extension function and returns `int8`.
- `spqrhash_murmur3(text)` is an extension function and returns `int8`.

### Requirements and Caveats

- The reviewed control file declares default version `1.2`.
- The control file marks the extension as relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
