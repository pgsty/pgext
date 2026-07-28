## 用法

来源：

- [官方上游 README](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/README.md)
- [官方扩展控制文件 (querymem.control)](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/querymem.control)
- [官方扩展 SQL (querymem--1.0.sql)](https://github.com/pgedge/querymem/blob/6ed191e4f65cadc22077734c68feb9944b133f80/querymem--1.0.sql)

`querymem` — > 不要将 work_mem 设置得过高；每个查询节点都可以使用这么多内存。在收集或解释相应的 PostgreSQL 统计信息时使用它。在目标 PostgreSQL 构建中测试上游提供的固定版本作为 API 边界。

### 核心工作流

```sql
CREATE EXTENSION querymem;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行它，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `get_query_mem(text)` 是一个扩展函数，返回 `INT`。

### 要求与注意事项

- 审查过的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用之前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况，与固定源进行比对。
