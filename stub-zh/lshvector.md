## 用法

来源：

- [官方上游 README](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/README.md)
- [官方扩展控制文件 (lshvector.control)](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/tools/ghidra_11.2.1_PUBLIC/Ghidra/Features/BSim/src/lshvector/lshvector.control)
- [官方扩展 SQL (lshvector--1.0.sql)](https://github.com/massinissadjellouli/ctfstuff/blob/ddf59b4bb6dab358f8cda26a5b09adef8c2d62d3/tools/ghidra_11.2.1_PUBLIC/Ghidra/Features/BSim/src/lshvector/lshvector--1.0.sql)

`lshvector` — 一种特征向量类型和局部敏感哈希索引。使用它来处理相应的向量、模型或检索工作流。请使用上述链接的上游固定版本作为 API 边界，并在目标 PostgreSQL 构建中进行测试。

### 核心工作流

```sql
CREATE EXTENSION lshvector;
```

在目标数据库中安装扩展，当可用时运行上游示例中的最小示例，并在将其集成到应用程序 SQL 中之前验证已安装的版本和返回值。

### 重要对象

- `lsh_getweight(lshvector)` 是一个扩展函数，返回 `float8`。
- `lsh_load()` 是一个扩展函数，返回 `int4`。
- `lsh_reload()` 是一个扩展函数，返回 `int4`。
- `lshvector_compare(lshvector,lshvector)` 是一个扩展函数，返回 `lshvector_comptype`。
- `lshvector_gin_consistent(internal, int2, lshvector, int4, internal, internal, internal, internal)` 是一个扩展函数。
- `lshvector_gin_extract_query(lshvector,internal,int2,internal,internal,internal,internal)` 是一个扩展函数，返回 `internal`。
- `lshvector_gin_extract_value(lshvector,internal)` 是一个扩展函数，返回 `internal`。
- `lshvector_hash(lshvector)` 是一个扩展函数，返回 `int8`。
- `lshvector_in(cstring)` 是一个扩展函数，返回 `lshvector`。
- `lshvector_out(lshvector)` 是一个扩展函数，返回 `cstring`。
- `lshvector_overlap(lshvector,lshvector)` 是一个扩展函数。
- `lshvector_recv(internal)` 是一个扩展函数，返回 `lshvector`。
- `lshvector_send(lshvector)` 是一个扩展函数，返回 `bytea`。
- `lshvector` 是一个扩展定义的类型。

### 要求与注意事项

- 审查后的控制文件声明默认版本为 `1.0`。
- 控制文件将扩展标记为可重定位。
- 控制文件不要求超级用户安装。
- 在生产使用前，请确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与固定源中的信息一致。
