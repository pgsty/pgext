## 用法

来源：

- [已复核 commit 的官方 README](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/README.md)
- [1.0 版本类型定义](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype--1.0.sql)
- [类型输入与输出实现](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype.c)

`pg_imputed_genotype` 增加固定大小的 `imputed_genotype` 基础类型，用于保存三个推断基因型概率。它把三个 10 位整数压入一个四字节按值传递的数据。只有在紧凑存储与 0.001 级量化均可接受时才应使用。

### 核心流程

```sql
CREATE EXTENSION pg_imputed_genotype;

CREATE TABLE sample_call (
  sample_id bigint PRIMARY KEY,
  probabilities imputed_genotype NOT NULL
);

INSERT INTO sample_call VALUES (1, '{0.2,0.5,0.1}');
SELECT probabilities FROM sample_call WHERE sample_id = 1;
-- {0.200,0.500,0.100}
```

输入是在花括号中写入三个逗号分隔的概率。各值先乘以 1000 并截断而非四舍五入，然后再打包。总和可以小于一，剩余概率保持隐含。扩展只定义基础类型的输入输出路径；没有比较操作符、类型转换或索引操作符类。

### 验证与兼容性边界

实现把整数化总和 1001 作为舍入容差，因此可能接受略大于一的输入。解析器也不能可靠拒绝尾随文本，且没有显式防止非有限浮点值。转换前，应在应用或 SQL 代码中验证恰好存在三个有限的非负值，并检查其总和。

量化会损失小数精度，打包表示也是扩展专属的二进制格式。该旧仓库没有发行标签或 PostgreSQL 兼容矩阵。在保存持久数据前，应针对确切的 PostgreSQL 大版本与 CPU 架构测试输入、输出、备份恢复和二进制行为。
