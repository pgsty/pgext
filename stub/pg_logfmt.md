## Usage

Sources:

- [Official README](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/README.md)
- [SQL function implementation](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/src/lib.rs)
- [Parser implementation](https://github.com/clowder/pg_logfmt/blob/aaa2da2b71c6864264ad71b8d1d4a198d3fb5d9c/src/parser.rs)

`pg_logfmt` parses logfmt key-value pairs inside text. It is useful for converting application log lines to `jsonb`, extracting their keys as rows, or collecting the keys as a text array before indexing or analysis.

### Core Workflow

```sql
CREATE EXTENSION pg_logfmt;

SELECT logfmt_to_jsonb(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);

SELECT *
FROM logfmt_keys(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);

SELECT logfmt_keys_array(
    'at=info method=POST path="/orders/42" status=204 bytes=0'
);
```

`logfmt_to_jsonb(text)` returns an object whose values are strings or JSON null. `logfmt_keys(text)` returns a set of keys, while `logfmt_keys_array(text)` returns the same keys as an array. All three functions are declared immutable and parallel safe.

### Parsing Behavior

The parser accepts bare values and double-quoted values, including spaces inside quotes. It can skip a non-logfmt prefix and begin at the first recognizable `key=value` pair, which is useful for tagged application logs. An empty unquoted value becomes null; an explicitly quoted empty string remains an empty string. Double-escaped quotes in values are unescaped by `logfmt_to_jsonb`.

If no recognizable pair exists, `logfmt_to_jsonb` and `logfmt_keys_array` return NULL, while `logfmt_keys` returns no rows. Duplicate keys collapse to the last value when converted to `jsonb`, although the key-returning functions preserve parsed occurrences.

### Caveats

The reviewed upstream version is `0.0.0` and its Cargo feature matrix covers PostgreSQL `14` and `15`. It is a focused parser rather than a schema-on-read system: values are not converted to numbers, booleans, or timestamps. Validate edge cases such as malformed quotes, duplicate keys, unknown escapes, and mixed free text before depending on it for ingestion. No preload or restart is required.
