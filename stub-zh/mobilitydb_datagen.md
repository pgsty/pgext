

## 用法

> [mobilitydb_datagen: MobilityDB 的合成移动数据生成器](https://github.com/MobilityDB/MobilityDB)

MobilityDB DataGen 提供用于生成合成移动数据的函数，用于测试和基准测试 MobilityDB 工作负载。它可以创建随机的时态值，包括行程、轨迹和时变测量数据。

### 生成随机时态值

```sql
-- 在时间跨度内生成随机时态浮点数
SELECT random_tfloat(
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',  -- 时间跨度
    0.0, 100.0,                                      -- 值范围
    10                                               -- 瞬时值数量
);

-- 生成随机时态几何点（轨迹）
SELECT random_tgeompoint(
    '2025-06-01 08:00+00', '2025-06-01 18:00+00',   -- 时间跨度
    ST_MakeEnvelope(2.2, 48.8, 2.4, 48.9, 4326),    -- 空间范围
    20                                               -- 瞬时值数量
);
```

### 生成测试数据集

批量创建测试数据用于行程查询基准测试：

```sql
INSERT INTO trips (vehicle_id, trip, trip_date)
SELECT
    i,
    random_tgeompoint(
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        ST_MakeEnvelope(2.2, 48.8, 2.5, 48.9, 4326),
        50
    ),
    '2025-06-01'
FROM generate_series(1, 1000) AS i;
```
