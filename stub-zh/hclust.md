## 用法

来源：

- [固定提交的扩展 control 文件](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust.control)
- [0.0.1 版安装 SQL](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust--0.0.1.sql)
- [固定提交的聚类实现](https://github.com/IrNesterova/hierarchicalClusteringForDiploma/blob/0e9890797682dced54eca9a857d9eb1c11664f6a/hclust.c)

hclust 0.0.1 版是一个毕业设计性质的 C 层次聚类函数。它通过 SPI 执行传入的 SELECT，把第一列作为数字 row ID、其余列作为 feature，给指定表增加 cluster_id 列，并更新每个选中行的聚类编号。

### 实验性调用

六个参数依次为查询文本、目标 cluster 数、目标表名、linkage selector、distance selector 和 hierarchy selector。一个最小数字实验如下：

```sql
CREATE EXTENSION hclust;

CREATE TABLE hclust_points (
    id integer PRIMARY KEY,
    x double precision,
    y double precision
);

INSERT INTO hclust_points VALUES
    (1, 0.0, 0.0),
    (2, 0.1, 0.1),
    (3, 9.0, 9.0),
    (4, 9.1, 9.1);

SELECT hclust(
    'SELECT id, x, y FROM hclust_points ORDER BY id',
    2,
    'hclust_points',
    1,
    1,
    'a'
);

SELECT id, cluster_id FROM hclust_points ORDER BY id;
```

selector 并不是有文档的 SQL 常量，而是直接编码在 C 中。审阅源码显示 linkage 1 至 4 对应 average、centroid、complete 和 single linkage；distance 1 至 6 对应 Euclidean、Chebyshev、Manhattan、cosine、Jaccard 或 Sorensen；a/A 表示 agglomerative clustering。

### 严重实现限制

项目没有 README、测试、发行说明或兼容矩阵。实现使用固定 100 feature 和短 label buffer，以极少验证从文本转换数值，输出大量服务器消息，并包含未完成或错误的数学路径。表名和 row label 通过固定大小的 sprintf buffer 拼进 ALTER/UPDATE，没有做 identifier 或 literal quoting；传入的 SELECT 也会原样执行。

只有超级用户才应在隔离的一次性数据库中测试 hclust，并使用硬编码的可信标识符和输入。不要向用户开放函数，不要对重要表运行，也不要把其聚类结果视为已经验证的分析输出。
