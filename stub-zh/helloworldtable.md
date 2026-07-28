## 用法

来源：

- [官方上游 README](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/old/README.md)
- [官方扩展控制文件 (helloworldtable.control)](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/helloworldtable/helloworldtable.control)
- [官方实现源代码](https://github.com/stephen-hilton/ugly.network/blob/e8297ceb5be82a8f662a4c8cb2fc993901f30aaa/helloworldtable/src/lib.rs)

`helloworldtable` — pgrx 示例，返回包含 Hello World 欢迎信息的一行表。当 SQL 需要这些特殊函数或聚合时使用它。在目标 PostgreSQL 构建中使用上述固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION helloworldtable;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `helloworldtable()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
