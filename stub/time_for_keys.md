## Usage

Sources:

- [Official upstream documentation](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/README.md)
- [Official extension SQL](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/time_for_keys--0.0.1.sql)
- [Official coverage aggregate implementation](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/completely_covers.c)

`time_for_keys` 0.0.1 implements trigger-based foreign keys between valid-time tables. A child reference is valid when parent rows with the same integer key collectively cover the child's entire `tstzrange`, even if no single parent row covers it. This is an early prototype for state-time history, not system-time auditing or a general SQL temporal-constraint implementation.

### Define a Temporal Relationship

Parent history normally needs a non-overlap exclusion constraint. After creating the tables, install the extension and create the four constraint triggers:

```sql
CREATE EXTENSION btree_gist;
CREATE EXTENSION time_for_keys;

CREATE TABLE houses (
  id integer NOT NULL,
  valid_at tstzrange NOT NULL,
  EXCLUDE USING gist (id WITH =, valid_at WITH &&)
);

CREATE TABLE rooms (
  id integer NOT NULL,
  house_id integer,
  valid_at tstzrange NOT NULL
);

SELECT create_temporal_foreign_key(
  'rooms_have_a_house',
  'rooms',  'house_id', 'valid_at',
  'houses', 'id',       'valid_at'
);
```

The function validates existing child rows, then creates deferred constraint triggers for child INSERT/UPDATE and parent UPDATE/DELETE. Violations use foreign-key SQLSTATE `23503`.

### Important Objects

- `completely_covers(tstzrange, tstzrange)` is an aggregate that tests whether input periods collectively cover a fixed target range.
- `create_temporal_foreign_key(text, text, text, text, text, text, text)` creates and validates the temporal relationship.
- `drop_temporal_foreign_key(text, text, text)` removes its four generated triggers.

### Prototype Limits

The implementation supports only integer keys, one key column, and `tstzrange`. Table names cannot include schemas. It does not implement CASCADE, SET NULL, SET DEFAULT, MATCH modes, multi-column keys, other range types, views, or a catalog tracking created relationships. Trigger names include the supplied constraint name and remain subject to PostgreSQL's identifier limit.

The checks are dynamic SQL and do not provide the mature locking, plan caching, DDL dependency management, or concurrency analysis of built-in foreign keys. Test concurrent parent/child changes, overlapping and adjacent ranges, infinite/empty bounds, deferred constraint behavior, table renames, dump/restore, and trigger cleanup. Use it only after verifying the exact semantics under the application's transaction isolation level.
