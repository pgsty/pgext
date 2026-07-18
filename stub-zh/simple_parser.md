## 用法

来源：

- [官方扩展控制文件](https://github.com/fairingrey/pg_simple_parser/blob/cdeaa0f1ed3b89d1961312c01e529f1c7c068084/simple_parser.control)

`simple_parser` — 仅按 ASCII 空格切分全文检索文本、保留下划线等标点在 token 内的 C 解析器扩展。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "simple_parser";
SELECT extversion
FROM pg_extension
WHERE extname = 'simple_parser';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
