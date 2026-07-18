## 用法

来源：

- [官方扩展控制文件](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/sql/output/ai.control)
- [官方上游文档](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/projects/extension/README.md)
- [官方上游 README](https://github.com/timescale/pgai/blob/47d74affd8a16d0b7f11c8f5bd6e35f7cd43abc1/README.md)

`ai` — 为 PostgreSQL AI 工作流提供辅助函数。

已复核目录快照记录的版本为 `0.11.3-dev`、类型为 `puresql`、实现语言为 `Python`。
应先安装并验证声明的扩展依赖：`plpython3u`, `vector`。
整理后的兼容版本集合为 `16,17,18`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "ai";
SELECT extversion
FROM pg_extension
WHERE extname = 'ai';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
