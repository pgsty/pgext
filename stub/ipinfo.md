## Usage

Sources:

- [Official documentation](https://github.com/FranckPachot/pg_ipinfo/blob/9eb3e7163d3d5a87ef0645b94f6f553a2657ba43/README.md)
- [Extension control file](https://github.com/FranckPachot/pg_ipinfo/blob/9eb3e7163d3d5a87ef0645b94f6f553a2657ba43/ipinfo.control)
- [Official repository](https://github.com/FranckPachot/pg_ipinfo)

`ipinfo` IP geolocation lookup functions backed by the ipinfo.io HTTP service.

### Enablement

Install the files for the intended server, then create `ipinfo` in the target database:

```sql
CREATE EXTENSION ipinfo;
```

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
SELECT getipinfo();
```

### Main Objects

The reviewed install surface includes these prominent objects:

| Object | Kind | Operational role |
|---|---|---|
| `getipinfo` | FUNCTION | Callable function from the reviewed install surface. |

### Operations and Boundaries

- Catalog lifecycle is abandoned; test upgrades, dump/restore, and server compatibility before production use.
