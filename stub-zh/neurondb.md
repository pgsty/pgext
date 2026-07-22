## 用法

来源：

- [官方 README](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/README.md)
- [官方控制文件](https://github.com/neurondb/neurondb/blob/66284d0476a864bbe074017b517a34daef454fba/neurondb.control)
- [官方文档](https://www.neurondb.ai/docs)

`neurondb` 是一款专有 PostgreSQL 扩展，提供向量类型与检索、库内机器学习与嵌入推理、混合检索及可选后台工作进程。本次核验的源码自称开发版本 `4.0.0-devel`；把生产数据存入其类型或索引前，必须确认许可证与实际二进制所实现的 SQL 接口。

### 核心流程

首先验证扩展构建，不要假定宣传的每个模块都存在：

```sql
CREATE EXTENSION neurondb;

SELECT neurondb.version();

SELECT n.nspname AS schema_name,
       p.proname AS function_name,
       pg_get_function_identity_arguments(p.oid) AS arguments
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
WHERE n.nspname = 'neurondb'
ORDER BY p.proname, arguments;
```

上游项目记录了 `vector`、`halfvec`、`binaryvec`、`sparsevec` 等向量类型，`<->`、`<=>`、`<#>` 等距离运算符，以及 HNSW 和 IVF 索引。采用运算符类或调优参数时应遵循版本匹配的索引文档，并测试实际召回率。

### 主要接口

- 向量存储、距离计算、近似 HNSW/IVF 检索及量化。
- 面向 SQL 的模型训练、预测、嵌入、混合检索、重排与 RAG 辅助函数。
- 名为 `neuranq`、`neuranmon`、`neurandefrag` 与 `neuranllm` 的可选工作进程。
- 面向 CUDA、ROCm 与 Metal 的可选平台特定 GPU 路径。

接口范围取决于构建与许可证。应以已安装目录及获得许可的制品对应厂商文档为准，不要把仓库概览中的函数数量或性能宣称当作证据。

### 前置条件与运维

控制文件要求超级用户安装、标记扩展为不可信，并说明使用后台工作进程时要把 `neurondb` 放入 `shared_preload_libraries` 后重启。README 声明支持 PostgreSQL 16 至 18，并采用限制商业使用、修改与再分发的专有许可证；须取得适用于实际部署二进制的条款。

向量索引是近似的，并增加构建内存、磁盘、WAL 与写维护成本。模型和 GPU 功能还会引入原生库、模型文件、后端内存，以及可能的外部提供商凭据或流量。应基准测试召回率与失败行为、限制资源、收紧执行权限，并在生产形态数据副本上验证升级和转储恢复。
