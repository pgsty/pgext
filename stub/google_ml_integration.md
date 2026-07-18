## Usage

Sources:

- [Cloud SQL model endpoint management reference](https://docs.cloud.google.com/sql/docs/postgres/model-endpoint-management-reference)
- [Cloud SQL model registration guide](https://docs.cloud.google.com/sql/docs/postgres/model-endpoint-register-model)

`google_ml_integration` is a provider-managed Cloud SQL for PostgreSQL extension for registering remote model endpoints, storing Secret Manager references, generating embeddings, and invoking row-level predictions. Model endpoint management is currently a Preview feature and requires the `google_ml_integration.enable_model_support` database flag.

```sql
CREATE EXTENSION google_ml_integration;
SELECT * FROM google_ml.model_info_view;
SELECT google_ml.embedding(
  model_id => 'textembedding-gecko@001',
  contents => 'PostgreSQL extension metadata'
);
```

Register endpoints with `google_ml.create_model(...)`. Supported provider metadata includes Google, OpenAI, Anthropic, Hugging Face, and custom endpoints; non-Google credentials should be referenced through Secret Manager rather than stored as SQL literals. Restrict model, secret, embedding, and prediction routines to controlled roles.

Calls send database content to an external model provider. Classify data, enforce residency and egress policy, minimize prompts, redact secrets and personal data, and account for provider retention, cost, rate limits, latency, outages, and nondeterminism. Validate request/response transforms and vector dimension. Because the feature is Preview and provider-managed, verify exact Cloud SQL edition, region, engine version, extension version, IAM service-agent permissions, and upgrade behavior before production dependence.
