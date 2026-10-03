## 用法

来源：

- [Official README.md](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/README.md)
- [Official pg_dbbackup.control](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/pg_dbbackup.control)
- [Official pg_dbbackup--0.0.1.sql](https://github.com/GerardSmit/pgBackupDatabase/blob/e1bbaf1cb09657986e00bcfe350601bad25b9141/sql/pg_dbbackup--0.0.1.sql)

`pg_dbbackup` 0.0.1 在 PostgreSQL 17 与 18 上提供按数据库执行的逻辑备份与恢复链。安装需要管理员权限，使用固定的 `dbbackup` 模式，并要求共享预加载与逻辑 WAL。请将预加载项合并到现有配置，重启后再创建扩展。

### 创建本地备份

```conf
shared_preload_libraries = 'pg_dbbackup'
wal_level = logical
```

```sql
CREATE EXTENSION pg_dbbackup;
SELECT dbbackup.pg_dbbackup_get_mode('app');
SELECT dbbackup.pg_dbbackup('app', '/var/backups/app-full.bak',
  type := 'full', compress := true);
SELECT * FROM dbbackup.pg_dbbackup_header('/var/backups/app-full.bak');
SELECT * FROM dbbackup.pg_dbbackup_verify('/var/backups/app-full.bak');
```

### 恢复模式与存储

默认 SIMPLE 恢复模式提供全量备份与累积差异备份。FULL 模式增加逻辑解码与 DDL 日志，支持日志备份和时间点恢复；通过 `dbbackup.pg_dbbackup_set_mode()` 选择。路径位于服务端文件系统，而非客户端。S3 目标使用 `dbbackup.create_s3_target()` 与 `dbbackup.pg_dbbackup_to_storage()`；目标记录不含秘密的配置，AWS 凭据来自服务端环境。热备库拒绝本地文件备份。

### 调度、观测与恢复

`dbbackup.create_schedule()` 为备份集建立计划。`dbbackup.pg_dbbackup_async()` 将任务排队；`dbbackup.pg_dbbackup_status()` 与 `dbbackup.pg_dbbackup_wait()` 报告进度。扩展模式中的相应表可检查任务、备份产物与事件。`dbbackup.pg_dbbackup_retention_plan()` 用于在应用保留策略前预览过期项。`dbbackup.pg_dbrestore()` 恢复文件链，`dbbackup.pg_dbrestore_at()` 恢复至指定时间。恢复在暂存替代数据库后会删除已存在的目标数据库，因此必须明确选择并保护目标。

### 限制

它不能替代集群级物理恢复。FULL 备份会取得表级 SHARE 锁；长期无人维护的逻辑槽可能持续保留 WAL。应维护复制槽容量与保留状态，高可用环境还需配置故障转移逻辑槽同步。保护备份文件、口令与云凭据，并在隔离目标上测试恢复。C 库使用 zstd、OpenSSL 和 libcurl，SQL 辅助函数依赖 PL/pgSQL。
