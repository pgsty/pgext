## Usage

Sources:

- [Official upstream documentation](https://database.dev/raminder/countries)

`countries` — Pure-SQL package that creates and seeds a public.countries reference table.

The reviewed catalog snapshot records version `0.0.4`, kind `puresql`, and implementation language `SQL`.

```sql
CREATE EXTENSION "countries";
SELECT extversion
FROM pg_extension
WHERE extname = 'countries';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
