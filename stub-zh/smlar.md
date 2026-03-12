

## 用法

> [smlar](https://github.com/jirutka/smlar)：PostgreSQL 数组的高效相似度搜索。
> 来源：[README](https://github.com/jirutka/smlar/blob/master/README)

`smlar` 扩展提供 PostgreSQL 数组上的高效相似度搜索，支持可配置的相似度公式、GiST 和 GIN 索引，以及 TF/IDF 加权。


--------

## 函数

```
float4 smlar(anyarray, anyarray)
```
计算两个数组的相似度。数组应为相同类型。

```
float4 smlar(anyarray, anyarray, bool useIntersect)
```
计算两个复合类型数组的相似度。复合类型格式如下：

```sql
CREATE TYPE type_name AS (element_name anytype, weight_name FLOAT4);
```

`useIntersect` 选项指定分母中仅使用交集元素。

```
float4 smlar(anyarray a, anyarray b, text formula)
```
通过给定公式计算两个数组的相似度。公式中的预定义变量：

- `N.i` -- 两个数组中的公共元素数量（交集）
- `N.a` -- 第一个数组中的唯一元素数量
- `N.b` -- 第二个数组中的唯一元素数量

示例：

```sql
SELECT smlar('{1,4,6}'::int[], '{5,4,6}');
SELECT smlar('{1,4,6}'::int[], '{5,4,6}', 'N.i / sqrt(N.a * N.b)');
-- 这两个调用是等价的。
```

```
anyarray % anyarray
```
如果数组的相似度大于阈值限制则返回 true。

```
text[] tsvector2textarray(tsvector)
```
将 tsvector 类型转换为文本数组。

```
anyarray array_unique(anyarray)
```
排序并去重数组。

```
float4 inarray(anyarray, anyelement)
```
如果第二个参数不存在于第一个参数中返回零，否则返回 1.0。

```
float4 inarray(anyarray, anyelement, float4, float4)
```
如果第二个参数不存在于第一个参数中返回第四个参数，否则返回第三个参数。


--------

## GUC 配置变量

```
smlar.threshold  FLOAT
```
相似度低于阈值的数组不被 `%` 运算视为相似。

```
smlar.persistent_cache  BOOL
```
全局统计缓存存储在事务无关的内存中。

```
smlar.type  STRING
```
相似度公式类型：`cosine`（默认）、`tfidf`、`overlap`。

```
smlar.stattable  STRING
```
存储集合级统计数据的表名。表应定义为：

```sql
CREATE TABLE table_name (
    value   data_type UNIQUE,
    ndoc    int4 (or bigint)  NOT NULL CHECK (ndoc > 0)
);
```

值为 null 的行表示文档总数。仅用于 `smlar.type = 'tfidf'`。

```
smlar.tf_method  STRING
```
词频计算方法。取值：
- `"n"` -- 简单计数（默认）
- `"log"` -- 1 + log(n)
- `"const"` -- TF 等于 1

仅用于 `smlar.type = 'tfidf'`。

```
smlar.idf_plus_one  BOOL
```
如果为 false（默认），idf 计算为 `log(d/df)`。如果为 true，计算为 `log(1+d/df)`。仅用于 `smlar.type = 'tfidf'`。

强烈建议在 `postgresql.conf` 中添加：

```
smlar.threshold = 0.6  # 或其他 > 0 且 < 1 的值
```


--------

## GiST/GIN 索引支持

`%` 和 `&&` 操作支持多种数组类型的 GiST 和 GIN 索引：

| 数组类型 | GIN 操作符类 | GiST 操作符类 |
|----------|-------------|--------------|
| `bit[]` | `_bit_sml_ops` | |
| `bytea[]` | `_bytea_sml_ops` | `_bytea_sml_ops` |
| `char[]` | `_char_sml_ops` | `_char_sml_ops` |
| `cidr[]` | `_cidr_sml_ops` | `_cidr_sml_ops` |
| `date[]` | `_date_sml_ops` | `_date_sml_ops` |
| `float4[]` | `_float4_sml_ops` | `_float4_sml_ops` |
| `float8[]` | `_float8_sml_ops` | `_float8_sml_ops` |
| `inet[]` | `_inet_sml_ops` | `_inet_sml_ops` |
| `int2[]` | `_int2_sml_ops` | `_int2_sml_ops` |
| `int4[]` | `_int4_sml_ops` | `_int4_sml_ops` |
| `int8[]` | `_int8_sml_ops` | `_int8_sml_ops` |
| `interval[]` | `_interval_sml_ops` | `_interval_sml_ops` |
| `macaddr[]` | `_macaddr_sml_ops` | `_macaddr_sml_ops` |
| `money[]` | `_money_sml_ops` | |
| `numeric[]` | `_numeric_sml_ops` | `_numeric_sml_ops` |
| `oid[]` | `_oid_sml_ops` | `_oid_sml_ops` |
| `text[]` | `_text_sml_ops` | `_text_sml_ops` |
| `time[]` | `_time_sml_ops` | `_time_sml_ops` |
| `timestamp[]` | `_timestamp_sml_ops` | `_timestamp_sml_ops` |
| `timestamptz[]` | `_timestamptz_sml_ops` | `_timestamptz_sml_ops` |
| `timetz[]` | `_timetz_sml_ops` | `_timetz_sml_ops` |
| `varbit[]` | `_varbit_sml_ops` | |
| `varchar[]` | `_varchar_sml_ops` | `_varchar_sml_ops` |
