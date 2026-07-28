## 用法

来源：

- [官方上游 README](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/ReadMe.md)
- [官方扩展控制文件 (kmea.control)](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/kmea.control)
- [官方扩展 SQL (kmea--1.0.sql)](https://github.com/pipcount/pg-dna-extension/blob/1340cb24d11f6657d71555cf7cfd85009f7914c3/kmea--1.0.sql)

`kmea` — KMEA[^1] 是一个支持各种 DNA 数据类型的 PostgreSQL 扩展，还包含一些操作符。使用它当应用程序需要这种类型、领域或其操作符时。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建上进行测试。

### 核心工作流

```sql
CREATE EXTENSION kmea;
```

在目标数据库中安装扩展，如果有可用的最小上游示例，请运行该示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `canonical(kmer)` 是一个扩展函数，返回 `kmer`。
- `contains(qkmer, kmer)` 是一个扩展函数，返回 `boolean`。
- `DNA(text)` 是一个扩展函数，返回 `DNA`。
- `dna_in(cstring)` 是一个扩展函数，返回 `DNA`。
- `dna_out(DNA)` 是一个扩展函数，返回 `cstring`。
- `dna_recv(internal)` 是一个扩展函数，返回 `DNA`。
- `dna_send(DNA)` 是一个扩展函数，返回 `bytea`。
- `equals(kmer, kmer)` 是一个扩展函数，返回 `boolean`。
- `generate_kmers(DNA, integer)` 是一个扩展函数，返回 `SETOF`。
- `kmer(text)` 是一个扩展函数，返回 `kmer`。
- `kmer_hash(kmer)` 是一个扩展函数，返回 `integer`。
- `kmer_in(cstring)` 是一个扩展函数，返回 `kmer`。
- `kmer_out(kmer)` 是一个扩展函数，返回 `cstring`。
- `kmer_recv(internal)` 是一个扩展函数，返回 `kmer`。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况。
