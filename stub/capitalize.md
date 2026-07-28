## Usage

Sources:

- [Official extension control file (capitalize.control)](https://github.com/darshan744/postgres-extension/blob/2a22caeecee3505a327ff37b3752f51bbb1c017a/CapitalizeExtension/capitalize.control)
- [Official extension SQL (capitalize--1.0.sql)](https://github.com/darshan744/postgres-extension/blob/2a22caeecee3505a327ff37b3752f51bbb1c017a/CapitalizeExtension/capitalize--1.0.sql)

`capitalize` — A function to capitalize all the given character. Use it when SQL needs these specialized functions or aggregates. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION capitalize;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `capitalize(text)` is an extension function and returns `TEXT`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
