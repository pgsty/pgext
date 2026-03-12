


## Usage

> [mongo_fdw: Foreign data wrapper for MongoDB access](https://github.com/EnterpriseDB/mongo_fdw)

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

mongo_fdw pushes down WHERE clauses, JOINs between foreign tables on the same server, aggregate functions, and ORDER BY to MongoDB for efficient query execution.

### Notes

- BSON only supports UTF-8; ensure PostgreSQL database uses UTF-8 encoding
- Column names with uppercase letters or dots (for nested documents) require double-quoting
- PostgreSQL limits column names to 63 characters by default
