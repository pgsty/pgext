## 用法

来源：

- [postgresql/contrib/temporalbox/temporalbox--1.0.sql](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporalbox/temporalbox--1.0.sql)
- [README.md](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/README.md)
- [postgresql/contrib/temporalbox/temporalbox.control](https://github.com/yash7312/DBIS_Assignment/blob/fc65587474a7e0c1275e3eb831478c5f2482602b/postgresql/contrib/temporalbox/temporalbox.control)

`temporalbox` 是学术时态索引项目中的小型 SQL 辅助扩展，将整数键和时间区间映射为二维 `cube` 值，用于空间索引实验。

### 核心用法

```sql
CREATE EXTENSION cube WITH SCHEMA public;
CREATE EXTENSION temporalbox;
SELECT temporalbox(7, '[2026-01-01,2026-02-01)'::tsrange);
SELECT temporalbox_point(7, TIMESTAMP '2026-01-15');
```

### 运行边界

应先将 `cube` 安装在 `public` 中：函数显式调用 `public.cube`，安装脚本中的依赖创建并未声明在 control 中。`temporalbox` 映射 `tsrange`，`temporalbox_point` 映射单个时间戳，`temporalbox_range` 映射两个边界。时间以浮点纪元秒数表示，无上界区间映射为无穷大，但立方体表示不会保留开闭区间语义；边界精确性重要时应再加原始区间条件。此扩展不含原生库，也不需要预加载。所在项目包含 PostgreSQL 索引内核修改，本条目仅描述 SQL 辅助扩展，不代表验证了整个内核改动或主版本兼容矩阵。
