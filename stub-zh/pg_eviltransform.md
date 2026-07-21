## 用法

来源：

- [Official v0.0.4 README](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/README.md)
- [v0.0.4 release notes](https://github.com/aiyou178/pg_eviltransform/releases/tag/v0.0.4)
- [v0.0.4 control file](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/pg_eviltransform.control)
- [v0.0.4 upgrade SQL](https://github.com/aiyou178/pg_eviltransform/blob/v0.0.4/pg_eviltransform--0.0.3--0.0.4.sql)

`pg_eviltransform` 扩展了 PostGIS，增加了涉及中国 GCJ-02 和 BD-09 系统的坐标转换。版本 `0.0.4` 还通过 `ST_JenksBins` 数组和聚合重载添加了精确的 Jenks 自然断点分类。

### 坐标转换

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_eviltransform;

-- WGS84 to GCJ-02 using a readable coordinate-system name.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120 30)'::geometry, 4326),
    'GCJ02'
);

-- BD-09 to Web Mercator.
SELECT ST_EvilTransform(
    ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002),
    3857
);
```

自定义 SRID 为：`990001` 对应 GCJ-02，`990002` 对应 BD-09。当两端均未使用自定义系统时，`ST_EvilTransform` 将委托给 PostGIS 的 `ST_Transform`；否则在必要时通过 WGS84 (`4326`) 转换。

### Jenks 自然断点

```sql
-- Array form; NULL elements are ignored.
SELECT ST_JenksBins(ARRAY[1, 2, NULL, 10, 11]::numeric[], 2);

-- Streaming aggregate form for a large table.
SELECT ST_JenksBins(value, 7)
FROM measurements;

-- Return lower rather than upper bin edges.
SELECT ST_JenksBins(value, 7, true)
FROM measurements;
```

数组输入支持 `numeric`, `double precision`, `real`, `bigint`, `integer`, 和 `smallint`。聚合输入为 `numeric` 或 `double precision`；当需要时请将其他数值列转换为这些类型。

### API 索引和注意事项

- `ST_EvilTransform(geometry, integer|text)` 和 `ST_EvilTransform(geometry, text, integer|text)`：四个重载对应于 PostGIS 的 `ST_Transform` 接口。
- `ST_JenksBins(values[], breaks [, invert])`：对数组进行分类并返回 `double precision[]` 边界值。
- `ST_JenksBins(value, breaks [, invert])`：流式聚合，避免生成 `array_agg`。
- PostGIS 是运行时先决条件，在安装 `pg_eviltransform` 之前必须已安装。
- Jenks 输入必须是有限的且 `breaks` 至少为一。`numeric` 值会被转换为有限的 `f64`，因此返回的边界值为浮点数。
- 当唯一值的数量不超过 `breaks` 时，结果将是排序后的唯一值集合；没有有效输入行将返回 `NULL`。
