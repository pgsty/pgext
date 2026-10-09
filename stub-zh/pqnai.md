## 用法

来源：

- [README.md](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/README.md)
- [extension/pqnai--0.3.0.sql](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/extension/pqnai--0.3.0.sql)
- [docs/architecture.md](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/docs/architecture.md)
- [.github/workflows/ci.yml](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/.github/workflows/ci.yml)
- [extension/pqnai.control](https://github.com/PgQueryNarrative/pqnai/blob/674537ce965eb995b7b8a9b57f676885cb2643a8/extension/pqnai.control)

`pqnai` 通过外部 Go 工作进程 `pqnaid` 提供预测和检索过程。此早期版本实现移动平均预测，以及基于 pgvector 和 Ollama 的分块检索；ARIMA、混合检索和重排序仍在规划中。

### 核心用法

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pqnai;
CALL pqnai.forecast(ARRAY[10,20,30,40,50]::double precision[], 3);
CALL pqnai.embed('The worker receives jobs using LISTEN and NOTIFY.');
CALL pqnai.ask('How does the worker receive jobs?', 1);
```

### 运行边界

先安装 `vector`，再由超级用户创建 `pqnai`。固定模式包含 `jobs`、`sources` 和 `chunks`。工作进程必须连接同一数据库并可访问 Ollama。未声明 PostgreSQL 预加载要求；只安装 SQL 扩展并不提供推理能力。上游 CI 目标为 PostgreSQL 14–17。

在显式事务块外使用顶层 `CALL`：过程先提交排队任务，再等待工作进程。超时不会撤销已经提交的工作。`pqnai.enqueue_job` 和 `pqnai.wait_for_job` 是底层队列接口。应限制任务和文档的访问权限；此版本尚未实现租户级访问控制。安装的向量列固定为 384 维，更换嵌入维度需要模式迁移。应规划任务和文档数据的保留策略。
