## 用法

来源：

- [官方 README](https://github.com/pgEdge/pgedge-vectorizer/blob/7d020be5399f3812fc38633cc287b5a613ed8c64/README.md)
- [官方控制文件](https://github.com/pgEdge/pgedge-vectorizer/blob/7d020be5399f3812fc38633cc287b5a613ed8c64/pgedge_vectorizer.control)
- [官方文档](https://docs.pgedge.com/pgedge-vectorizer/)

`pgedge_vectorizer` 异步切分表中文本，并通过 OpenAI、Voyage AI 或 Ollama 生成 pgvector 嵌入。触发器创建队列任务，后台工作进程填充生成的分块表，因此源行提交与嵌入可用是两个独立事件。

### 前置条件与配置

扩展要求 PostgreSQL 14 或更高版本、`vector`、libcurl 及嵌入提供商。重启 PostgreSQL 前加入库与目标数据库：

```ini
shared_preload_libraries = 'pgedge_vectorizer'
pgedge_vectorizer.databases = 'appdb'
pgedge_vectorizer.provider = 'openai'
pgedge_vectorizer.api_key_file = '/run/secrets/pgedge-vectorizer-key'
pgedge_vectorizer.model = 'text-embedding-3-small'
pgedge_vectorizer.num_workers = 2
```

PostgreSQL 服务账号必须能读取密钥文件；不要把它放入 SQL 设置、备份或日志。预加载或工作进程数量变更需要重启。

### 核心流程

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pgedge_vectorizer;

CREATE TABLE articles (
    id bigserial PRIMARY KEY,
    title text NOT NULL,
    content text NOT NULL
);

SELECT pgedge_vectorizer.enable_vectorization(
    source_table := 'articles',
    source_column := 'content',
    chunk_strategy := 'token_based',
    chunk_size := 400,
    chunk_overlap := 50
);

INSERT INTO articles (title, content)
VALUES ('Example', 'Text to split and embed asynchronously.');

SELECT * FROM pgedge_vectorizer.queue_status;
SELECT * FROM pgedge_vectorizer.failed_items;
```

`enable_vectorization` 会创建辅助触发器、队列状态与分块表。生命周期操作可使用 `disable_vectorization`、`retry_failed` 与 `clear_completed`；禁用前应审查 `drop_chunk_table` 选项。

### 运维注意事项

提供商调用是异步外部副作用。应预期速率限制、重试、重复尝试、模型故障，以及源数据存在但嵌入尚不可用的时间窗口。监控队列年龄与失败，让下游读取容忍延迟；模型、维度或切分方式变化后，应有计划地重新嵌入。

控制文件要求超级用户安装、不可迁移，并依赖 `vector`。生成的分块表、触发器、队列元数据、GUC、密钥文件、提供商连通性及后台工作进程都必须纳入备份恢复与灾难恢复测试。文本分块、错误与外发请求可能包含敏感数据，因此要限制配置和监控 API 权限。
