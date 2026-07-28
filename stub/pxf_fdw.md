## Usage

Sources:

- [Official upstream README](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/README.md)
- [Official extension control file (pxf_fdw.control)](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/pxf_fdw.control)
- [Official extension SQL (pxf_fdw--1.0.sql)](https://github.com/greengagedb/pxf/blob/4cf00f3908fd63481666646fcc3bf9a5c3df68e6/fdw/pxf_fdw--1.0.sql)

`pxf_fdw` — This Greengage extension implements a Foreign Data Wrapper (FDW) for PXF. Use it when PostgreSQL must access the corresponding external data source through a foreign-data interface. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pxf_fdw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `pxf_fdw_handler()` is an extension function and returns `fdw_handler`.
- `pxf_fdw_validator(text[], oid)` is an extension function and returns `void`.
- `adl_pxf_fdw` is an extension-defined foreign data wrapper.
- `file_pxf_fdw` is an extension-defined foreign data wrapper.
- `gs_pxf_fdw` is an extension-defined foreign data wrapper.
- `hbase_pxf_fdw` is an extension-defined foreign data wrapper.
- `hdfs_pxf_fdw` is an extension-defined foreign data wrapper.
- `hive_pxf_fdw` is an extension-defined foreign data wrapper.
- `jdbc_pxf_fdw` is an extension-defined foreign data wrapper.
- `s3_pxf_fdw` is an extension-defined foreign data wrapper.
- `wasbs_pxf_fdw` is an extension-defined foreign data wrapper.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
