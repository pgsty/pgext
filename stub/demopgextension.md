## Usage

Sources:

- [Official upstream README](https://github.com/jwdeitch/pg_cmake_template/blob/6eafd52da9790db53fb32b669f77350b8476a1e5/readme)
- [Official extension control file (demopgextension.control)](https://github.com/jwdeitch/pg_cmake_template/blob/6eafd52da9790db53fb32b669f77350b8476a1e5/demopgextension.control)

`demopgextension` — This project provides a CMake template for developing postgres extensions, and is intended to replace PGXS. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION demopgextension;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
