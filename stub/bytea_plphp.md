## Usage

Sources:

- [PGXN 2.6.0 README](https://pgxn.org/dist/plphp/2.6.0/README.html)
- [bytea_plphp control file](https://api.pgxn.org/src/plphp/plphp-2.6.0/bytea_plphp/bytea_plphp.control)
- [bytea_plphp 1.0 SQL definitions](https://api.pgxn.org/src/plphp/plphp-2.6.0/bytea_plphp/bytea_plphp--1.0.sql)
- [PL/php 2.6.0 changelog](https://api.pgxn.org/src/plphp/plphp-2.6.0/CHANGELOG.md)

`bytea_plphp` transforms PostgreSQL `bytea` values to and from binary-safe PHP strings for functions written in `plphp`. Use it when code must preserve raw bytes, including embedded NULs, without passing through PostgreSQL's textual `\x...` representation.

### Core Workflow

```sql
CREATE EXTENSION bytea_plphp CASCADE;

CREATE FUNCTION bytea_roundtrip(bytea)
RETURNS bytea
TRANSFORM FOR TYPE bytea
LANGUAGE plphp
AS $$
    return $args[0];
$$;

SELECT encode(bytea_roundtrip(decode('000102ff', 'hex')), 'hex');
```

`CASCADE` creates `plphp` when available. Only functions that declare `TRANSFORM FOR TYPE bytea` receive the raw binary PHP string. The same transform can apply in supported nested composite, array, set-returning, and trigger contexts.

### Boundaries

Binary strings may contain NULs and invalid text encodings; do not pass them to PHP text APIs that assume UTF-8 or C-style termination. Returning a PHP string writes its bytes verbatim, so code that intends to decode hexadecimal or base64 must do that explicitly. Test large values for backend memory pressure and avoid logging binary payloads or secrets.

`bytea_plphp` inherits the untrusted, superuser-only security model of `plphp`. The transform's control version is `1.0`, introduced in the PL/php 2.6.0 distribution, and requires the PHP embed SAPI plus `plphp` on PostgreSQL 11 through 18. It needs no preload.
