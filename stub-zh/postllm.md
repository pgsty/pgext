## 用法

来源：

- [Official extension control file](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/postllm.control)
- [Official upstream documentation](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/README.md)
- [Official API reference](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/docs/reference.md)
- [Official configuration guide](https://github.com/mrcsparker/postllm/blob/0f7d67b574b9d676efaf1024dc201fa563b2c111/docs/configuration.md)

`postllm` 0.1.0 让 PostgreSQL SQL 能调用托管或本地 LLM 工作流。它涵盖聊天与补全、嵌入、分块、检索和重排，以及配置档、密钥、角色权限、模型别名与出站主机策略。推理会同步运行在数据库后端中。

### 核心流程

```sql
CREATE EXTENSION postllm;

SELECT postllm.configure(
  runtime => 'openai',
  model => 'gpt-4o-mini',
  base_url => 'http://127.0.0.1:11434/v1/chat/completions'
);

SELECT postllm.chat_text(ARRAY[
  postllm.system('You are concise.'),
  postllm.user('Explain MVCC in one sentence.')
]);
```

提供服务前先调用 `postllm.runtime_discover()` 与 `postllm.runtime_ready()`。`postllm.settings()` 和 `postllm.capabilities()` 描述当前环境。API 参考按类别列出消息构造器、生成函数、结构化输出与工具、嵌入、`postllm.ingest_document`、重排、混合排序和 RAG 辅助函数。

### 治理与运维

应使用 `postllm.profile*`、`postllm.secret*`、`postllm.permission*` 和 `postllm.model_alias*` 函数组，而不是把配置散落在应用 SQL 中。限制 `http_allowed_hosts` 和 `http_allowed_providers`。模型调用可能把提示、数据行、文档和图像发送到 PostgreSQL 外部，并在等待网络或本地推理时占用后端；投入生产前必须考虑隐私、授权、超时、重试、费用、供应商保留策略与事务时长。
