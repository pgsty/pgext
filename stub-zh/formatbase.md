## 用法

来源：

- [官方扩展控制文件](https://github.com/ttfkam/pg_formatbase/blob/a456be91e662a8e9b89b005f07e5f6256e3e637b/formatbase.control)

`formatbase` — formatbase：函数工具类 PostgreSQL 扩展；上游说明为“使用 C 函数按数值进制格式化数字并解析字符串”。

已复核目录快照记录的版本为 `2.0.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "formatbase";
SELECT extversion
FROM pg_extension
WHERE extname = 'formatbase';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
