## 用法

来源：

- [官方上游 README](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/README.md)
- [官方扩展控制文件 (maludb_core.control)](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/maludb_core.control)
- [官方实现源代码](https://github.com/maludb/maludb-core/blob/b4d6d521cb94cb0c05cfbadd2a9958c2e5ddfc4f/src/maludb_core.c)

`maludb_core` — MaluDB 是一种用于长期机构记忆、人类与人工智能知识共享及情境回忆的内存 DBMS。它以 **C** 语言构建，作为在 **Ubuntu 24.04 LTS** 上的 PostgreSQL 扩展，以 **PostgreSQL 17**（PGDG）为基础。使用它来进行相应的向量、模型或检索工作流。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION maludb_core;
```

在目标数据库中安装扩展，当可用时运行上游最小示例，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 要求与注意事项

- 控制文件声明默认版本为 `0.104.0`。
- 先安装确认的扩展依赖项：`vector`、`btree_gist`、`pg_trgm`、`pgcrypto`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源代码一致。
