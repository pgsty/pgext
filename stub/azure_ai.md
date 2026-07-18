## Usage

Sources:

- [Azure AI extension overview](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-overview)
- [Azure Database for PostgreSQL product page](https://azure.microsoft.com/en-us/products/postgresql/)

`azure_ai` is a Microsoft-managed extension for Azure Database for PostgreSQL Flexible Server. It exposes the `azure_ai`, `azure_openai`, `azure_cognitive`, and `azure_ml` schemas for calling Azure AI, OpenAI, Cognitive Services, and Azure Machine Learning from SQL.

```sql
CREATE EXTENSION IF NOT EXISTS azure_ai;
SELECT azure_ai.version();
```

The server must allow the extension through the Azure `azure.extensions` setting. Configure service endpoints and authentication with `azure_ai.set_setting`; only superusers and members of `azure_ai_settings_manager` may manage those settings. Prefer managed identity over stored API keys. Functions such as `azure_ai.generate`, `azure_ai.is_true`, `azure_ai.extract`, and `azure_ai.rank` make outbound provider calls, so account for latency, cost, data residency, and prompt-data disclosure.

This is a provider-only component, not a portable open-source package. Availability and supported functions depend on the Azure region, server version, and enabled services.
