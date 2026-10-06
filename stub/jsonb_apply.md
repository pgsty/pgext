## Usage

Sources:

- [README.md](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/README.md)
- [sql/jsonb_apply--0.1.0.sql](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/sql/jsonb_apply--0.1.0.sql)
- [jsonb_apply.control](https://github.com/Florents-Tselai/jsonb_apply/blob/206fcae653f52ebc3a55310cafe7eceb1841ca15/jsonb_apply.control)

`jsonb_apply` recursively transforms JSONB strings by invoking a named PostgreSQL text function. It preserves the document structure and non-string values.

### Core Workflow

```sql
CREATE EXTENSION jsonb_apply;
SELECT jsonb_apply('{"name":"Alice","tags":["hello"]}'::jsonb, 'upper');
SELECT jsonb_apply('{"message":"hello"}'::jsonb, 'replace', 'hello', 'bye');
```

### Operational Boundaries

The function takes a JSONB document, a function name and optional variadic arguments. The called function must accept text as its first argument and return text; its other argument types participate in overload resolution. Use reviewed function names and appropriate schema visibility, since this invokes SQL functions rather than a fixed list of transformations.

Install the C library and create the extension as a superuser. No preload or restart is specified. The reviewed source publishes neither a PostgreSQL-major support matrix nor a license declaration; the catalog does not assert either.
