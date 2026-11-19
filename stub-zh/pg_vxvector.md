## 用法

来源：

- [官方来源](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/README.md)
- [官方来源](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/pg_vxvector.control)
- [官方来源](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/sql/pg_vxvector--0.1.0.sql)
- [官方来源](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/docs/pg_vxvector.md)
- [官方来源](https://github.com/minzyfox/VxVector/blob/8e4b294b74e12ca411af3b17b531dadd5ab1aac5/pg_vxvector/Cargo.toml)

`pg_vxvector` 0.1.0 通过外部 VxVector 工作进程训练聚类中心，并可据此构建 VectorChord 索引。它要求 PostgreSQL 17、`vector` 与 `vchord`。这是实验性的训练集成，不提供查询阶段 GPU 搜索。

### 训练聚类中心

安装匹配的扩展与工作进程文件，配置 `vxvector.worker_path`，预加载 `vchord` 并重启。安装需要超级用户；调用者需要源列的 SELECT 权限以及创建输出对象的普通权限。示例假定已有向量列，并显式选择只生成聚类中心、不创建索引。

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vchord;
CREATE EXTENSION pg_vxvector;
SELECT vxvector.create_vector_index(
    'public.items'::regclass, 'embedding', ARRAY[100],
    distance_operator => 'l2', skip_index_build => true
);
```

### 控制参数与故障边界

`vxvector.create_vector_index` 返回描述输出的 JSON。参数包括 `lists`、`sampling_factor`、`batch_size`、`kmeans_iterations`、`kmeans_nredo`、`seed`、`memory_budget_bytes`、`worker_timeout_ms`、`centroid_table` 与 `index_name`。大规模运行前应核对维度、样本数与内存预算。调用遵循 RLS，不支持 `halfvec`、`sparsevec` 与 `bit`。

超级用户设置 `vxvector.backend_mode` 接受 `auto`、`cpu` 与 `gpu`。没有可用的受支持 GPU 时，自动模式可在训练前选择 CPU；GPU 训练开始后的故障会报错，不会静默转为 CPU 重试。当前带日期的路线图与 README 已取代旧文档中的原生 macOS 阻碍：上游验证过一套 macOS/Apple Metal PostgreSQL 17 配置。Linux Docker 使用 CPU，CUDA 后端尚未经过 NVIDIA 硬件验证。应保持文档指定的编译器和工作进程修订匹配；这不代表生产就绪或普遍的 GPU 兼容性。
