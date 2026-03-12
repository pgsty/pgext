


## 用法

> [imgsmlr: 使用 Haar 小波变换实现 PostgreSQL 相似图片搜索](https://github.com/postgrespro/imgsmlr)

`imgsmlr` 扩展基于 Haar 小波变换实现了相似图片搜索功能。它提供了两种数据类型和用于将图片转换为可搜索签名的函数。

```sql
CREATE EXTENSION imgsmlr;
```

### 数据类型

| 数据类型 | 存储长度 | 描述 |
|---------|-------:|------|
| `pattern` | 16388 字节 | 图片的 Haar 小波变换结果 |
| `signature` | 64 字节 | pattern 的短表示，用于快速 GiST 索引搜索 |

### 函数

| 函数 | 返回类型 | 描述 |
|------|---------|------|
| `jpeg2pattern(bytea)` | pattern | 将 JPEG 图片数据转换为 pattern |
| `png2pattern(bytea)` | pattern | 将 PNG 图片数据转换为 pattern |
| `gif2pattern(bytea)` | pattern | 将 GIF 图片数据转换为 pattern |
| `pattern2signature(pattern)` | signature | 从 pattern 创建 signature |
| `shuffle_pattern(pattern)` | pattern | 打乱 pattern 以降低对图片位移的敏感度 |

### 运算符

| 运算符 | 左操作数 | 右操作数 | 返回类型 | 描述 |
|--------|---------|---------|---------|------|
| `<->` | pattern | pattern | float8 | 两个 pattern 之间的欧氏距离 |
| `<->` | signature | signature | float8 | 两个 signature 之间的欧氏距离 |

`signature` 类型支持基于 `<->` 运算符的 GiST KNN 索引。

### 示例

从 JPEG 图片创建 pattern 和 signature 表：

```sql
CREATE TABLE pat AS (
    SELECT
        id,
        shuffle_pattern(pattern) AS pattern,
        pattern2signature(pattern) AS signature
    FROM (
        SELECT id, jpeg2pattern(data) AS pattern
        FROM image
    ) x
);

ALTER TABLE pat ADD PRIMARY KEY (id);
CREATE INDEX pat_signature_idx ON pat USING gist (signature);
```

搜索与给定图片最相似的前 10 张图片：

```sql
SELECT id, smlr
FROM (
    SELECT
        id,
        pattern <-> (SELECT pattern FROM pat WHERE id = :id) AS smlr
    FROM pat
    WHERE id <> :id
    ORDER BY signature <-> (SELECT signature FROM pat WHERE id = :id)
    LIMIT 100
) x
ORDER BY x.smlr ASC
LIMIT 10;
```

内层查询利用 GiST 索引按 signature 选出前 100 个候选项，外层查询按 pattern 距离精炼到前 10 个。
