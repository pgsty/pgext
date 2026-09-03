## Usage

Sources:

- [Official documentation](https://github.com/ZhengYang/dc_fdw/blob/b362a57b8b7b934575688e1f428e8c5b3be2b861/README.md)
- [Extension control file](https://github.com/ZhengYang/dc_fdw/blob/b362a57b8b7b934575688e1f428e8c5b3be2b861/dc_fdw.control)
- [Official repository](https://github.com/ZhengYang/dc_fdw)

`dc_fdw` Historical foreign data wrapper for document collections stored on local disk.

### Enablement

Install the files for the intended server, then create `dc_fdw` in the target database:

```sql
CREATE EXTENSION dc_fdw;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE SERVER dc_server
  FOREIGN DATA WRAPPER dc_fdw;

CREATE FOREIGN TABLE documents (id integer, content text)
SERVER dc_server
OPTIONS (
  data_dir '/srv/documents',
  index_dir '/srv/documents-index',
  index_method 'SPIM',
  buffer_size '10',
  id_col 'id',
  text_col 'content'
);

SELECT id, content FROM documents LIMIT 10;
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `dc_fdw_handler` | FUNCTION | Callable function from the reviewed install surface. |
| `dc_fdw_validator` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 9.1, 9.2, 9.3; do not infer unlisted majors.
- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
