## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/README.md)
- [官方扩展控制文件 (soundex-function.control)](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/soundex-function.control)
- [官方扩展 SQL (soundex-function--1.0.0.sql)](https://api.pgxn.org/src/soundex-function/soundex-function-1.0.0/soundex-function--1.0.0.sql)

`soundex-function` — 该仓库包含一个创建 PostgreSQL 函数的模板。此模板的目标是为 PostgreSQL 创建函数提供一个起点。使用它来实现相应的文本搜索、解析或语言工作流。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION "soundex-function";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `soundex(input TEXT)` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
