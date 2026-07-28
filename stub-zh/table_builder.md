## 用法

来源：

- [官方上游 README](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/Readme.md)
- [官方扩展控制文件 (table_builder.control)](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/extension/table_builder.control)
- [官方实现源代码](https://github.com/jlockerman/table_builder_poc/blob/b24fb265a70d8dcabc0b9d8331f9d256b212f740/extension/src/lib.rs)

`table_builder` — 一个安全的 Rust 表构建器和查询器的概念验证。当应用程序需要此特定数据库功能时使用它。上游将其描述为一个概念验证。

### 核心工作流

```sql
CREATE EXTENSION table_builder;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 控制文件不要求超级用户安装。
- 上游将该项目描述为一个概念验证。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，以匹配固定源代码。
