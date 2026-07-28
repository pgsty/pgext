## 用法

来源：

- [官方上游 README](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/README.md)
- [官方扩展控制文件 (pggeojson.control)](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/pggeojson.control)
- [官方扩展 SQL (pggeojson--1.0.sql)](https://github.com/oncox/pggeojson/blob/e1a1c64218ab215b08325f2f8b38cb333342a786/pggeojson--1.0.sql)

`pggeojson` — pgGeoJSON 是一个 PostgreSQL 模块，提供生成 GeoJSON 输出的额外功能。使用它来处理相应的空间数据或地理空间工作流。在目标 PostgreSQL 构建中测试链接的上游修订版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION pggeojson;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `PGG_AsGeoJSON` 是一个扩展函数。
- `PGG_AsGeoJSON` 是一个扩展函数。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件标记该扩展为不可重定位。
- 在生产使用之前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源进行验证。
