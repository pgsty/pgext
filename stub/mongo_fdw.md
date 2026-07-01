


## Usage

Sources: [README](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_5_3/README.md), [REL-5_5_3 release](https://github.com/EnterpriseDB/mongo_fdw/releases/tag/REL-5_5_3)

### Create Server

```sql
CREATE EXTENSION mongo_fdw;

CREATE SERVER mongo_server FOREIGN DATA WRAPPER mongo_fdw
  OPTIONS (address '127.0.0.1', port '27017');
```

**Server Options:** `address` (default `127.0.0.1`), `port` (default `27017`), `authentication_database`, `replica_set`, `read_preference` (`primary`, `secondary`, `primaryPreferred`, `secondaryPreferred`, `nearest`), `ssl` (default `false`), `pem_file`, `pem_pwd`, `ca_file`, `ca_dir`, `crl_file`, `weak_cert_validation`, `use_remote_estimate` (default `false`), `enable_join_pushdown` (default `true`), `enable_aggregate_pushdown` (default `true`), `enable_order_by_pushdown` (default `true`).

### Create User Mapping

```sql
CREATE USER MAPPING FOR pguser SERVER mongo_server
  OPTIONS (username 'mongouser', password 'mongopass');
```

### Create Foreign Table

```sql
CREATE FOREIGN TABLE warehouse (
  _id name,
  warehouse_id int,
  warehouse_name text,
  warehouse_created timestamptz
)
SERVER mongo_server
OPTIONS (database 'mydb', collection 'warehouse');
```

**Important:** The first column must be `_id` of type `name` (MongoDB's object identifier).

**Table Options:** `database` (default `test`), `collection` (defaults to table name), `enable_join_pushdown`, `enable_aggregate_pushdown`, `enable_order_by_pushdown`.

### CRUD Operations

```sql
SELECT warehouse_id, warehouse_name FROM warehouse WHERE warehouse_id > 10;
INSERT INTO warehouse VALUES ('new_id', 100, 'New Warehouse', now());
UPDATE warehouse SET warehouse_name = 'Updated' WHERE warehouse_id = 100;
DELETE FROM warehouse WHERE warehouse_id = 100;
```

### Pushdown Features

mongo_fdw pushes down WHERE clauses, joins between foreign tables on the same server, aggregate functions, ORDER BY, LIMIT, and OFFSET to MongoDB for efficient query execution. Use the `mongo_fdw.enable_join_pushdown`, `mongo_fdw.enable_aggregate_pushdown`, `mongo_fdw.enable_order_by_pushdown`, and `mongo_fdw.log_remote_query` GUCs when diagnosing remote plans.

### Version Notes

`mongo_fdw` 5.5.3, released upstream as `REL-5_5_3`, adds PostgreSQL 18 support, updates bundled `mongoc-driver` and `json-c` libraries for MongoDB 8, adds the `mongo_fdw.log_remote_query` debug GUC, and fixes nested-field, WHERE, ORDER BY, and unsafe join-pushdown cases. Upstream dropped PostgreSQL 12 support in this line.

### Notes

- BSON only supports UTF-8; ensure PostgreSQL database uses UTF-8 encoding
- Column names with uppercase letters or dots (for nested documents) require double-quoting
- PostgreSQL limits column names to 63 characters by default
