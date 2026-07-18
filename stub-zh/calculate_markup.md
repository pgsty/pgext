## 用法

来源：

- [官方扩展控制文件](https://github.com/chobostar/calculate_markup/blob/c70a994c8a215e11d3fd444f0d1412289bcadf3b/calculate_markup.control)

`calculate_markup` — C 语言函数，根据价格与加价率数组通过线性插值计算加价百分比。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "calculate_markup";
SELECT extversion
FROM pg_extension
WHERE extname = 'calculate_markup';
```

整理后的生命周期为 `archived`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
