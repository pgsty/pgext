## 用法

来源：

- [官方上游 README](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/README.md)
- [官方扩展控制文件 (pgtokdb.control)](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/pgtokdb.control)
- [官方扩展 SQL (pgtokdb--0.0.1.sql)](https://github.com/hughhyndman/pgtokdb/blob/5e79929cbc13fcb69cdb36d90ce2af0723687ab5/pgtokdb--0.0.1.sql)

`pgtokdb` — 该项目实现了允许 PostgreSQL 进程通过其 SQL 接口访问 kdb+ 数据的扩展。当 PostgreSQL 需要通过外部数据接口访问相应的外部数据源时，请使用此扩展。请在目标 PostgreSQL 构建中测试链接的上游固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pgtokdb;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgtokdb.genddl(varchar, varchar, varchar, varchar)` 是一个扩展函数，返回 `setof`。
- `pgtokdb.getstatus(varchar)` 是一个扩展函数，返回 `setof`。
- `pgtokdb.genddl_t` 是一个由扩展定义的类型。
- `pgtokdb.getstatus_t` 是一个由扩展定义的类型。
- `pgtokdb` 是一个由扩展创建的模式。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `0.0.1`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行比对。
