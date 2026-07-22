## 用法

来源：

- [Project status and deprecation notice](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/README.md)
- [PostgreSQL extension README](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/README.md)
- [Extension control file](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/sql/output/ai.control)
- [API-key and privilege guidance](https://github.com/timescale/pgai/tree/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/docs/security)

ai 是 pgai 中已归档的 PostgreSQL 扩展组件。它可以调用模型提供商、装载数据集、切分文本，并用于数据库内的 AI/RAG 实验。上游说明 pgai 自 2026 年 2 月起不再维护或提供支持，因此该扩展只适合维持或迁移已有部署。

### 核心流程

扩展固定安装在 `ai` 模式中，并依赖 `vector` 和不受信任语言 `plpython3u`。提供商函数在数据库后端进程内运行。以下最小示例调用本地 Ollama，从而避免放入托管提供商密钥：

```sql
CREATE EXTENSION IF NOT EXISTS ai CASCADE;

SELECT answer->>'response'
FROM ai.ollama_generate(
  'tinyllama',
  'Summarize why database transactions are useful.'
);
```

可使用 OpenAI、Anthropic、Cohere、Voyage AI、LiteLLM 或 Ollama 对应的提供商函数，并在构建检索流程时把返回的嵌入持久化到 `vector` 列。网络调用是同步的，而且外部副作用不会随外围 SQL 事务回滚。

### 重要对象

- `ai.ollama_generate` 向 Ollama 端点发送提示词并返回提供商响应。
- `ai.openai_embed`、`ai.cohere_embed` 及其他提供商函数族允许从 SQL 调用模型。
- `ai.load_dataset` 由 PostgreSQL 服务端进程导入 Hugging Face 数据集。
- `ai.grant_ai_usage` 是 security-definer 辅助函数，会授予较宽泛的模式、表、序列和函数权限；管理员模式还会授予转授权限。

### 安全与生命周期说明

该扩展会在 PostgreSQL 中执行不受信任的 Python，并可发起出站请求，因此只应让可信角色安装和执行其函数。密钥应放在提供商密钥存储或服务端进程环境变量中；上游警告直接传入密钥可能令其暴露在日志、SQL 文本或监控数据中。使用 `ai.grant_ai_usage` 前要审查其宽泛授权。目录中的开发版本和已归档的上游源码不代表仍有兼容性或安全修复；应规划把模型调用迁移到仍受维护的外部工作进程或库。
