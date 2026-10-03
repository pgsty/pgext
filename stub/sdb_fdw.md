## Usage

Sources:

- [Official README](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/README.md)
- [Extension control file](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/sdb_fdw.control)
- [Installation SQL](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/sdb_fdw--1.0.sql)
- [Build configuration](https://github.com/SequoiaDB/SequoiaDB/blob/f14a3bccb4639e920a376bb65f20540a60206f1c/driver/postgresql/Makefile)

`sdb_fdw` maps SequoiaDB collections to PostgreSQL foreign tables. The authoritative connector documentation targets the historical SequoiaSQL/PostgreSQL 9.3.4 environment; current stock PostgreSQL compatibility is unverified.

### Core Workflow

After the matching connector and SequoiaDB client libraries are installed, an administrator can create the wrapper and map an existing collection.

```sql
CREATE EXTENSION sdb_fdw;
CREATE SERVER sdb_server FOREIGN DATA WRAPPER sdb_fdw
  OPTIONS (address 'localhost', service '11810');
CREATE FOREIGN TABLE records (a integer, b integer, c text)
  SERVER sdb_server OPTIONS (collectionspace 'cs', collection 'cl');
SELECT * FROM records;
```

### Mapping and Boundaries

Server options identify the coordinator address and service. Table options `collectionspace` and `collection` identify the remote collection. Quoted dotted column names can map nested document fields; compatible PostgreSQL arrays map array values. The adapter and its cryptographic/client dependencies are built within the SequoiaDB source tree, not supplied as an independent modern PGXS package. Review remote credentials and write semantics for the deployed SequoiaDB version; no cross-database atomicity guarantee is established here.
