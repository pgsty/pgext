## Usage

Sources:

- [Upstream README](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/README.md)
- [Extension SQL](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils--0.1.sql)
- [Extension control file](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils.control)
- [PGXN distribution](https://pgxn.org/dist/json_utils/)

`json_utils` is a SQL/PLpgSQL extension for PostgreSQL's legacy `json` type. It adds key-existence, key/value containment, and equality operators:

```sql
CREATE EXTENSION json_utils;

SELECT '{"field1":"value1","field2":341}'::json ? 'field2';
SELECT '{"field1":"value1","field2":341}'::json @> '"field2":341';
SELECT '{"field1":"value1"}'::json = '{"field1":"value1"}'::json;
```

### Version and caveats

The PGXN package is labeled `0.1.1`, while its control file and install script expose extension version `0.1`. The implementation uses textual matching rather than PostgreSQL's structural `jsonb` semantics, and the upstream README explicitly warns that the published version has known bugs. Prefer built-in `jsonb` operators for new applications; use `json_utils` only where its legacy behavior has been tested against representative inputs.
