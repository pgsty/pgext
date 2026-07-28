## Usage

Sources:

- [Official upstream README](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/README.md)
- [Official extension control file (pg_gpt.control)](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/pg_gpt.control)
- [Official implementation source](https://github.com/cloudquery/pg_gpt/blob/1b1e6178e2f9dd0caca666e2abd6a4c428403ab0/src/lib.rs)

`pg_gpt` — Experimental PostgreSQL extension that enables the use of OpenAI GPT API inside PostgreSQL, allowing for queries to be written using natural language. Use it for the corresponding vector, model, or retrieval workflow. Upstream describes this capability as experimental.

### Core Workflow

```sql
CREATE EXTENSION pg_gpt;
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `gpt` is an extension function.
- `gpt_tables` is an extension function.

### Requirements and Caveats

- The catalog records version `0.0.0`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- Upstream labels part or all of the project experimental.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
