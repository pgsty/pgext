## 用法

来源：

- [README.md](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/README.md)
- [contrib/alohadb_cron/alohadb_cron--1.0.sql](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron--1.0.sql)
- [contrib/alohadb_cron/alohadb_cron.c](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron.c)
- [contrib/alohadb_cron/alohadb_cron.control](https://github.com/Vern-AllWorks-LLC/alohadb/blob/ce440857c0d0982f684a1dd2adfd5fa6514d31a9/contrib/alohadb_cron/alohadb_cron.control)

`alohadb_cron` 1.0 是 AlohaDB 的 SQL 任务调度器，后台进程轮询配置的数据库，执行到期命令并记录运行明细。

### 核心用法

```ini
shared_preload_libraries = 'alohadb_cron'
alohadb.cron_database = 'postgres'
```

```sql
CREATE EXTENSION alohadb_cron;
SELECT cron_schedule_named('health-check', '*/5 * * * *', 'SELECT 1');
SELECT * FROM cron_job_status();
SELECT cron_unschedule_named('health-check');
```

### 运行边界

将库追加到现有预加载列表并重启，然后由超级用户在指定数据库创建扩展。`alohadb.cron_database` 默认为 postgres，`alohadb.cron_check_interval` 控制轮询周期；后台进程要求相关表位于 public。

`cron_schedule` 与 `cron_schedule_named` 创建任务，`cron_unschedule` 与 `cron_unschedule_named` 删除任务，`cron_job_status` 汇总状态。可检查 `alohadb_cron_job_run_details`，并为持续增长的历史设置保留策略。

核对的后台进程以默认后台进程超级用户身份连接，并通过 SPI 执行命令；每个任务的 database 和 username 参数未用于切换连接或角色。应将建任务及改表权限限制在可信管理员范围内，此实现不提供租户隔离。它是 AlohaDB 组件，并非经过验证的原生 PostgreSQL 软件包。
