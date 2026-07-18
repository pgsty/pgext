## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/README.md)
- [1.0 版本 SQL 对象](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/jsonbx--1.0.sql)
- [C 实现](https://github.com/erthalion/jsonbx/blob/3b1bd19b49b7e5f9a4b952e18082165b8c1e1e7b/jsonbx.c)
- [PGXN 1.0.0 文档](https://pgxn.org/dist/jsonbx/1.0.0/README.html)

`jsonbx` 是面向 PostgreSQL 9.4 的 JSONB 操作兼容扩展。在这些函数与运算符进入 PostgreSQL 9.5 核心之前，它提供美化输出、拼接、按键/索引/路径删除，以及按路径替换。

```sql
SELECT jsonb_pretty('{"a":1,"b":2}'::jsonb);
SELECT '{"a":1}'::jsonb || '{"b":2}'::jsonb;
SELECT '{"a":1,"b":2}'::jsonb - 'a';
```

不要在仍受支持的 PostgreSQL 版本上安装 `jsonbx`。相同对象名已经存在于核心中，因此扩展脚本执行 `CREATE EXTENSION` 时可能冲突；应直接使用上面展示的内置函数与运算符。

1.0 是绑定 PostgreSQL 9.4 内部 JSONB API 的历史代码，只应保留用于隔离旧版 9.4 实例的取证。升级此类数据库前应移除扩展依赖，同时保留应用 SQL，使其改为解析到核心对象。
