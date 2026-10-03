## Usage

Sources:

- [18.2.0 README](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/README.md)
- [18.2.0 control file](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/db2_fdw.control)
- [18.2.0 SQL API](https://github.com/pg-fdw/db2_fdw/blob/18.2.0/sql/db2_fdw--18.2.0.sql)

`db2_fdw` 18.2.0 queries and modifies IBM Db2 tables through PostgreSQL foreign tables. It pushes down supported filters and only the required columns. Upstream requires PostgreSQL 10.1 or later and an IBM Db2 client 11.1 or later with the same architecture as PostgreSQL. The server process must be able to load that client and access the configured database; installing this extension does not supply the external client or a Db2 server.

### Connect and Import Tables

The example assumes that the Db2 client can connect to the catalogued SAMPLE database and that DB2INST1.EMPLOYEE exists. Install the extension as a superuser, then create the server and a role-specific mapping. Use the actual Db2 credentials instead of the example password:

```sql
CREATE EXTENSION db2_fdw;
CREATE SERVER db2srv FOREIGN DATA WRAPPER db2_fdw
  OPTIONS (dbserver 'SAMPLE');
CREATE USER MAPPING FOR CURRENT_USER SERVER db2srv
  OPTIONS (user 'db2inst1', password 'change-me');
CREATE SCHEMA db2_remote;
IMPORT FOREIGN SCHEMA "DB2INST1" LIMIT TO ("EMPLOYEE")
  FROM SERVER db2srv INTO db2_remote;
SELECT empno, firstname, lastname, salary
FROM db2_remote.employee
WHERE empno = '000010';
```

Importing obtains the Db2 column metadata required by the wrapper. Default smart case folding lowercases all-uppercase names. A separate application role needs USAGE on the foreign server, its own user mapping, and appropriate schema/table privileges. Avoid a PUBLIC mapping when credentials are meant for one role. Empty user/password strings select the upstream external-authentication path and depend on the Db2 client environment.

### Options and Writes

- Server `dbserver` selects the Db2 connection; `no_encoding_error` controls handling of encoding-conversion errors. `batch_size` is reserved in this release and should not be treated as working batch insertion.
- Table `schema` and `table` identify the remote object; `readonly` prohibits modifications. `prefetch` defaults to 100 and accepts 0–1024; `fetch_size` is accepted but currently fixed to 1. `sample_percent` controls ANALYZE sampling.
- Column `key` must identify every remote primary-key column for `UPDATE` and `DELETE`. Imported metadata includes `db2type`, `db2size`, `db2bytes`, `db2chars`, `db2scale`, `db2null` and `db2ccsid`; preserve it when altering imported tables.
- IMPORT FOREIGN SCHEMA options `case` and `readonly` control name folding and whether imported tables permit writes.

INSERT, UPDATE and DELETE additionally require Db2-side privileges. Review column mapping and remote key definitions before enabling writes; PostgreSQL declarations alone do not create a remote primary key. Unsupported filters are evaluated locally, so inspect EXPLAIN before assuming pushdown.

### Diagnostics and Transactions

```sql
SELECT db2_diag();
SELECT db2_diag('db2srv');
SELECT db2_close_connections();
```

`db2_diag()` reports local client/build diagnostics and optionally remote-server details. `db2_close_connections()` closes connections cached by the current backend; do not call it within a transaction that has modified Db2 data. Long-lived PostgreSQL sessions can retain remote connections and transaction resources.

### Types and Maintenance

Common mappings include character types to text/character, BLOB to bytea, integer types to PostgreSQL integer types, and DATE/TIMESTAMP/TIME to their corresponding types. The declared PostgreSQL type and width must accommodate the remote values; failed conversions surface at query time. The control version is 18.2.0 and is relocatable. Normal use requires no shared preload. Follow the versioned README for IBM client environment and connectivity requirements, and distinguish a package upgrade from validation against the actual external database.
