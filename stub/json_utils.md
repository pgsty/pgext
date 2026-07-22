## Usage

Sources:

- [Official README at the catalog source revision](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/README.md)
- [Extension control file at the catalog source revision](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils.control)
- [Version 0.1 extension SQL](https://github.com/asotolongo/pg_json_utils/blob/3910a970506f1f9978f8925322e0a75e391b8962/json_utils--0.1.sql)
- [PGXN distribution page](https://pgxn.org/dist/json_utils/)

json_utils is an early pure-SQL extension for PostgreSQL's textual JSON type. It adds three comparison or containment-style operators and a hash operator class, but its own README warns that version 0.1 contains bugs and the implementation compares serialized text rather than JSON structure.

### Core Workflow

After installation, the operators work on JSON values and character strings. The following examples illustrate their literal behavior, not semantic JSON containment.

```sql
CREATE EXTENSION json_utils;

SELECT '{"a":1}'::json = '{"a":1}'::json;
SELECT '{"a":1,"b":2}'::json ? 'a'::varchar;
SELECT '{"a":1,"b":2}'::json @> '"b":2'::varchar;
```

### Installed Objects

- The equality function casts both documents to character strings and compares those strings.
- The key test searches for the exact serialized pattern consisting of a quoted key followed immediately by a colon.
- The containment-style test performs an unrestricted substring search.
- The hash support function derives an integer from the MD5 text of the serialized document, and the extension registers a default hash operator class for JSON.

### Safety and Compatibility

These operations are lexical, so whitespace, key order, formatting, and substrings can change results without changing—or while misrepresenting—the logical JSON value. The containment-style operator can match text inside strings or nested content and should not be treated as structural containment. All helper functions are declared volatile in the cataloged SQL.

Current PostgreSQL provides structural JSONB equality, containment, existence operators, and index support. Prefer those built-in JSONB facilities for new schemas. If maintaining a legacy installation, test every predicate with adversarial formatting and duplicate or nested keys before trusting it, and check for operator conflicts in the target database before creating the extension.

The extension is relocatable, does not require superuser installation according to its control file, and declares no preload or restart requirement.
