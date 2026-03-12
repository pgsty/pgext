

## Usage

> [table_version: PostgreSQL table versioning extension](https://github.com/linz/postgresql-tableversion)

PostgreSQL table versioning extension, recording row modifications and its history. The extension provides APIs for accessing snapshots of a table at certain revisions and the difference generated between any two given revisions. It uses a PL/PgSQL trigger based system to record and provide access to row revisions.

### Quick Start

```sql
CREATE EXTENSION table_version;

CREATE SCHEMA foo;
SET search_path TO foo, public;

CREATE TABLE foo.bar (
    id INTEGER NOT NULL PRIMARY KEY,
    baz TEXT
);

-- Enable versioning
SELECT table_version.ver_enable_versioning('foo', 'bar');

-- Create a revision and insert data
SELECT table_version.ver_create_revision('Insert data');
INSERT INTO foo.bar (id, baz) VALUES
  (1, 'foo bar 1'),
  (2, 'foo bar 2'),
  (3, 'foo bar 3');
SELECT table_version.ver_complete_revision();

-- Show differences between revisions
SELECT * FROM table_version.ver_get_foo_bar_diff(1001, 1002);
```


## How It Works

When a table is versioned, the original table data is left untouched and a new revision table is created with all the same fields plus `_revision_created` and `_revision_expired` fields. A row-level trigger is set up on the original table to record every insert, update and delete in the revision data table. A statement-level trigger is set up to forbid TRUNCATE.

### Table Prerequisites

- The table must have a unique non-composite integer, bigint, text or varchar column
- The table must not be temporary


## Auto Revisions

If you don't want to call `ver_create_revision` and `ver_complete_revision` explicitly, auto-revision mode groups edits by transactions:

```sql
SELECT table_version.ver_enable_versioning('foo', 'bar');

BEGIN;
INSERT INTO foo.bar (id, baz) VALUES (1, 'foo bar 1');
INSERT INTO foo.bar (id, baz) VALUES (2, 'foo bar 2');
COMMIT;

BEGIN;
UPDATE foo.bar SET baz = 'foo bar 1 edit' WHERE id = 1;
COMMIT;

SELECT * FROM table_version.foo_bar_revision;
```

The revision message will be automatically created based on the transaction ID.


## Replicate Data Using Table Differences

To maintain a copy of table data on a remote system:

```sql
-- 1. Determine which tables are versioned
SELECT * FROM table_version.ver_get_versioned_tables();

-- 2. Get the base revision
SELECT table_version.ver_get_table_base_revision('foo', 'bar');

-- 3. Create a base snapshot
CREATE TABLE foo_bar_copy AS
SELECT * FROM table_version.ver_get_foo_bar_revision(
    table_version.ver_get_table_base_revision('foo', 'bar')
);

-- 4. Get differences to apply incremental updates
SELECT * FROM table_version.ver_get_foo_bar_diff(
    my_last_revision,
    table_version.ver_get_table_last_revision('foo', 'bar')
);
```


## Security Model

- Anyone can create revisions
- Revisions can only be completed by their creators
- Only those who have ownership privileges on a table can enable/disable versioning
- Only empty revisions can be deleted
- Only the creator of a revision can delete it

Note: Disabling versioning on a table results in all history for that table being deleted.


## Key Functions

| Function | Description |
|----------|-------------|
| `ver_enable_versioning(schema, table)` | Enable versioning on a table |
| `ver_disable_versioning(schema, table)` | Disable versioning and remove history |
| `ver_create_revision(comment)` | Create a new revision |
| `ver_complete_revision()` | Mark current revision as complete |
| `ver_get_<schema>_<table>_diff(rev1, rev2)` | Get differences between two revisions |
| `ver_get_<schema>_<table>_revision(rev)` | Get snapshot at a specific revision |
| `ver_get_versioned_tables()` | List all versioned tables |
| `ver_get_last_revision()` | Get the last revision number |
