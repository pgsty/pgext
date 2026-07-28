## 用法

来源：

- [官方上游 README](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/README.md)
- [官方扩展控制文件 (plsci.control)](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/plsci.control)
- [官方实现源码](https://github.com/borkdude/plsci/blob/ac7d249b7182c65ec462da4541b6287819a726ff/src/lib.rs)

`plsci` — PostgreSQL 过程语言处理程序，通过 SCI 为 Clojure 提供支持。当数据库代码需要在该过程语言中运行或与其进行交互时，请使用此扩展。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION plsci;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `plsci` 是一个扩展函数。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源码中的信息一致。
