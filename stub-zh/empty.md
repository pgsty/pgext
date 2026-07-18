## 用法

来源：

- [官方扩展控制文件](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/empty.control)
- [官方上游文档](https://github.com/tvondra/empty/blob/ac8d6da58df3763a462a5fdd1f781269cff7d67d/README.md)

`empty` — empty：演示型 C 扩展，包含示例函数、matrix 类型、FDW、逻辑解码回调与共享内存 hook。

已复核目录快照记录的版本为 `1.0`、类型为 `preload`、实现语言为 `C`。

```sql
CREATE EXTENSION "empty";
SELECT extversion
FROM pg_extension
WHERE extname = 'empty';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
