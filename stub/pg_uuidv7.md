## Usage

Sources:

- [README](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/README.md)
- [Control](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/pg_uuidv7.control)
- [Changelog](https://api.pgxn.org/src/pg_uuidv7/pg_uuidv7-1.8.0/CHANGELOG.md)

`pg_uuidv7` generates time-ordered UUIDv7 values and converts between their timestamp component and `timestamptz`. Release 1.8.0 supports PostgreSQL 13–19 and uses SQL extension version 1.8. PostgreSQL 18 has a built-in `uuidv7()` function, but this extension also supplies the conversion helpers below.

### Core Workflow

```sql
CREATE EXTENSION pg_uuidv7;
SELECT uuid_generate_v7();
SELECT uuid_v7_to_timestamptz('018570bb-4a7d-7c7e-8df4-6d47afd8c8fc');
SELECT uuid_timestamptz_to_v7('2023-01-02 04:26:40.637+00', true);
CREATE TABLE events (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v7(),
    data jsonb
);
```

`uuid_timestamptz_to_v7` accepts an optional boolean that zeros the random bits when true, useful for constructing time-range bounds. Ordinary generation retains random bits. Millisecond timestamp ordering does not guarantee strict ordering of every UUID generated within the same millisecond.

### Installation and Upgrades

Install the library matching the PostgreSQL server major. The control file is relocatable and has no declared extension dependency; installation requires the default superuser privilege. No preload or server restart is specified.

Release 1.8.0 adds PostgreSQL 19 support. After installing the matching files, follow the supplied update path with `ALTER EXTENSION pg_uuidv7 UPDATE`; the SQL extension receipt is 1.8.
