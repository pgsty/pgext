## 用法

来源：

- [官方上游 README](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/README.md)
- [官方扩展控制文件 (xata_id_extension.control)](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/xata_id_extension.control)
- [官方实现源代码](https://github.com/tsirysndr/xata_id_extension/blob/bf2071f81c25d29e67b11d821a1f81300f78a6fc/src/lib.rs)

`xata_id_extension` — 生成带有 rec_ 前缀和 20 个随机字符（a-z0-9）的 24 位唯一 ID。当 SQL 需要这些特殊功能或聚合时使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION xata_id_extension;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `xata_id()` 是一个扩展函数。

### 要求与注意事项

- 该目录记录了扩展的版本 `0.0.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 控制文件将扩展标记为不受信任。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源代码进行比对。
