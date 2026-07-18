## 用法

来源：

- [锁定提交的上游 README](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/README.md)
- [1.0 版 SQL API](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos--1.0.sql)
- [锁定提交的事务状态实现](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos.c)
- [锁定提交的扩展控制文件](https://github.com/bdrouvot/pg_subtrans_infos/blob/021336ba9d31e41d581f8d55a6d4ab79da4cf5de/pg_subtrans_infos.control)

pg_subtrans_infos 1.0 报告近期事务 ID 的内部状态和父子关系。它唯一的函数返回 xid、状态、直接父事务、顶层父事务、子事务深度和可选提交时间戳。这是一个底层诊断辅助工具，不是持久事务历史目录。

### 检查近期锁事务

```sql
CREATE EXTENSION pg_subtrans_infos;

SELECT *
FROM pg_subtrans_infos(txid_current());

SELECT l.pid, l.transactionid, s.*
FROM pg_locks AS l
CROSS JOIN LATERAL
  pg_subtrans_infos(l.transactionid::text::bigint) AS s
WHERE l.transactionid IS NOT NULL
ORDER BY l.pid, l.transactionid;
```

状态可以是 committed、aborted 或 in progress。只有在启用 track_commit_timestamp 且对应提交时间戳记录仍被保留时，commit_timestamp 才有值。早于已保留 pg_xact 状态的 ID 会产生 NULL 详情，未来 ID 则会报错。

### 保留与兼容性边界

函数会把提供的 32 位事务 ID 与当前 epoch 组合。它只适用于 pg_xact 和 pg_subtrans 条目尚未回绕或截断的近期 ID。父事务信息可能比其他事务证据更早消失，README 也指出，事务不再处于进行中后，top_parent_xid 可能为空。不得把结果用于取证完整性、业务逻辑或恢复决策。

C 模块复制了 PostgreSQL 后端逻辑，并调用私有子事务、提交时间戳、事务状态、共享内存和锁接口。这些 API 和不变量可能在每个主版本中改变。README 记录的测试范围为 PostgreSQL 10 至 12.2，锁定代码增加了 13 的分支；它不支持 14 或更新版本。

目录将该扩展标记为已废弃。只应针对受支持的历史服务器编译，短暂用于诊断后立即移除。在当前 PostgreSQL 版本上，应优先使用受支持的监控视图和日志，或在完成上游后端测试的前提下移植代码，不要加载旧二进制文件。
