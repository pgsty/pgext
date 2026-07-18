## Usage

Sources:

- [Official extension control file](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/gdpr_anatomizer.control)
- [Official upstream documentation](https://github.com/pavilkau/postgresql_anatomy/blob/431772c17f1a7931414d8066b744150d41b00770/README.md)

`gdpr_anatomizer` — Prototype GDPR/anatomy anonymization extension implemented with PL/Python and PL/pgSQL helpers.

The reviewed catalog snapshot records version `0.0.1`, kind `puresql`, and implementation language `Python`.
Install and validate the declared extension dependencies first: `plpython3u`.
The curated compatibility set is `12`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "gdpr_anatomizer";
SELECT extversion
FROM pg_extension
WHERE extname = 'gdpr_anatomizer';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
