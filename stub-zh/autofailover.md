## 用法

来源：

- [1.0 版本控制文件](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover.control)
- [1.0 版本 SQL 对象](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover--1.0.sql)
- [命令实现](https://github.com/gfphoenix78/autofailover/blob/6494a6ed6b8b882f287331996b2ab075e334afa0/autofailover_funcs.c)

`autofailover` 是一个实验性的 C 扩展，用于暴露 PostgreSQL 角色/WAL 状态，并可直接提升备库或修改同步复制设置。它使用服务器私有内部接口，因此必须针对完全相同的 PostgreSQL 版本构建和测试；它并不是完整的故障转移控制器。

### 检查状态

```sql
CREATE EXTENSION autofailover;

SELECT *
FROM autofailover_execute('status', NULL);
```

`autofailover_execute(cmd, last_role)` 返回 `role`、`syncrep`、`sync_state`、`lsn` 和 `walconn`。实现可识别 `status`、`info`、`promote`、`syncrep` 与 `unsyncrep`；虽然接受 `last_role`，但不会用它做校验。

### 运维边界

`promote` 会写入 PostgreSQL 提升信号并通知 postmaster。`syncrep` 与 `unsyncrep` 会执行 `ALTER SYSTEM`，将 `synchronous_standby_names` 设为 `*` 或空值，然后重载配置。这些操作影响整个集群，但不提供隔离、仲裁、选主、脑裂防护或回滚。

两个安装函数都是 `SECURITY DEFINER`；`test_udf()` 只会产生合成调试行。应撤销不可信角色的访问权限，且不要通过应用连接暴露这些函数：

```sql
REVOKE EXECUTE ON FUNCTION autofailover_execute(text, text) FROM PUBLIC;
REVOKE EXECUTE ON FUNCTION test_udf() FROM PUBLIC;
```

只能将状态操作用作低层探针。考虑任何变更命令前，应在隔离集群中验证准确的 PostgreSQL ABI，并在外部提供健康检查、隔离、持久编排与恢复流程。
