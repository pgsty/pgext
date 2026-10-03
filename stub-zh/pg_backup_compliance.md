## 用法

来源：

- [README.md](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/README.md)
- [pg_backup_compliance.control](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance.control)
- [sql/pg_backup_compliance--1.0.sql](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/sql/pg_backup_compliance--1.0.sql)
- [pg_backup_compliance.c](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance.c)
- [pg_backup_compliance_capture.c](https://github.com/klouddb/pg_backup_compliance/blob/40876bd5d60c33dfc3451375cb472bba516e5cf1/pg_backup_compliance_capture.c)

`pg_backup_compliance` 1.0 观察备份相关会话，并通过 SQL 视图提供活动记录。它用于监测活动；记录中的成功状态不能证明备份可恢复。

### 启用

将库加入现有预加载列表并重启，再由超级用户在需要查询记录的数据库中创建 SQL 对象。README 声明支持 PostgreSQL 13+，但所引用捕获钩子使用 PostgreSQL 14+ 的 ProcessUtility 签名且没有 PG13 分支，不能据此认定兼容 PG13。

```ini
shared_preload_libraries = 'pg_backup_compliance'
```

```sql
CREATE EXTENSION pg_backup_compliance;
SELECT application_name, backup_type, status, start_time, end_time
FROM pg_backup_compliance ORDER BY start_time DESC LIMIT 20;
SELECT * FROM pg_backup_compliance_failed;
```

### 对象与配置

所核验的 SQL 中，`pg_backup_compliance` 合并内存与归档的操作记录。`pg_backup_compliance_running`、`pg_backup_compliance_failed`、`pg_backup_compliance_last_24h` 以及月度、季度视图进一步筛选记录。视图向 `pg_monitor` 授予访问权限。`pg_backup_compliance_info()` 返回缓存计数；`pg_backup_compliance_reset()` 仅限超级用户。

主要配置项为 `pg_backup_compliance.enabled`、`pg_backup_compliance.save`、`pg_backup_compliance.max_entries` 与 `pg_backup_compliance.track_apps`；最后一项按应用名称前缀匹配。

### 限制

同一主机、同一用户发起的独立 `pg_dump` 可能被合并到同时运行的 `pg_dumpall` 记录中。部分 pgBackRest 失败无法准确分类。应用名称由客户端提供，缓存淘汰、持久化与归档保留需要结合部署评估；README 与 SQL 对历史记录的描述并不一致。不能把这些视图当作完整保留或恢复能力的保证。上游提供了版权声明，但未发现明确的许可授权。
