## 用法

来源：

- [官方扩展控制文件](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/setrank.control)
- [官方上游文档](https://github.com/obartunov/setrank/blob/b8492a5968c33aa1c2d87d6aecf4a35477d75e10/README)

`setrank` — 基于统计表中的文档频率，为 tsvector/tsquery 提供 TF-IDF 与覆盖密度排序函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "setrank";
SELECT extversion
FROM pg_extension
WHERE extname = 'setrank';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
