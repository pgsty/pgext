

## 用法

> [pg_eviltransform: WGS84、GCJ02 和 BD09 之间的坐标转换](https://github.com/aiyou178/pg_eviltransform)

`pg_eviltransform` 为 PostGIS 的 `ST_Transform` 扩展了 BD09/GCJ02 中国坐标系支持。它暴露了 `ST_EvilTransform` 函数，与 `ST_Transform` 具有相同的重载接口。

自定义 SRID：
- `990001`：GCJ02（火星坐标系）
- `990002`：BD09（百度坐标系）

### 函数

```sql
ST_EvilTransform(geometry, to_srid integer)
ST_EvilTransform(geometry, to_proj text)
ST_EvilTransform(geometry, from_proj text, to_srid integer)
ST_EvilTransform(geometry, from_proj text, to_proj text)
```

如果双方都不涉及自定义坐标系，则直接委托给 `ST_Transform`。如果涉及 BD09/GCJ02，会在需要时通过 WGS84（`4326`）进行中转。

### 示例

```sql
-- WGS84 转 GCJ02（使用文本字面量）
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'GCJ02');

-- WGS84 转 BD09（使用文本字面量）
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 'BD09');

-- WGS84 转 GCJ02（使用数字 SRID）
SELECT ST_EvilTransform(ST_SetSRID('POINT(120 30)'::geometry, 4326), 990001);

-- BD09 转 Web Mercator
SELECT ST_EvilTransform(
  ST_SetSRID('POINT(120.011070620552 30.0038830555128)'::geometry, 990002), 3857
);

-- from_proj / to_proj 重载
SELECT ST_EvilTransform('POINT(120 30)'::geometry, 'EPSG:4326', 'GCJ02');
```

### 性能

在 PG18 上处理 200,000 行数据时，`ST_EvilTransform` 比基于正则表达式的 SQL 方法快约 30-45 倍。
