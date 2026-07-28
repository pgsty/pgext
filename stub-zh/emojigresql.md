## 用法

来源：

- [官方上游 README](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/README.md)
- [官方扩展控制文件 (emojigresql.control)](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/emojigresql.control)
- [官方扩展 SQL (emojigresql--1.0.sql)](https://github.com/vinsmokesomya/emojigresql/blob/55d4b4ab81e93179b36b455f9960385d5f849db8/emojigresql--1.0.sql)

`emojigresql` — EmojigreSQL 是一个 **纯 SQL** 的 PostgreSQL 扩展，旨在无缝地将字节数据（二进制数据）和文本编码/解码为表情符号。使用它来对应相应的 SQL 或数据库实用程序工作流。请使用上述链接的上游修订版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION emojigresql;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `emojigresql.decode(text)` 是一个扩展函数，返回 `bytea`。
- `emojigresql.encode(bytea)` 是一个扩展函数，返回 `text`。
- `emojigresql.from_text(text)` 是一个扩展函数，返回 `text`。
- `emojigresql.to_text(text)` 是一个扩展函数，返回 `text`。
- `emojigresql.chars` 是一个由扩展安装或管理的表。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况。
