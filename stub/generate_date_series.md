## Usage

Sources:

- [Official extension control file](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/generate_date_series.control)
- [Official upstream documentation](https://github.com/corlogic-code/generate_date_series/blob/3cfe1013764521fbd6e4fc300f20208ff1f709b4/doc/generate_date_series.md)
- [Official PGXN distribution page](https://pgxn.org/dist/generate_date_series/)

`generate_date_series` — generate_series()-style set-returning function for date ranges with integer day steps.

The reviewed catalog snapshot records version `1.0.0`, kind `standard`, and implementation language `C`.
The curated compatibility set is `10,11,12,13,14,15,16,17,18`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "generate_date_series";
SELECT extversion
FROM pg_extension
WHERE extname = 'generate_date_series';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
