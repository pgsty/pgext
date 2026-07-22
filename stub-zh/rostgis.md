## 用法

来源：

- [官方入门指南](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/docs/user-guide/GETTING_STARTED.md)
- [导出的 SQL 函数](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/src/lib.rs)
- [当前几何实现](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/src/functions.rs)
- [上游 GiST 故障报告](https://github.com/tumluliu/rostgis/blob/5558428d8cf260e1f5e5f4286dca97c6983c5a42/pgrx_gist_cbor_issue.md)

`rostgis` 版本 `0.1.0` 是实验性的 Rust/pgrx 空间扩展，模仿了 PostGIS 命名接口的一小部分。只能在隔离的开发数据库中使用：它不是 PostGIS 的兼容替代品，而且 SQL 对象名称可能与 PostGIS 对象冲突。

### 核心流程

```sql
CREATE EXTENSION rostgis;

SELECT rostgis_version();
SELECT ST_AsText(ST_MakePoint(1.0, 2.0));
SELECT ST_AsText(ST_GeomFromText('LINESTRING(0 0, 3 4)'));
SELECT ST_Distance(ST_MakePoint(0, 0), ST_MakePoint(3, 4));
SELECT ST_SRID(ST_SetSRID(ST_MakePoint(1, 2), 4326));
```

源码解析器处理基本的 `POINT`、`LINESTRING` 和单环 `POLYGON` WKT。常见导出名称包括 `ST_MakePoint`、`ST_GeomFromText`、`ST_AsText`、`ST_AsGeoJSON`、`ST_X`、`ST_Y`、`ST_SRID`、`ST_SetSRID`、`ST_Distance`、`ST_Area`、`ST_Length`、`ST_Perimeter` 和 `ST_Equals`。

### 不完整语义

若干导出名称只是占位或近似实现：

- `ST_GeomFromWKB` 总是报告尚未实现 WKB 解析，而 `ST_AsWKB` 只是给 WKT 文本加前缀，并不生成 WKB。
- `ST_MakePointZ` 会丢弃 Z，`ST_Z` 返回 NULL。
- `ST_Distance` 只计算点到点距离，对其他几何对象组合返回零。
- `ST_Intersects` 与 `ST_Contains` 使用包围盒测试，并非精确几何谓词。

不要把这些函数用于要求正确性的空间分析。

### 索引与升级边界

虽然仓库包含 GiST 支持代码与示例，但其自身故障报告记录了 pgrx/CBOR 序列化导致的索引创建失败。此版本不得声称或依赖可工作的 GiST 索引。扩展控制文件要求超级用户安装且不可重定位。应把已存储的自定义 `geometry` 值视为与该实验构建严格绑定，并在升级或删除前导出数据。
