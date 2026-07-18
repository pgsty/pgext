## Usage

Sources:

- [Official extension control file](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/pgotel.control)
- [Official upstream documentation](https://github.com/ongres/pgotel/blob/c1789946ae93169fd8a9e1a1ea8ac4f7d19e804e/README.md)

`pgotel` — PostgreSQL extension exposing OpenTelemetry functionality for other extensions.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C++`.

```sql
CREATE EXTENSION "pgotel";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgotel';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
