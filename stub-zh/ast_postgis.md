## 用法

来源：

- [已复核 commit 的 AST-PostGIS README](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/README.md)
- [已复核 commit 的 ast_postgis.control](https://github.com/lizardoluis/ast_postgis/blob/a7b71d3b350cd7c12dfb574ea886bd14c870ce6c/ast_postgis.control)

`ast_postgis` 在 `postgis` 之上增加更丰富的空间类型与空间完整性检查。其类型包括 `ast_point`、`ast_polygon` 和 `ast_uniline`，它们在 PostGIS 几何或栅格表示之上加入适合地理对象与地理场模型的约束。

该扩展还提供 `ast_spatialrelationship`、`ast_arcnodenetwork` 和 `ast_aggregation` 等触发器过程。`ast_isTopologicalRelationshipValid` 等验证函数可在实施约束前检查既有数据，并把不一致记录到 `ast_validation_log`。

### 基本设置

```sql
CREATE EXTENSION IF NOT EXISTS postgis;
CREATE EXTENSION ast_postgis;

CREATE TABLE landmarks (
  id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  name text NOT NULL,
  geom ast_point
);
```

将高级空间类型用作表列类型，然后针对模型必须实施的空间关系添加文档所列的触发器过程。

### 注意事项

- control 文件将 `postgis` 声明为依赖，因此数据库中必须能使用它。
- 上游当前标识版本 1.1.1，并说明在 PostgreSQL 15 上做过测试。请测试部署所用的确切 PostgreSQL/PostGIS 组合。
- 约束过程接收表名和列名。应严格复现上游触发器签名，并在生产数据上启用约束前验证既有行。
