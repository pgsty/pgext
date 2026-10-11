## Usage

Sources:

- [README.md](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/README.md)
- [fkarray--1.0.0.sql](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray--1.0.0.sql)
- [fkarray.c](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray.c)
- [fkarray.control](https://github.com/pjungwir/fkarray/blob/434290df52c2163676ee7f7d73888ab1b1ec19c8/fkarray.control)

`fkarray` 1.0.0 checks that each value in an array column references an existing unique or primary-key value. It uses triggers and functions rather than a native array foreign-key constraint.

### Core Workflow

```sql
CREATE EXTENSION fkarray;
CREATE EXTENSION intarray;
CREATE TABLE colors (id integer PRIMARY KEY);
CREATE TABLE products (id integer PRIMARY KEY, color_ids integer[]);
SELECT fkarray_create('public','products','color_ids','public','colors','id');
INSERT INTO colors VALUES (1), (2);
INSERT INTO products VALUES (1, ARRAY[1,2]);
```

### Operational Boundaries

`fkarray_create()` accepts the referencing and referenced schema, table and column, followed by optional `on_delete` and `on_update` actions. Both default to `no_action`; `cascade` removes the referenced value from arrays, including on key update. `fkarray_drop()` removes the generated enforcement triggers. Array element and target types must match and the target must have a unique or primary-key index. Integer arrays need `intarray`, bigint arrays need `bigintarray`; other types use ordinary PostgreSQL operators. These are type-dependent requirements, not unconditional control dependencies. The control marks installation trusted and non-relocatable; callers still need the privileges to create the required table triggers.
