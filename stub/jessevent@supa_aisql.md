## Usage

Sources:

- [Official database.dev package page](https://database.dev/jessevent/supa_aisql)

`jessevent@supa_aisql` — AI inference SQL functions for Supabase. Self-deploys its own edge function via ai.deploy(). Use it for the corresponding vector, model, or retrieval workflow. Use the pinned upstream revision linked above as the API boundary and test it on the target PostgreSQL build.

### Core Workflow

```sql
CREATE EXTENSION "jessevent@supa_aisql";
```

Install the extension in the intended database, run the smallest upstream example above when available, and verify the installed version and returned values before integrating it into application SQL.

### Important Objects

- `ai._edge_function_source` is an extension function.
- `ai.deploy(project_ref text)` is an extension function and returns `jsonb`.
- `ai.status()` is an extension function and returns `jsonb`.
- `public._ai_call` is an extension function.
- `public._ai_extract_agg_ffunc` is an extension function.
- `public._ai_summarize_agg_ffunc` is an extension function.
- `public._ai_text_accum` is an extension function.
- `public.ai_classify(input_text text, categories text[] default array['positive','negative','neutral'], model_name text default null, provider text default 'openai')` is an extension function and returns `text`.
- `public.ai_complete(input_text text, system_text text default null, model_name text default null, provider text default 'openai')` is an extension function and returns `text`.
- `public.ai_embed(input_text text, model_name text default 'text-embedding-3-small', provider text default 'openai')` is an extension function and returns `real[]`.
- `public.ai_extract(input_text text, schema_hint text default null, model_name text default null, provider text default 'openai')` is an extension function and returns `jsonb`.
- `public.ai_redact(input_text text, entity_types text[] default array['email','phone','ssn','name'], model_name text default null, provider text default 'openai')` is an extension function and returns `text`.
- `public.ai_sentiment(input_text text, model_name text default null, provider text default 'openai')` is an extension function and returns `text`.
- `public.ai_translate(input_text text, target_language text default 'Spanish', model_name text default null, provider text default 'openai')` is an extension function and returns `text`.

### Requirements and Caveats

- The catalog records version `1.0.0`.
- This is a database.dev/pg_tle package; register or generate its package migration before issuing the quoted `CREATE EXTENSION` identity.
- Confirm privileges, supported PostgreSQL versions, upgrade behavior, and failure cases against the pinned source before production use.
