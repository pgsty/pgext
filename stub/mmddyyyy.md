## Usage

Sources:

- [README.md](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/README.md)
- [mmddyyyy.control](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/mmddyyyy.control)
- [mmddyyyy--0.1.0.sql](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/sql/mmddyyyy--0.1.0.sql)
- [mmddyyyy.c](https://github.com/FranckPachot/pg-mm-dd-yyyy/blob/6d68f142d9777c2d6ee851d181746e77c6a71d6f/src/mmddyyyy.c)

`mmddyyyy` is a teaching extension that stores a date as ten literal MM/DD/YYYY bytes. It demonstrates C base types, operators, B-tree ordering, and a GiST operator class. Upstream explicitly excludes production use; ordinary application dates should use PostgreSQL’s built-in date type.

### Core Workflow

```sql
CREATE EXTENSION mmddyyyy;
CREATE TABLE events (id integer PRIMARY KEY, happened_on mmddyyyy NOT NULL);
INSERT INTO events VALUES (1, '09/15/2026'), (2, '12/31/2025');
CREATE INDEX events_date_btree ON events (happened_on);
CREATE INDEX events_date_gist ON events USING gist (happened_on);
SELECT * FROM events WHERE happened_on <@ '09/*/*'::mmddyyyy_pattern;
SELECT * FROM events
ORDER BY happened_on <-> '09/15/2026'::mmddyyyy_pattern LIMIT 5;
```

### Type and Index Semantics

`mmddyyyy` validates real Gregorian dates and canonical zero-padded text. Casts convert to and from `date`. `mmddyyyy_month`, `mmddyyyy_day`, and `mmddyyyy_year` expose components. Ordinary comparisons and the default B-tree order dates chronologically, despite the month-first storage spelling.

`mmddyyyy_pattern` permits a wildcard for any date component; `<@` tests component matching. `mmddyyyy_gist_ops` supports those patterns and `<->` nearest-neighbor search. Its distance prioritizes cyclic month difference, then day, then year; it is not elapsed time in days. Index effectiveness depends on the query shape.

### Requirements and Limits

The container and correctness suite target PostgreSQL 17. Install the C extension with superuser privileges; it is relocatable and needs no preload. Supported years are 0001–9999, with no BC, infinite dates, or time zones. Binary send/receive functions and upgrade scripts are absent. No license file was present in the reviewed source; a public repository alone does not establish redistribution rights.
