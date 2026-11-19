## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/pg_embeddings.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/pg_embeddings--1.1.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_embeddings/meson.build)

`pg_embeddings` `1.1` 提供基于 AlohaDB 向量的嵌入生命周期与混合检索。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_embeddings;
SELECT embedding_binary_quantize('[1,-1,0.5]'::vector);
```

控制文件未要求超级用户，但安装者仍需数据库和目标模式的创建权限。

### 运行边界

先通过 embedding_model_register 注册真实的 SQL 文本转向量函数，再使用 embedding_track、embedding_process 和 embedding_hybrid_search。必须明确安排队列处理，CREATE EXTENSION 不会自动启用后台处理。模型维数以及调用者的表权限、行级安全规则仍适用。模块使用 AlohaDB 内核向量函数，而非独立的 pgvector 扩展。升级旧版 1.0 前，需按官方流程将 embedding_model.embed_fn 从 regprocedure 迁移为文本。

此模块对应上游旧名 `alohadb_embeddings`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
