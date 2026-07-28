## 用法

来源：

- [官方上游 README](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/README.md)
- [官方扩展控制文件 (pgx-s3sign.control)](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/pgx-s3sign.control)
- [官方实现源代码](https://github.com/thebf/pgx-s3sign/blob/90a0d59f9655470df8c3a6700b94ed711207bf4e/src/lib.rs)

`pgx-s3sign` — 快速签名 S3 请求：由 pgx 创建。在实现相应的安全、审计或访问控制工作流时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION "pgx-s3sign";
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgx_s3sign_pre_get` 是一个扩展函数。
- `pgx_s3sign_pre_put` 是一个扩展函数。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
