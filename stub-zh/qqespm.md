## 用法

来源：

- [README.md](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/README.md)
- [qqespm--1.0.sql](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/qqespm--1.0.sql)
- [qqespm.control](https://github.com/viniciuscva/qqespm_postgres_extension/blob/9c7d03a51a831e7066096c22e4888fe0736cb06f/qqespm.control)

`qqespm` 针对地理文本对象构造并执行空间模式查询。它使用 PostGIS，由 PL/pgSQL 实现；control 中的库路径并不对应实际提供的二进制库。

### 核心用法

```sql
CREATE EXTENSION postgis;
CREATE EXTENSION qqespm;
SELECT * FROM match_spatial_pattern(
  ARRAY[distance_constraint('school', 'pharmacy', 10, 10000, true, false)],
  'pois', ARRAY['school', 'pharmacy']);
```

### 运行边界

查询前先安装 PostGIS 并准备要求的 POI 表：默认约定包含 osm_id、geometry、centroid 以及存放关键词的列。`distance_constraint` 和 `connectivity_constraint` 构造 JSONB 约束；`match_spatial_pattern` 将约束转成 SQL，并返回匹配对象的标识符。查询生成辅助函数可展示中间 SQL。文档示例中的距离单位是米。

README 展示了 PostgreSQL 16 的安装路径，但没有声明支持的大版本矩阵。无需预加载；control 允许迁移且未标记为可信扩展。实现会构造动态 SQL，因此表名、关键词列和查询约束应来自可信输入。使用代表性数据检查生成计划及适当的 PostGIS 索引。项目在 README 中声明采用 CC BY 4.0。
