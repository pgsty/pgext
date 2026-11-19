## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_pipeline/pg_pipeline.c)

`pg_pipeline` `1.0` 提供DAG 任务调度与执行历史。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_pipeline;
SELECT pipeline_create('demo');
SELECT pipeline_task_add('demo', 'step1', 'sql', 'SELECT 1');
SELECT pipeline_run('demo');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

SQL 函数负责入队并管理依赖和历史。执行队列任务需预加载 pg_pipeline、重启并设置 pg_pipeline.database，pg_pipeline.check_interval 控制轮询。任务 SQL 在高权限后台工作进程中运行，因此仅可信操作人员应能编辑定义或启动任务。max_parallel 或已入队的运行标识不代表任务已完成，需检查任务、运行历史和取消状态。

此模块对应上游旧名 `alohadb_pipeline`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
