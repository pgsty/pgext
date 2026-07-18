## Usage

Sources:

- [Pinned upstream README](https://github.com/durch/google-bigtable-postgres-fdw/blob/8337808ab544a202755f38f506cde8def8bfcabb/README.md)
- [Pinned extension control file](https://github.com/durch/google-bigtable-postgres-fdw/blob/8337808ab544a202755f38f506cde8def8bfcabb/bigtable.control)

`bigtable` is a Rust foreign data wrapper for Google Cloud Bigtable. The reviewed implementation represents a remote row as one PostgreSQL `json` value and configures the cloud instance, project, table, and service-account credential path through foreign-server, foreign-table, and user-mapping options.

```sql
CREATE EXTENSION bigtable;

CREATE SERVER bigtable_srv
  FOREIGN DATA WRAPPER bigtable
  OPTIONS (instance 'instance_id', project 'project_id');

CREATE FOREIGN TABLE bt_rows (bt json)
  SERVER bigtable_srv
  OPTIONS (name 'table_name');

CREATE USER MAPPING FOR CURRENT_USER
  SERVER bigtable_srv
  OPTIONS (credentials_path '/secure/service-account.json');

SELECT bt->>'rowKey' FROM bt_rows LIMIT 10;
```

The pinned README documents `SELECT`, `SELECT LIMIT`, and `INSERT`. Its roadmap leaves general `WHERE`, `OFFSET`, `DELETE`, and native `UPDATE` incomplete; inserting an existing key is described as the update mechanism. Treat predicate behavior as implementation-specific and verify whether filtering is pushed down or performed locally.

The reviewed control version is `0.1.0`, and the upstream setup targets the much older PostgreSQL 9.6/Rust 1.16 era. Keep the credential file outside database-readable locations, restrict its filesystem permissions and the user mapping, test Bigtable schema/column-family assumptions, and validate current API, TLS, retry, quota, and transaction semantics before use.
