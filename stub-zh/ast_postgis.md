## 用法

来源：

- [官方扩展控制文件](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/ast_postgis.control)
- [官方上游文档](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/README.md)
- [官方一致性检查实现](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/sql/consistency_functions.sql)

`ast_postgis` 1.1.1 为 PostGIS 添加受 OMT-G 启发的空间域和数据库端完整性检查。空间模型需要在数据附近执行简单几何、网络、拓扑或聚合规则时可以使用它。它依赖 `postgis` 并安装事件触发器，因此安装和 DDL 集成需要具备适当权限的管理员。

### 核心流程

创建扩展，使用其域定义列，再附加关系触发器：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION ast_postgis;

CREATE TABLE bus_stop (
  stop_id integer PRIMARY KEY,
  geom ast_point
);

CREATE TABLE route_segment (
  segment_id integer PRIMARY KEY,
  geom ast_uniline
);

CREATE TRIGGER route_segment_nodes
AFTER INSERT OR UPDATE ON route_segment
FOR EACH STATEMENT
EXECUTE PROCEDURE ast_arcnodenetwork(
  'route_segment', 'geom', 'bus_stop', 'geom'
);
```

关系触发器会在写入后验证受影响模型，并在违反配置规则时抛出错误。

### 重要对象

- 几何域包括 `ast_polygon`、`ast_line`、`ast_point`、`ast_node`、`ast_isoline`、`ast_planarsubdivision`、`ast_tin`、`ast_sample`、`ast_uniline` 和 `ast_biline`；`ast_tesselation` 包装 PostGIS raster。
- `ast_spatialrelationship()` 支持 `contains`、`containsproperly`、`covers`、`coveredby`、`crosses`、`disjoint`、`distant`、`intersects`、`near`、`overlaps`、`touches` 和 `within` 触发器规则。
- `ast_arcnodenetwork()`、`ast_arcarcnetwork()` 和 `ast_aggregation()` 执行网络与部分或整体关系。
- `ast_isSpatialRelationshipValid(...)`、`ast_isNetworkValid(...)` 和 `ast_isSpatialAggregationValid(...)` 检查现有数据，并将违规写入 `ast_violation_log`。

### 注意事项

扩展通过语句级触发器、事件触发器、PL/pgSQL 和动态构造的 SQL 实现约束。应只允许可信角色使用检查和触发器管理 API，只传入由管理员控制的标识符，并测试写入并发和大表成本。已复核的 `ast_isSpatialAggregationValid(...)` 源码包含针对 `tablea` 与 `tableb` 的硬编码查询，因此在为实际模型修补并完成回归测试前，不能信任该检查器。在现有数据上启用强制约束前，应在副本中验证每个域和关系；验证记录会进入扩展配置转储，所以还应管理 `ast_violation_log` 的保留策略。
