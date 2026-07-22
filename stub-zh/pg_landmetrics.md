## 用法

来源：

- [官方 README 与指标清单](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/README.md)
- [指标类型定义](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/sql/types/metric.sql)
- [`p_area` 实现与示例](https://github.com/andrearosado/pg_landmetrics/blob/20ea7350111e3b8bfbd15e921b8f239503893801/sql/functions/p_Area.sql)

`pg_landmetrics` 0.0.1 为 PostGIS 矢量数据提供 FRAGSTATS 风格的景观指标，并为 `geometry` 与 `geography` 提供重载。它可用于计算单个斑块、某一类别斑块或整个景观的指标。多数结果是复合指标值，应读取其中字段，而不是将结果当作单个数字。

### 核心流程

先安装 PostGIS，再创建 `pg_landmetrics`：

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION pg_landmetrics;
```

斑块函数每次接收一个几何对象。例如，斑块面积返回 `metric` 复合值；使用默认除数时，`value` 的单位为公顷：

```sql
SELECT id,
       (p_area(geom)).id    AS metric_id,
       (p_area(geom)).value AS hectares
FROM habitat_patches;
```

类别与景观函数是聚合函数。类别指标应按每行代表的类别分组，景观指标则对完整输入集合执行：

```sql
SELECT habitat_class, (c_totalarea(geom)).value AS class_area
FROM habitat_patches
GROUP BY habitat_class;

SELECT (l_numpatches(geom)).value AS patch_count
FROM habitat_patches;
```

### 指标索引

- 斑块函数以 `p_` 开头：`p_area`、`p_corearea`、`p_coreareaindex`、`p_euclideanearestneighbourdistance`、`p_numcoreareas`、`p_perimeter`、`p_perimarearatio` 与 `p_shapeindex`。
- 类别聚合以 `c_` 开头：总面积/核心面积、景观占比、总边长、边缘密度、斑块数与斑块密度。
- 景观聚合以 `l_` 开头：总面积/边长、边缘与斑块密度、斑块数/丰富度、丰富度密度，以及 Shannon 或 Simpson 多样性。
- `metric` 保存元数据标识与数值；`metric_labeled` 和 `metric_labeled_pair` 保存部分指标使用的带标签集合。

### 依赖与注意事项

该扩展依赖 PostGIS。对 `geometry` 而言，指标单位取决于坐标参考系：以米或公顷描述的计算需要适当的投影 CRS，经纬度单位的几何对象不能产生这些单位。若干核心面积与最近邻函数还需要额外参数或有意义的斑块集合；选择阈值前应查看上游对应函数页面。这是早期的 0.0.1 项目，应使用已知数据集核验每个选定指标，并检查复合结果定义，不要假定所有函数签名相同。
