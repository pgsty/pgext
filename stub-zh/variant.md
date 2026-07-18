## 用法

来源：

- [官方扩展控制文件](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/variant.control)
- [官方 PGXN 分发页](https://pgxn.org/dist/variant/)
- [官方上游 README](https://github.com/BlueTreble/variant/blob/527cb8b2a85bec118de157bcee46d83fa98ec45a/README.md)

`variant` — 可保存任意 PostgreSQL 类型值并记录其原始类型的 variant 数据类型。

已复核目录快照记录的版本为 `1.0.1`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`plpgsql`。

```sql
CREATE EXTENSION "variant";
SELECT extversion
FROM pg_extension
WHERE extname = 'variant';
```

整理后的生命周期为 `preview`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
