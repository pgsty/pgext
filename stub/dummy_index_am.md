## Usage

Sources:

- [Official upstream README](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/README.md)
- [Official extension control file (dummy_index_am.control)](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/dummy_index_am.control)
- [Official implementation source](https://github.com/karajan1001/dummy_index_am/blob/5468e7325bd930af02152b0bc0c6f99eae5e64af/src/lib.rs)

`dummy_index_am` — Dummy index AM is a module for testing any facility usable by an index access method, whose code is kept a maximum simple. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION dummy_index_am;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
