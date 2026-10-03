## 用法

来源：

- [README.rst](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/README.rst)
- [pgq_coop.control](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/pgq_coop.control)
- [Makefile](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/Makefile)
- [mk/common-pgxs.mk](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/mk/common-pgxs.mk)
- [sql/pgq_coop_test.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/sql/pgq_coop_test.sql)
- [functions/pgq_coop.next_batch.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/functions/pgq_coop.next_batch.sql)
- [functions/pgq_coop.finish_batch.sql](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/functions/pgq_coop.finish_batch.sql)
- [.github/workflows/ci.yml](https://github.com/pgq/pgq-coop/blob/d6b27d536ccffbd13822b7578a79d6fa701c4fa6/.github/workflows/ci.yml)

`pgq_coop` 3.4 允许多个连接协作消费同一个 PgQ 消费者的事件。当前上游 CI 配置覆盖 PostgreSQL 14–18。

### 核心工作流

```sql
CREATE EXTENSION pgq;
CREATE EXTENSION pgq_coop;
SELECT pgq.create_queue('work');
SELECT pgq_coop.register_subconsumer('work', 'workers', 'worker_a');
SELECT pgq.insert_event('work', 'example', 'payload');
SELECT pgq.ticker();
SELECT pgq_coop.next_batch('work', 'workers', 'worker_a');
```

### 消费与维护

`pgq_coop.next_batch` 返回批次 ID；暂无任务时返回 NULL。通过普通 PgQ 批次 API 读取事件，完成应用处理后，将返回的 ID 传给 `pgq_coop.finish_batch(batch_id)`。`pgq_coop.next_batch_custom` 控制批次阈值；四参数形式的 next-batch 函数接受用于接管任务的不活跃时间间隔。`pgq_coop.unregister_subconsumer` 删除工作进程的注册信息。

安装需要超级用户，并要求已有可用的 `pgq` 依赖。该 SQL 扩展不可重定位，会创建 `pgq_coop` 模式；控制文件选择 `pg_catalog` 作为安装模式。无需额外共享库或预加载。PgQ 仍需要正常的 ticker 与维护任务。工作进程失败与任务接管可能触发重新处理，因此处理器必须能够安全重试；此扩展不能替代工作进程。
