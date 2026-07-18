## Usage

Sources:

- [Official extension control file](https://github.com/airpage-app/pg-utils/blob/2d56c14862dcf60d83cb79f1ebe0a80273d9e58d/packages/verify/skitch-extension-verify.control)

`skitch-extension-verify` — PL/pgSQL catalog-verification functions for checking PostgreSQL objects, roles, privileges, policies, and other schema properties.

The reviewed catalog snapshot records version `0.0.7`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `skitch-extension-utils`, `uuid-ossp`.

```sql
CREATE EXTENSION "skitch-extension-verify";
SELECT extversion
FROM pg_extension
WHERE extname = 'skitch-extension-verify';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
