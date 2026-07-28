## 用法

来源：

- [官方上游 README](https://github.com/oxgraph/oxgraph/blob/12705eaffc29940595f53b61a58992e1382cc2bc/crates/oxgraph-pgrx/README.md)
- [官方扩展控制文件 (oxgraph_pgrx.control)](https://github.com/oxgraph/oxgraph/blob/12705eaffc29940595f53b61a58992e1382cc2bc/crates/oxgraph-pgrx/oxgraph_pgrx.control)
- [官方实现源代码](https://github.com/oxgraph/oxgraph/blob/12705eaffc29940595f53b61a58992e1382cc2bc/crates/oxgraph-pgrx/src/lib.rs)

`oxgraph_pgrx` — PostgreSQL 扩展界面，用于 OxGraph (pgrx SQL/SPI 粘合剂)。当应用程序需要此特定数据库功能时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION oxgraph_pgrx;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码中的信息一致。
