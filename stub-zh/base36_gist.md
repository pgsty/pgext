## 用法

来源：

- [官方上游 README](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/README.md)
- [官方扩展控制文件 (base36_gist.control)](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/base36_gist.control)
- [官方扩展 SQL (base36_gist--1.0.sql)](https://github.com/dimitri/base36/blob/d1bc4175148c80b125f566c3cc577383687a6162/base36_gist--1.0.sql)

`base36_gist` — 这种数据类型只是一个快速的 hack，只有演示质量，距离生产就绪还差很远。如果你需要它来做严肃的工作，请考虑向该项目提交问题，甚至可以附带一个补丁（他们称之为“拉取请求”）。在应用程序需要这种特定数据库功能时使用它。在安装扩展及其依赖项并验证安装版本和返回值之前，请勿将其集成到应用程序 SQL 中。

### 核心工作流

```sql
CREATE EXTENSION base36_gist;
```

在目标数据库中安装扩展，运行可用的最小上游示例，并在将其集成到应用程序 SQL 之前验证安装的版本和返回值。

### 重要对象

- `base36_dist(base36, base36)` 是一个扩展函数，返回 `base36`。
- `gist_base36_ops` 是一个扩展定义的操作类。

### 要求与注意事项

- 控制文件声明默认版本为 `1.0`。
- 首先安装确认的扩展依赖项：`btree_gist`，`base36`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
