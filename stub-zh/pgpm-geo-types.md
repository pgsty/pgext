## 用法

来源：

- [packages/geotypes/README.md](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/README.md)
- [packages/geotypes/sql/pgpm-geo-types--0.47.0.sql](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/sql/pgpm-geo-types--0.47.0.sql)
- [packages/geotypes/Makefile](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/Makefile)
- [packages/geotypes/pgpm-geo-types.control](https://github.com/constructive-io/pgpm-modules/blob/65abdcd9379edfdc9287ce0d64516d63a7837294/packages/geotypes/pgpm-geo-types.control)

`pgpm-geo-types` 在 PostGIS 类型之上定义四个域：`geo_point`、`geo_polygon`、`geography_point` 与 `geography_polygon`。底层点/多边形类型修饰符使用 SRID 4326。

### 核心用法

```sql
CREATE EXTENSION "pgpm-geo-types";
CREATE TABLE geo_demo (id integer PRIMARY KEY, location geo_point);
INSERT INTO geo_demo VALUES (1, ST_SetSRID(ST_MakePoint(-122.4194, 37.7749), 4326));
```

### 运行边界

几何与地理类型保留各自的 PostGIS 计算语义。SQL 只定义类型域和 Graphile 命名注释，不会补充通用的拓扑有效性检查，也不会转换坐标。赋予不同 SRID 前须先转换坐标。安装脚本中的域名称没有模式限定，引用时应使用实际安装模式。 依赖为 `plpgsql`, `citext`, `postgis`, `pgpm-types`, `pgpm-verify`。 这是纯 SQL 模块，无需预加载动态库或重启；该修订未公布 PostgreSQL 主版本支持矩阵。模式布局固定，不可迁移。启用前须部署相匹配的完整依赖。
