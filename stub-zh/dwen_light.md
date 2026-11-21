## 用法

来源：

- [README.md](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/README.md)
- [dwen_light.control](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/dwen_light.control)
- [dwen_light--0.2.sql](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/dwen_light--0.2.sql)
- [LICENSE](https://github.com/kiramontus/dwen_light/blob/6dc268dc623682c53787bcdd4a20ebbe17745b73/LICENSE)

`dwen_light` 是仍处于早期、不稳定阶段的 SQL 与 PL/pgSQL 后台任务执行器。版本 0.2 的开发测试环境是 Ubuntu 上的 PostgreSQL 18。它安装到固定的 `dwen` 模式，使用配置的数据库角色执行队列中的 SQL。

### 配置与启用

官方流程要求预加载库、配置工作进程和连接，重启 PostgreSQL，创建扩展后再次重启：

```ini
shared_preload_libraries = 'dwen_light'
dwen.max_execution_workers = 2
dwen.database = 'postgres'
dwen.role = 'dwen'
```

应保留原有预加载列表中的其他库。以超级用户连接配置指定的数据库：

```sql
CREATE EXTENSION dwen_light;
```

安装脚本在角色 `dwen` 不存在时创建它，并授予扩展模式内的权限，包括后续对象的默认权限。启用工作进程前，应核对配置角色及任务表的访问权限；任务字符串会作为 SQL 执行。

### 提交与查看

扩展安装时会创建上游示例过程：

```sql
INSERT INTO dwen.task (status, proc, execute_string, run_timestamp)
VALUES ('READY', 'X', 'call dwen.dev_test_proc()', NULL);

SELECT event_at, type, message, task_id, errcode
FROM dwen.log
ORDER BY event_at DESC;
```

文档中的轮询机制会在约 15 秒内启动该任务。任务表还包含 `timeout_minutes` 列，默认值为 60。`dwen.task_execute_string()` 执行保存的命令并记录成败；`dwen.get_engine_status_active()` 和 `dwen.set_engine_status_active(boolean)` 用于读取和设置引擎活动标志。

### 边界

这是行为仍在变化的预览框架，上游尚未给出持久化队列保证。所审阅源码提供 0.2 安装脚本，没有版本化升级路径，还会安装演示对象。仅允许获准以工作进程身份执行操作的角色提交任务；投入运行前应测试失败和重启行为。
