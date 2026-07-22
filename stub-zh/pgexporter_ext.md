## 用法

来源：

- [官方 README](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/README.md)
- [官方入门指南](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/doc/GETTING_STARTED.md)
- [扩展控制文件与 SQL 脚本](https://github.com/pgexporter/pgexporter_ext/tree/c73074eb5b7aa2768e1a41045905c01ebcdab197/sql)
- [指标实现](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/src/pgexporter_ext/lib.c)
- [文件系统与日志实现](https://github.com/pgexporter/pgexporter_ext/blob/c73074eb5b7aa2768e1a41045905c01ebcdab197/src/pgexporter_ext/utils.c)

`pgexporter_ext` 通过 SQL 暴露 Linux 主机、文件系统、PostgreSQL 日志与 FIPS 信息，供 pgexporter 采集。其函数在 PostgreSQL 内部以服务器操作系统账户的访问权运行，因此应把获授权的监控角色视为主机观察权限，而不是普通数据库只读权限。

### 配置与核心流程

遵循上游支持的配置：预加载模块、重启 PostgreSQL、在 `postgres` 数据库安装，并把内置监控角色授予 exporter 登录角色。

```conf
shared_preload_libraries = 'pgexporter_ext'
```

```sql
CREATE EXTENSION pgexporter_ext;
GRANT pg_monitor TO pgexporter;

SET ROLE pgexporter;
SELECT pgexporter_ext_version();
SELECT * FROM pgexporter_ext_get_functions();
SELECT * FROM pgexporter_ext_os_info();
SELECT * FROM pgexporter_ext_load_avg();
```

安装脚本会撤销公众执行权限，并把函数授予 `pg_monitor`。授予该预定义角色前应评估影响；它本身已包含广泛的 PostgreSQL 监控能力。

### 指标族

- `pgexporter_ext_version()`、`pgexporter_ext_is_supported(text)` 和 `pgexporter_ext_get_functions()` 提供发现与能力元数据。
- `pgexporter_ext_os_info()`、`pgexporter_ext_cpu_info()`、`pgexporter_ext_memory_info()`、`pgexporter_ext_network_info()` 和 `pgexporter_ext_load_avg()` 暴露主机信息。
- `pgexporter_ext_used_space(text)`、`pgexporter_ext_free_space(text)` 和 `pgexporter_ext_total_space(text)` 检查调用者提供的文件系统路径。
- `pgexporter_ext_fips()` 报告链接的 OpenSSL 上下文是否处于 FIPS 模式。
- `pgexporter_ext_log_*()` 函数族在配置的日志目录中统计从 `DEBUG5` 到 `PANIC` 的 PostgreSQL 严重级别字符串出现次数。

### 成本与安全边界

路径函数会以 PostgreSQL 操作系统用户身份遍历目录或调用文件系统统计 API。`pg_monitor` 成员可以探测该账户可访问的任意路径，并推断其存在性或大小。应相应限制数据库登录和网络访问。

每个日志计数函数都会同步扫描 `log_directory` 中的所有普通文件，包括 gzip、bzip2、LZ4 与 Zstandard 文件，并统计严重级别子串，而不是解析结构化 PostgreSQL 日志记录。频繁抓取或大量保留日志可能造成严重的 CPU、I/O、解压与查询延迟开销。当前源码定义了 `pgexporter.log_cache_refresh_interval`，但日志函数仍直接调用完整扫描；不要假定该 GUC 提供了有效缓存。

### 版本与打包注意事项

核验的控制文件声明版本 `0.3.1`，但同一源码树只有基础脚本 `pgexporter_ext--0.1.0.sql` 和逐级升级脚本，没有基础脚本 `pgexporter_ext--0.3.1.sql`。因此，从未经修改的当前源码安装后直接执行 `CREATE EXTENSION pgexporter_ext` 可能失败。上线前应确认发行包提供有效基础脚本，或者在一次性数据库中安装版本 `0.1.0` 并测试完整的 `ALTER EXTENSION ... UPDATE` 链。

上游报告支持 Linux 上的 PostgreSQL 13+；当前构建声明需要 C17、OpenSSL 与压缩库。应使用针对实际 PostgreSQL 主版本和操作系统构建的软件包，并在真实主机上验证函数可用性与采集成本。
