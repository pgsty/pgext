## 用法

来源：

- [Azure AI 扩展概览](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-overview)
- [Azure Database for PostgreSQL 产品页](https://azure.microsoft.com/en-us/products/postgresql/)

`azure_ai` 是 Azure Database for PostgreSQL Flexible Server 上由 Microsoft 托管的扩展。它提供 `azure_ai`、`azure_openai`、`azure_cognitive` 和 `azure_ml` schema，用于从 SQL 调用 Azure AI、OpenAI、Cognitive Services 与 Azure Machine Learning。

```sql
CREATE EXTENSION IF NOT EXISTS azure_ai;
SELECT azure_ai.version();
```

服务器必须通过 Azure 的 `azure.extensions` 设置允许该扩展。服务端点与认证信息使用 `azure_ai.set_setting` 配置；只有超级用户和 `azure_ai_settings_manager` 成员可以管理这些设置。应优先使用托管身份，而不是保存 API 密钥。`azure_ai.generate`、`azure_ai.is_true`、`azure_ai.extract` 和 `azure_ai.rank` 等函数会发起外部服务调用，需评估延迟、费用、数据驻留与提示数据泄露风险。

这是供应商专有组件，并非可移植的开源软件包。实际可用性和函数范围取决于 Azure 区域、服务器版本及已启用服务。
