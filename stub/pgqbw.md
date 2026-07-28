## Usage

Sources:

- [Official extension control file (pgqbw.control)](https://github.com/rekgrpth/pgqbw/blob/a104ef679aa83fc16913ce4b9dea3b4cf98259b3/pgqbw.control)
- [Official implementation source](https://github.com/rekgrpth/pgqbw/blob/a104ef679aa83fc16913ce4b9dea3b4cf98259b3/src/pgqbw.c)

`pgqbw` — postgres queue backgound worker. Use it when an application needs this specific database capability. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pgqbw;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as non-relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
