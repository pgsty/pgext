## 用法

来源：

- [官方扩展控制文件](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/pltoolbox.control)
- [官方上游文档](https://github.com/okbob/pltoolbox/blob/6d9efc22d1071a071d54d6db9bd22d1672bf28de/README.md)

`pltoolbox` — 为 PostgreSQL 存储过程提供字符串、日期、记录与位图集合等 C 工具函数。

已复核目录快照记录的版本为 `1.0.3`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "pltoolbox";
SELECT extversion
FROM pg_extension
WHERE extname = 'pltoolbox';
```

整理后的生命周期为 `abandoned`。采用前应固定已复核构建并确认维护状态。

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
