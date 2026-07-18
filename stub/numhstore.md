## Usage

Sources:

- [Official extension control file](https://github.com/adjust/pg-numhstore/blob/754e2487940321f93344f1cb04244c8234097532/numhstore.control)
- [Official upstream documentation](https://github.com/adjust/pg-numhstore/blob/754e2487940321f93344f1cb04244c8234097532/README.md)

`numhstore` — Adds inthstore and floathstore numeric hstore types and helpers

The reviewed catalog snapshot records version `0.1.7`, kind `standard`, and implementation language `C`.
Install and validate the declared extension dependencies first: `hstore`.

```sql
CREATE EXTENSION "numhstore";
SELECT extversion
FROM pg_extension
WHERE extname = 'numhstore';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
