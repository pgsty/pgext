## Usage

Sources:

- [Official upstream README](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/README.md)
- [Official extension control file (doom.control)](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/doom.control)
- [Official extension SQL (doom--0.0.1.sql)](https://github.com/marvinirwin/terminal-doom/blob/053d773e13c9f8bc9c2b9fb1b09f7da1c549640c/build/doom--0.0.1.sql)

`doom` — I'm going to try and make it run as a postgres extension using tables for input and output. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION doom;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `doom()` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `0.0.1`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
