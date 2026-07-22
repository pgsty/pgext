## 用法

来源：

- [官方扩展控制文件](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm.control)
- [官方 README](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/README.md)
- [官方扩展 SQL](https://github.com/rtosm/rtosm/blob/fb8fbec1cfb127f6ea07abc7db3f79b9a2b2a7f7/rtosm--0.1.sql)

`rtosm` 0.1 为 `openstreetmap-website` 所用的 OpenStreetMap Rails API 数据库实现带误差界的实时空间对象简化。它为 way 和 relation 构建 EB-tree 元数据，再为指定边界框返回数量受限的节点 ID；它不是通用的 PostGIS 几何简化器。

### 前置条件与构建

只能安装到已具备准确 OpenStreetMap API schema 的数据库中。扩展 SQL 会直接引用 `nwr_enum`、`current_nodes`、`way_nodes`、`relation_members` 和 `relation_tags` 等对象；在普通 PostgreSQL 数据库中执行 `CREATE EXTENSION rtosm` 会失败。还应先安装声明的 `intarray` 依赖。

```sql
CREATE EXTENSION intarray;
CREATE EXTENSION rtosm;

SELECT build_way_trees();
SELECT build_relation_trees();
SELECT build_indexes();
SELECT tileid_c(30000);

SELECT unnest(
    vquery_c(-0.15, 51.48, -0.05, 51.55, 5000, 0.2)
) AS node_id;
```

这四个构建调用可能扫描并物化 OSM 数据集的很大一部分。应在维护窗口执行，监控空间与锁，并先在副本上验证结果，再用于在线 API 数据库。

### 主要对象

- 持久元数据表：`tile_nums`、`node_vis_errors`、`way_trees` 和 `relation_trees`。
- 构建函数：`build_way_trees()`、`build_relation_trees()`、`build_indexes()`，以及包括 `tileid_c(integer)` 在内的 tile 辅助函数。
- 查询：`vquery_c(x1, y1, x2, y2, k, e) RETURNS bigint[]`；前四个参数以度表示边界框，`k` 限制返回节点数，`e` 是误差上限。
- 底层 C 与 PL/pgSQL 辅助函数实现 EB-tree 构建、空间过滤和序列编辑。

主查询结果是组装简化 OSM 对象所需的节点 ID。调用方必须把它们连接回 OSM schema，并自行构造所需响应或可视化。

### 模式与维护边界

0.1 版是 2016 年的历史研究代码，假定当时的 Rails schema 与坐标约定。安装前必须对照实际部署的 `openstreetmap-website` 版本，核对每个被引用的类型、表、列和成员编码。

源码提供了编辑辅助函数，但没有记录与 Rails 写入路径自动集成的触发器。不要假定 OSM 编辑后 `way_trees`、`relation_trees` 或 `node_vis_errors` 会保持同步；应用必须定义并测试更新/重建流程。应备份这些派生表，或把重建时间纳入恢复规划。
