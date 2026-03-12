

## 用法

> [tzf: PostgreSQL 快速时区查找](https://github.com/ringsaturn/pg-tzf)

根据坐标查找时区名称。基于 [tzf-rs](https://github.com/ringsaturn/tzf-rs) 构建，使用来自 [timezone-boundary-builder](https://github.com/evansiroky/timezone-boundary-builder) 的时区边界数据。

```sql
CREATE EXTENSION tzf;
```

### 函数

根据坐标（经度、纬度）查找时区：

```sql
SELECT tzf_tzname(116.3883, 39.9289) AS timezone;
-- Asia/Shanghai
```

批量查找坐标的时区：

```sql
SELECT unnest(
  tzf_tzname_batch(
    ARRAY[-74.0060, -118.2437, 139.6917],
    ARRAY[40.7128, 34.0522, 35.6895]
  )
) AS timezones;
-- America/New_York
-- America/Los_Angeles
-- Asia/Tokyo
```

根据点查找时区：

```sql
SELECT tzf_tzname_point(point(-74.0060, 40.7128)) AS timezone;
-- America/New_York
```

批量查找点的时区：

```sql
SELECT unnest(
  tzf_tzname_batch_points(
    ARRAY[
      point(-74.0060, 40.7128),
      point(-118.2437, 34.0522),
      point(139.6917, 35.6895)
    ]
  )
) AS timezones;
-- America/New_York
-- America/Los_Angeles
-- Asia/Tokyo
```

### 性能

| 函数 | TPS | 备注 |
|------|-----|------|
| `tzf_tzname` | ~17,700 | 单坐标查找 |
| `tzf_tzname_point` | ~17,600 | 单点查找 |
| `tzf_tzname_batch` | ~51 | 1000 个一批 ≈ ~51,000 TPS |
| `tzf_tzname_batch_points` | ~32 | 1000 个一批 ≈ ~32,000 TPS |
