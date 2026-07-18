## Usage

Sources:

- [Alibaba Cloud rds_ai documentation](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/ai-rds-ai)
- [Alibaba Cloud RDS for PostgreSQL supported extensions](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`rds_ai` is an Alibaba Cloud extension for invoking large-language and embedding models from SQL on ApsaraDB RDS for PostgreSQL. It provides `rds_ai.prompt`, `rds_ai.embed`, `rds_ai.retrieve`, and `rds_ai.rag` for generation, vectorization, similarity retrieval, and retrieval-augmented generation, and can call configured HTTP-compatible custom models.

### Install, Configure, and Prompt

A privileged RDS account can install the provider plugin. `CASCADE` also installs its `pgvector` and `pgsql-http` dependencies.

```sql
CREATE EXTENSION IF NOT EXISTS rds_ai CASCADE;

SELECT rds_ai.update_model('qwen-plus', 'token', :'model_api_key');
SELECT rds_ai.update_model(
  'qwen-plus',
  'uri',
  'https://dashscope-intl.aliyuncs.com/api/v1/services/aigc/text-generation/generation'
);

SELECT rds_ai.prompt('Summarize the purpose of PostgreSQL WAL.');
```

Model configuration is stored in `rds_ai.model_list`. Default provider models include Qwen prompt models and `text-embedding-v3`; `rds_ai.add_model` can register another HTTP-compatible endpoint.

### Caveats

- This is an Alibaba Cloud provider extension, not a portable open-source package. Availability and the reported `1.0` or `1.0.0` version depend on RDS PostgreSQL major/minor engine version and edition.
- The provider requires a supported RDS engine, a privileged database account, and an active Model Studio API key. External model access also requires outbound VPC networking, commonly through NAT.
- Prompts, retrieved context, and embedding inputs are sent through HTTP to configured model endpoints. Apply data-classification, residency, logging, and egress controls before sending database content.
- API tokens and endpoint templates live in `rds_ai.model_list` and are set through SQL. Restrict table and function privileges, avoid literal secrets in query logs, and rotate exposed credentials.
- Preloading is not required for ordinary calls, but Alibaba Cloud documents adding `rds_ai` to `shared_preload_libraries` before changing `rds_ai.default_prompt_model`. That change requires the managed-instance parameter workflow.
- Dense and sparse embedding dimensions, default models, and supported minor versions can change with the service. Size vector columns from the live provider documentation rather than assuming catalog defaults.
