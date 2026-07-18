## Usage

Sources:

- [Pinned upstream README](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/README.md)
- [Version 0.1 installation SQL](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/jsonfun--0.1.sql)
- [Pinned JSON parser implementation](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/json_extract_keys_array.c)
- [Pinned extension control file](https://github.com/everilae/jsonfun/blob/fb658f71be3f7f9810bd613ae0cb4e3a5bc1af89/jsonfun.control)

jsonfun version 0.1 exposes two C functions that extract several object paths from a json value into one text array. The default function uses a dot separator; the second accepts a delimiter. If a selected value is a JSON array, its immediate elements are appended to the result.

### Extraction example

```sql
CREATE EXTENSION jsonfun;

SELECT json_extract_keys_array(
    '{"a":{"b":"hello"},"c":{"d":[1,2,null]}}'::json,
    'a.b',
    'c.d',
    'missing.path'
);
```

The result is a text array containing hello, 1, and 2. Missing paths and JSON null values are ignored. Paths traverse objects only: array indexes in a path are unsupported, even though a terminal array is unpacked.

### Compatibility and input limits

The implementation copies and uses PostgreSQL's internal JSON lexer and semantic-action interfaces from 2014. Those are not a stable extension API, and the repository has no compatibility matrix for modern majors. It accepts json, not jsonb, and reparses the complete document once for every requested path, so work grows with both document size and key count.

The delimiter variant converts the delimiter text to a C string but uses only its first byte; its own source warns that unusual multibyte encodings behave badly. Results are normalized to text, so JSON type distinctions are lost. Deep input, many keys, large terminal arrays, duplicate paths, empty path components, multibyte delimiters, and current server headers all need regression tests.

The repository has not changed since 2014. Treat jsonfun as legacy source code; prefer current jsonb operators and SQL/JSON path features unless exact historical behavior is required.
