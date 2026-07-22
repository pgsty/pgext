## 用法

来源：

- [官方 README](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/README.md)
- [版本 1.0 扩展 SQL](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota--1.0.sql)
- [后台工作进程与 GUC 实现](https://github.com/hlinnaka/pg_quota/blob/0e86613bc39c1653f4c562b2cfc2b79aa3eeebe6/pg_quota.c)

`pg_quota` 跟踪关系文件所有权，并按数据库角色施加软磁盘空间配额。它适合在角色超过配置配额后检测并阻止部分新写入；它不是严格的存储边界。

### 预加载与启用

把库与目标数据库加入 `postgresql.conf`，然后重启 PostgreSQL。数据库列表是 postmaster 级设置，每个列出的数据库都会注册一个后台工作进程。

```conf
shared_preload_libraries = 'pg_quota'
pg_quota.databases = 'appdb'
pg_quota.refresh_naptime = '10 s'
```

重启后，在每个已配置数据库中创建扩展，并以字节数设置配额：

```sql
CREATE EXTENSION pg_quota;

INSERT INTO quota.config(roleid, quota)
VALUES ('app_writer'::regrole, pg_size_bytes('10 GB'))
ON CONFLICT (roleid)
DO UPDATE SET quota = EXCLUDED.quota;

SELECT rolname,
       pg_size_pretty(space_used) AS used,
       pg_size_pretty(quota) AS quota
FROM quota.status;
```

`quota.config` 是持久配置表。`quota.status` 报告每个被跟踪角色当前占用的关系空间和配置配额；quota 为 NULL 表示该角色被跟踪但不受限。

### GUC

- `pg_quota.databases` 选择数据库，修改后必须重启。
- `pg_quota.refresh_naptime` 控制完整数据目录扫描间隔，可通过重载生效。
- `pg_quota.restart_interval` 控制工作进程重启延迟，属于 postmaster 级设置。

应确保 `max_worker_processes` 除每个配置数据库的一个工作进程外，还为其他扩展与并行工作保留足够容量。

### 软配额限制

空间归属关系的直接所有者；不支持组配额。临时文件、临时关系与目录对象不会计入。所有权变更和新建关系只有在提交并完成后续扫描后才会反映，因此状态存在延迟；关系或分区很多时扫描也可能很昂贵。

执行只在 `INSERT` 与 `COPY` 开始时检查配额。语句开始时低于配额仍可能在执行后超额，而 `UPDATE`、`CREATE INDEX` 及其他工具命令不会被配额钩子拦截。应仔细测试失败行为与恢复流程，不要把 `pg_quota` 当作安全或计费级硬限制。
