## Usage

Sources:

- [Extension control file](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.control)
- [Version 1.0 SQL API](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw--1.0.sql)
- [Official foreign-table example](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/example.sql)
- [FDW implementation](https://github.com/goldenhelix/tsf-fdw/blob/bf8d839bf447d3300cd371b9e059b612b9330923/tsf_fdw.c)

`tsf_fdw` version `1.0` is a Golden Helix foreign data wrapper for reading fields from local TSF genomic data files as PostgreSQL foreign tables. It implements scans only and is not a writer for TSF files.

### Core Workflow

```sql
CREATE EXTENSION tsf_fdw;
CREATE SERVER tsf_server FOREIGN DATA WRAPPER tsf_fdw;

CREATE FOREIGN TABLE low_level (
  _id integer,
  Chr text,
  Start integer,
  Stop integer,
  FloatField real,
  StringField text
) SERVER tsf_server
  OPTIONS (filename '/data/Low_level.tsf', sourceid '1');

SELECT Chr, Start, Stop FROM low_level LIMIT 20;
```

The PostgreSQL server process must be able to read the file path. Foreign-table options include `filename`, optional base `path`, `sourceid`, and `fieldtype`; per-column options can override the file/source and select `fieldidx` or mapping sources. Match column names and types to the TSF source exactly.

For discovery, `tsf_generate_schemas(prefix, tsf_path, source_id)` returns generated DDL for one source or all sources. Review its output before executing it, especially the generated server name and absolute paths.

### Query and Operational Boundaries

The FDW skips unreferenced fields and implements local restriction processing against TSF iterators, but it does not provide write callbacks. File contents, source IDs, field kinds, and mappings must remain consistent with the foreign-table definition.

This project has no substantive upstream README and depends on the bundled TSF C library. Validate file-format compatibility, permissions, concurrent access, array/mapping fields, query plans, and corrupt-file errors with representative Golden Helix data. Back up the source files independently; PostgreSQL backup does not include external TSF contents.
