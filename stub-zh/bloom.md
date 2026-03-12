


## 用法

> [bloom: bloom 访问方法 - 基于签名文件的索引](https://www.postgresql.org/docs/current/bloom.html)

`bloom` 扩展提供基于[布隆过滤器](https://en.wikipedia.org/wiki/Bloom_filter)的索引访问方法。布隆过滤器是一种空间高效的数据结构，用于测试元素是否属于集合，可能出现假阳性但不会有假阴性。

布隆索引特别适用于有很多列的表，且查询会测试列的任意组合的场景。单个布隆索引可以替代多个 btree 索引，同时占用显著更少的空间。

### 创建布隆索引

```sql
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3, i4, i5, i6);
```

使用自定义参数：

```sql
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3)
       WITH (length=80, col1=2, col2=2, col3=4);
```

### 参数

| 参数 | 默认值 | 最大值 | 描述 |
|------|--------|--------|------|
| `length` | 80 | 4096 | 每个签名的位长度（取最近的 16 的倍数） |
| `col1` - `col32` | 2 | 4095 | 每个索引列生成的位数 |

### 运算符类支持

布隆索引仅支持等值运算符（`=`）。内置的运算符类适用于 `int4` 和 `text`。你可以为有哈希函数的其他类型定义自定义运算符类：

```sql
CREATE OPERATOR CLASS text_ops
DEFAULT FOR TYPE text USING bloom AS
    OPERATOR    1   =(text, text),
    FUNCTION    1   hashtext(text);
```

### 示例

```sql
-- 创建一个多列表
CREATE TABLE tbloom AS
   SELECT
     (random() * 1000000)::int as i1,
     (random() * 1000000)::int as i2,
     (random() * 1000000)::int as i3,
     (random() * 1000000)::int as i4,
     (random() * 1000000)::int as i5,
     (random() * 1000000)::int as i6
   FROM generate_series(1, 10000000);

-- 单个布隆索引覆盖所有列组合
CREATE INDEX bloomidx ON tbloom USING bloom (i1, i2, i3, i4, i5, i6);

-- 对任意列子集的查询都可以使用该索引
SELECT * FROM tbloom WHERE i2 = 898732 AND i5 = 123451;
```

### 限制

- 仅支持等值（`=`）查询（不支持范围查询）
- 不支持 `UNIQUE` 索引
- 不支持搜索 `NULL` 值
- 由于假阳性，结果需要回表复核
