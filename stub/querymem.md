## Usage

Sources:

- [Official upstream README](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/README.md)
- [Official extension control file (querymem.control)](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/querymem.control)
- [Official extension SQL (querymem--1.0.sql)](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/querymem--1.0.sql)

`querymem` — > Don't set work_mem too high; every query node can use that much memory. Use it when collecting or interpreting the corresponding PostgreSQL statistics. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION querymem;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `get_query_mem(text)` is an extension function and returns `INT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
