## 用法

来源：

- [项目 README](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/README.txt)
- [扩展 control 文件](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/postbis.control)
- [1.0 版 SQL API](https://github.com/no0p/postbis/blob/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/sql/postbis--1.0.sql)
- [序列回归测试](https://github.com/no0p/postbis/tree/ce454ebfbc27e0b6c8357ef6bfc8da1c4b2967c8/test/sql)

`postbis` 1.0 为 DNA、RNA、氨基酸及比对序列提供紧凑的原生数据类型，并提供可配置字母表与类型修饰符、类型转换、序列操作、生物学变换、比较操作符，以及 B-tree 和哈希操作符类。

### 存储强类型序列

```sql
CREATE EXTENSION postbis;

CREATE TABLE specimen (
  specimen_id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
  dna dna_sequence(SHORT, FLC, CASE_SENSITIVE) NOT NULL,
  rna rna_sequence(IUPAC, CASE_SENSITIVE),
  protein aa_sequence(IUPAC, CASE_SENSITIVE)
);

INSERT INTO specimen (dna, rna, protein)
VALUES ('AACCGGTT', 'AACGUU', 'ACDEFG');

SELECT specimen_id,
       char_length(dna) AS bases,
       substr(dna, 3, 4)::text AS fragment
FROM specimen;
```

输入验证取决于所选字母表、大小写敏感性和类型修饰符。应验证转换能拒绝所需生物学约定之外的符号，并避免意外混用对齐与未对齐类型。

### 变换与翻译序列

```sql
SELECT complement('ACGTN'::dna_sequence)::text;
-- TGCAN

SELECT reverse_complement('ACGTN'::dna_sequence)::text;
-- NACGT

SELECT transcribe('AACGTT'::dna_sequence)::text;
-- AACGUU

SELECT translate('AUGGCCUAA'::rna_sequence)::text;
-- MA
```

扩展还提供 `reverse_transcribe()`、`six_frame()`、`get_alphabet()`、`entropy()`、`gc_content()` 和序列生成函数。标准遗传密码不适用时，翻译函数可以接收显式翻译表。

### 检查压缩并添加索引

```sql
SELECT char_length(sequence) AS symbols,
       octet_length(sequence) AS storage_bytes,
       compression_ratio(sequence) AS storage_ratio
FROM (
  SELECT repeat('ACGT', 256)::dna_sequence AS sequence
) AS sample;

CREATE INDEX specimen_dna_btree ON specimen USING btree (dna);
CREATE INDEX specimen_dna_hash  ON specimen USING hash  (dna);
```

这些序列类型支持等值、排序、拼接、子串、搜索和长度函数。在生产负载中依赖索引之前，应使用真实数据分布检查执行计划。

### 打包与持久性风险

Pigsty 应用下游兼容补丁，将 PostBIS 1.0 打包到 PostgreSQL 14–18。这一打包结果并不改变上游生命周期：项目已经停止维护，也没有 1.0 以后的扩展升级路径。

自定义类型使用原生压缩磁盘表示。应把已存储值和索引视为绑定到经过测试的准确构建。采用或迁移前，必须验证转储恢复、二进制和逻辑升级、复制、驱动解码、索引重建、错误输入处理和大型序列内存行为。

`reverse()`、`char_length()` 和 `substr()` 等函数重载了常见名称，因此模式限定和受控 `search_path` 设置很重要。对新的持久数据集，除非已经在本地审计、打包该扩展并指定明确的长期迁移负责人，否则应优先使用受维护的序列工具或普通 PostgreSQL 类型。
