## 用法

来源：

- [官方上游 README](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/README.md)
- [官方扩展控制文件 (pg_deltax.control)](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/pg_deltax.control)
- [官方实现源代码](https://github.com/xataio/deltax/blob/3e4d6034ecefd8d8ef58f55f550e6c2d3dea488c/src/lib.rs)

`pg_deltax` — DeltaX (δx) 是一个为时间序列数据提供压缩和列存储的 PostgreSQL 扩展。使用它来进行相应的调度、时间或时间序列工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_deltax;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 表记录版本 `0.2.1`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
