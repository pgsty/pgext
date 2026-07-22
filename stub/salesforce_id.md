## Usage

Sources:

- [Official README](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/README.md)
- [Official extension SQL](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/salesforce_id--1.0.sql)
- [Official extension control file](https://github.com/sylvainv/pg-salesforce-id/blob/1f821d8b491a09ed946349436d842d85dc75e9c5/salesforce_id.control)

`salesforce_id` provides a compact 12-byte PostgreSQL type for Salesforce object identifiers. It accepts the case-sensitive 15-character form and the case-insensitive 18-character form, normalizes output to 18 characters, and supplies comparison and index support for databases that store large numbers of Salesforce IDs.

### Core Workflow

Create the extension, use the type at the schema boundary, and confirm that the two textual representations compare as the same identifier:

```sql
CREATE EXTENSION salesforce_id;

SELECT '0012800000CXn0kAAD'::salesforce_id;
SELECT '0012800000CXn0kAAD'::salesforce_id
     = '0012800000CXn0k'::salesforce_id AS same_id;

CREATE TABLE sf_objects (
  id salesforce_id PRIMARY KEY
);
```

Input must use the Salesforce alphabet `0-9A-Za-z` and have exactly 15 or 18 characters. The type defines implicit casts to and from `text`, so review overloaded expressions and application bindings carefully rather than relying on silent coercion.

### Object Index

- `salesforce_id` is the fixed-width identifier type.
- `=` and `<>` test equality and inequality.
- `#<#`, `#<=#`, `#>#`, and `#>=#` provide ordering comparisons.
- `salesforce_id_ops` supplies default B-tree and hash operator classes.
- `gen_random_salesforce_id()` and `check_salesforce_id_internal()` are diagnostic helpers from the upstream SQL, not a substitute for identifiers issued by Salesforce.

### Caveats

The upstream README describes this release as experimental and not production-tested; the catalog also treats it as abandoned. Its stated storage saving is modest, and upstream does not establish a performance advantage. Validate canonical 18-character values at external boundaries, keep the original source identifier where reconciliation matters, and test dumps, restores, drivers, casts, sorting, and indexes on the exact PostgreSQL build before adopting this custom type.
