## Usage

Sources:

- [Official upstream README](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/README.md)
- [Official extension control file (pg_ai_sdk.control)](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/pg_ai_sdk.control)
- [Official extension SQL (pg_ai_sdk--1.0.sql)](https://github.com/deeox/pg_ai_sdk/blob/24498354cb766ac274157d42127ff6702fa66649/pg_ai_sdk--1.0.sql)

`pg_ai_sdk` — This project is a PostgreSQL extension that translates natural language queries into SQL using the ClickHouse/ai-sdk-cpp. It allows users to query their database by asking questions in plain English. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_ai_sdk;

-- Generate SQL query from natural language
SELECT generate_sql_from_text('how many matches were won by Royal Challengers Bangalore is 2023');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `generate_sql_from_text(natural_language_query TEXT)` is an extension function and returns `TEXT`.
- `generate_sql_from_text(natural_language_query TEXT, model_name TEXT)` is an extension function and returns `TEXT`.
- `pg_ai_sdk_execute_json(natural_language_query text)` is an extension function and returns `text`.
- `pg_ai_sdk_execute_json(natural_language_query text, model_name TEXT)` is an extension function and returns `text`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
