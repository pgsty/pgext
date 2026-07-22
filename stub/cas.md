## Usage

Sources:

- [Extension control file](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.control)
- [Version 0.1.0 extension SQL](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas--0.1.0.sql)
- [C implementation](https://github.com/philopon/cas/blob/e23971e2a65e5c9e50fa238cce9bb4209f041b36/cas.c)

`cas` defines a fixed-width PostgreSQL type for Chemical Abstracts Service register-number-shaped values. It parses three numeric components separated by hyphens, exposes each component, and supplies comparison operators plus a default B-tree operator class. It stores the supplied check digit but does not validate the CAS checksum algorithm.

### Core Workflow

After confirming that the package installs on the target PostgreSQL build:

```sql
CREATE EXTENSION cas;

CREATE TABLE substance (
    registry_number cas PRIMARY KEY,
    name text NOT NULL
);

INSERT INTO substance VALUES ('57-09-0', 'example');

SELECT registry_number,
       cas_first(registry_number),
       cas_second(registry_number),
       cas_checkdigit(registry_number)
FROM substance;
```

Values sort numerically by the first component, then the second component, then the stored check digit.

### Object Index

- `cas` is the 12-byte scalar type with `first-second-check` text I/O.
- `cas_first(cas)`, `cas_second(cas)`, and `cas_checkdigit(cas)` return the stored components.
- Operators `=`, `!=`, `<`, `<=`, `>`, and `>=` compare component tuples.
- `cas_ops` is the default B-tree operator class.

### Operational Notes

Version `0.1.0` is relocatable and declares no preload requirement. The parser enforces only the numeric, hyphenated shape; invalid real-world check digits can still be stored. Validate registry numbers separately when checksum correctness matters.

The pinned installation SQL declares functions that reference `cas` before defining the shell type normally required by PostgreSQL's base-type creation sequence. Test `CREATE EXTENSION cas` on a disposable database before rollout; a package may need a corrected extension script for a fresh installation. The repository provides no README or release documentation, so do not infer casts, hash indexing, or compatibility beyond the published control, SQL, C source, and regression files.
