## 用法

来源：

- [indices developer guide](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/documents/dev_guide.md)
- [pg_extension control file](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/pg_extension.control)
- [SQL-facing Rust API](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/lib.rs)
- [Inference and SPI implementation](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/bindings/inference.rs)
- [Cargo features and dependencies](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/Cargo.toml)

`pg_extension` 0.1.0 是 `indices` 项目的数据库内研究组件。它通过 pyo3 嵌入 Python，并暴露实验性的模型选择、性能分析、模型初始化和推理函数。它与仓库自己的数据集、Python 模块、模型产物、配置文件和文件系统布局紧密耦合，并不是可移植的通用机器学习扩展。

### 仓库特定工作流

control 文件要求超级用户，默认 Cargo 特性为 PostgreSQL 14 加嵌入式 Python。开发指南要求先复现其容器、Python 依赖、数据库表和 `/project/...` 绝对路径，再创建扩展：

```sql
CREATE EXTENSION pg_extension;

SELECT inference(
  'frappe',
  '{"2":977}',
  '/project/Trails/internal/ml/model_selection/config.ini',
  '/project/Trails/frappe_col_cardinalities',
  '/project/tensor_log/frappe/dnn_K16_alpha4',
  'WHERE col2=''977:1''',
  10000
);
```

该示例只是复现上游实验的目标，不能作为接受任意应用输入的安全模板。

### 接口分组

- 模型选择：`model_selection`、`model_selection_workloads`、`model_selection_trails` 及 coordinator/filter/refinement 辅助函数。
- 性能分析与基准：`profiling_filtering_phase`、`profiling_refinement_phase` 和延迟基准函数。
- 推理：`inference`、`inference_shared`、`inference_shared_write_once`、整数/共享内存变体及 `model_init`。
- 构建矩阵：Cargo 声明 PostgreSQL 11–15 特性；默认是 PostgreSQL 14 加 `python`。

### 安全性与正确性边界

名称 `pg_extension` 与 PostgreSQL 的 `pg_catalog.pg_extension` 系统目录冲突；查询系统目录时必须显式限定。

许多函数被声明为 `IMMUTABLE PARALLEL SAFE`，但实现会读取数据库表、加载并修改 Python/模型状态、访问文件、创建 `my_shared_memory` 等固定名称的操作系统共享内存、写日志和测量时间。这些规划器标记与实现不符。没有完整审计并修正 volatility/parallel 声明前，不要在索引、生成列、约束或并行生产查询中使用这些函数。

数据集名称和调用方提供的 SQL 片段会被拼接进 SPI 查询。固定共享内存名称和未检查的 `unwrap` 路径也使并发调用不安全。所有参数都应视为管理员控制，并把实验与生产数据隔离；执行前必须审查文件系统、Python 导入、本地库、内存及 SQL 注入边界。
