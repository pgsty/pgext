## 用法

来源：

- [官方 README](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/README.md)
- [扩展 SQL](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo--1.0.sql)
- [libpq 包装实现](https://github.com/ibarwick/conninfo/blob/6eefa30a916761963dc2727a08bf6c958df5d8fe/conninfo.c)

`conninfo` 将服务端所链接 libpq 的解析器和默认值暴露为 SQL 函数。它可用于拆分连接字符串、检查 libpq 能否解析字符串，或查看有效的默认连接参数；它不会实际尝试数据库连接。

### 核心流程

```sql
CREATE EXTENSION conninfo;

SELECT *
FROM conninfo_parse(
    'host=db.internal port=5432 dbname=app user=reader connect_timeout=5'
);

SELECT conninfo_validate('host=db.internal connect_timeout=5');

SELECT keyword, envvar, compiled, value
FROM conninfo_defaults()
WHERE keyword IN ('host', 'port', 'user', 'dbname');
```

`conninfo_parse(text)` 为每个给定选项返回一行 `keyword`/`value`，遇到格式错误或未知选项时抛错。`conninfo_validate(text)` 对同类解析失败返回 `false` 并发出警告。结果为 `true` 只表示 libpq 接受了语法和选项名；它不验证可达性、凭据、TLS，也不判断参数值在语义上是否合适。

### 函数索引

- `conninfo_parse(text)` 返回 `SETOF record`，字段为 `keyword text` 和 `value text`。
- `conninfo_validate(text)` 返回 `boolean`。
- `conninfo_defaults()` 返回来自 `PQconndefaults()` 的 `keyword`、`envvar`、`compiled`、`value`、`label`、`dispchar` 和 `dispsize`。

### 安全与兼容性

解析结果和默认值可能含有密码，包括从 `PGPASSWORD` 或 libpq 服务及环境设置继承的值。不要记录查询结果，不要通过不可信 API 暴露这些函数，也不要在未考虑凭据泄露时授予访问权。

这些函数使用 PostgreSQL 服务端进程所链接的 libpq 及其环境，它们可能不同于客户端应用的 libpq 构建和环境。如果行为取决于客户端版本，仍应使用相同的客户端运行时验证。
