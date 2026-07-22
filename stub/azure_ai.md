## Usage

Sources:

- [Azure AI extension overview](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-overview)
- [Azure AI extension functions](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-ai-functions)
- [Azure Database for PostgreSQL product page](https://azure.microsoft.com/en-us/products/postgresql/)

`azure_ai` is a Microsoft-managed extension for Azure Database for PostgreSQL Flexible Server. It provides SQL access to Azure OpenAI, Azure AI Language, and Azure Machine Learning through the `azure_ai`, `azure_openai`, `azure_cognitive`, and `azure_ml` schemas. It is a provider feature rather than a portable open-source extension, so availability depends on the Azure region and server configuration.

### Enable and Inspect

An Azure administrator must first allow `azure_ai` through the server's `azure.extensions` parameter. Then enable it separately in each database where it is needed.

```sql
CREATE EXTENSION IF NOT EXISTS azure_ai;
SELECT azure_ai.version();
```

Upgrade an existing installation with `ALTER EXTENSION azure_ai UPDATE` after confirming that the target Azure server exposes the intended version.

### Configure a Service

Use `azure_ai.set_setting` and `azure_ai.get_setting` to manage service endpoints and authentication. Only superusers and members of `azure_ai_settings_manager` can change these settings; `azure_pg_admin` receives that role by default.

Common setting keys include `azure_openai.endpoint`, `azure_openai.auth_type`, `azure_openai.subscription_key`, `azure_cognitive.endpoint`, `azure_cognitive.auth_type`, `azure_cognitive.subscription_key`, `azure_ml.scoring_endpoint`, `azure_ml.auth_type`, and `azure_ml.endpoint_key`. Prefer `managed-identity` authentication to `subscription-key` where the service supports it.

The extension includes embedding and model-call functions in its service schemas. The preview convenience functions `azure_ai.generate()`, `azure_ai.is_true()`, `azure_ai.extract()`, and `azure_ai.rank()` provide higher-level text operations.

### Operational Notes

- Model and scoring calls are synchronous outbound requests. They add network latency, can fail independently of PostgreSQL, incur provider cost, and are not undone by a database rollback.
- Treat prompts, source rows, embeddings, and model responses as data leaving the database boundary. Review data residency, private networking, access policy, and personally identifiable information before use.
- Avoid exposing stored subscription keys through broad `SELECT` access, query logs, or diagnostic output. Managed identity reduces secret distribution but still requires correctly scoped Azure permissions.
- Model output is nondeterministic and can be influenced by untrusted prompt content. Validate results before using them in security-sensitive or irreversible workflows.
- Catalog version `2.0.0` does not guarantee that every Azure region exposes the same surface; verify the server's reported version, supported models, quotas, and preview status.
