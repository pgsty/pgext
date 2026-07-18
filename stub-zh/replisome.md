## 用法

来源：

- [官方扩展控制文件](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/replisome.control)
- [官方上游文档](https://github.com/dvarrazzo/replisome/blob/775fe3766b8e433470c1048ac79334ed04b8caa2/README.rst)

`replisome` — 通过逻辑解码将 PostgreSQL 行变更以 JSON 流输出，并提供 Python 消费框架。

已复核目录快照记录的版本为 `0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "replisome";
SELECT extversion
FROM pg_extension
WHERE extname = 'replisome';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
