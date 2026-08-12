## 用法

来源：

- [pg_statviz v1.1 发行说明](https://github.com/vyruss/pg_statviz/releases/tag/v1.1)
- [pg_statviz v1.1 README](https://github.com/vyruss/pg_statviz/blob/v1.1/README.md)
- [pg_statviz v1.1 安装 SQL](https://github.com/vyruss/pg_statviz/blob/v1.1/pg_statviz--1.1.sql)
- [pg_statviz v1.1 控制文件](https://github.com/vyruss/pg_statviz/blob/v1.1/pg_statviz.control)
- [pg_statviz v1.1 元数据](https://github.com/vyruss/pg_statviz/blob/v1.1/META.json)
- [pg_statviz v1.1 Python 软件包元数据](https://github.com/vyruss/pg_statviz/blob/v1.1/pyproject.toml)
- [pg_statviz v1.1 AI 服务商实现](https://github.com/vyruss/pg_statviz/blob/v1.1/src/pg_statviz/libs/ai.py)
- [正式 PGXN 分发](https://pgxn.org/dist/pg_statviz/)

`pg_statviz` v1.1 由一个纯 SQL 与 PL/pgSQL 的统计快照扩展和一个单独安装的 Python 可视化工具组成。扩展把 PostgreSQL 的累积及动态统计保存在固定的 `pgstatviz` 模式中；工具读取选定时间范围，并生成图表或可选的 AI 辅助 HTML 报告。它要求 PostgreSQL 13 或以上版本，不需要 `shared_preload_libraries`，也无需重启。工具要求 Python 3.11 或以上版本。

### 采集并保留快照

由管理员安装扩展，然后让专用采集角色继承 `pg_monitor`，再通过 cron 或其他外部作业运行器定期调用 `pgstatviz.snapshot()`。

```sql
CREATE EXTENSION pg_statviz;

GRANT pg_monitor TO stats_collector;

SELECT pgstatviz.snapshot();

DELETE FROM pgstatviz.snapshots
WHERE snapshot_tstamp < CURRENT_DATE - 90;
```

删除父表行会级联删除相应样本；`pgstatviz.delete_snapshots()` 则会截断全部历史。应根据需要观测的最短事件与相应表增长量选择采集间隔和保留窗口；PostgreSQL 原始计数器是累积值且可能独立重置，因此应分析带时间戳的增量，不能把存储值直接当作速率。

### 存储数据与版本边界

主要关系包括 `pgstatviz.snapshots`、`pgstatviz.buf`、`pgstatviz.conf`、`pgstatviz.conn`、`pgstatviz.db`、`pgstatviz.io`、`pgstatviz.lock`、`pgstatviz.repl`、`pgstatviz.slru`、`pgstatviz.wait` 和 `pgstatviz.wal`。样本会包含配置值、连接用户名与时长、复制应用及槽名称、等待、锁、I/O、数据库计数器和 WAL 计数器。应把这些表、转储、图表与报告作为运维数据加以保护。

配置只在发生变化时保存，因此 `pgstatviz.conf` 不一定对应每次快照都有一行。PostgreSQL 14 及以上版本采集 `pg_stat_wal`，PostgreSQL 16 及以上版本采集 `pg_stat_io`，并单独处理 PostgreSQL 18 基于字节的字段。较早的受支持版本仍会创建这些表，但会跳过不可用的采集器。

扩展把快照表标记为可感知扩展的转储对象，因此可以用 `pg_dump` 搬迁历史，但仍需主动限制保留量与备份大小。

### 可视化时间范围

可视化工具需要单独安装，并接受常规 libpq 连接选项。`analyze` 命令会运行全部分析模块；只需要较窄的报告时，可以选择 `conn`、`io`、`wait` 和 `wal` 等单个模块。

```bash
pip install pg_statviz

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports
```

应限制数据库凭据与报告目录的访问权限。可视化角色只需读取已采集模式，不需要采集或删除快照的权限。

### 权限边界

v1.1 安装 SQL 会向 `pg_monitor` 的所有成员授予模式使用权、函数执行权，以及全部 `pgstatviz` 表上的 `SELECT`、`INSERT`、`DELETE` 与 `TRUNCATE`。因此，该成员身份同时允许采集快照，并能通过 `pgstatviz.delete_snapshots()` 删除全部历史；它并不是只读可视化角色。

如果必须分离采集、可视化和保留管理，应在安装后修订默认授权，只向专用角色授予所需函数与表权限。扩展升级后应再次检查这些授权。

### 可选 AI 与云端数据审查

普通图表生成不会请求 LLM。AI 模式需要可选的 `pg_statviz[ai]` 依赖，并显式使用 `--ai` 参数。Claude 是默认云服务商并读取 `ANTHROPIC_API_KEY`；Gemini 读取 `GOOGLE_API_KEY`；`--ai local` 使用本地 Ollama 服务。当前默认模型为 `claude-sonnet-4-6`、`gemini-2.5-flash` 与 `gemma4:e4b`；这些只是实现默认值，并不保证服务商账户或本地运行时会持续提供相应模型。

```bash
pip install 'pg_statviz[ai]'

pg_statviz analyze \
  -h /var/run/postgresql -d mydb -U stats_reader \
  -D 2026-08-01T00:00 2026-08-02T00:00 \
  -O /srv/pg_statviz/reports \
  --ai gemini
```

使用云服务商时，请求可能包含图表图像和汇总时间序列，以及采集到的 PostgreSQL 版本、主库/备库角色、主机名、相关配置值、确定性检查结果、用户或角色名称和复制标识符。应把这视为一次明确的运维数据导出：审查服务商保留与区域政策，缩小所选时间范围，保护生成的 HTML 与 PNG 文件，并使用获准的出站路径。提示词中的数据封装可以降低提示注入风险，但不提供机密性或授权能力，也不能替代服务商治理。
