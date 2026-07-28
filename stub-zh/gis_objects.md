## 用法

来源：

- [官方上游 README](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/README.md)
- [官方扩展控制文件 (gis_objects.control)](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/gis_objects.control)
- [官方扩展 SQL (gis_objects--1.0.sql)](https://github.com/mmiranda2/c-postgres-gis/blob/93ef09fbee9bb6b30282b85001d968f60dd7ec7e/gis_objects--1.0.sql)

`gis_objects` 是一个用 C 实现的 PostgreSQL 空间对象扩展，支持点、多边形和点在多边形内判断等操作。它适用于相应的空间数据或地理空间工作流。请以上方链接的固定上游修订为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION gis_objects;
```

在目标数据库中安装扩展；如果上游提供了最小示例，请运行该示例，并在集成到应用 SQL 前验证安装版本和返回值。

### 重要对象

- `my_log()` 是扩展函数，返回 `varchar`。
- `my_point_in_polygon(mypoint, mypoint[])` 是扩展函数，返回 `boolean`。
- `mypoint_in(cstring)` 是扩展函数，返回 `mypoint`。
- `mypoint_out(mypoint)` 是扩展函数，返回 `cstring`。
- `mypoint` 是扩展定义的类型。

### 要求与注意事项

- 审阅的控制文件声明默认版本为 `1.0`。
- 控制文件将该扩展标记为可重定位。
- 生产使用前，请根据固定版本源码确认权限、支持的 PostgreSQL 版本、升级行为和失败情形。
