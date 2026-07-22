## Usage

Sources:

- [Official README](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/README.md)
- [Version 0.6 SQL notes](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/sql/pg_lexo--0.6.0.sql)
- [Public function implementation](https://github.com/blad3mak3r/pg_lexo/blob/521605b23f5c5cfa80f7a743c536b8040bb81e5a/src/schema.rs)

`pg_lexo` 0.6.0 provides a Base62 `lexo` type for stable user-controlled ordering. It generates short values before, after, or between existing positions, allowing an item to move without renumbering every row. The type has comparison operators plus default B-tree and hash operator classes.

### Core Workflow

```sql
CREATE EXTENSION pg_lexo;

CREATE TABLE items (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title text NOT NULL,
    position lexo NOT NULL UNIQUE
);

INSERT INTO items (title, position)
VALUES ('first', lexo_first());

INSERT INTO items (title, position)
SELECT 'last', lexo_after(max(position)) FROM items;

INSERT INTO items (title, position)
VALUES ('middle', lexo_between('H'::lexo, 'I'::lexo));

SELECT * FROM items ORDER BY position;
```

`lexo_first()` starts at `H`. `lexo_after(lexo)`, `lexo_before(lexo)`, and `lexo_between(lexo, lexo)` generate neighboring positions; either argument to `lexo_between` may be NULL to represent an open end. Inputs contain only `0-9`, `A-Z`, and `a-z`, whose ASCII order matches the type's ordering.

### Table Helpers

`lexo_next(table, position_column, key_column, key_value)` finds the current maximum and returns the next position, optionally within a group. `lexo_add_column(table, column)` adds a `lexo` column. `lexo_rebalance(table, position_column, key_column, key_value)` rewrites matching rows with evenly distributed positions and returns the number updated.

The helpers issue dynamic SQL and require the caller to have the corresponding table privileges. Rebalancing updates every selected row and should run with application writers coordinated. Position generation itself does not reserve a value: concurrent sessions can derive the same position, so enforce a suitable unique constraint and retry or serialize conflicting moves. When no Base62 value exists between adjacent positions (or before the minimum), rebalance first.

### Version 0.6 Boundary

Version 0.6 removed the dedicated `lexo` schema. The former `lexo.lexorank` type became `lexo`, and schema-qualified functions such as `lexo.first()` became names such as `lexo_first()`. Older examples are incompatible until migrated. This release's Cargo features target PostgreSQL 16–18 through pgrx 0.16.1.
