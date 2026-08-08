## 用法

来源：

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [hstore_plruby v1.0 控制文件](https://github.com/commandprompt/plruby/blob/v2.5.0/hstore_plruby/hstore_plruby.control)
- [hstore_plruby v1.0 扩展 SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/hstore_plruby/hstore_plruby--1.0.sql)

`hstore_plruby` 为 `plruby` 语言安装 PostgreSQL `hstore` 与 Ruby `Hash` 值之间的转换。键会变为 Ruby 字符串，值会变为字符串或 `nil`；兼容的 Ruby 哈希也可以直接作为 `hstore` 返回。

### 安装并使用转换

```sql
CREATE EXTENSION hstore;
CREATE EXTENSION plruby;
CREATE EXTENSION hstore_plruby;

CREATE FUNCTION ruby_add_hstore_key(hstore)
RETURNS hstore
LANGUAGE plruby
TRANSFORM FOR TYPE hstore
AS $$
  value = args[0]
  value['processed'] = 'yes'
  value
$$;

SELECT ruby_add_hstore_key('id=>42'::hstore);
```

只有声明了 `TRANSFORM FOR TYPE hstore` 的函数才会使用该转换。

### 对象与注意事项

- `hstore_to_plruby(internal)` 实现从 SQL 到 Ruby 的转换。
- `plruby_to_hstore(internal)` 实现从 Ruby 到 SQL 的转换。
- 扩展版本为 `1.0`，同时依赖 `hstore` 和 `plruby`，并且可重定位。
- `hstore` 是从字符串到字符串或 NULL 的扁平映射。它不会保留嵌套的 Ruby 哈希、数组或有类型的数值；如果这些数据形态很重要，请使用 `jsonb_plruby`。
- PL/Ruby 仍是不受信任的语言。安装此转换不会为 Ruby 代码提供沙箱，也不会降低创建 PL/Ruby 函数所需的权限。
