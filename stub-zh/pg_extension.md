## 用法

来源：

- [已复核 commit 的 indices 开发指南](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/documents/dev_guide.md)
- [已复核 commit 的 pg_extension Rust SQL API](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/src/lib.rs)
- [已复核 commit 的 pg_extension Cargo 配置](https://github.com/NLGithubWP/indices/blob/93ce9d11a66b8f2ae4cb5422a874de49e116e0b4/internal/pg_extension/Cargo.toml)

`pg_extension` 0.1.0 是 indices 项目的数据库内研究组件。它通过 Rust/pgrx 暴露模型选择、模型切片、性能分析与推理函数。文档中的 `inference` 调用接收数据集、JSON 条件、配置与基数文件、模型路径、SQL 过滤片段和批大小。

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

该示例只适用于仓库自身环境；必须复现其 Python 环境、数据表、模型文件和配置布局后才能运行。

### 注意事项

- 名称与 PostgreSQL 的 `pg_extension` 系统目录冲突；同时讨论或查询二者时，应限定系统目录引用。
- control 文件要求超级用户。默认构建嵌入 Python 支持，只声明 PostgreSQL 11 至 15 的特性，默认是 PostgreSQL 14。
- 源码和指南包含部署专用绝对路径，并假定特定的数据集表名与列数；这不是可移植的通用推理 API。
- 多个 API 被声明为 immutable 和 parallel safe，但实现会查询表、加载模型或使用共享内存。独立审计之前，应把这些规划器标记视为不安全。
- SQL 过滤参数会拼入动态构造的查询；只应接受来自可信且已验证配置的值。
