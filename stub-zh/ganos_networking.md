## 用法

来源：

- [阿里云 GanosBase Networking 路径模型](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/path-model)
- [阿里云 RDS PostgreSQL 支持的扩展列表](https://www.alibabacloud.com/help/en/rds/apsaradb-rds-for-postgresql/extensions-supported-by-apsaradb-rds-for-postgresql)

`ganos_networking` 是阿里云 GanosBase Networking 扩展，用于在 ApsaraDB RDS PostgreSQL 上分析道路与交通网络。它提供兼容 pgRouting 的路径函数，包括 `pgr_dijkstra`、`pgr_astar` 和 `pgr_trsp`，可计算最短、最快、启发式及带转向限制的路径。

### 最小路径查询

在 RDS 数据库中安装云厂商提供的扩展。`CASCADE` 允许 PostgreSQL 安装云厂商声明的依赖项。

```sql
CREATE EXTENSION ganos_networking WITH SCHEMA public CASCADE;

SELECT *
FROM pgr_dijkstra(
  'SELECT id, source, target, cost, reverse_cost FROM edge_table',
  2,
  3
);
```

边查询需要提供边标识符、源与目标节点标识符以及通行成本。负数 `cost` 表示正向不可通行；负数 `reverse_cost` 表示反向不可通行。路径结果包含序号、节点、边、单边成本和累计成本列。

### 注意事项

- 这是阿里云的专有云厂商扩展，并非可独立安装的开源软件包。请在受支持实例上通过 RDS 插件市场或特权 `CREATE EXTENSION` 命令安装。
- 可用性与版本取决于 RDS PostgreSQL 主/次引擎版本和实例版本。部署前应查询实时支持表；目录版本 `7.0` 并不保证每个实例都提供该版本。
- 传给路径函数的 SQL 文本定义图与成本模型。应验证其结果结构，并限制哪些用户能够提供动态 SQL。
- A* 还要求坐标列，而转向限制路由需要限制数据。请查阅云厂商的路径模型参考，确认当前引擎版本使用的确切签名。
