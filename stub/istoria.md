## Usage

Sources:

- [Official extension control file](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria.control)
- [Official upstream documentation](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/README)

`istoria` — Table history management with non-linear undo/redo

The reviewed catalog snapshot records version `1.1`, kind `puresql`, and implementation language `SQL`.
Install and validate the declared extension dependencies first: `plpgsql`.

```sql
CREATE EXTENSION "istoria";
SELECT extversion
FROM pg_extension
WHERE extname = 'istoria';
```

The curated lifecycle is `active`. Pin the reviewed build and verify maintenance status before adoption.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
