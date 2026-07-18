## 用法

来源：

- [固定提交的当前上游 README](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/README.md)
- [1.1 版安装 SQL](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/pg_statviz--1.1.sql)
- [固定提交的扩展控制文件](https://github.com/vyruss/pg_statviz/blob/b9af6214a10d9295de54b27842a7e0da1e3066fa/pg_statviz.control)
- [正式 PGXN 发行](https://pgxn.org/dist/pg_statviz/)

pg_statviz 1.1 版是纯 SQL 的统计快照扩展，并配套一个外部 Python 可视化工具。每次快照都会在固定的 pgstatviz 模式中保存选定的缓冲区、配置、连接、数据库、I/O、锁、复制、SLRU、等待事件和 WAL 计数器。

### 采集快照

```sql
CREATE EXTENSION pg_statviz;

SELECT pgstatviz.snapshot();

SELECT snapshot_tstamp, conn_total, conn_active,
       max_query_age_seconds, max_xact_age_seconds
FROM pgstatviz.conn
ORDER BY snapshot_tstamp DESC
LIMIT 10;
```

应使用外部作业运行器，按所需分辨率安排快照调用。单独安装的 pg_statviz 命令读取一个时间范围并生成图表。当前 README 要求工具使用 Python 3.11+，并声明支持截至 18 的近期 PostgreSQL 版本。

### 保留、权限与可选 AI

快照会持续累积，直到删除行或由 delete_snapshots 截断所有快照表。频繁采集前应设置保留作业并估算存储。统计值是累积值，也可能被单独重置；应分析增量，而不是把原始计数器直接当作速率。

安装 SQL 给内置 pg_monitor 角色授予模式使用权、所有函数的执行权，以及所有 pgstatviz 表的 SELECT、INSERT、DELETE 和 TRUNCATE 权限。这意味着任一成员既能采集也能删除历史，包括调用全量删除函数。若采集与保留管理需要分离，应使用专用角色或修改授权。

工具的可选 AI 模式可以把图表数据、采集的配置、主机/角色上下文和确定性规则结果发送给 Claude 或 Gemini，也支持本地 Ollama。普通可视化并不需要该模式。启用云服务商前必须审查数据分类、服务商保留策略、提示词内容、凭据和出站网络策略。
