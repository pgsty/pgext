## Usage

Sources:

- [Official extension control file](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/damm.control)
- [Official upstream documentation](https://github.com/fusiongyro/damm/blob/0829d6bc4e51f724a9186a55e26208bcf2152836/README.md)

`damm` — SQL implementation of the Damm check digit algorithm with a damm_code domain and helper functions.

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "damm";
SELECT extversion
FROM pg_extension
WHERE extname = 'damm';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
