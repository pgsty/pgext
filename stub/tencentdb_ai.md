## Usage

Sources:

- [Tencent Cloud tencentdb_ai function documentation](https://cloud.tencent.com/document/product/409/116227)
- [TencentDB for PostgreSQL supported extension versions](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_ai` is a TencentDB for PostgreSQL managed extension that calls network-reachable model APIs from SQL. Version 1.3 manages model definitions and credentials and provides generic calls plus chat, embedding, reranking, and image-to-text helpers. Requests leave the database and can incur provider cost, latency, rate limits, and data-governance obligations.

### Register and Call a Model

On a supported TencentDB instance:

```sql
CREATE EXTENSION tencentdb_ai CASCADE;

SELECT tencentdb_ai.add_model(
  'hunyuan-lite',
  '2023-09-01',
  NULL,
  NULL
);

SELECT * FROM tencentdb_ai.list_models();

SELECT tencentdb_ai.chat_completions(
  'hunyuan-lite',
  'Summarize the current transaction isolation level.'
);

SELECT tencentdb_ai.get_embedding(
  'hunyuan-embedding',
  ARRAY['PostgreSQL', 'vector search']
);
```

Creation also installs `pgcrypto`. Before calling a registered model, use `update_model_attr(model_name, attr_name, attr_value)` to configure the provider's `SecretId` and `SecretKey`, plus `version`, `region`, or `json_path` when required. Do not embed credentials in reusable SQL, application logs, or shared shell history; tightly restrict function and model-catalog access.

### Function Index

`add_model` registers a name, version, region, and optional JSONPath used like `jsonb_path_query_first` to select a response field. `list_models`, `update_model_attr`, and `delete_model` manage registrations. `call_model(model_name, common_params, api_params)` exposes the generic text-array request form. `chat_completions`, `get_embedding`, `run_rerank`, and `image_to_text` provide fixed-scene interfaces.

### Service, Version, and Security Boundaries

The detailed function page states PostgreSQL 17 as its prerequisite, while the current provider extension matrix lists `tencentdb_ai` 1.3 for PostgreSQL 14–18. Its older walkthrough output also shows an extension version 1.0. Treat the instance's `pg_available_extension_versions` and current Tencent support response as authoritative, and query `pg_extension` after installation.

Calls are synchronous external network operations and are not made transactional by SQL rollback. Prompts, documents, image URLs, and returned content may contain sensitive data or untrusted model output. Apply egress controls, least-privilege Tencent credentials, time/cost budgets, output validation, retry/idempotency rules, and explicit consent/retention policy before using results in database writes or automated decisions.
