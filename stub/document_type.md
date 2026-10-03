## Usage

Sources:

- [Official document_type.control](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/document/document_type.control)
- [Official document_type--1.0.0.sql](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/src/postgres/document/document_type--1.0.0.sql)
- [Official README.md](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/README.md)
- [Official bugs.txt](https://github.com/danieltahara/sinew/blob/585f707f156020ab3d71cdff364c3e2400f2d701/documentation/bugs.txt)

`document_type` 1.0.0 is the document representation used by the historical Sinew research system. It stores structured key/value data and exposes typed accessors and mutators. The published source dates to 2014; compatibility with current PostgreSQL is unverified.

### Document Values

```sql
CREATE EXTENSION document_type;
SELECT document_get_int('{"n":7}'::document, 'n');
SELECT document_put_int('{"n":7}'::document, 'n', 8);
```

### API and Metadata

The `document` type has string input/output conversion. `document_get()` and `document_delete()` accept key/type information; typed get/put variants cover integer, float, boolean, text and nested documents. Mutators return a new document value, so callers must store that result to persist a change. `document_schema._attributes` tracks key names and types for the internal representation. `schema_analyzer` is the optional Sinew companion for key-frequency tracking.

### Research-System Limits

Installation requires administrator privileges and the matching C library; no preload is documented. Although the control declares relocation support, SQL creates the fixed document metadata schema. Upstream records incomplete loading/verification and query failures on malformed records. Treat this as preserved research software, not a maintained substitute for JSONB; validate the complete Sinew environment before handling important data.
