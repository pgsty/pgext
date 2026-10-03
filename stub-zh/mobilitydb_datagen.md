## 用法

来源：

- [1.3.1 数据生成器手册](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/doc/data_generator.xml)
- [扩展控制文件](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/mobilitydb_datagen.in.control)
- [时态值生成器](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/temporal/random_temporal.sql)
- [时态点生成器](https://github.com/MobilityDB/MobilityDB/blob/v1.3.1/mobilitydb/datagen/geo/random_tpoint.sql)

`mobilitydb_datagen` 1.3.1 提供 PL/pgSQL 函数，用来生成合成的 PostgreSQL、PostGIS 和 MobilityDB 值。它主要适用于需要随机时态值或轨迹的回归数据、演示和基准测试数据集。

```sql
-- After the main MobilityDB extension is loaded:
CREATE EXTENSION mobilitydb_datagen;
```

### 生成随机时态值

```sql
-- A random temporal float sequence.
SELECT random_tfloat_seq(
    -100.0, 100.0,                                  -- value bounds
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',  -- time bounds
    10.0,                                           -- max value delta
    10,                                             -- max minutes between instants
    5, 10                                           -- min/max instants
);

-- Step interpolation instead of the default linear interpolation.
SELECT random_tfloat_seq(
    -100.0, 100.0,
    '2025-06-01 00:00+00', '2025-06-02 00:00+00',
    10.0, 10, 5, 10,
    false
);

-- A random temporal geometry point sequence.
SELECT asEWKT(
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,                  -- x/y bounds
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,                           -- max delta, max minutes, min/max instants
        srid => 4326
    )
);
```

其他已确认的生成器家族包括 `random_bool`、`random_int`、`random_float`、`random_text`、`random_timestamptz` 等标量辅助函数；数组、集合、span 和 range 辅助函数；`random_tbool_inst`、`random_tint_discseq`、`random_tfloat_seq`、`random_tfloat_seqset` 等 temporal 辅助函数；以及 `random_geom_point`、`random_geom_linestring`、`random_tgeompoint_contseq`、`random_tgeompoint_seqset`、`random_tgeogpoint_contseq`、`random_tgeogpoint_seqset` 等空间/temporal-point 辅助函数。

### 生成测试数据集

为 trip 查询基准测试创建批量测试数据：

```sql
CREATE TABLE trip_samples AS
SELECT
    vehicle_id,
    random_tgeompoint_contseq(
        2.20, 2.50, 48.80, 48.95,
        '2025-06-01 08:00+00', '2025-06-01 18:00+00',
        0.02, 5, 20, 40,
        srid => 4326
    ) AS trip
FROM generate_series(1, 1000) AS vehicle_id;
```

### 注意事项

- control file 要求主 `mobilitydb` 扩展已存在；`mobilitydb_datagen` 不是独立扩展。
- 生成器随 `mobilitydb` 版本 `1.3.1` 分发。应保持依赖与生成器文件一致；可选类型族是否存在取决于构建开关及软件包可用性。
- 上游文档有意省略许多生成器函数的详细参数列表，并让用户查看 SQL 源文件确认精确签名。

生成器默认产生随机合成值，并不保证应用数据集可重现。比较多次运行时应设置明确的随机种子，并核对边界、时区、插值方式和 SRID 是否符合工作负载。
