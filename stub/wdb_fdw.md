## Usage

Sources:

- [Upstream README](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/README.md)
- [Extension control file](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/wdb_fdw.control)
- [SQL installation script](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/sql/wdb_fdw.sql)
- [C implementation](https://github.com/Kentik-Archive/wdb_fdw/blob/fe6a1992449cdaeb6ec8f222daf720f4ae03a413/src/wdb_fdw.c)

`wdb_fdw` version `0.0.1` is a historical foreign data wrapper for an attached WhiteDB shared-memory database. A foreign server accepts `address` and `size` options, and foreign tables map their columns to WhiteDB record fields. The implementation advertises scan, insert, update, and delete callbacks.

### Historical example

```sql
CREATE EXTENSION wdb_fdw;
CREATE SERVER white_local FOREIGN DATA WRAPPER wdb_fdw
  OPTIONS (address '1000', size '1000000');
CREATE USER MAPPING FOR CURRENT_USER SERVER white_local;
CREATE FOREIGN TABLE white_items (key text, value text) SERVER white_local;
SELECT * FROM white_items;
```

The repository is archived and the code targets 2013-era PostgreSQL FDW APIs and WhiteDB. It supports only a narrow set of numeric and text conversions, has incomplete locking/error semantics, and cannot be assumed crash-safe or transactionally equivalent to PostgreSQL storage. Do not deploy it on a current production server. If historical data must be recovered, pin the matching legacy toolchain and work on disposable copies in an isolated environment.
