## 用法

来源：

- [官方上游 README](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/readme.txt)
- [官方扩展控制文件 (pgrao.control)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/vagrant/postgres17/pgrao/pgrao.control)
- [官方扩展 SQL (pgrao--1.0.0.sql)](https://github.com/raogaru/devops/blob/d431d4f2a67e4d146602db0a7d59dfd11e7abe8a/vagrant/postgres17/pgrao/pgrao--1.0.0.sql)

`pgrao` — PGRAO Postgres DB Catalog 扩展。在管理或自动化上述数据库行为时使用它。在集成到应用程序 SQL 之前，必须先安装并验证其扩展依赖项。

### 核心工作流

```sql
CREATE EXTENSION pgrao;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 之前验证已安装的版本和返回值。

### 重要对象

- `rao_database` 是一个扩展定义视图。
- `rao_domain` 是一个扩展定义视图。
- `rao_event_trigger` 是一个扩展定义视图。
- `rao_extension` 是一个扩展定义视图。
- `rao_fdw` 是一个扩展定义视图。
- `rao_index` 是一个扩展定义视图。
- `rao_language` 是一个扩展定义视图。
- `rao_mview` 是一个扩展定义视图。
- `rao_part_tables` 是一个扩展定义视图。
- `rao_role` 是一个扩展定义视图。
- `rao_routine` 是一个扩展定义视图。
- `rao_schema` 是一个扩展定义视图。
- `rao_sequence` 是一个扩展定义视图。
- `rao_table` 是一个扩展定义视图。

### 要求与注意事项

- 审核的控制文件声明默认版本为 `1.0.0`。
- 先安装并验证确认的扩展依赖项：`plpgsql`。
- 控制文件将扩展标记为可重定位。
- 控制文件要求超级用户进行安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
