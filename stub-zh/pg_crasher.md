## 用法

来源：

- [官方 README](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/README)
- [官方扩展 SQL](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/pg_crasher--1.0.sql)
- [官方 C 实现](https://github.com/davidcrawford/pg_crasher/blob/fdfb24d7b7087a211c940cf96473069103c42a95/pg_crasher.c)

`pg_crasher` 会故意让调用它的 PostgreSQL 后端进程崩溃。它唯一合理的用途是在受控环境中测试客户端断线处理、事务恢复、连接池、监控与故障转移行为；绝不能在共享生产数据库中安装版本 `1.0`。

### 安全安装边界

按照 PostgreSQL 的常规默认权限，扩展创建的函数可由 `PUBLIC` 执行。应在安装扩展的同一事务中撤销该权限，并仅在确有需要时授予专用测试角色：

```sql
BEGIN;
CREATE EXTENSION pg_crasher;
REVOKE ALL ON FUNCTION pg_crasher() FROM PUBLIC;
GRANT EXECUTE ON FUNCTION pg_crasher() TO crash_test_role;
COMMIT;
```

### 主动崩溃测试

从连接到可丢弃实例的隔离客户端调用一次该函数：

```sql
SELECT pg_crasher();
```

C 函数会通过空指针写入。该会话应立即终止，未提交事务应回滚，postmaster 随后会为新连接启动替代后端。

### 运维注意事项

这不是普通的报错辅助函数：它会造成原生进程崩溃，并可能触发崩溃告警、核心转储、连接池抖动或更广泛的恢复行为。不要在包含必须提交数据的事务中运行。测试前应确认确切的主机、集群、数据库、角色与维护窗口，监控服务器日志与恢复过程，并在结束后移除扩展。复核的项目没有兼容矩阵，源码最后变更于 2013 年，因此即使实验室使用也必须针对确切版本构建，并使用可丢弃环境。
