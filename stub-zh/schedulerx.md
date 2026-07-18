## 用法

来源：

- [官方扩展控制文件](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/schedulerx.control)
- [官方上游文档](https://github.com/okbob/generic-scheduler/blob/3fc103f12a276656e59d3bc07b0f81137979b6c2/README.md)

`schedulerx` — 通过后台工作进程运行周期、单次或 NOTIFY 触发 SQL 作业的原型调度器。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "schedulerx";
SELECT extversion
FROM pg_extension
WHERE extname = 'schedulerx';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
