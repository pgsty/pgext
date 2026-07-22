## 用法

来源：

- [官方归档 README](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/README.md)
- [官方扩展 control 模板](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/cartodb.control.in)
- [官方 0.37.1 发行说明](https://github.com/CartoDB/cartodb-postgresql/blob/e0a0597061bef6cdd7d7239ffa07ab44491817f1/NEWS.md)
- [官方 SQL 源码目录](https://github.com/CartoDB/cartodb-postgresql/tree/e0a0597061bef6cdd7d7239ffa07ab44491817f1/scripts-available)

`cartodb` 是历史上的服务端扩展，用于把 PostgreSQL 数据库转换成 CartoDB 用户数据库。它把 PostGIS 辅助函数与表元数据、配额、CartoDB 角色约定、表准备、空间网格/分箱、overview、同步和联邦服务器管理组合在一起；并非小型通用地理空间库。

### 安装与基础几何操作

生成的 control 文件把对象固定在 `cartodb` schema，要求 PostGIS 与不受信任的 PL/Python，并要求超级用户安装。PostgreSQL 12+ 构建替换为 `plpython3u`，PostgreSQL 11 构建使用旧语言名称。

```sql
CREATE EXTENSION cartodb CASCADE;

SELECT cartodb.CDB_LatLng(40.4168, -3.7038);

SELECT cartodb.CDB_TransformToWebmercator(
  cartodb.CDB_LatLng(40.4168, -3.7038)
);
```

`CASCADE` 可能安装声明的 PostGIS 与 PL/Python 依赖。在已有数据库启用前，应检查最终扩展集合和可信语言策略。

### 主要接口

- `CDB_CartodbfyTable` 按历史 CartoDB 约定准备表；它可能添加或重建数据库对象，应先在副本上测试。
- `CDB_UserTables`、`CDB_ColumnNames`、`CDB_ColumnType` 与 `CDB_EstimateRowCount` 提供目录/报告辅助能力。
- `CDB_XYZ_Extent`、`CDB_RectangleGrid`、`CDB_HexagonGrid`、分箱函数和 `CDB_TransformToWebmercator` 支持地图查询。
- 同步、ghost table、groups/OAuth、overview、配额与联邦服务器函数会修改元数据，或在配置后连接 Redis、HTTP 与远程数据库。

该 API 假设外围采用历史 CartoDB 角色与配置模型，包括组织/public-user 约定。只安装扩展并不会得到当前 CARTO 平台部署。

### 安全与生命周期

多个管理路径使用动态 SQL、不受信任的 PL/Python、配置表中的凭据以及外连 Redis/HTTP/数据库。应限制 schema/函数权限、保护配置行、约束出站网络，并避免让普通 SQL 用户调用管理辅助函数。

版本 `0.37.1` 把支持基线提升到 PostgreSQL 11，并依赖 PostGIS 与 Python Redis 模块。仓库最后一次功能更新在 2022 年，目前已归档为只读且不提供支持。应把它视为迁移/兼容组件；若没有维护中的分支和完整集成测试，不能假定它兼容当前 PostgreSQL、PostGIS、Python、Redis 或 CARTO 服务。
