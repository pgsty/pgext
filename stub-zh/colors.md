## 用法

来源：

- [官方 README](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/README.md)
- [扩展 SQL](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/sql/colors.sql)
- [C 实现](https://github.com/zejn/pg-colors/blob/f963a30f408bf8cbbb07782478a1ba0dc93ab190/src/colors.c)

`colors` 0.0.1 为两组 CIE L*a*b* 坐标提供四种感知色差计算。它适合比较已转换到相同 CIELAB 白点与观察者约定的颜色；该扩展不会转换 RGB/XYZ 数值，也不携带色彩管理元数据。

### 核心流程

安装扩展后，将两组 L*、a* 与 b* 坐标作为 `double precision` 数值传入：

```sql
CREATE EXTENSION colors;

SELECT delta_e_cie_1976(50, 2.5, -20, 50, 0, -18);
SELECT delta_e_cie_2000(50, 2.5, -20, 50, 0, -18);
```

结果越小，表示两种样本在所选公式下越接近。一个工作流应固定使用同一公式及参数；不同公式的结果不能直接互换。

### 函数索引

- `delta_e_cie_1976(...)` 实现六坐标的 CIE76 欧氏距离。
- `delta_e_cie_1994(...)` 提供采用图形艺术默认值的六坐标重载，以及暴露权重常数的十一参数重载。
- `delta_e_cmc(...)` 提供采用 2:1 可接受性默认值的六坐标重载，以及暴露明度和色度权重的八参数重载。
- `delta_e_cie_2000(...)` 提供六坐标重载，以及暴露明度、色度和色相权重的九参数重载。

所有函数均为严格、不可变函数，并返回 `double precision`。

### 运维说明

实现既不验证坐标的物理范围，也不执行色彩空间转换。比较前应规范化数据，并用领域参考样本验证所选公式。其源码是较早期的小型 C 实现，因此应在目标 PostgreSQL 构建上核验，并将边界值纳入数值回归测试。
