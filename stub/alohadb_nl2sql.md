## Usage

Sources:

- [Official upstream README](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/README)
- [Official extension control file (alohadb_nl2sql.control)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_nl2sql/alohadb_nl2sql.control)
- [Official extension SQL (alohadb_nl2sql--1.0.sql)](https://github.com/vern-allworks-llc/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_nl2sql/alohadb_nl2sql--1.0.sql)

`alohadb_nl2sql` — Natural language to SQL translation via LLM API. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION alohadb_nl2sql;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `alohadb_explain_query(sql text)` is an extension function and returns `text`.
- `alohadb_nl2sql(question text)` is an extension function and returns `text`.
- `alohadb_nl2sql_execute(question text)` is an extension function and returns `SETOF`.

### Requirements and Caveats

- The reviewed control file declares default version `1.0`.
- The control file marks the extension as relocatable.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
