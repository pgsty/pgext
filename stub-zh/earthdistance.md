

## 用法

> [earthdistance: 地球表面大圆距离计算](https://www.postgresql.org/docs/current/earthdistance.html)

`earthdistance` 模块提供两种计算地球表面大圆距离的方法。假设地球为完美球体（需要更高精度请使用 PostGIS）。

```sql
CREATE EXTENSION earthdistance CASCADE; -- 需要 cube
```

## 基于 Cube 的地球距离

数据以 cube 存储，使用 3 个坐标表示距地球中心的 x、y、z 距离。提供了基于 `cube` 类型的 `earth` 域。

### 函数

| 函数 | 说明 |
|------|------|
| `earth()` → `float8` | 返回假定的地球半径 |
| `sec_to_gc(float8)` → `float8` | 将直线（割线）距离转换为大圆距离 |
| `gc_to_sec(float8)` → `float8` | 将大圆距离转换为直线（割线）距离 |
| `ll_to_earth(float8, float8)` → `earth` | 根据经纬度（度）返回位置 |
| `latitude(earth)` → `float8` | 返回纬度（度） |
| `longitude(earth)` → `float8` | 返回经度（度） |
| `earth_distance(earth, earth)` → `float8` | 返回两点之间的大圆距离 |
| `earth_box(earth, float8)` → `cube` | 返回用于 cube `@>` 运算符索引搜索的框 |

### 示例

```sql
-- 纽约到伦敦的距离（米）
SELECT earth_distance(
  ll_to_earth(40.7128, -74.0060),
  ll_to_earth(51.5074, -0.1278)
);

-- 查找某位置 1000 米范围内的所有点（可使用索引）
SELECT *
FROM places
WHERE earth_box(ll_to_earth(40.7128, -74.0060), 1000) @> ll_to_earth(lat, lon);
```


## 基于 Point 的地球距离

将地球位置表示为 `point` 类型值，其中第一个分量是经度，第二个是纬度（度）。

### 运算符

| 运算符 | 说明 |
|--------|------|
| `point <@> point` → `float8` | 两点之间的法定英里距离 |

### 示例

```sql
-- 法定英里距离
SELECT point(-74.0060, 40.7128) <@> point(-0.1278, 51.5074);
```

注意：单位固定为法定英里。基于 point 的方法在极点和 ±180° 经线附近有边界条件问题；基于 cube 的表示方式避免了这些不连续性。
