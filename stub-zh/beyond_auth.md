## 用法

来源：

- [官方上游 README](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/README.md)
- [官方扩展控制文件 (beyond_auth.control)](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/beyond_auth.control)
- [官方实现源码](https://github.com/beyondoss/auth/blob/50766c44c0d5a06037741b6c1e80b54f13bb54e0/beyond-auth-extension/src/lib.rs)

`beyond_auth` — 在 PostgreSQL 中评估传递权限。通过使用广度优先搜索（BFS）遍历 auth.authz_relations 图来替代 N×depth 次往返查询，将深度加一的查询次数。在实现相应的安全、审计或访问控制工作流时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION beyond_auth;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `authz_check_array` 是一个扩展函数。
- `authz_check_batch` 是一个扩展函数。
- `authz_check_multi` 是一个扩展函数。
- `authz_check_parallel_batch` 是一个扩展函数。
- `authz_check_path_batch` 是一个扩展函数。
- `authz_check_single` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
