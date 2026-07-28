## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/README.md)
- [官方扩展控制文件 (graph_component.control)](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/graph_component.control)
- [官方扩展 SQL (graph_component--1.0.0.sql)](https://api.pgxn.org/src/graph_component/graph_component-1.0.1/graph_component--1.0.0.sql)

`graph_component` — 在纯 PostgreSQL 上计算图组件非常困难。使用此扩展，您可以非常高效地完成此任务。扩展通过指针在最小的内存消耗下进行计算。这使得您能够基于数以百万计的边对构建数以十万计顶点的组件，仅需几秒钟。当 SQL 需要这些特殊函数或聚合时，请使用它。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION graph_component;

SELECT
  get_component(graph_components(array[1,2,3,4,5]))
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证安装的版本和返回值。

### 重要对象

- `get_component(state graph_component_hashmap)` 是一个扩展函数，返回 `SETOF int4`。
- `get_component_id(state graph_component_hashmap)` 是一个扩展函数，返回 `TABLE`。
- `graph_components_final(state internal)` 是一个扩展函数，返回 `graph_component_hashmap`。
- `graph_components_step_arr(state internal, vertex int[])` 是一个扩展函数，返回 `internal`。
- `graph_components` 是由扩展公开的聚合函数。
- `graph_component_hashmap` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为以及失败情况，与固定源进行比对。
