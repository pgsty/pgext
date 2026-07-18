## 用法

来源：

- [官方扩展控制文件](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/event_manager.control)
- [官方上游文档](https://bitbucket.org/neadwerx/event_manager/src/760428da0d8e7a33c24e202ca83a3e1789ad552d/README.md)
- [官方上游来源](https://bitbucket.org/neadwerx/event_manager)

`event_manager` — event_manager：基于触发器与队列的 PostgreSQL 事件处理系统，支持同步或异步处理。

已复核目录快照记录的版本为 `0.1`、类型为 `puresql`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。
整理后的兼容版本集合为 `10,11,12,13,14,15,16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "event_manager";
SELECT extversion
FROM pg_extension
WHERE extname = 'event_manager';
```

整理后的生命周期为 `active`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
