

## Usage

> [btree_gist: B-tree equivalent GiST operator classes](https://www.postgresql.org/docs/current/btree-gist.html)

Provides GiST index operator classes for data types that normally only support B-tree indexing. Enables exclusion constraints combining equality with range operators.

```sql
CREATE EXTENSION btree_gist;
```

### Supported Data Types

`int2`, `int4`, `int8`, `float4`, `float8`, `numeric`, `timestamp with time zone`, `timestamp without time zone`, `time with time zone`, `time without time zone`, `date`, `interval`, `oid`, `money`, `char`, `varchar`, `text`, `bytea`, `bit`, `varbit`, `macaddr`, `macaddr8`, `inet`, `cidr`, `uuid`, `bool`, and all `enum` types.

### Distance Operator

The `<->` operator is provided for nearest-neighbor searches on numeric and temporal types.

### Examples

```sql
-- GiST index on integer column
CREATE INDEX idx ON test USING GIST (a);
SELECT * FROM test WHERE a < 10;

-- Nearest-neighbor search
SELECT *, a <-> 42 AS dist FROM test ORDER BY a <-> 42 LIMIT 10;

-- Exclusion constraint: each cage can only contain one type of animal
CREATE TABLE zoo (
  cage   integer,
  animal text,
  EXCLUDE USING GIST (cage WITH =, animal WITH <>)
);

INSERT INTO zoo VALUES (1, 'lion');    -- OK
INSERT INTO zoo VALUES (1, 'tiger');   -- ERROR: conflicting key value
INSERT INTO zoo VALUES (2, 'tiger');   -- OK

-- Exclusion constraint for non-overlapping time ranges per room
CREATE TABLE reservations (
  room int,
  during tsrange,
  EXCLUDE USING GIST (room WITH =, during WITH &&)
);
```
