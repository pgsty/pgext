## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/pg_vector.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/pg_vector--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_vector/vector_ivfflat.c)

`pg_vector` `1.0` 提供基于 AlohaDB 内核向量的命名 ANN 索引函数。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_vector;
SELECT vector_l2_distance('[1,2]'::vector, '[2,3]'::vector);
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

vector 类型和内核 HNSW/IVFFlat 访问方法已属于 AlohaDB 内核。本模块增加 vector_index_create/insert/rebuild/search 及自己的索引元数据，与 pgvector 不同。重建操作训练辅助索引，应用需明确维护辅助索引条目。应认真选择维数、度量和探测、列表数量；近似搜索可能漏掉邻居，辅助状态应与源数据一起备份。

此模块对应上游旧名 `alohadb_vector`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
