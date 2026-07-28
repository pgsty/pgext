## 用法

来源：

- [官方上游 README](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/readme.md)
- [官方扩展控制文件 (pgtfs.control)](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/pgtfs.control)
- [官方扩展 SQL (pgtfs--0.0.1.sql)](https://github.com/adrianprelipcean/pgtfs/blob/8e2512ce6dd95db1f74b513e366a588fc6ae8ceb/sql/pgtfs--0.0.1.sql)

`pgtfs` — PGTFS 是一个设计用于在 GTFS（通用公交信息格式）格式基础上进行路由的 PostgreSQL 扩展。它提供了查询和分析存储在 PostgreSQL 数据库中的公交数据的功能，使用 GTFS 传输数据。适用于相应的空间数据或地理空间工作流。上游明确表示该项目尚未准备好用于生产环境。

### 核心工作流

```sql
CREATE EXTENSION pgtfs;
```

在目标数据库中安装扩展，当可用时运行上游提供的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `pgtfs_csa(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT)` 是一个扩展函数，返回 `TABLE`。
- `pgtfs_csa(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT, minimize_transfers BOOLEAN DEFAULT FALSE)` 是一个扩展函数，返回 `TABLE`。
- `pgtfs_raptor(origin TEXT, destination TEXT, departure_time DOUBLE PRECISION, network TEXT, max_rounds INT DEFAULT 5)` 是一个扩展函数，返回 `TABLE`。
- `pgtfs_version()` 是一个扩展函数，返回 `TEXT`。

### 要求与注意事项

- 控制文件声明默认版本为 `0.0.4`。
- 控制文件标记该扩展为可重定位。
- 上游明确表示该项目尚未准备好用于生产环境。
- 上游描述该项目仍处于开发阶段。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
