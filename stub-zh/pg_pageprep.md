## 用法

来源：

- [上游 WIP 文档](https://github.com/postgrespro/pg_pageprep/blob/981b279f1af697adb954bcc6de87297371c0bcaf/README.md)
- [扩展控制文件](https://github.com/postgrespro/pg_pageprep/blob/981b279f1af697adb954bcc6de87297371c0bcaf/pg_pageprep.control)

`pg_pageprep` 是尚未完成的迁移辅助工具，用于为 PostgresPro Enterprise 的 64 位事务 ID 页面格式预留普通堆页面空间。已复核代码面向从 PostgreSQL 或 PostgresPro Standard 10/11 迁移到 PostgresPro Enterprise 11。

```conf
shared_preload_libraries = 'pg_pageprep'
pg_pageprep.database = 'postgres'
pg_pageprep.role = 'postgres'
```

重启后需在每个数据库创建扩展，并通过附带脚本或函数管理工作进程。处理时可能把关系 fillfactor 改为 90；升级后应停止工作进程并恢复原 fillfactor。这是侵入性且版本特定的迁移软件：必须在恢复副本上演练，盘点不支持的关系，监控膨胀与 I/O，准备已验证备份，并且只配合匹配的 PostgresPro 迁移流程使用。
