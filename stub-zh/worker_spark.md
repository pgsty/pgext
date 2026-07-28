## 用法

来源：

- [官方上游 README](https://api.pgxn.org/src/worker_spark/worker_spark-0.0.1/README.md)

`worker_spark` — 一个用于 PostgreSQL 9.3 的后台工作者，定期执行一个过程。使用它来进行相应的调度、时间序列或时间工作流。在目标 PostgreSQL 构建中测试链接的上游固定版本作为 API 边界。

### 核心工作流

经过审核的分发使用了过时的 SQL 或非控制文件安装布局，因此没有建立现代独立的 `CREATE EXTENSION` 和升级工作流。遵循链接的上游安装机制，并在隔离数据库中验证安装的对象。

### 要求与注意事项

- 该目录记录了版本 `0.0.1`。
- 在生产使用前，确认权限、支持的 PostgreSQL 版本、升级行为和失败情况与链接的源代码一致。
