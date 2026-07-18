## 用法

来源：

- [官方扩展控制文件](https://github.com/nate1001/chess_index/blob/4d6d473ea9c8238f06d8759f65862a382621f33e/chess_index.control)

`chess_index` — 用于存储并索引国际象棋棋局与局面的 PostgreSQL 数据类型。

已复核目录快照记录的版本为 `0.0.1`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "chess_index";
SELECT extversion
FROM pg_extension
WHERE extname = 'chess_index';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
