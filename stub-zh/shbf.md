## 用法

来源：

- [官方 README](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/README.md)
- [1.0.0 版扩展 SQL](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf--1.0.0.sql)
- [C 实现](https://github.com/sdsmart/shifting-bloom-filter/blob/e7c2d5de0217dae14b49079956be7f26a0a3b0db/shbf.c)
- [原始 Shifting Bloom Filter 论文](https://www.vldb.org/pvldb/vol9/p408-yang.pdf)

`shbf` 把 Count-Min Sketch、普通 Bloom filter 和三种 Shifting Bloom Filter 变体实现为紧凑的 PostgreSQL 概率值。这些结构以准确性换取固定内存：成员与关联的肯定结果可能是假阳性，频率估计也可能偏高。

### 核心流程

按规划位数 `m` 和预期元素数 `n` 创建过滤器。插入函数会返回新值，因此必须始终把结果赋回；它们不会原地持久修改已存储 datum。

```sql
CREATE EXTENSION shbf;

CREATE TABLE seen_key (
  id integer PRIMARY KEY,
  filter bf NOT NULL
);

INSERT INTO seen_key VALUES (1, new_bf(1000000, 10000));

UPDATE seen_key
SET filter = insert_bf(filter, 'alpha')
WHERE id = 1;

SELECT query_bf(filter, 'alpha')
FROM seen_key
WHERE id = 1;
```

Count-Min Sketch 可用于近似频率。

```sql
CREATE TEMP TABLE frequency(sketch cms);
INSERT INTO frequency VALUES (new_cms(0.001, 0.99));

UPDATE frequency SET sketch = insert_cms(sketch, 'alpha');
UPDATE frequency SET sketch = insert_cms(sketch, 'alpha');

SELECT query_cms(sketch, 'alpha') FROM frequency;
```

### 重要对象

- `cms`：Count-Min Sketch；`new_cms(error_bound, confidence)` 默认参数为 `0.001` 和 `0.99`，配合 `insert_cms` 与 `query_cms` 提供近似计数。
- `bf`：普通 Bloom filter；`new_bf(m, n)`、`insert_bf` 与 `query_bf` 提供成员操作。
- `new_shbf_m(m, n)`、`insert_shbf_m`、`query_shbf_m`：shifting 成员过滤器。
- `new_shbf_a(m, n)`、`insert_shbf_a`、`query_shbf_a`：两个集合的关联过滤器。
- `new_shbf_x(m, n, max_x)`、`insert_shbf_x`、`query_shbf_x`：不超过规划上限的多重性过滤器。

`query_shbf_a` 返回编码状态：`-1` 表示不存在；`0` 表示第二集合；`1` 表示第一集合；`2` 表示同时属于两者；`3` 表示第一或第二；`4` 表示两者或第二；`5` 表示两者或第一；`6` 表示三种成员状态均可能。含糊编码来自哈希碰撞。

### 容量规划与安全

应根据实测负载选择正的 `m`、`n` 与 `max_x`，并把插入的多重性限制在规划范围内。所复核构造器几乎不做校验：零或无效尺寸可能进入 C 中的除法、取模、内存分配或偏移运算。存储过滤器前，应在生产之外测试容量与饱和行为。

这些类型没有删除、合并聚合、索引运算符类、自动扩容或精确信心报告。权威数据必须保存在其他位置，过滤器只能用作加速或遥测摘要。上游项目是紧凑的研究型代码；把存储值作为持久数据前，应验证跨架构与 PostgreSQL 升级时的序列化兼容性。
