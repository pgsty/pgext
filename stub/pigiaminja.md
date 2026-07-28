## Usage

Sources:

- [Official upstream README](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/README.md)
- [Official extension control file (pigiaminja.control)](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/pigiaminja.control)
- [Official implementation source](https://github.com/leomos/pigiaminja/blob/a9b476b1405deb2926f57ac577ddf716a3f34b9b/src/lib.rs)

`pigiaminja` — A PostgreSQL extension that adds Jinja template format support to the COPY TO command. Use it when moving, transforming, or integrating the corresponding data from PostgreSQL. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pigiaminja;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
