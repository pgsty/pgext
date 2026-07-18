## 用法

来源：

- [上游固定版本 README](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/README.md)
- [1.0 版本安装 SQL](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/pg_feedback--1.0.sql)
- [固定版本 PostgreSQL 集成实现](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/pg_feedback.c)
- [固定版本操作系统辅助实现](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/helpers.c)
- [固定版本发行元数据](https://github.com/ligurio/pg_feedback/blob/8d11c8ca08b23bf0e8ae7957c506ddd9ba37b56b/META.json)

pg_feedback 在本地生成 PostgreSQL 安装情况的 JSON 报告。报告包含集群系统标识、一个标为 postmaster 运行时长的值、数据目录文件系统类型、服务器版本、部分设置、可用扩展和数据库总大小。扩展本身不会发送报告；是否发送以及发送到哪里完全由调用方决定。

### 在本地生成并审查

```sql
CREATE EXTENSION pg_feedback;

SELECT jsonb_pretty(feedback()::jsonb);
```

在导出前应把结果当作敏感运维清单审查。SQL 对象版本为 1.0，与本目录一致；固定源码中的 PGXN 风格发行元数据版本则是 0.2.2，两者属于不同版本空间。

### 数据质量与暴露面

实现中存在实质性的正确性缺陷。它把 64 位集群标识写入堆缓冲区时使用了指针本身的大小，在常见平台会截断标识。名为 postmaster_uptime 的函数实际返回父进程 PID。Linux 下独立的 sysinfo 辅助函数把系统运行时间写进内存字段，又把内存大小写进运行时间字段；主报告目前没有使用它，但该函数仍可直接调用。

安装 SQL 没有撤销默认的 PUBLIC 执行权。能够调用函数的角色可看到设置、已安装能力、数据库总大小、平台细节和派生自集群的标识。应撤销公共执行，只授权监控角色；未经明确的隐私与目标端审查，不要上传输出。

C 代码读取 PostgreSQL 控制文件内部结构和操作系统 API，因此必须针对精确大版本与平台测试。应把报告字段视为提示而非权威遥测，对缺失或截断值明确标注；生产清单优先采用受维护的 PostgreSQL 视图与主机监控。
