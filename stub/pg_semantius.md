## Usage

Sources:

- [Official documentation](https://github.com/semantius/semantius/blob/eb80b5f1f596ad2d206a6688bd362b5db3aa85d3/extension/README.md)
- [Extension control file](https://github.com/semantius/semantius/blob/eb80b5f1f596ad2d206a6688bd362b5db3aa85d3/extension/pg_semantius.control)
- [Official repository](https://github.com/semantius/semantius)

`pg_semantius` Database-first application backend with RBAC, RLS, data dictionary, and message queue.

### Enablement

Install the files for the intended server, then create `pg_semantius` in the target database:

```sql
CREATE EXTENSION pg_semantius CASCADE;
```

The reviewed control or official workflow requires `pgcrypto`. `CASCADE` only succeeds when those extension files are already installed on the server.

### Core Workflow

The following example is taken from the reviewed upstream documentation. Adapt object names, paths, credentials, and workload values before use.

```sql
BEGIN;
SET LOCAL ROLE authenticated;                              -- the identity role (member of semantius_user)
SELECT set_config('request.jwt.claims', $1::text, true);  -- LOCAL; inject BEFORE any rbac call
-- … queries …
COMMIT;
```

### Main Objects

The official sources define the extension surface dynamically or through provider tooling; inspect the installed version before granting access.

### Operations and Boundaries

- The verified PostgreSQL-major evidence covers 18; do not infer unlisted majors.
