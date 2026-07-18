## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/README.md)
- [1.0 版本类型定义](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype--1.0.sql)
- [C 实现](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/pg_imputed_genotype.c)
- [PostgreSQL 许可证](https://github.com/rlichtenwalter/pg_imputed_genotype/blob/4473d6b38b8a98c40176cfb15120823eb10e0458/LICENSE)

`pg_imputed_genotype` 增加用于紧凑存储三个基因型概率的 `imputed_genotype` 基础类型。每个分量量化为 0 到 1000 的 10 位整数，完整三元组只占四字节。

```sql
CREATE EXTENSION pg_imputed_genotype;
SELECT '{0.2,0.5,0.1}'::imputed_genotype;
-- {0.200,0.500,0.100}
```

输入必须恰好包含三个概率，且总和不得超过 1.0。该类型适合密集的推断基因型数组，其设计侧重固定精度与存储尺寸，而不是保留任意小数精度。

仓库没有发行标签或兼容性矩阵，已复核源码也较旧。迁移数据前，应在实际使用的 PostgreSQL 大版本与 CPU 架构上测试二进制兼容性，并确认 0.001 量化精度满足后续统计分析要求。
