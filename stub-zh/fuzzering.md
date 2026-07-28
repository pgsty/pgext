## 用法

来源：

- [官方上游 README](https://gitlab.com/shanenoi/fuzzering/-/blob/master/README.md)
- [官方扩展控制文件](https://gitlab.com/shanenoi/fuzzering/-/blob/master/fuzzering.control)
- [官方扩展 SQL](https://gitlab.com/shanenoi/fuzzering/-/blob/master/fuzzering--0.0.1.sql)

`fuzzering` — > 你需要安装一些用于构建扩展的包：Clang。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION fuzzering;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `fuzzering(text, text)` 是一个扩展函数，返回 `float8`。
- `levenshtein(text, text)` 是一个扩展函数，返回 `int32`。
- `simular(int32, int32, int32)` 是一个扩展函数，返回 `float8`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.1`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配固定源代码。
