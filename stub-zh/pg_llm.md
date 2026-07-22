## 用法

来源：

- [官方 README](https://github.com/onelazyteam/pg_llm/blob/master/README.md)
- [官方 1.1 版 SQL](https://github.com/onelazyteam/pg_llm/blob/master/sql/pg_llm--1.1.sql)
- [官方 control 文件](https://github.com/onelazyteam/pg_llm/blob/master/pg_llm.control)

`pg_llm` 把远程和本地语言模型集成到 PostgreSQL。版本 `1.1` 提供模型登记、聊天与流式调用、多轮会话、并行模型选择、文本转 SQL、SQL 分析/报告、向量知识检索、反馈、审计日志和请求追踪。它依赖 `vector`，并从数据库后端进程发送模型提示词。

### 核心流程

预加载该库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_llm'
```

安装 `vector`，创建扩展、登记模型并调用。示例使用 psql 变量，避免把密钥字面量写入命令历史：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_llm;

SELECT pg_llm_add_model(
  'qwen',
  'qianwen-chat',
  :'api_key',
  '{"model_name":"qwen-turbo","api_endpoint":"https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"}'
);

SELECT pg_llm_chat('qwen', 'Explain PostgreSQL MVCC briefly.');
```

多轮使用时，先通过 `pg_llm_create_session(max_messages)` 创建会话，再把返回的 ID 传给 `pg_llm_multi_turn_chat`。

### 主要 API 分组

- 模型/聊天：`pg_llm_add_model`、`pg_llm_remove_model`、`pg_llm_chat`、`pg_llm_chat_stream`、`pg_llm_chat_json` 与并行聊天变体。
- 会话：`pg_llm_create_session`、`pg_llm_multi_turn_chat`，以及会话查询、状态更新、删除和清理函数。
- 数据工作：`pg_llm_text2sql[_json]`、`pg_llm_execute_sql_with_analysis`、`pg_llm_generate_report`。
- 检索：`pg_llm_get_embedding`、向量存储/搜索、`pg_llm_add_knowledge`、`pg_llm_search_knowledge`。
- 可观测性：`pg_llm_record_feedback`、`pg_llm_get_audit_log`、`pg_llm_get_trace`。

模型、会话、向量、审计、追踪、报告、知识与反馈状态保存在 `_pg_llm_catalog`；版本 `1.1` 为内部检索表创建 64 维 pgvector 字段与 IVFFlat 索引。

### 注意事项

扩展需要 PostgreSQL 14+、C++20 构建、libcurl、OpenSSL、JsonCpp 与 pgvector。预加载和外发 HTTP 都发生在数据库服务器进程内；服务商延迟、限流、TLS/DNS 故障和大型/流式响应可能占用后端与事务。

提示词、SQL 文本、查询结果、API 密钥、审计记录和知识文档都可能敏感。应限制网络、文本转 SQL、SQL 执行函数的 `EXECUTE` 权限，保护 `_pg_llm_catalog`，采用托管密钥路径，并审查保留与日志策略。生成 SQL 和模型建议必须作为不可信输入：验证权限，尽可能使用只读角色，绝不能在高权限会话中自动执行。
