## 用法

来源：

- [docs/getting-started/pqn-installation.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/getting-started/pqn-installation.md)
- [docs/integrations/postgres-extension.md](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/docs/integrations/postgres-extension.md)
- [infra/pqn-extension/pqn.control](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn.control)
- [infra/pqn-extension/pqn--1.0.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn--1.0.sql)
- [infra/pqn-extension/pqn--1.0--1.1.sql](https://github.com/pgquery-narrative/pgquerynarrative/blob/b0d91876f50ea2f3519ff12ee27c01c5e253e7a3/infra/pqn-extension/pqn--1.0--1.1.sql)

`pqn` 1.1 通过选定的授权视图和独立证据台账，在 PostgreSQL 内进行受限查询调查，无需外部 PgQueryNarrative 服务。

### 启用

受支持的工作流覆盖 PostgreSQL 16–18。由超级用户创建扩展并初始化台账。非超级用户安装前，需要 DBA 执行上游角色脚本并授予文档要求的角色成员资格。

```sql
CREATE EXTENSION pqn;
SELECT pqn_api.init();
SELECT pqn_api.verify_setup();
```

### 调查流程

`pqn_api.expose_sql` 输出拟执行的授权视图 SQL；审阅后再调用 `pqn_api.expose`。`pqn_api.enroll` 将已有登录角色加入查看者、分析者或管理员组，并记录限制。`pqn_api.plan`、`pqn_api.run`、`pqn_api.investigate` 与 `pqn_api.prove` 提供执行计划、受限查询、发现项和比较证据。独立命令行工具还可提出查询改写。

默认授权范围限制可见列；完整范围允许对隐藏列进行规划和测量，可能通过计数或计划暴露信息。加入用户前应审查授权与函数所有者角色。

### 运维

`pg_stat_statements` 及其预加载和重启仅用于工作负载排名，并非基础扩展的要求。用户可以修改登录超时设置；严格执行限制需要外部调度器以适当的监测和取消权限调用 `pqn_api.enforce_limits()`。`DROP EXTENSION pqn` 删除扩展代码，但保留独立初始化的 `pqn_ledger` 证据数据；视图、角色与保留数据的清理应遵循上游卸载流程。
