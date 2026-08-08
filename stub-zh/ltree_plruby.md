## 用法

来源：

- [PL/Ruby v2.5.0 README](https://github.com/commandprompt/plruby/blob/v2.5.0/README.md)
- [ltree_plruby v1.0 控制文件](https://github.com/commandprompt/plruby/blob/v2.5.0/ltree_plruby/ltree_plruby.control)
- [ltree_plruby v1.0 扩展 SQL](https://github.com/commandprompt/plruby/blob/v2.5.0/ltree_plruby/ltree_plruby--1.0.sql)

`ltree_plruby` 为 `plruby` 语言安装 PostgreSQL `ltree` 路径与 Ruby 数组之间的转换。`ltree` 参数会成为由标签字符串组成的数组；由有效标签组成的数组也可以直接作为 `ltree` 值返回。

### 安装并使用转换

```sql
CREATE EXTENSION ltree;
CREATE EXTENSION plruby;
CREATE EXTENSION ltree_plruby;

CREATE FUNCTION ruby_append_label(ltree, text)
RETURNS ltree
LANGUAGE plruby
TRANSFORM FOR TYPE ltree
AS $$
  path = args[0]
  path << args[1]
  path
$$;

SELECT ruby_append_label('Top.Science'::ltree, 'Astronomy');
```

只有声明了 `TRANSFORM FOR TYPE ltree` 的函数才会使用该转换。

### 对象与注意事项

- `ltree_to_plruby(internal)` 实现从 SQL 到 Ruby 的转换。
- `plruby_to_ltree(internal)` 实现从 Ruby 到 SQL 的转换。
- 扩展版本为 `1.0`，同时依赖 `ltree` 和 `plruby`，并且可重定位。
- 返回数组中的每个元素都必须是有效的 `ltree` 标签。PostgreSQL 会拒绝无效字符、空标签或超过 `ltree` 限制的路径。
- PL/Ruby 仍是不受信任的语言。安装此转换不会为 Ruby 代码提供沙箱，也不会降低创建 PL/Ruby 函数所需的权限。
