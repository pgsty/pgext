## 用法

来源：

- [上游配置、schema 与警告](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/README.md)
- [扩展控制文件](https://github.com/asotolongo/stat_record_old/blob/4c9d43aba5168f4dcefdd843932b9c8ecf4112e7/stat_record.control)

`stat_record` 运行后台工作进程，定期把集群、数据库、对象、查询、WAL 与配置统计快照写入固定的 `_stat_record` schema，并生成历史报表。它依赖 `pg_stat_statements`。

```conf
shared_preload_libraries = 'pg_stat_statements,stat_record'
stat_record.database_name = 'postgres'
stat_record.interval = 3600
stat_record.retention = 7
```

重启后用 `CREATE EXTENSION stat_record CASCADE` 安装。仓库已废弃、没有声明许可证，而且上游明确警告存在缺陷。工作进程以高权限默认值连接，保存查询文本和大量元数据，还提供服务端 CSV 导出。应审查权限与敏感数据保留、限制路径、监控表增长，并优先使用受维护的监控体系。
