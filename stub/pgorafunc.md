## Usage

Sources:

- [Official extension control file](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/pgorafunc.control)
- [Official upstream documentation](https://github.com/pgoracle/pgora-builtin/blob/d7f836d619031a5d180b5c8b48b68994ff8a555c/README.md)

`pgorafunc` — Oracle compatibility functions and packages for PostgreSQL

The reviewed catalog snapshot records version `9.4`, kind `standard`, and implementation language `C`.

```sql
CREATE EXTENSION "pgorafunc";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgorafunc';
```

The curated lifecycle is `abandoned`. Pin the reviewed build and verify maintenance status before adoption.
The official material contains an experimental, deprecated, unsupported, or explicit warning boundary; read it in full and test failure cases before non-lab use.

Before production use, review the linked control/SQL or provider documentation, verify privileges and compatibility, and test the actual API and failure behavior on the target PostgreSQL build.
