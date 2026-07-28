## 用法

来源：

- [官方扩展控制文件（pboutput.control）](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/pboutput.control)
- [官方实现源代码](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/src/lib.rs)
- [官方 Rust 包清单](https://github.com/semtexzv/pboutput/blob/2a1062eee71acec95866a15086edb5060caf1d66/Cargo.toml)

`pboutput` — 二进制逻辑解码输出插件，将 PostgreSQL 的更改编码为 Protocol Buffers。在移动、转换或集成相应的数据时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION pboutput;
```

在目标数据库中安装扩展，当可用时运行最小的上游示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 要求与注意事项

- 该目录记录了版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
