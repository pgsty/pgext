## 用法

来源：

- [官方上游 README](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/README.md)
- [官方扩展控制文件 (jsonb_plruby.control)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/jsonb_plruby/jsonb_plruby.control)
- [官方扩展 SQL (jsonb_plruby--1.0.sql)](https://github.com/commandprompt/plruby/blob/0720d8e72522c5196db062a1610eb2031a832246/jsonb_plruby/jsonb_plruby--1.0.sql)

`jsonb_plruby` — 在数据库代码需要在或与此过程语言进行交互时，用于在 jsonb 和 Ruby 数据之间进行转换。在安装扩展之前，必须先安装并验证其依赖项。

### 核心工作流

```sql
CREATE EXTENSION jsonb_plruby;

CREATE EXTENSION plruby;

CREATE FUNCTION hello(text) RETURNS text LANGUAGE plruby AS $$
    "Hello, #{args[0]}!"
$$;

SELECT hello('world');   -- Hello, world!
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，然后在集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `jsonb_to_plruby(val internal)` 是一个扩展函数，返回 `internal`。
- `plruby_to_jsonb(val internal)` 是一个扩展函数，返回 `jsonb`。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 首先安装并验证确认的扩展依赖项：`plruby`。
- 控制文件标记该扩展为可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
