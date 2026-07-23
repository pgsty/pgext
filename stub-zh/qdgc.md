## 用法

来源：

- [PGXN qdgc 0.1.0 发布页](https://pgxn.org/dist/qdgc/0.1.0/)
- [官方 0.1.0 README](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/README.md)
- [官方 qdgc 控制文件](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc.control)
- [官方 qdgc 0.1.0 扩展 SQL](https://api.pgxn.org/src/qdgc/qdgc-0.1.0/qdgc--0.1.0.sql)

`qdgc` 0.1.0 是 QDGC 扩展家族中可信、可迁移、完全由 SQL 实现的核心扩展。它可以把经纬度编码为扩展四分之一度网格单元编码，解码单元边界，沿前缀层级导航，查询层级指标，并按经纬度包围盒生成网格。它不依赖 PostGIS 或本地动态库；geometry、geography 与多边形填充能力由伴生扩展 `qdgc_postgis` 提供。

### 核心流程

```sql
CREATE EXTENSION qdgc;

-- qdgc_encode uses (longitude, latitude, level).
SELECT qdgc_encode(31.4, 2.7, 5);
-- E031N02ADBAC

-- The h3-style alias reverses the coordinate arguments.
SELECT qdgc_latlng_to_cell(2.7, 31.4, 5);

SELECT *
FROM qdgc_cell_to_bounds('E031N02ADBAC');

SELECT qdgc_cell_to_parent('E031N02ADBAC', 3);
SELECT * FROM qdgc_cell_to_children('E031N02AD', 5);
```

QDGC 的层级直接编码在文本中：子单元编码以前缀形式包含父单元编码，因此可以直接进行汇总和后代单元查询：

```sql
CREATE TABLE observations (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    qdgc_code text NOT NULL
);

CREATE INDEX observations_qdgc_idx ON observations (qdgc_code);

SELECT qdgc_cell_to_parent(qdgc_code, 3) AS level_3_cell,
       count(*)
FROM observations
GROUP BY 1;

SELECT *
FROM observations
WHERE qdgc_code LIKE 'E031N02AB%';
```

### 包围盒与层级指标

核心扩展不依赖 PostGIS 即可枚举矩形区域。跨越反子午线时，令 `min_lon > max_lon`。

```sql
SELECT qdgc_bbox_cell_count(30.0, 1.0, 32.0, 3.0, 7);

SELECT *
FROM qdgc_bbox_to_cells(30.0, 1.0, 32.0, 3.0, 7);

SELECT qdgc_level_degrees(7);
SELECT qdgc_get_num_cells(7);
SELECT qdgc_average_cell_area(7, 2.0, 'km^2');
SELECT qdgc_version();
```

`qdgc_average_cell_area` 给出球面近似值。若需要针对具体单元按 WGS84 椭球计算面积，应使用 `qdgc_postgis` 提供的 `qdgc_cell_area_km2`。

### 重要对象

- `qdgc_encode(lon, lat, level)` 与 `qdgc_latlng_to_cell(lat, lng, level)` 用于生成编码；两者的坐标参数顺序有意保持不同。
- `qdgc_is_valid_cell`、`qdgc_get_level`、`qdgc_cell_to_bounds`、`qdgc_cell_to_lonlat` 与 `qdgc_cell_to_latlng` 用于校验和解码。
- `qdgc_cell_to_parent` 与 `qdgc_cell_to_children` 用于沿四叉前缀层级导航。
- `qdgc_bbox_to_cells` 枚举与包围盒相交的单元；`qdgc_bbox_cell_count` 只计算数量，不生成结果集。
- `qdgc_level_degrees`、`qdgc_get_num_cells` 与 `qdgc_average_cell_area` 返回网格层级指标。

### 运维说明

- 上游要求 PostgreSQL 13 或更高版本，并测试了 PostgreSQL 13–17；PostgreSQL 18 不在 0.1.0 已发布的测试矩阵内。
- 控制文件设置了 `trusted = true` 与 `relocatable = true`，不需要 `shared_preload_libraries`、`LOAD`、重启服务或本地动态库。
- 可迁移函数之间使用未限定名称互相调用，因此应把 `qdgc` 安装到当前 `search_path` 包含的模式中；默认的 `public` 模式满足这一条件。
- 坐标单位是经纬度角度。`qdgc_encode` 先接收经度，`qdgc_latlng_to_cell` 则先接收纬度。
- 每增加一级，后代结果数量会扩大四倍。生成包围盒网格前先计算数量，不要在没有明确结果规模上限时请求很深的后代层级。
