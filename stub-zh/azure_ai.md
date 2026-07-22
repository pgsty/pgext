## 用法

来源：

- [Azure AI 扩展概览](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-overview)
- [Azure AI 扩展函数](https://learn.microsoft.com/en-us/azure/postgresql/azure-ai/generative-ai-azure-ai-functions)
- [Azure Database for PostgreSQL 产品页](https://azure.microsoft.com/en-us/products/postgresql/)

`azure_ai` 是 Azure Database for PostgreSQL Flexible Server 上由 Microsoft 托管的扩展。它通过 `azure_ai`、`azure_openai`、`azure_cognitive` 和 `azure_ml` schema，从 SQL 访问 Azure OpenAI、Azure AI Language 与 Azure Machine Learning。它属于服务商功能，而不是可移植的开源扩展，因此实际可用性取决于 Azure 区域和服务器配置。

### 启用与检查

Azure 管理员必须先通过服务器参数 `azure.extensions` 允许 `azure_ai`，然后在每个需要使用它的数据库中分别启用。

```sql
CREATE EXTENSION IF NOT EXISTS azure_ai;
SELECT azure_ai.version();
```

确认目标 Azure 服务器提供预期版本后，可用 `ALTER EXTENSION azure_ai UPDATE` 升级已有安装。

### 配置服务

使用 `azure_ai.set_setting` 和 `azure_ai.get_setting` 管理服务端点与认证。只有超级用户和 `azure_ai_settings_manager` 成员可以修改这些设置；`azure_pg_admin` 默认获得该角色。

常用设置键包括 `azure_openai.endpoint`、`azure_openai.auth_type`、`azure_openai.subscription_key`、`azure_cognitive.endpoint`、`azure_cognitive.auth_type`、`azure_cognitive.subscription_key`、`azure_ml.scoring_endpoint`、`azure_ml.auth_type` 和 `azure_ml.endpoint_key`。在服务支持时，应优先采用 `managed-identity` 认证，而不是 `subscription-key`。

扩展在各服务 schema 中提供嵌入与模型调用函数。预览版便捷函数 `azure_ai.generate()`、`azure_ai.is_true()`、`azure_ai.extract()` 和 `azure_ai.rank()` 提供更高层的文本操作。

### 运维注意事项

- 模型与评分调用都是同步的外部请求。它们会增加网络延迟，可能独立于 PostgreSQL 失败，会产生服务商费用，而且不会随数据库回滚而撤销。
- 应把提示词、源数据行、嵌入和模型响应视为离开数据库边界的数据。使用前要评估数据驻留、私有网络、访问策略和个人身份信息。
- 不要通过宽泛的 `SELECT` 权限、查询日志或诊断输出暴露已保存的订阅密钥。托管身份可减少密钥分发，但仍需正确限定 Azure 权限。
- 模型输出具有非确定性，也可能受到不可信提示内容影响。将结果用于安全敏感或不可逆工作流前必须验证。
- 目录版本 `2.0.0` 不保证所有 Azure 区域提供相同接口；应核对服务器报告的版本、支持模型、配额和预览状态。
