## 用法

来源：

- [项目 README](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/README.txt)
- [扩展 control 文件](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/postbis.control)
- [1.0 版 SQL API](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/sql/postbis--1.0.sql)
- [序列回归测试](https://github.com/no0p/postbis/tree/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/test/sql)

`postbis` 1.0 是一个已废弃的原生压缩生物序列扩展。它定义 `dna_sequence`、`rna_sequence`、`aa_sequence` 及相应对齐类型、可配置字母表和类型修饰符，还提供转换、子串与长度操作、序列变换、比较操作符以及 B-tree 和哈希支持。

### 存储并变换 DNA 序列

```sql
CREATE EXTENSION postbis;

CREATE TABLE specimen (
  specimen_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  sequence dna_sequence NOT NULL
);

INSERT INTO specimen (sequence) VALUES ('ACGTACGT'::dna_sequence);

SELECT specimen_id,
       char_length(sequence) AS bases,
       reverse_complement(sequence)::text AS reverse_complement
FROM specimen;
```

输入验证取决于所选字母表、大小写敏感性和类型修饰符。应验证转换能拒绝所需生物学约定之外的符号，并避免意外混用对齐与未对齐类型。

### 持久性与兼容风险

自定义类型使用原生压缩磁盘表示。仓库包含陈旧的 Subversion 导入元数据，没有当前 release、受支持 PostgreSQL 矩阵，也没有 1.0 以后的升级路径。应把已存储值和索引视为绑定到经过测试的准确构建。采用或迁移前，必须验证转储恢复、二进制和逻辑升级、复制、驱动解码、索引重建、错误输入处理和大型序列内存行为。

`reverse()`、`char_length()` 和 `substr()` 等函数重载了常见名称，因此模式限定和受控 `search_path` 设置很重要。对新的持久数据集，除非已经在本地审计、打包该扩展并指定明确的长期迁移负责人，否则应优先使用受维护的序列工具或普通 PostgreSQL 类型。
