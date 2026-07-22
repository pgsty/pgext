## 用法

来源：

- [官方配置、schema 与警告](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/README.md)
- [0.4.0 版本安装 SQL](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record--0.4.0.sql)
- [扩展控制文件](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record.control)

`stat_record` 运行一个后台工作进程，定期把集群、数据库、对象、查询、WAL 和配置统计快照写入固定的 `_stat_record` schema。它无需外部调度器即可生成历史报表，但上游仓库已废弃，且明确警告存在缺陷。

### 核心流程

安装扩展及其 `pg_stat_statements` 依赖，配置两个库使用共享预加载，然后重启 PostgreSQL：

```sql
CREATE EXTENSION stat_record CASCADE;
```

```conf
shared_preload_libraries = 'pg_stat_statements,stat_record'
stat_record.database_name = 'postgres'
stat_record.interval = 3600
stat_record.retention = 7
```

重启后，`stat_record worker` 会按配置间隔采集快照。也可手动采集：

```sql
SELECT _stat_record.take_record();
SELECT * FROM _stat_record.lastest_records(10);
```

### 存储数据与报表

- `_stat_record._record_number`、`_global_stat`、`_db_stat` 和 `_query_stat` 保存快照标识符与统计信息。
- `_stat_record.detail_record(bigint)`、`global_report_record(bigint, bigint)` 和 `total_report_record(bigint, bigint)` 提供快照之上的报表视图。
- `_stat_record.delete_record(bigint)` 删除一个快照；`truncate_record(boolean)` 删除全部历史，并可重置编号。
- `_stat_record.export_total_report_record(bigint, bigint, text)` 在服务器端把 CSV 文件写入指定路径。

### 安全与兼容性

上游声明支持 PostgreSQL 10 或更高版本，但没有说明是否针对现代统计目录变更做过测试。版本 `0.4.0` 没有声明许可证，并带有明确缺陷警告。工作进程使用一个已配置数据库，保存查询文本和广泛服务器元数据，且可创建服务器端文件。应审查工作进程权限、查询文本敏感性、输出路径、保留与表增长、重启行为和 PostgreSQL 大版本兼容性。生产历史应优先使用受维护的监控体系，并在工作进程停止后测试删除或升级。
