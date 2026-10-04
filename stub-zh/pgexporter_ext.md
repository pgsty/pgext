## 用法

来源：

- [0.2.5 版 README](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/README.md)
- [0.2.5 版配置指南](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/doc/GETTING_STARTED.md)
- [0.2.5 版 SQL 定义](https://github.com/pgexporter/pgexporter_ext/tree/0.2.5/sql)
- [0.2.3 版 SQL 定义](https://github.com/pgexporter/pgexporter_ext/tree/0.2.3/sql)
- [0.2.4 版 SQL 定义](https://github.com/pgexporter/pgexporter_ext/tree/0.2.4/sql)
- [0.2.5 版文件系统实现](https://github.com/pgexporter/pgexporter_ext/blob/0.2.5/src/pgexporter_ext/utils.c)

`pgexporter_ext` 通过 SQL 提供 Linux 主机与文件系统指标，供 pgexporter 采集。以下示例适用于已打包的 0.2.3–0.2.5 版本 API。函数以 PostgreSQL 操作系统账户的权限访问资源，因此监控角色应只授予可信用户。

### 配置与核心流程

上游配置指南将模块加入 `shared_preload_libraries`，然后重启 PostgreSQL。配置时应保留已有的其他库。

```ini
shared_preload_libraries = 'pgexporter_ext'
```

由有权限的管理员在 `postgres` 数据库安装扩展，并将监控权限授予已有的 exporter 登录角色：

```sql
CREATE EXTENSION pgexporter_ext;
GRANT pg_monitor TO pgexporter;

SET ROLE pgexporter;
SELECT pgexporter_version_ext();
SELECT * FROM pgexporter_get_functions();
SELECT * FROM pgexporter_os_info();
SELECT * FROM pgexporter_load_avg();
RESET ROLE;
```

SQL 脚本撤销公众执行权限，并将执行权限授予 `pg_monitor`。这个预定义角色也提供广泛的 PostgreSQL 监控权限。发行包包含基础安装脚本和升级链；使用 `CREATE EXTENSION` 安装默认版本时，PostgreSQL 可以自动使用这条升级链。

### 指标函数

- `pgexporter_information_ext()` 和 `pgexporter_version_ext()` 返回扩展信息与版本文本。
- `pgexporter_get_functions()` 列出指标名称、是否接受参数、说明与类型；`pgexporter_is_supported(text)` 检查指标名称。
- `pgexporter_os_info()`、`pgexporter_cpu_info()`、`pgexporter_memory_info()`、`pgexporter_network_info()` 和 `pgexporter_load_avg()` 返回主机指标。
- `pgexporter_used_space(text)`、`pgexporter_free_space(text)` 和 `pgexporter_total_space(text)` 返回文件系统路径对应的字节数。

```sql
SELECT pgexporter_is_supported('pgexporter_load_avg');
SELECT pgexporter_free_space('/var/lib/postgresql');
```

应选择服务器上实际存在且操作系统账户可访问的路径。这些版本不提供后续版本中的 FIPS 或日志计数 API。

### 运行边界

文件系统调用使用服务器进程所见的文件系统，容器中看到的可能并非宿主机。已用空间函数会遍历目录，大目录树可能导致较高的采集开销。`pg_monitor` 成员可以检查可访问的路径，因此应限制数据库访问并谨慎选择采集路径。

上游列出的支持范围为 Linux 与 PostgreSQL 13+；本目录提供 PostgreSQL 14–18 软件包。安装时应匹配 PostgreSQL 主版本。RPM 主要由 PGDG 提供 0.2.4（EL8 PostgreSQL 14–16 使用 0.2.3），DEB 则由 Pigsty 提供 0.2.5；按照跨平台安装说明操作时，也需要启用 Pigsty 仓库。
