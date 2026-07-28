## Usage

Sources:

- [Official upstream README](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/README.md)
- [Official extension control file (otlet.control)](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/crates/otlet_pg/otlet.control)
- [Official implementation source](https://github.com/joshmeek/otlet/blob/5bac507a37862fb305c3a9d5a731d90f146ce946/crates/otlet_pg/src/lib.rs)

`otlet` — Local LLM inference inside PostgreSQL with background workers, queued tasks, receipts, and semantic data operations. Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION otlet;

SELECT output
FROM otlet.ask(
  'qwen35_4b',
  'Summarize these customer notes in one sentence.',
  (SELECT jsonb_agg(to_jsonb(n))
   FROM customer_notes n WHERE customer = 'Riverline Labs'),
  '{"type":"object","required":["summary"],"additionalProperties":false,"properties":{"summary":{"type":"string"}}}'
);
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Requirements and Caveats

- The reviewed control file declares default version `0.1.0`.
- The control file marks the extension as non-relocatable.
- The control file does not require superuser-only installation.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
