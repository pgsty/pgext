## 用法

来源：

- [阿里云 oss_fdw 文档](https://www.alibabacloud.com/help/en/polardb/polardb-for-oracle/oss-fdw)

`oss_fdw` 是阿里云 PolarDB 的服务商扩展，可把对象存储服务中的目录或文件前缀映射成 PostgreSQL 外部表。它面向受支持的 PolarDB PostgreSQL 兼容服务版本上的追加式归档与冷数据，不是供自建 PostgreSQL 使用的可移植社区 FDW。

### 核心流程

启用服务商扩展，定义 OSS 服务器，然后把一个目录映射为外部表：

```sql
CREATE EXTENSION oss_fdw;

CREATE SERVER archive_oss
FOREIGN DATA WRAPPER oss_fdw
OPTIONS (
    host 'oss-cn-hangzhou-internal.aliyuncs.com',
    bucket 'example-archive',
    id '<access-key-id>',
    key '<access-key-secret>'
);

CREATE FOREIGN TABLE public.events_archive (
    event_id bigint,
    occurred_at timestamptz,
    payload text
)
SERVER archive_oss
OPTIONS (dir 'events/', format 'csv', compressiontype 'gzip');

INSERT INTO public.events_archive
SELECT event_id, occurred_at, payload
FROM public.events
WHERE occurred_at < current_date - 90;

SELECT count(*) FROM public.events_archive;
SELECT * FROM oss_fdw_list_file('events_archive', 'public');
```

对象映射必须在 `dir` 与 `prefix` 中恰好选择一个。`format` 默认为 CSV。`compressiontype` 可以选择 gzip，或在受支持的 PostgreSQL 14 服务版本上选择 Zstandard；`compressionlevel` 控制服务商支持的压缩级别范围。

### 写入与删除语义

支持 `SELECT`、`INSERT` 和 `TRUNCATE`，不支持 `UPDATE` 与 `DELETE`。每次执行 `INSERT` 都会新建一个 OSS 文件，而不是修改现有文件。`TRUNCATE public.events_archive` 会删除该外部表映射的全部 OSS 文件，因此它是破坏性的对象存储删除操作，不能当成普通本地表清理。

删除外部表或扩展与截断映射数据是不同操作。移除定义前应确认服务商行为，并保留独立的 OSS 留存策略或备份。

### 服务商与安全边界

服务器选项包含 OSS 访问凭据。应使用只拥有最低桶权限的专用 RAM 身份，限制外部服务器上的 `USAGE`，并避免不受信角色查看目录元数据或执行 DDL。服务支持时，应优先使用与 PolarDB 集群同地域的内网端点；网络位置会同时影响可达性与性能。

可用性、精确版本、支持的 PostgreSQL 版本、地域、压缩算法与权限都由阿里云控制。运行示例前，应检查目标集群的扩展列表和服务文档。访问 OSS 的延迟与事务行为和本地 PostgreSQL 存储差异很大；用于归档流程前，应测试重试、部分失败、文件累积以及下游读取端。
