## 用法

来源：

- [官方上游 README](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/README)
- [官方扩展控制文件 (ftree.control)](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/ftree.control)
- [官方实现源代码](https://github.com/xiayingyin/atree/blob/8d5f94d51d02c90114ae47beb6cd1a2fc8f83ca2/ftree.c)

`ftree` — 为了将 L&Y 算法集成到 Postgres 中，我们进行了以下更改：在需要此特定数据库功能的应用程序中使用它。使用上方链接的已固定上游版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION ftree;
```

在目标数据库中安装扩展，当可用时运行上方的最小上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 已审核的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，以匹配已固定的源代码。
