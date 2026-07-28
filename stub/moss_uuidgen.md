## Usage

Sources:

- [Official upstream README](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/README)
- [Official extension control file (moss_uuidgen.control)](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/moss_uuidgen.control)
- [Official extension SQL (moss_uuidgen--1.0.sql)](https://github.com/openuruunofficial/moss/blob/ace66c80de4e3388b94db8f0f95b5a615bc92027/postgresql/moss_uuidgen--1.0.sql)

`moss_uuidgen` — If you are using >= 9.1, you need to create an extension instead: psql -U postgres moss -c "CREATE EXTENSION moss_uuidgen". Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION moss_uuidgen;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `uuid()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
