## Usage

Sources:

- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/README.md)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/mongo_fdw.control)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/blob/REL-5_6_0/mongo_fdw--1.1.sql)
- [Official source](https://github.com/EnterpriseDB/mongo_fdw/releases/tag/REL-5_6_0)

`mongo_fdw` 5.6.0 reads and writes MongoDB collections from PostgreSQL 14–19. Its SQL/control version remains 1.1. Install the matching MongoDB client libraries and create the extension as a superuser; shared preload is unnecessary.

### Map a Collection

For an existing collection, configure an account with the intended remote privileges:

```sql
CREATE EXTENSION mongo_fdw;
CREATE SERVER mongo_server FOREIGN DATA WRAPPER mongo_fdw
  OPTIONS (address '127.0.0.1', port '27017');
CREATE USER MAPPING FOR CURRENT_USER SERVER mongo_server
  OPTIONS (username 'mongo_user', password 'replace_with_remote_password');
CREATE FOREIGN TABLE warehouse (
  _id name, warehouse_id integer, warehouse_name text
) SERVER mongo_server OPTIONS (database 'mydb', collection 'warehouse');
SELECT warehouse_id, warehouse_name FROM warehouse;
```

The first foreign-table column must be `_id` of type `name`. Let MongoDB create valid identifiers rather than inventing short ObjectId strings. Column types must match the BSON representation. Use `INSERT`, `UPDATE` and `DELETE` only with the expected remote permissions and identifier mapping; remote writes are not a substitute for a distributed transaction protocol.

### Connection and Planning

Server options include `authentication_database`, `replica_set`, `read_preference` and TLS certificate settings. Restrict user-mapping credentials and verify the remote certificate. The foreign-table options `database` and `collection` select data. This FDW does not implement `IMPORT FOREIGN SCHEMA`.

Join, aggregate, ordering and limit pushdown are controlled by the server/table options and session GUCs. `mongo_fdw.log_remote_query` helps diagnose pushed queries. BSON text requires UTF8. Quote nested or uppercase field names; field names starting with a dollar sign or containing one after a dot are unsupported.

### Changes and Upgrade Checks

5.6.0 fixes URI-option injection and short-identifier out-of-bounds reads. Cursor failures now raise errors rather than silently returning no rows. BSON doubles converted to integer types round to the nearest integer; non-finite and out-of-range input raises an error. Test applications that depended on truncation or hidden errors.

Aggregate pushdown can bypass that integer-input validation: non-finite source values may survive in a remote aggregate. Set `mongo_fdw.enable_aggregate_pushdown` to `false` when local validation is required. Install the new binary and reconnect sessions; the SQL version stays 1.1. Verify row counts, identifiers and query plans against representative remote data after upgrading.
