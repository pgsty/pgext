## 用法

来源：

- [官方用法文档](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/doc/kmeans.md)
- [扩展 SQL](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.sql.in)
- [C 实现](https://github.com/umitanuki/kmeans-postgresql/blob/d6a57eda9e4908ef5768ad2a5c70bd92d0abb3ae/kmeans.c)

`kmeans` 版本 `1.1.0` 把 K-means 聚类实现为 PostgreSQL 窗口函数。它为每个输入行返回从零开始的簇编号，适合规模相对较小的库内向量分类。

### 核心流程

```sql
CREATE EXTENSION kmeans;

SELECT id,
       kmeans(ARRAY[x, y, z]::float8[], 3) OVER () AS cluster
FROM samples;
```

同一窗口分区中的全部向量必须是一维且长度相同。第二个参数是所需簇数。

如需可重复的初始化，可显式提供质心：

```sql
SELECT id,
       kmeans(
         ARRAY[x, y]::float8[],
         2,
         ARRAY[0.5, 0.5, 1.0, 1.0]::float8[]
       ) OVER (PARTITION BY group_key) AS cluster
FROM samples;
```

第三个参数可以是二维质心数组，也可以是长度等于簇数乘以向量维度的一维数组。

### 行为与注意事项

- `kmeans(float[], int) returns int`：先在各维度最小值与最大值之间插值，再从输入向量中选择初始质心。
- `kmeans(float[], int, float[]) returns int`：使用调用方提供的初始质心。

K-means 结果依赖初始化与量纲。应按需要归一化各维度，在重视可重复性时提供固定质心，并测试 NULL 与退化分区。该上游版本发布于 2011 年，面向 PostgreSQL 8.4 时代的 API；应在目标 PostgreSQL 版本上确认源码兼容性与结果。不需要预加载。
