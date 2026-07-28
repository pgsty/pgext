## 用法

来源：

- [官方上游 README](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/README.md)
- [官方扩展控制文件 (hstore_plruby.control)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/hstore_plruby/hstore_plruby.control)
- [官方扩展 SQL (hstore_plruby--1.0.sql)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/hstore_plruby/hstore_plruby--1.0.sql)

`hstore_plruby` — 用于在数据库代码中运行或与此过程语言进行交互时，将 hstore 转换为 Ruby Hash。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION hstore_plruby;

CREATE EXTENSION plruby;

CREATE FUNCTION hello(text) RETURNS text LANGUAGE plruby AS $$
    "Hello, #{args[0]}!"
$$;

SELECT hello('world');   -- Hello, world!
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码片段，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `hstore_to_plruby(val internal)` 是一个扩展函数，返回 `internal`。
- `plruby_to_hstore(val internal)` 是一个扩展函数，返回 `hstore`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 首先安装并验证确认的扩展依赖项：`hstore`, `plruby`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源保持一致。
