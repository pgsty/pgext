## Usage

Sources:

- [Official upstream README](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/README.md)
- [Official extension control file (pg_synapse_pgrx.control)](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/crates/pg-synapse-pgrx/pg_synapse_pgrx.control)
- [Official implementation source](https://github.com/yonk-labs/pg-synapse/blob/b8adaba13b1ff7c91777aee7473dd37045db99a0/crates/pg-synapse-pgrx/src/lib.rs)

`pg_synapse_pgrx` — pg synapse: Postgres-native agent-loop runtime. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION pg_synapse_pgrx;

SELECT synapse.execute('notes_agent', 'Add a note that says "Hello"');
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The catalog records version `0.1.1`.
- The control file marks the extension as non-relocatable.
- The control file requires a superuser for installation.
- The control file marks the extension as not trusted.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
