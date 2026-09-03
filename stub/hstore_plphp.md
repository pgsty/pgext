## Usage

Sources:

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [hstore_plphp control file](https://api.pgxn.org/src/plphp/plphp-2.6.0/hstore_plphp/hstore_plphp.control)
- [hstore_plphp 1.0 SQL definitions](https://api.pgxn.org/src/plphp/plphp-2.6.0/hstore_plphp/hstore_plphp--1.0.sql)
- [PL/php language reference](https://pgxn.org/dist/plphp/2.6.0/doc/plphp.html)

`hstore_plphp` transforms PostgreSQL `hstore` values to and from native PHP associative arrays for functions written in `plphp`. PHP string keys and string-or-null values map to the corresponding `hstore` entries.

### Core Workflow

```sql
CREATE EXTENSION hstore_plphp CASCADE;

CREATE FUNCTION add_label(hstore)
RETURNS hstore
TRANSFORM FOR TYPE hstore
LANGUAGE plphp
AS $$
    $value = $args[0];
    $value['label'] = 'reviewed';
    return $value;
$$;

SELECT add_label('id=>1'::hstore);
```

`CASCADE` creates the required `hstore` and `plphp` extensions when available. Only functions that declare `TRANSFORM FOR TYPE hstore` use native PHP arrays. A PHP null value maps to an SQL NULL value inside `hstore`, not to the text string `NULL`.

### Boundaries

`hstore_plphp` inherits the untrusted, superuser-only security model of `plphp`; it does not sandbox PHP. The transform covers supported arguments, results, nested composite or array values, set-returning rows, and trigger rows, while SPI result rows keep normal text conversion. Test key/value encoding and null behavior before replacing explicit text parsing.

The transform's control version is `1.0` and it ships in PL/php 2.6.0. It needs no preload, but requires a working PHP embed SAPI and `plphp` on PostgreSQL 11 through 18.
