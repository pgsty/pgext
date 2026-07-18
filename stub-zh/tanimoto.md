## 用法

来源：

- [官方扩展控制文件](https://github.com/ergo70/tanimoto/blob/df31a86da4098c94c684167a789e34008d051e81/tanimoto.control)

`tanimoto` — 面向 PostgreSQL bit varying 指纹的快速 C 语言 Tanimoto 相似系数函数。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "tanimoto";
SELECT extversion
FROM pg_extension
WHERE extname = 'tanimoto';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
