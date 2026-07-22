## 用法

来源：

- [Official README](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/README)
- [Extension SQL](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/resources--0.1.0.sql)
- [Extension control file](https://github.com/asotolongo/resources/blob/911dae0439eeed46760e2ef294047b079a509fb4/resources.control)

resources 是一组 Linux 专用的 `file_fdw` program 外部表及视图，用于读取 CPU、内存、操作系统、文件系统和 PostgreSQL 进程信息。每次扫描都会在数据库服务器上运行固定的 shell 管道。上游警告当前版本存在缺陷。

### 核心流程

只能由超级用户在兼容 Linux 主机上安装，然后查询预先创建的对象：

```sql
CREATE EXTENSION resources CASCADE;

SELECT * FROM resources.cpu_loadavg;
SELECT * FROM resources.meminfo;
SELECT * FROM resources.partitions;
SELECT * FROM resources.v_ps_postgresv2;
```

这些外部表不存储应用数据；查询时会解析 `cat`、`awk`、`top`、`lscpu`、`df` 和 `ps` 等命令的结果。

### 重要对象

- `resources.cpu_loadavg` 和 `resources.cpu_use` 暴露负载均值及 CPU 状态输出。
- `resources.distribution` 和 `resources.lscpu` 暴露内核及 CPU 信息。
- `resources.meminfo` 把 Linux 内存信息换算为 GB。
- `resources.partitions` 解析文件系统利用率。
- `resources.ps_postgres` 和 `resources.ps_postgresv2` 解析硬编码操作系统用户 postgres 所拥有的进程。
- `resources.v_ps_postgres` 和 `resources.v_ps_postgresv2` 把进程输出与 PostgreSQL 活动数据连接起来。

### 安全与可移植性说明

program 型 `file_fdw` 表会以 PostgreSQL 服务账户执行操作系统命令，因此设置和访问都需要高权限。应撤销不可信角色对外部表及服务器的访问。输出格式会随 Linux 发行版、locale、命令版本、cgroup/容器环境及服务账户名变化，解析可能失败或静默错标列。某些命令在每次查询时休眠或扫描进程/文件系统状态。不要把这些视图作为可移植监控 API；应优先使用 PostgreSQL 外部仍受维护的主机 exporter。
