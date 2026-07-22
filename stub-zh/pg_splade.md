## 用法

来源：

- [官方 README](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/README.md)
- [官方编码实现](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/encode.rs)
- [官方 SQL 函数实现](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/datatype/function.rs)
- [官方扩展入口](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/src/lib.rs)
- [官方扩展控制文件](https://github.com/silver-ymz/pg_splade/blob/f6ed591e4a5ba4717ce0f78a7209d0e377bc774d/pg_splade.control)

`pg_splade` 是一个开发中的神经稀疏嵌入扩展。它在 PostgreSQL 内运行 SPLADE 模型，并为文档与查询文本返回 pgvector `sparsevec` 值，但模型文件、后端内存与服务器预加载也会成为数据库操作的一部分。

### 设置与核心流程

扩展依赖 pgvector 的 `vector` 扩展，并且必须由 postmaster 加载：

```conf
shared_preload_libraries = 'pg_splade'
splade.preload_models = 'distill'
```

安装和使用扩展前先重启 PostgreSQL：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION pg_splade;

SELECT list_model();
SELECT encode_document('distill', 'PostgreSQL full text document');
SELECT encode_query('distill', 'database search');
```

`encode_document(text, text)` 与 `encode_query(text, text)` 生成稀疏向量。`truncate_sparsevec(sparsevec, integer)` 保留更小的表示。模型按后端加载，也可以惰性加载。

### 模型管理与注意事项

`download_model(text, text)` 将 Hugging Face 仓库下载到服务器共享数据目录下的 `splade` 目录，`delete_model(text)` 递归删除指定模型目录，`list_model()` 报告已安装模型。这些函数会执行外部网络与文件系统操作；只能授权给可信管理员，并应独立维护模型来源与备份策略。

已核验代码面向小端 64 位构建以及 PostgreSQL 14 到 17。容量与延迟取决于所选模型和 CPU、CUDA 或 Metal 执行路径。虽然编码函数声明为不可变且并行安全，但替换或删除模型文件仍会改变结果或使调用失败，因此应固定模型内容，并让更新与使用嵌入的会话及索引保持协调。
