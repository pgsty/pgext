## Usage

Sources:

- [Installation SQL](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/mot_fdw--1.0.sql)
- [Extension control file](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/mot_fdw.control)
- [Build configuration](https://github.com/opengauss-mirror/openGauss-server/blob/9c01442e7d6811e5d3ab4c621d9153b6ce1af678/src/gausskernel/storage/mot/fdw_adapter/Makefile)
- [MOT table workflow in openGauss 3.0](https://docs.opengauss.org/en/docs/3.0.0/docs/Developerguide/creating-dropping-an-mot-table.html)

`mot_fdw` connects the openGauss SQL engine to its memory-optimized table engine. It requires an openGauss build with MOT support and does not provide a stock PostgreSQL in-memory table implementation.

### Enablement and Objects

Use the MOT server and foreign-table workflow documented for the deployed openGauss version. The installation SQL registers `mot_fdw_handler()`, `mot_fdw_validator(text[], oid)` and the `mot_fdw` wrapper. First verify that the kernel installed the extension and wrapper.

```sql
SELECT extname, extversion FROM pg_extension WHERE extname = 'mot_fdw';
SELECT srvname FROM pg_foreign_server
WHERE srvfdw = (SELECT oid FROM pg_foreign_data_wrapper WHERE fdwname = 'mot_fdw');
```

### Operational Boundary

MOT table memory, indexes, transaction integration, checkpointing and recovery are governed by the openGauss MOT engine, not by remote-server options of an ordinary network FDW. Provision memory and follow the kernel backup/recovery procedure before using it for durable data. Control version `1.0` does not identify a PostgreSQL major or a modern portable binary. The SQL uses kernel-specific language syntax; do not install these files into stock PostgreSQL.

### MOT Table Workflow

On a MOT-enabled deployment that supplies `mot_server`, the official 3.0 guide uses this table syntax; follow the deployed kernel's type and capacity limits:

```sql
CREATE FOREIGN TABLE mot_example (x integer) SERVER mot_server;
INSERT INTO mot_example VALUES (1);
SELECT * FROM mot_example;
```
