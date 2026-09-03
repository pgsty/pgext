## Usage

Sources:

- [Official documentation](https://github.com/mrayva/pg_zerialize/blob/2aa98a11c2c51ee468336b702eebb571428f84d2/README.md)
- [Extension control file](https://github.com/mrayva/pg_zerialize/blob/2aa98a11c2c51ee468336b702eebb571428f84d2/pg_zerialize.control)
- [Official repository](https://github.com/mrayva/pg_zerialize)

`pg_zerialize` Binary serialization and deserialization for MessagePack, CBOR, ZERA, FlexBuffers, Ion, BSON, and BEVE.

### Enablement

Install the files for the intended server, then create `pg_zerialize` in the target database:

```sql
CREATE EXTENSION pg_zerialize;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
CREATE TYPE employee AS (id int, name text, active boolean);

SELECT msgpack_populate_record(NULL::employee, row_to_msgpack(ROW(1, 'Ada', true)::employee));
SELECT cbor_populate_record(ROW(1, 'Ada', true)::employee, cbor_build_object('active', false));
SELECT zera_populate_record(NULL::employee, zera_from_jsonb('{"id":2,"name":"Grace"}'::jsonb));
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 16, 17, 18; do not infer unlisted majors.
