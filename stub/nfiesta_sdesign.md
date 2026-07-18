## Usage

Sources:

- [Official upstream source](https://gitlab.com/1on.cz/nfiesta/nfiesta_sdesign)

`nfiesta_sdesign` — PostgreSQL extension for manipulating sampling design data.

The reviewed catalog snapshot records version `1.1.17`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`, `postgis`.
The curated compatibility set is `16`; confirm the exact build against the target server.

```sql
CREATE EXTENSION "nfiesta_sdesign";
SELECT extversion
FROM pg_extension
WHERE extname = 'nfiesta_sdesign';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
