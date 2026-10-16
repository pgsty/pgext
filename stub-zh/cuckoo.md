## 用法

来源：

- [README.md](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/README.md)
- [cuckoo--1.0.sql](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/cuckoo--1.0.sql)
- [LICENSE.md](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/LICENSE.md)
- [cuckoo.control](https://github.com/JerrySievert/pg_cuckoo/blob/f4218b4cfb37c92f5c97427356cc5aa28869cc33/cuckoo.control)

`cuckoo` 1.0 提供面向等值谓词的概率索引访问方法。项目名为 pg_cuckoo，SQL 扩展与访问方法均为 `cuckoo`。

### 核心用法

```sql
CREATE EXTENSION cuckoo;
CREATE TABLE cuckoo_example (label text);
CREATE INDEX ON cuckoo_example USING cuckoo (label);
INSERT INTO cuckoo_example VALUES ('alpha');
SELECT * FROM cuckoo_example WHERE label = 'alpha';
```

### 运行边界

上游要求 PostgreSQL 16 或更高版本，control 默认要求超级用户安装。索引支持等值查询，不支持范围查询，需通过堆表复核排除误报。最简单的用法是单列索引；多列索引会合并指纹，需要对所有索引列提供谓词。选项 `bits_per_tag`、`tags_per_bucket`、`max_kicks` 默认分别为 12、4、500。增加指纹位数会以空间换取更低误报率，上游基准结果不代表实际负载保证。
