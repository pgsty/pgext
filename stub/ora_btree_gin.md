

## Usage

> [ora_btree_gin: Support for indexing oracle datatypes in GIN](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ora_btree_gin)

The `ora_btree_gin` extension adds GIN (Generalized Inverted Index) operator class support for Oracle-compatible data types provided by IvorySQL.

### Enabling

```sql
CREATE EXTENSION ora_btree_gin;
```

### Creating GIN Indexes on Oracle Types

This extension allows you to create GIN indexes on Oracle-compatible data types such as `NUMBER`, `VARCHAR2`, and Oracle-style `DATE`, similar to how the standard `btree_gin` extension works for native PostgreSQL types.

```sql
CREATE TABLE t (val NUMBER);
CREATE INDEX t_val_gin_idx ON t USING gin (val);
```

### Use Cases

GIN indexes with Oracle-compatible types are particularly useful for:

- Multi-column index queries where some columns use Oracle data types
- Queries combining full-text search with Oracle-type column filters
- Any scenario requiring inverted index support for Oracle-compatible data types

This extension is part of the IvorySQL Oracle compatibility suite and requires `ivorysql_ora` to be available.
