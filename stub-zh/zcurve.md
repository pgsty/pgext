## 用法

来源：

- [官方扩展控制文件](https://github.com/bmuratshin/zcurve/blob/76b88d80794a4ccde95c01111453ad8d7fbdbf36/zcurve.control)

`zcurve` — 提供二维和三维 Z 序位交错编码，并可直接在 numeric B-tree 索引上执行范围查找。

已复核目录快照记录的版本为 `1.3`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "zcurve";
SELECT extversion
FROM pg_extension
WHERE extname = 'zcurve';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
