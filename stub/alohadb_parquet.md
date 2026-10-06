## Usage

Sources:

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_parquet/alohadb_parquet--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet--1.0.sql)
- [contrib/alohadb_parquet/alohadb_parquet.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet.c)
- [contrib/alohadb_parquet/parquet_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/parquet_reader.c)
- [contrib/alohadb_parquet/csv_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/csv_reader.c)
- [contrib/alohadb_parquet/json_reader.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/json_reader.c)
- [contrib/alohadb_parquet/alohadb_parquet.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_parquet/alohadb_parquet.control)

`alohadb_parquet` reads CSV and JSON Lines into JSONB rows inside AlohaDB. In version 1.0, the Parquet entry point validates the file and returns metadata; it does not decode Parquet rows.

### Core Workflow

```sql
CREATE EXTENSION alohadb_parquet;
SET alohadb.parquet_allowed_paths = '/srv/import/';
SELECT * FROM read_csv('/srv/import/example.csv', ',', true);
SELECT * FROM read_json('/srv/import/example.jsonl');
```

### Operational Boundaries

`read_csv` takes a server-local path, delimiter and header flag. `read_json` reads JSON Lines. `read_parquet` checks magic/footer information and reports file metadata; full Parquet decoding is not implemented in this source.

Install as a superuser; no preload is required. `alohadb.parquet_allowed_paths` is a superuser-settable directory list, defaulting to /tmp. Reads use the PostgreSQL operating-system account. The path check is a string-prefix check and is not a complete isolation boundary: restrict function privileges and the server files it can reach. Validate file size, encoding and CSV parsing limits before feeding untrusted input.
