## Usage

Sources:

- [Official upstream README](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/README.md)
- [Official extension control file (myext.control)](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/myext.control)
- [Official extension SQL (myext--1.0--noop.sql)](https://github.com/diraneyya/pgext_noop_pattern/blob/6b464bcda27957cf356de5db9def8e279759cbd9/myext--1.0--noop.sql)

`myext` — After installing this demo extension, run these commands using psql, pgAdmin, or an IPython notebook using a SQL kernel:. Use it when administering or automating the database behavior described above. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION myext;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `myext_reload()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `noop`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
