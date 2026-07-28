## 用法

来源：

- [官方上游 README](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/README.md)
- [官方扩展控制文件 (garbagemap.control)](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/garbagemap.control)
- [官方扩展 SQL (garbagemap--1.0.sql)](https://github.com/masahikosawada/walker/blob/4246967426b8c4c86077fc4bef53b638ef3e6124/garbagemap/garbagemap--1.0.sql)

`garbagemap` — GarbageMap 在执行 CHECKPOINT 时会记录所有垃圾回收的总结信息到 PostgreSQL 服务器日志（日志级别为 LOG）。示例如下。在管理或自动化上述数据库行为时使用它。请使用上述链接的固定上游版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION garbagemap;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小代码，验证安装版本和返回值，然后再将其集成到应用程序 SQL 中。

### 重要对象

- `gs(rel regclass, OUT rangeno INT, OUT freespace INT, OUT n_tuples INT, OUT n_dead_tuples INT, OUT n_all_visible INT, OUT dead_tuple_ratio NUMERIC(15,4))` 是一个扩展函数，返回 `SETOF`。
- `gs_rank(rel regclass, OUT rownum INT, OUT percent_blocks NUMERIC(15,4), OUT rangeno INT, OUT freespace INT, OUT n_tuples INT, OUT n_dead_tuples INT, OUT n_all_visible INT, OUT dead_tuple_ratio NUMERIC(15,4), OUT percent NUMERIC(15,4))` 是一个扩展函数，返回 `SETOF`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况与固定源代码进行比对。
