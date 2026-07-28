## 用法

来源：

- [官方上游 README](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/README.md)
- [官方扩展控制文件 (pg_rvbbit.control)](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/crates/pg_rvbbit/pg_rvbbit.control)
- [官方实现源代码](https://github.com/ryrobes/rvbbit-sql/blob/6c82cb49a85937ca1ebc0361d703101c1300ae52/crates/pg_rvbbit/src/lib.rs)

`pg_rvbbit` — 基于可配置的大语言模型的语义 SQL 操作符，带有缓存、收据和跨多个执行引擎的路由功能。使用它来进行相应的向量、模型或检索工作流。在目标 PostgreSQL 构建上使用上述链接的固定上游版本作为 API 边界并进行测试。

### 核心工作流

```sql
CREATE EXTENSION pg_rvbbit;
```

在目标数据库中安装扩展，当可用时运行上方最小的上游示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `backend_probe` 是一个扩展函数。
- `backend_probe_with_input` 是一个扩展函数。
- `env_present` 是一个扩展函数。
- `reload_backends()` 是一个扩展函数。
- `rvbbit_build_info()` 是一个扩展函数。
- `rvbbit_version()` 是一个扩展函数。

### 要求与注意事项

- 控制文件声明默认版本为 `4.1.4`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
