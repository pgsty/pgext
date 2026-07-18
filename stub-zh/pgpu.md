## 用法

来源：

- [上游 README](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/README.md)
- [扩展 control 文件](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/pgpu.control)
- [Cargo 软件包元数据](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/Cargo.toml)
- [VectorChord 索引创建实现](https://github.com/EnterpriseDB/pgpu/blob/b39cc5e3fa49efb673c524d722ea7c812228bc7e/vectorchord-indexing/src/vectorchord_index.rs)

`pgpu` `2.1.0` 版使用 NVIDIA GPU 计算 VectorChord 索引的向量质心。它分批读取 pgvector 数据，运行 cuVS k-means，写入质心表，并可选择用这些外部质心删除/重建 VectorChord 索引。

### 冒烟测试

构建需要文档指定的 CUDA 13.1/cuVS 环境，以及匹配的 PostgreSQL、`vector` 和 `vchord` 安装。安装后执行：

```sql
CREATE EXTENSION vector;
CREATE EXTENSION vchord;
CREATE EXTENSION pgpu;
SELECT pgpu.ping();
SET pgpu.gpu_acceleration = 'enable';
```

主入口是 `pgpu.create_vector_index_on_gpu(...)`；它会修改模式且大量使用 GPU 和内存，因此应先在副本上测试参数和召回率。`2.1.0` 的 Cargo 特性支持 PostgreSQL 14–18，但复核的索引构建器会删除生成的索引名称，并在最终 VectorChord DDL 中把索引列硬编码为 `embedding`。应只允许表所有者执行，并核对生成名称/DDL、GPU 资源上限、距离运算符、维度和恢复行为。
