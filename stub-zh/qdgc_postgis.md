## 用法

来源：

- [PGXN qdgc 0.1.0 发布页](https://pgxn.org/dist/qdgc/0.1.0/)
- [官方 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [官方 qdgc_postgis 控制文件](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis.control)
- [官方 qdgc_postgis 0.1.0 扩展 SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc_postgis--0.1.0.sql)

`qdgc_postgis` 0.1.0 是纯 SQL 核心扩展 `qdgc` 的 PostGIS 伴生扩展。它可以在 QDGC 单元与 PostGIS 点、多边形之间转换，按 WGS84 椭球计算单元面积，并用 QDGC 单元填充任意 geometry。该扩展同时依赖 `qdgc` 与 `postgis`，不能替代其中任何一个。

### 核心流程

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION qdgc;
CREATE EXTENSION qdgc_postgis;

SELECT qdgc_latlng_to_cell(
    ST_SetSRID(ST_MakePoint(31.4, 2.7), 4326),
    5
);

SELECT qdgc_cell_to_geometry('E031N02ADBAC');
SELECT qdgc_cell_to_boundary_geometry('E031N02ADBAC');
SELECT qdgc_cell_area_km2('E031N02ADBAC');
```

点 geometry 重载会把非零且非 4326 的 SRID 转换到 EPSG:4326；SRID 为零时，则假定输入已经是经纬度坐标。

### 填充关注区域

执行深层级填充前，应先估算结果规模：

```sql
WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT qdgc_estimate_cell_count(geom, 7)
FROM area;

WITH area AS (
    SELECT ST_GeomFromText(
        'POLYGON((31.0 2.0, 31.5 2.0, 31.5 2.5, 31.0 2.5, 31.0 2.0))',
        4326
    ) AS geom
)
SELECT cell
FROM area
CROSS JOIN LATERAL qdgc_polygon_to_cells(
    geom,
    7,
    'intersects'
) AS cell;
```

谓词可选值如下：

- `intersects` 是默认值，返回与输入 geometry 相交的单元；
- `centroid` 返回中心点位于输入 geometry 内的单元；
- `contains` 返回完全位于输入 geometry 内的单元。

实现采用可剪枝四叉树逐层下降，而不是测试 geometry 完整包围盒中的每个单元。多部件 geometry 会按部件分别填充，再合并单元集合。

### 重要对象

- `qdgc_latlng_to_cell(geometry, level)` 及其 `geography` 重载用于编码 PostGIS 点。
- `qdgc_cell_to_geometry` 与 `qdgc_cell_to_geography` 返回单元中心点。
- `qdgc_cell_to_boundary_geometry` 与 `qdgc_cell_to_boundary_geography` 返回矩形单元边界。
- `qdgc_cell_area_km2` 在 WGS84 椭球上测量单元边界对应的 geography 面积。
- `qdgc_polygon_to_cells` 按三种已记录谓词之一填充区域。
- `qdgc_estimate_cell_count` 在真正生成填充结果前提供受包围盒上限约束的低成本估算。

### 运维说明

- `qdgc_postgis.control` 声明了 `requires = 'qdgc,postgis'` 与 `relocatable = true`。应先由具备相应权限的角色安装 PostGIS，再把伴生扩展的使用交给普通用户。
- 不需要 `shared_preload_libraries`、`LOAD` 或重启。该扩展自身只有 SQL，但其 PostGIS 依赖包含本地代码。
- `qdgc`、`qdgc_postgis` 及其被调用依赖应安装到当前 `search_path` 可见的模式中，因为这些可迁移 SQL 函数使用未限定名称调用彼此。
- 上游测试了 PostgreSQL 13–17；不能因为该扩展没有编译代码就推断 PostgreSQL 18 已获支持。
- 即使采用剪枝，深层级区域填充仍可能生成巨量结果。应把 `qdgc_estimate_cell_count` 作为运维保护，并在调用 `qdgc_polygon_to_cells` 前施加应用侧规模限制。

