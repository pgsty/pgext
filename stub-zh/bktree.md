## 用法

来源：

- [官方 README](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/README.md)
- [1.0 版 SQL API](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree--1.0.sql)
- [汉明距离实现](https://github.com/meniam/pg_bktree/blob/59db54ff5cd0b9f1b60c95be6f04414282699237/bktree.c)

`bktree` 为 64 位整数哈希提供 BK-tree SP-GiST 操作符类。其距离是两个位模式之间的汉明距离，因此适合对感知哈希进行近邻式过滤，而不是处理普通数值范围。

### 核心流程

把每个 64 位哈希保存为 `bigint`，用 `bktree_ops` 建立索引，再用包含中心值与最大距离的 `bktree_area` 查询。

```sql
CREATE EXTENSION bktree;
CREATE TABLE images (id bigint PRIMARY KEY, phash bigint NOT NULL);
CREATE INDEX images_phash_bktree ON images USING spgist (phash bktree_ops);

SELECT id, phash <-> 123456789::bigint AS distance
FROM images
WHERE phash <@ ROW(123456789::bigint, 8::bigint)::bktree_area;
```

`<@` 操作符检查汉明距离是否不大于给定半径。`<->` 操作符返回距离，而扩展提供的 `=` 操作符检查位相等。`int64_to_bitstring()` 与 `bitstring_to_int64()` 是转换辅助函数。

### 注意事项

该操作符类只适用于 `int8`，实现假定 64 位汉明距离介于 0 到 64。这里的半径不是数值容差：数值接近的整数可能有很多不同位，数值相距很远的整数也可能只有少量不同位。所审阅分支专门面向 PostgreSQL 15；生产采用前应在目标服务端版本上验证索引创建、执行计划、相等语义以及有符号哈希转换。
