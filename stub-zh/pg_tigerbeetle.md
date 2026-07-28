## 用法

来源：

- [官方上游 README](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/README.md)
- [官方扩展控制文件 (pg_tigerbeetle.control)](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/extension/pg_tigerbeetle.control)
- [官方扩展 SQL (pg_tigerbeetle--0.1.sql)](https://github.com/matdehaast/pg_tigerbeetle/blob/c8e8b581e4b344a6cc084db55d1d19f4aa1ff442/extension/pg_tigerbeetle--0.1.sql)

`pg_tigerbeetle` — 实验性 Zig 扩展，用于从 PostgreSQL 查找 TigerBeetle 账户。使用它进行相应的 SQL 或数据库实用程序工作流。在目标 PostgreSQL 构建中测试上游链接的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pg_tigerbeetle;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装版本和返回值。

### 重要对象

- `lookup_account()` 是一个扩展函数，返回 `TEXT`。
- `query_by_id(int4)` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 审查的控制文件声明默认版本为 `0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
