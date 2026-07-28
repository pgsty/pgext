## Usage

Sources:

- [Official extension control file](https://api.pgxn.org/src/openbarter/openbarter-0.8.2/openbarter.control)
- [Official project page](https://pgxn.org/dist/openbarter/)

`openbarter` — Multilateral agreement engine. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION openbarter;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `’0.8.0’`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
