## Usage

Sources:

- [Official upstream README](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/README.md)
- [Official extension control file (numpy_plpython3u.control)](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/numpy_plpython3u.control)
- [Official extension SQL (numpy_plpython3u--1.0.sql)](https://github.com/tarkmeper/numpy_plpython/blob/acd66d925f205227acec6e87fcb5c49eae770abb/numpy_plpython3u--1.0.sql)

`numpy_plpython3u` — Library to transform Postgres Array's to numpy array's directly without going through python lists. Use it when database code must run in or interoperate with this procedural language. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION numpy_plpython3u;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `numpy_to_plpython3(val internal)` is an extension function and returns `internal`.
- `plpython3_to_numpy(val internal)` is an extension function and returns `real[]`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- Install the confirmed extension dependencies first: `plpython3u`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
