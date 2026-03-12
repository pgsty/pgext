

## Usage

> [ora_btree_gist: Support for oracle indexing common datatypes in GiST](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ora_btree_gist)

The `ora_btree_gist` extension adds GiST (Generalized Search Tree) operator class support for Oracle-compatible data types provided by IvorySQL.

### Enabling

```sql
CREATE EXTENSION ora_btree_gist;
```

### Creating GiST Indexes on Oracle Types

This extension allows you to create GiST indexes on Oracle-compatible data types such as `NUMBER`, `VARCHAR2`, and Oracle-style `DATE`, similar to how the standard `btree_gist` extension works for native PostgreSQL types.

```sql
CREATE TABLE t (val NUMBER);
CREATE INDEX t_val_gist_idx ON t USING gist (val);
```

### Use Cases

GiST indexes with Oracle-compatible types enable:

- Exclusion constraints using Oracle data types
- Nearest-neighbor searches on Oracle-type columns
- Range queries with GiST optimizations for Oracle-compatible types

This extension is part of the IvorySQL Oracle compatibility suite and requires `ivorysql_ora` to be available.
