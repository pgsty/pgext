## Usage

Sources:

- [Official extension control file](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/italian_codes.control)
- [Official upstream documentation](https://github.com/dvarrazzo/italian_codes/blob/3d4d0604bfe12addf55d19683ed28d9a36187fb6/README.rst)
- [Official PGXN distribution page](https://pgxn.org/dist/italian_codes/)

`italian_codes` — Validation domains for Italian fiscal and VAT codes

The reviewed catalog snapshot records version `1.0`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "italian_codes";
SELECT extversion
FROM pg_extension
WHERE extname = 'italian_codes';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
