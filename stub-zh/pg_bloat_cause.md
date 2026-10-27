## 用法

来源：

- [README](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/README.md)
- [Control](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/pg_bloat_cause.control)
- [Install and grants](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_bloat_cause/blob/b5de883b65784eccf6f9d0c8e65d68d7252349d2/pg_bloat_cause--0.7.0.sql)

`pg_bloat_cause` 0.7.0 根据 PostgreSQL 目录和统计信息，诊断死亡行未被清理的原因。它是纯 SQL 扩展，测试覆盖 PostgreSQL 14–19rc1，需要超级用户安装，不要求预加载或重启。

### 读取诊断

```sql
CREATE EXTENSION pg_bloat_cause;
SELECT * FROM bloat_cause_text();
```

诊断区分被固定的清理边界、无效的 autovacuum 和非 HOT 更新压力。结果属于估计和观察，不是精确的物理膨胀测量。逻辑槽的 catalog_xmin 阻碍目录清理，并不自动固定普通用户表行。操作辅助函数只展示供审阅的命令，不执行修复。

### 权限、历史与告警

读取器需要显式视图、函数授权及 pg_read_all_stats 才能查看其他会话。采样、修剪和 NOTIFY 函数还需历史与告警表的写权限。这些函数由外部调度调用，扩展不安装 worker。记录持有者查询文本需要显式启用，可能保留敏感信息；恢复历史应用数据时应同时恢复相关角色和授权。
