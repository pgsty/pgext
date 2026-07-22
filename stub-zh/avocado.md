## 用法

来源：

- [固定提交的上游 PostgreSQL 扩展指南](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/README.md#postgresql-extension-new-in-v22)
- [固定提交的扩展 control 文件](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/avocado.control)
- [固定提交的扩展 schema](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/sql/schema.sql)
- [固定提交的 SQL 函数源码](https://github.com/avocadodb/avocadodb/blob/18d30fd558a4c7f538f9667f788291e2ad6d43f9/avocado-pgext/src/lib.rs)

`avocado` `2.2.0` 是一个 pgrx PostgreSQL 扩展，用于摄取文本制品、生成嵌入、检索相关片段，并为 AI 代理编译带引用的上下文。它把数据存放在 `avocado` schema 下的普通 PostgreSQL 表中，并使用 `vector` 建立 1,024 维 HNSW 索引。

### 核心流程

control 文件要求先有 `vector`，安装需要超级用户，并且扩展不可重定位。安装后先初始化表，再摄取内容：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION avocado;
SELECT avocado_init();

SELECT avocado_ingest_artifact(
  'docs/auth.md',
  'Authentication uses short-lived access tokens and rotating refresh tokens.',
  '{"type":"markdown"}'::jsonb
);

SELECT avocado_search_spans('How are access tokens refreshed?', 5);
SELECT avocado_compile(
  'How does authentication work?',
  '{"token_budget":2000,"enable_mmr":true}'::jsonb
);
```

默认嵌入提供者是使用 `all-MiniLM-L6-v2` 的 FastEmbed。也可以在摄取或搜索前，为当前后端选择 Ollama 端点：

```sql
SELECT avocado_set_embedding_provider('ollama');
SELECT avocado_set_ollama_config('http://localhost:11434', 'bge-m3');
SELECT avocado_embedding_config();
SELECT avocado_test_embedding('configuration check');
```

### 重要对象

- `avocado_ingest_artifact(path, content, metadata)`：按行切分文本，为各片段生成嵌入并保存制品与片段。
- `avocado_search_spans(query, limit)`：以 `jsonb` 返回匹配片段及余弦相似度分数。
- `avocado_compile(query, config)`：执行向量搜索、可选的 MMR 去冗余、token 预算装箱和确定性的路径/行号排序，并以 `jsonb` 返回上下文与引用。
- `avocado_create_session()`、`avocado_add_message()` 和 `avocado_get_conversation_history()`：维护会话状态。
- `avocado_register_agent()` 与关系函数：维护代理元数据和同意/反对记录。
- `avocado_stats()` 与 `avocado_version()`：报告已存对象数量、嵌入配置和扩展版本。

### 运维说明

已复核构建只暴露 PostgreSQL 16 的 pgrx feature，其容器也基于 PostgreSQL 16；其他服务器版本必须单独验证。默认 FastEmbed 模型可能需要下载模型并占用较多后端内存，Ollama 则会从 PostgreSQL 后端同步发起 HTTP 请求。应相应收紧执行权限、出站访问与提供者配置，并按嵌入延迟设置语句超时。

嵌入会被补齐或截断为 `vector(1024)`；更换提供者可能在同一存储中混入不同模型与维度，应重新摄取数据，而不是比较不兼容的嵌入。已复核的 `2.2.0` 摄取路径会先生成新的制品标识，再按唯一 path 执行 upsert，因此在依赖更新前必须测试同一路径的重复摄取和片段替换。`avocado` 表应作为普通数据库数据备份，并在真实语料上验证 HNSW 构建成本、模型可复现性与检索质量。
