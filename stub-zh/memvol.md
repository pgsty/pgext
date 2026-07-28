## 用法

来源：

- [官方扩展控制文件 (memvol.control)](https://github.com/grundez/pg_extensions/blob/59ffdb504865ff0838a8fce127cd66a4c1952121/MemoryVolume/memvol.control)
- [官方扩展 SQL (memvol--1.0.sql)](https://github.com/grundez/pg_extensions/blob/59ffdb504865ff0838a8fce127cd66a4c1952121/MemoryVolume/memvol--1.0.sql)

`memvol` — 内存单位。当 SQL 需要这些特殊函数或聚合时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION memvol;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `display_vlm_ref()` 是一个扩展函数，返回 `SETOF`。
- `init_vlm_ref()` 是一个扩展函数，返回 `void`。
- `vlm2vlm(numeric, text, text)` 是一个扩展函数，返回 `numeric`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
