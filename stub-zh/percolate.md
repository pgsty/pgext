## 用法

来源：

- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/README.md)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate/percolate.control)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/extension/percolate/sql/percolate--0.2.30.sql)
- [Official source](https://github.com/Percolation-Labs/percolate/blob/aa472a0b7e9f3393dccb5999620de355eb0654c3/dev/image-init/01-percolate.sh)

`percolate` 0.2.30 是面向 PostgreSQL 19 的纯 SQL 工作流与智能体记忆系统。工作流、运行记录、身份、文档和图数据存储在数据库模式中。完整服务还需要独立工作进程与 HTTP 服务；创建扩展不会启动这些进程。

### 初始化与对象所有权

首先遵循官方初始化流程创建应用角色与数据库权限。该流程会修改数据库级权限，包括内置配置设置函数的权限，应在专用数据库中审查后使用。前置扩展包括 `pgcrypto`、`pg_trgm`、`vector` 与 `percolate_parser`，其中不可信的编译型依赖必须由超级用户预先安装。

初始化完成角色与授权后：

```sql
SET ROLE app_owner;
CREATE EXTENSION percolate CASCADE;
RESET ROLE;
SELECT extname, extversion FROM pg_extension WHERE extname = 'percolate';
```

安装器明确拒绝超级用户调用，因为应用角色持有对象与行级安全是其安全约定的一部分。它还会检查必要角色、成员关系和权限限制，不能作为可随意替换角色名称的通用模式包安装。

### 工作流与维护

运行时通过工作进程执行数据库中的任务，调度计划保存在 `workflow.schedules`。`workflow.tick()` 推进调度，`workflow.reap_stale_tasks()` 回收遗留任务，`workflow.promote_due_timers()` 推进定时任务。标准部署通过可选的 `pg_cron` 运行这些函数，需要预加载并指定正确数据库；没有它时，调度、任务回收和定时任务不会自动运行。

扩展不可迁移模式。使用匹配的解析器与升级脚本，在备份中保留应用角色和权限，并以真实应用数据测试恢复。SQL 安装、工作进程部署和定时执行是独立步骤。
