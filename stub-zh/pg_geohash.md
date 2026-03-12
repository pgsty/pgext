

## 用法

> [pg_geohash: PostgreSQL 的 Geohash 函数](https://github.com/jistok/pg_geohash)

基于 C 实现的 PostgreSQL Geohash 函数（同时支持 HAWQ 和 Greenplum）。基于 [libgeohash](https://github.com/lyokato/libgeohash) C 库。

Geohash 背景知识：[维基百科：Geohash](http://en.wikipedia.org/wiki/Geohash)

### 函数

将经纬度编码为指定精度的 geohash 字符串：

```sql
SELECT LAT_LON_TO_GEOHASH_WITH_LEN(latitude, longitude, 5) AS geohash;
```

将经纬度编码为全精度的 geohash：

```sql
SELECT LAT_LON_TO_GEOHASH(latitude, longitude) AS geohash;
```

将 geohash 解码回经纬度：

```sql
SELECT GEOHASH_TO_LAT_LON('dp3w7') AS lat_lon;
```

### 示例

使用 5 字符精度（约 2.4km x 4.9km 网格）计算基于 geohash 的聚合：

```sql
SELECT LAT_LON_TO_GEOHASH_WITH_LEN(latitude, longitude, 5) AS geohash,
       COUNT(*)
FROM crimes
GROUP BY 1
ORDER BY 2 DESC
LIMIT 10;
```

```
 geohash | count
---------+-------
 dp3w7   | 72404
 dp3tt   | 70713
 dp3tw   | 63642
 dp3wm   | 62332
 dp3wk   | 56467
```

从 geohash 恢复坐标：

```sql
SELECT location,
       GEOHASH_TO_LAT_LON(LAT_LON_TO_GEOHASH(latitude, longitude))
FROM crimes
LIMIT 5;
```
