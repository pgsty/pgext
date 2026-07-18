## Usage

Sources:

- [Pinned upstream README](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/README.md)
- [Version 1.0 installation SQL](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/dump_fdw--1.0.sql)
- [Pinned C implementation](https://github.com/MeetMe/dump_fdw/blob/1af2c525d1526442abca4722d3e9b3db55a6426f/dump_fdw.c)

dump_fdw is a read-only foreign data wrapper for querying relations stored in PostgreSQL custom-format dump files. Its intended uses are ad hoc backup analytics and selecting individual rows for a controlled restore.

### Setup and example

Install the extension, create its server object, and map a foreign table to one relation in a server-local dump file:

```sql
CREATE EXTENSION dump_fdw;

CREATE SERVER dump_server
FOREIGN DATA WRAPPER dump_fdw;

CREATE FOREIGN TABLE dumped_authors (
    id integer,
    last_name text,
    first_name text
)
SERVER dump_server
OPTIONS (
    file_name '/srv/backups/booktown.dump',
    schema_name 'public',
    relation_name 'authors'
);

SELECT * FROM dumped_authors ORDER BY id;
```

All three foreign-table options are required. The column definition must agree with the dumped relation. Only a superuser should be allowed to create or alter these mappings, because the file name can reference arbitrary files visible to the PostgreSQL server process.

### Limits and safety

The wrapper implements scans, not insert, update, or delete callbacks. It parses old custom-dump internals and materializes rows through a CSV-style path; compatibility therefore depends on the dump format, compression, data types, and PostgreSQL build.

Upstream calls version 1.0 a prototype, says it was tested only on PostgreSQL 9.1 through 9.4, and explicitly warns against production use. The source has not been updated since 2015, although the repository is not archived. Use an isolated restoration host, a copy of the dump, a low-privilege database, and representative validation before relying on results.
