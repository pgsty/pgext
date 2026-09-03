## 用法

来源：

- [EDB Clone Schema 官方文档](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/)
- [官方设置指南](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/setting_up_edb_clone_schema/)
- [官方本地复制指南](https://www.enterprisedb.com/docs/epas/latest/database_administration/14_edb_clone_schema/copying_a_schema/)

`edb_cloneschema` 使用 foreign server、scheduler job 与 parallel clone worker，在本地或 EPAS 数据库之间复制 schema 及其数据库对象。

### 前置条件与启用

安装并预加载 EPAS worker 软件包、重启，然后在每个参与数据库中创建全部依赖扩展：

```ini
shared_preload_libraries = 'parallel_clone,edb_job_scheduler'
```

```sql
CREATE EXTENSION postgres_fdw SCHEMA public;
CREATE EXTENSION dblink SCHEMA public;
CREATE EXTENSION edb_job_scheduler;
CREATE EXTENSION dbms_job;
CREATE EXTENSION parallel_clone;
CREATE EXTENSION edb_cloneschema;
CREATE TRUSTED LANGUAGE plperl;
SELECT edb_util.create_clone_log_dir();
```

Source/target user mapping 必须认证具备读取与创建全部被复制对象权限的角色。

### 复制 Schema

按照 local 或 remote 模式文档定义 foreign server 与 mapping，再调用对应的 `localcopyschema`/`localcopyschema_nb` 或 `remotecopyschema`/`remotecopyschema_nb` 函数。Nonblocking variant 通过 EDB Job Scheduler 调度工作；使用 `process_status_from_log` 跟踪状态文件。

### 运维边界

Schema copy 可能产生大型事务、WAL、lock、index、foreign key 与 background-worker 需求。应根据计划对象集调整 `work_mem`、`maintenance_work_mem`、`max_worker_processes`、checkpoint 设置、WAL 大小与 `max_locks_per_transaction`。切换前应保护 credential 与日志文件，测试不支持的对象类型，并验证复制后的 owner、privilege、dependency 与数据一致性。

