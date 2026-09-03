## 用法

来源：

- [官方文档](https://github.com/sid2364/dna-sequences-pg-extension/blob/dc6e4f105cbc8bdcacab8ca0ba1f3ae005c42efb/README.md)
- [扩展控制文件](https://github.com/sid2364/dna-sequences-pg-extension/blob/dc6e4f105cbc8bdcacab8ca0ba1f3ae005c42efb/dna.control)
- [官方仓库](https://github.com/sid2364/dna-sequences-pg-extension)

`dna` 提供紧凑 DNA 序列与 k-mer 类型，以及比较、校验和分析函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `dna`：

```sql
CREATE EXTENSION dna;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
SELECT k.kmer FROM generate_kmers('ACGTACGCACGT', 6) AS k(kmer) WHERE 'DNMSRN' @> k.kmer ;
--  kmer
----------
-- GTACGC
-- GCACGT
--(2 rows)
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `kmer` | TYPE | 扩展创建的用户数据类型。 |
| `generate_kmers` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `length` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `contains` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `qkmer` | TYPE | 扩展创建的用户数据类型。 |
| `starts_with` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `dna_construct` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `dna_in` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 应将扩展升级视为数据库变更：先审查上游升级路径、权限、锁以及备份恢复行为。
