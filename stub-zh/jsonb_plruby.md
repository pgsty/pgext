## 用法

来源：

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [jsonb_plruby v1.0 控制文件](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby.control)
- [jsonb_plruby v1.0 扩展 SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/jsonb_plruby/jsonb_plruby--1.0.sql)

`jsonb_plruby` 为 `plruby` 语言安装 PostgreSQL `jsonb` 与原生 Ruby 值之间的转换。经过转换的 `jsonb` 参数会成为 Ruby `Hash`、`Array`、`String`、`Integer`、`Float`、`true`、`false` 或 `nil`；兼容的 Ruby 值也可以直接作为 `jsonb` 返回。

### 安装并使用转换

```sql
CREATE EXTENSION plruby;
CREATE EXTENSION jsonb_plruby;

CREATE FUNCTION ruby_mark_processed(jsonb)
RETURNS jsonb
LANGUAGE plruby
TRANSFORM FOR TYPE jsonb
AS $$
  value = args[0]
  value['processed'] = true
  value
$$;

SELECT ruby_mark_processed('{"id": 42}'::jsonb);
```

只有声明了 `TRANSFORM FOR TYPE jsonb` 的函数才会使用该转换。其他 PL/Ruby 函数仍采用该语言通常的 JSONB 转换行为。

### 对象与注意事项

- `jsonb_to_plruby(internal)` 实现从 SQL 到 Ruby 的转换。
- `plruby_to_jsonb(internal)` 实现从 Ruby 到 SQL 的转换。
- 扩展版本为 `1.0`，依赖 `plruby`，并且可重定位。
- 返回 PostgreSQL 的 Ruby `Hash` 键必须是有效的 JSON 对象键，数值及特殊值也必须能够由 PostgreSQL `jsonb` 表示。请显式测试嵌套值和数值边界。
- PL/Ruby 仍是不受信任的语言。安装此转换不会为 Ruby 代码提供沙箱，也不会降低创建 PL/Ruby 函数所需的权限。
