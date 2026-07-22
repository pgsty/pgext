## Usage

Sources:

- [Extension control file](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion.control)
- [Version 1.0 installation SQL](https://github.com/JeanChristopheMorinPerso/data_experiments/blob/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/condaversion--1.0.sql)
- [Upstream regression tests](https://github.com/JeanChristopheMorinPerso/data_experiments/tree/2e7140763ab395fe22a2c66b083bb34c51f23a93/conda_pgsql_ext/tests/sql)

`condaversion` version `1.0` adds a Conda-aware scalar type implemented in C++ with libmamba version parsing. It is useful for sorting and indexing Conda package versions whose precedence differs from lexical text order.

### Core Workflow

```sql
CREATE EXTENSION condaversion;

CREATE TABLE conda_release (
    name text NOT NULL,
    version condaversion NOT NULL
);

INSERT INTO conda_release VALUES
    ('demo', '1.0'),
    ('demo', '1.0.0'),
    ('demo', '1.1a1');

SELECT version::text
FROM conda_release
ORDER BY version;

CREATE INDEX conda_release_version_idx
    ON conda_release (version);
```

The extension defines input/output and binary send/receive functions, comparison operators `=`, `!=`, `<`, `>`, `<=`, and `>=`, implicit or assignment casts involving text types, and default B-tree and hash operator classes. The upstream tests explicitly demonstrate that `1.0` and `1.0.0` compare as equal rather than as different strings.

### Operational Notes

The control file is relocatable and declares no preload dependency. Building the module requires the upstream C++/libmamba environment; the SQL extension alone is not sufficient without its matching shared library. Validate dump/restore, index behavior, and ABI compatibility on the target PostgreSQL and libmamba versions before production use.
