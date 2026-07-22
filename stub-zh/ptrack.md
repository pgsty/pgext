## 用法

来源：

- [Pinned official documentation](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/README.md)
- [Pinned control file](https://github.com/postgrespro/ptrack/blob/62bd42d7308ec7ed786abf2ddc1af24fc3dd72dd/ptrack.control)

`ptrack` 是用于 PostgreSQL 增量备份的块级变更跟踪引擎，主要由 `pg_probackup` 使用。它在映射中标记发生变化的关系页面，使备份工具只复制某个 LSN 之后改变的页面。它可能产生误报，但设计目标是不漏掉已变更页面，只有仅修改 hint bit 的情况例外。

### 前提与配置

仅安装扩展软件包并不够。2.4 版要求 PostgreSQL 核心构建应用匹配的 `ptrack` 补丁；上游仓库分别提供 PostgreSQL 11–18 补丁。补丁后的服务端工具知道如何处理 `ptrack.map.*` 服务文件。

配置共享库和非零映射大小，然后重启并创建扩展：

```ini
shared_preload_libraries = 'ptrack'
ptrack.map_size = 64
wal_level = 'replica'
```

```sql
CREATE EXTENSION ptrack;

SELECT ptrack_version();
SELECT ptrack_init_lsn();
```

`ptrack.map_size` 以 MiB 为单位。上游的大小估算规则约为预期 `PGDATA` 每 1 GiB 配置 1 MiB。默认值 `0` 会禁用跟踪并清理服务文件。

### 面向备份的 API

```sql
SELECT *
FROM ptrack_get_change_stat('0/285C8C8'::pg_lsn);

SELECT path, pagecount, pagemap
FROM ptrack_get_pagemapset('0/185C8C0'::pg_lsn);
```

- `ptrack_version()` 返回扩展版本。
- `ptrack_init_lsn()` 返回最近一次初始化映射时的 LSN。
- `ptrack_get_change_stat(pg_lsn)` 返回某个 LSN 之后变化的文件数、页面数和大小汇总。
- `ptrack_get_pagemapset(pg_lsn)` 为备份工具返回每个变更数据文件的路径、块数与位图。

### 运维边界

必须使用 `wal_level >= replica`；若设为 `minimal`，崩溃恢复可能漏掉未写 WAL 的操作所产生的变更。上游文档把 `pg_probackup` 列为完整支持该引擎、可用于生产的备份工具。

映射大小不能在运行时改变。重启时调整大小会丢失已累计的跟踪状态，应同时执行新的全量备份。检查点持久化会使用临时映射文件，因此最多需预留 `ptrack.map_size` 两倍的额外磁盘空间。映射过小会增加误报。

升级可能使旧映射文件失效或要求删除它们，尤其是 2.1→2.2 和 2.2→2.3。应停止服务，遵循对应版本的升级顺序，同时升级 PostgreSQL 核心补丁和扩展二进制，重启后再执行 `ALTER EXTENSION ptrack UPDATE`。依赖它进行恢复前，必须在完全一致的构建上证明备份与恢复行为。
