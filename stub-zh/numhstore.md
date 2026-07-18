## 用法

来源：

- [官方扩展控制文件](https://github.com/adjust/pg-numhstore/blob/754e2487940321f93344f1cb04244c8234097532/numhstore.control)
- [官方上游文档](https://github.com/adjust/pg-numhstore/blob/754e2487940321f93344f1cb04244c8234097532/README.md)

`numhstore` — numhstore：为 hstore 增加 inthstore 与 floathstore 数值类型及辅助函数。

已复核目录快照记录的版本为 `0.1.7`、类型为 `standard`、实现语言为 `C`。
应先安装并验证声明的扩展依赖：`hstore`。

```sql
CREATE EXTENSION "numhstore";
SELECT extversion
FROM pg_extension
WHERE extname = 'numhstore';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
