## 用法

来源：

- [官方扩展控制文件](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/pgcov.control)
- [官方上游文档](https://github.com/johto/pgcov/blob/5f536efe07c91540bd149be152b65e80b75f3cf0/README.md)

`pgcov` — pgcov：通过共享预加载跟踪 PL/pgSQL 函数的测试覆盖率。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。
整理后的兼容版本集合为 `10,11,12`；仍需针对目标服务器确认实际构建。

```sql
CREATE EXTENSION "pgcov";
SELECT extversion
FROM pg_extension
WHERE extname = 'pgcov';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。
官方材料包含实验性、弃用、不支持或明确警告边界；用于非实验环境前必须阅读全文并测试失败场景。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
