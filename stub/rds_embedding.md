## Usage

Sources:

- [Official Alibaba Cloud rds_embedding guide](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/rds-embedding)
- [Official ApsaraDB RDS for PostgreSQL overview](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/what-is-apsaradb-rds-for-postgresql)

`rds_embedding` is an Alibaba Cloud ApsaraDB RDS for PostgreSQL extension that calls configured external HTTP models to turn text into vectors. Version `1.0` is provider-only, requires PostgreSQL 14 or later plus `vector`, and requires outbound network access from the managed instance.

### Prerequisites and Setup

Use a privileged RDS account and enable the vector type before the embedding extension:

```sql
CREATE EXTENSION vector;
CREATE EXTENSION rds_embedding;
```

The exact minor engine must list the extension. Alibaba Cloud's guide also requires a model-service API key and NAT/SNAT configuration because RDS PostgreSQL has no Internet access by default.

### Core Workflow

Register a model with its request and response templates, then request an embedding:

```sql
SELECT rds_embedding.add_model(
  'text-embedding-v3',
  'https://dashscope-intl.aliyuncs.com/api/v1/services/embeddings/text-embedding/text-embedding',
  'Authorization: Bearer <API_KEY>',
  '{"input":{"texts":["%s"]},"model":"text-embedding-v3","parameters":{"text_type":"query"}}',
  '->''output''->''embeddings''->0->>''embedding'''
);

SELECT rds_embedding.get_embedding_by_model(
  'text-embedding-v3',
  '<API_KEY>',
  'Endless Yangtze River rolls on'
)::real[]::vector;
```

Store the result in a dimension-compatible `vector(n)` column and use the operators supplied by `vector` for similarity search.

### Important Objects

- `rds_embedding.add_model(...)`, `rds_embedding.update_model(...)`, and `rds_embedding.del_model(name)` manage rows in `rds_embedding.models`.
- `rds_embedding.show_models()` lists configured models.
- `rds_embedding.get_embedding_by_model(model, api_key, text)` calls the provider and extracts the configured embedding path.
- The provider lists `rds_embedding.get_response_by_model(...)`, but its documentation says that function is not yet available.

### Operational Caveats

Model calls are remote, synchronous operations inside database statements: latency, rate limits, DNS/NAT failure, provider outages, response-shape changes, and token cost directly affect SQL. The request template and examples contain authorization material; restrict access to `rds_embedding.models`, function execution, query logs, and monitoring output, and rotate exposed keys. Validate `%s` substitution and the extraction path against the exact model response, enforce vector dimensions, and do not retry writes blindly. The extension exists only on supported ApsaraDB RDS minor engines and is governed by the managed-service lifecycle rather than an open-source release.
