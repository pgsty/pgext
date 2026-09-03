## Usage

Sources:

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [jsonb_plphp control file](https://api.pgxn.org/src/plphp/plphp-2.6.0/jsonb_plphp/jsonb_plphp.control)
- [jsonb_plphp 1.0 SQL definitions](https://api.pgxn.org/src/plphp/plphp-2.6.0/jsonb_plphp/jsonb_plphp--1.0.sql)
- [PL/php language reference](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)

`jsonb_plphp` adds a PostgreSQL transform between `jsonb` and native PHP arrays, scalars, booleans, and nulls for functions written in `plphp`. It is a companion to the untrusted PL/php runtime, not an independent procedural language.

### Core Workflow

```sql
CREATE EXTENSION jsonb_plphp CASCADE;

CREATE FUNCTION add_seen(jsonb)
RETURNS jsonb
TRANSFORM FOR TYPE jsonb
LANGUAGE plphp
AS $$
    $value = $args[0];
    $value['seen'] = true;
    return $value;
$$;

SELECT add_seen('{"id": 1}'::jsonb);
```

`CASCADE` creates `plphp` when available. Only functions that declare `TRANSFORM FOR TYPE jsonb` receive native PHP values. The transform also applies in supported nested composite, array, set-returning, and trigger contexts; SPI result rows remain on their normal text-conversion path.

### Boundaries

`jsonb_plphp` requires the same superuser installation and trust model as `plphp`: PHP code can access the PostgreSQL server account's files, network, and process environment. JSON objects become associative PHP arrays, so review key coercion and the distinction among arrays, objects, numeric values, and nulls. Test large or deeply nested documents for memory use and conversion behavior.

The transform's control version is `1.0` and it is shipped in the PL/php 2.6.0 distribution. It requires no preload, but the PHP embed SAPI and the `plphp` shared library must already be usable on PostgreSQL 11 through 18.
