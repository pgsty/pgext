## 用法

来源：

- [README](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/README.md)
- [Control](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence.control)
- [SQL](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence--1.0.sql)
- [Source](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/dna_sequence/dna_sequence.c)
- [SP-GiST tests](https://github.com/RahmanTektas/postgresql-dna-extension/blob/8b5d85bcc5f3f75d2ca8f8b4b33f3cf265bdab2d/tests/test_spgist.sql)

`dna_sequence` 是一个用 C 编写的教学项目扩展，用于存储基因序列并通过索引查询 k-mer。本源码快照声明扩展版本为 1.0；上游未公布 PostgreSQL 主版本兼容矩阵，也没有明确的许可证声明。采用这些磁盘数据类型前，应先在临时数据库中验证。

### 核心用法

安装共享库和 SQL 文件后，由超级用户创建扩展。控制文件允许迁移模式，未声明预加载设置或其他扩展依赖。

```sql
CREATE EXTENSION dna_sequence;
SELECT length('ACGTACGT'::dna);
SELECT * FROM generate_kmers('ACGTACGT'::dna, 3);

CREATE TABLE dna_kmers (value kmer);
INSERT INTO dna_kmers
SELECT * FROM generate_kmers('ACGTACGT'::dna, 3);
CREATE INDEX dna_kmers_spgist ON dna_kmers USING spgist (value);
SELECT value FROM dna_kmers WHERE value ^@ 'AC'::kmer;
SELECT value FROM dna_kmers WHERE value <@ 'ANG'::qkmer;
```

`dna` 存储由四种确定 DNA 碱基组成的序列，`kmer` 接受 1–32 个确定碱基，`qkmer` 则使用 IUPAC 简并碱基编码描述查询模式。`generate_kmers(dna, integer)` 以集合形式返回相互重叠的 k-mer，`length` 为三种类型分别提供重载。

### 查询接口

| 对象 | 含义 |
| --- | --- |
| `equals(kmer,kmer)`、`=` 和 `<>` | k-mer 精确相等与不等比较 |
| `starts_with(kmer,kmer)`、`^@` | 判断 k-mer 是否以给定前缀开头 |
| `contains(qkmer,kmer)`、`contained(kmer,qkmer)`、`@>` 和 `<@` | 判断确定的 k-mer 是否匹配简并碱基查询模式 |
| `kmer_hash_ops` | 支持哈希索引和分组 |
| `kmer_btree_ops` | 支持有序比较与 B-tree 索引 |
| `kmer_spgist_ops` | 为等值、前缀和模式谓词提供 SP-GiST 支持 |

### 使用边界

类型会拒绝无效碱基和无效 k-mer 长度。扩展包含原生类型和索引代码，更换构建版本前应保留逻辑备份并验证恢复路径。在同一模式中，其未限定类型名可能与其他基因组扩展冲突；应使用实际规范扩展名 `dna_sequence`，不要与目录中的另一个扩展 `dna` 混淆。

上游测试和构建辅助命令包含删除扩展及其依赖对象的操作，不应针对已有应用数据执行这类开发重置流程。本文说明源码声明的 SQL 接口，不代表生产支持承诺或运行测试认证。
