## 用法

来源：

- [上游 README](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/README.md)
- [扩展 control 文件](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype.control)
- [1.0 版 SQL 定义](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype--1.0.sql)
- [类型实现](https://github.com/rlichtenwalter/pg_genotype/blob/e7fa2e88e33c3f4fe4d912c8f86884d904af7a1b/pg_genotype.c)

`pg_genotype` `1.0` 版增加了紧凑的 `genotype` 基础类型，用于两个字符的等位基因对。其输入例程允许两个位置使用 A、C、G、T、I、D，并保留顺序，所以 `AT` 与 `TA` 是不同值。

### 存储等位基因对

```sql
CREATE EXTENSION pg_genotype;

CREATE TABLE sample_genotypes (
    sample_id bigint PRIMARY KEY,
    call genotype NOT NULL
);

INSERT INTO sample_genotypes VALUES
    (1, 'AT'),
    (2, 'TA'),
    (3, 'ID');

SELECT sample_id, call::text
FROM sample_genotypes
ORDER BY sample_id;
```

该扩展只定义类型输入/输出，没有提供基因型运算符、比较语义、类型转换或索引运算符类。README 虽描述为单字节压缩表示，SQL 定义却声明了四字节按值传递数据。C 解析器读取前两个字符，但不证明输入在此结束，因此应用应在转换前强制执行精确的 `^[ACGTID]{2}$` 检查。该仓库早于当前 PostgreSQL 版本；依赖其存储表示前，应针对确切服务器版本完成编译和回归测试。
