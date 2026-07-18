## 用法

来源：

- [上游 README](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/README.md)
- [扩展安装 SQL](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/main.sql)
- [共享缓冲区实现](https://github.com/postgrespro/pg_logging/blob/f7fd66ef4931d74ce2082c49dc0e14afcd2656f1/pg_logging.c)

pg_logging 将 PostgreSQL 错误报告事件捕获到共享内存环形缓冲区，再以行形式对外提供。它可用于短期诊断，或供需要结构化日志级别、消息、查询、详细信息和位置字段的外部收集器使用。

### 服务器配置

该模块会拒绝普通的延迟加载，因此必须预加载并重启 PostgreSQL。缓冲区大小在启动时分配：

```ini
shared_preload_libraries = 'pg_logging'
pg_logging.buffer_size = 10240
pg_logging.enabled = on
```

安装 SQL 对象后，可先在不推进共享读位置的情况下检查记录，再显式清空：

```sql
CREATE EXTENSION pg_logging;

SELECT level::error_level, message, position
FROM get_log(false)
ORDER BY position;

SELECT flush_log();
```

使用默认的 true 参数调用 get_log() 会推进共享读位置。收集器也可保存 position，并调用整数重载从已知位置续读。

### 注意事项

- 环形缓冲区绕回时会覆盖尚未读取的旧记录，它不是持久化审计存储。
- 以清空模式读取或调用 flush_log 都会改变共享消费者状态。多个收集器必须做好协调，避免一方丢弃另一方未读的区间。
- 捕获字段可能包含 SQL 文本、绑定相关上下文、对象名、用户数据和错误详情。应限制函数访问权，并在下游存储中实施明确的短保留策略。
- 更大的缓冲区设置会消耗共享内存，且需要重启。应在错误峰值下测量日志开销与溢出情况。
- 上游没有发布当前 PostgreSQL 大版本兼容性矩阵；应在确切的目标构建上验证这个服务器内部模块。
