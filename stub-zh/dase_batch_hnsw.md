## 用法

来源：

- [README.md](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/README.md)
- [dase_batch_hnsw.control](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/dase_batch_hnsw.control)
- [dase_batch_hnsw--1.0.sql](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/dase_batch_hnsw--1.0.sql)
- [Makefile](https://github.com/ljmzlh/DASE/blob/b2d0d4e9e2627867bdc9c607df7d3324c316b500/batch_hnsw/Makefile)

`dase_batch_hnsw` 为 DASE 研究系统提供批量 HNSW 范围探测，SQL 版本为 1.0。它编译同一检出版本中修改过的 pgvector HNSW 代码，必须与对应分叉及 PostgreSQL 服务器匹配。

### 核心用法

先安装匹配的扩展文件，再以超级用户启用依赖并检查函数：

```sql
CREATE EXTENSION IF NOT EXISTS vector;
CREATE EXTENSION IF NOT EXISTS dase_batch_hnsw;
SELECT n.nspname, p.proname
FROM pg_proc p JOIN pg_namespace n ON n.oid = p.pronamespace
WHERE n.nspname = 'bh';
SELECT bh.last_stats();
```

固定的 `bh` 模式提供 `range_join_all` 和 `last_stats`。范围探测函数接受 `left_sql`、`right_index`、`radius`、`ef`、`mode` 和 `n_seeds`。左侧查询按处理顺序返回 int8 ID、向量和 int4 分组。结果包含 `left_id`、`right_tid` 和 `dist`。统计信息对应当前后端最近一次范围探测调用。

### 集成与边界

上游文档要求 PostgreSQL 15 或更高版本，并提供 PostgreSQL 18 Docker 环境。随项目提供的 `vector` 分叉增加了 HNSW 行为，普通 pgvector 安装不能直接替代。扩展通过共享遍历构建语义连接候选，近似探测得到的配对集合可能与 DASE 的另一种方法不同。

控制文件将扩展设为不可迁移，函数由原生 C 代码实现，文档中的工作流使用超级用户。只应传入可信 SQL、匹配的向量维度和预期的 HNSW 索引；升级与恢复时需保留对应分叉。上游未要求强制预加载。DASE 数据集、Python 运行器和物化连接表与扩展安装相互独立；仓库未为该组件提供明确许可证。
