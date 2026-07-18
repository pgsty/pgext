## Usage

Sources:

- [Official upstream documentation](https://github.com/guedes/validadores/blob/63ea247a692f6995c54a24e80f1a39802ede59b6/doc/validadores.md)
- [Official PGXN distribution page](https://pgxn.org/dist/validadores/)

`validadores` — Brazilian CPF check-digit validator, operator, and domain

The reviewed catalog snapshot records version `0.0.2`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "validadores";
SELECT extversion
FROM pg_extension
WHERE extname = 'validadores';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
