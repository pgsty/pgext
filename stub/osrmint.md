## Usage

Sources:

- [Official upstream README](https://github.com/fpacheco/osrmint/blob/c1937d4153cf8a38712e9e996e603564a4cd8e1e/README.md)
- [Official extension control file (osrmint.control)](https://github.com/fpacheco/osrmint/blob/c1937d4153cf8a38712e9e996e603564a4cd8e1e/control/osrmint.control)

`osrmint` — PostgreSQL route OSRM integration. Use it for the corresponding spatial data or geospatial workflow. Its extension dependencies must be installed and validated first.

### Core Workflow

```sql
CREATE EXTENSION osrmint;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1`.
- Install the confirmed extension dependencies first: `postgis`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
