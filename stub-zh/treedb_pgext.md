## 用法

来源：

- [官方上游 README](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/README.md)
- [官方扩展控制文件 (treedb_pgext.control)](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/treedb_pgext.control)
- [官方扩展 SQL (treedb_pgext--1.0.sql)](https://github.com/jbylund/gomap_tam/blob/a06f376f5f8b10e8e0daa3ede248edd03eec0178/treedb_pgext--1.0.sql)

`treedb_pgext` — 一个基于 TreeDB 的 PostgreSQL 表访问方法，TreeDB 是一个用 Go 编写的内存映射 B+树存储引擎。当应用程序需要这种特定的数据库功能时使用它。使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION treedb_pgext;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `treedb_am_handler(internal)` 是一个扩展函数，返回 `table_am_handler`。
- `treedb` 是一个扩展定义的访问方法。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为不可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
