## 用法

来源：

- [官方扩展控制文件](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw.control)
- [官方上游文档](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/README.md)

`hello_fdw` — hello_fdw：用于访问或集成 hello world 的 PostgreSQL 外部数据包装器。

已复核目录快照记录的版本为 `1.0`、类型为 `standard`、实现语言为 `C`。

```sql
CREATE EXTENSION "hello_fdw";
SELECT extversion
FROM pg_extension
WHERE extname = 'hello_fdw';
```

投入生产前，应复核所链接的 control/SQL 或服务商文档，验证权限与兼容性，并在目标 PostgreSQL 构建上测试实际 API 和故障行为。
