## Usage

Sources:

- [Official extension control file](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/postgresml.control)
- [Official upstream documentation](https://github.com/yieldsfalsehood/postgresml/blob/35d3d10f3de964eefb6d27204d4b55d8d0359d69/README.md)

`postgresml` — Exposes tsvector lexemes and position counts for an experimental naive Bayes text-classification demo.

The reviewed catalog snapshot records version `1.0`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "postgresml";
SELECT extversion
FROM pg_extension
WHERE extname = 'postgresml';
```

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
