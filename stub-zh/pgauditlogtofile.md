## 用法

来源：

- [pgauditlogtofile v1.8.5 README](https://github.com/fmbiete/pgauditlogtofile/blob/v1.8.5/README.md)
- [Changes from v1.8.4 to v1.8.5](https://github.com/fmbiete/pgauditlogtofile/compare/v1.8.4...v1.8.5)

pgauditlogtofile 是 pgAudit 的一个扩展，用于将 pgAudit 记录路由到专用的 CSV 或 JSON 文件。使用它可以在保留审计记录和访问控制的同时，与普通 PostgreSQL 服务器日志分开。

### 预加载并创建扩展

首先加载 pgAudit，然后是 pgauditlogtofile：

    shared_preload_libraries = 'pgaudit,pgauditlogtofile'

重启 PostgreSQL，然后在 postgres 数据库中安装这两个扩展：

    CREATE EXTENSION pgaudit;
    CREATE EXTENSION pgauditlogtofile;

上游建议仅在 postgres 数据库中创建 pgauditlogtofile，而不是独立地在每个应用程序数据库中创建。

### 配置审计文件

    pgaudit.log_directory = 'log'
    pgaudit.log_filename = 'audit-%Y%m%d_%H%M.log'
    pgaudit.log_format = 'csv'
    pgaudit.log_rotation_age = 1440
    pgaudit.log_file_mode = 0600

如果未设置 pgaudit.log_directory 或 pgaudit.log_filename，则会禁用单独的目标，记录将回退到普通服务器日志。相对路径将在 PostgreSQL 数据目录下解析。

### 压缩

版本 1.8 支持压缩审计文件：

    pgaudit.log_compression = 'zstd'
    pgaudit.log_compression_level = 6

pgaudit.log_compression 接受 off、gzip、lz4 或 zstd，具体取决于相应的支持。级别范围是 0 到 22，但有效且有用的级别取决于所选的算法。压缩会消耗后端 CPU，请测试日志吞吐量和旋转延迟。

### 参数索引

- pgaudit.log_format: csv 或 json 输出。
- pgaudit.log_directory 和 pgaudit.log_filename: 目标及 strftime 样式的文件名。
- pgaudit.log_file_mode: 新创建文件的权限。
- pgaudit.log_rotation_age: 基于时间的旋转间隔（分钟）。
- pgaudit.log_compression 和 pgaudit.log_compression_level: 压缩方法和努力程度。
- pgaudit.log_connections 和 pgaudit.log_disconnections: 包含连接生命周期事件，前提是 PostgreSQL 的相应日志设置已启用。
- pgaudit.log_execution_time 和 pgaudit.log_execution_memory: 添加执行测量；这些需要重启。
- pgaudit.log_autoclose_minutes: 基于不活动的文件处理程序关闭（实验性）。

### 旋转和操作

PostgreSQL 配置重新加载会旋转审计文件。该扩展的后台工作进程可以向后端发送信号以关闭审计文件句柄；pg_rotate_logfile() 不会旋转独立的审计文件。

版本 1.8.5 改进了后台工作进程的信号、钩子恢复和 PostgreSQL 19 版本兼容性。它没有引入从 1.8.4 到 1.8.5 的必要配置迁移。

### 注意事项

- 文件分离不是保留管理。请显式地发送、旋转、保护和过期审计文件。
- 确保 PostgreSQL 操作系统帐户可以创建目标，并且文件权限符合审计策略。
- 突然的后端或主机故障可能会使最后一个压缩文件不完整；验证摄取行为。
- 启用时间、内存、连接或详细 pgAudit 类别会显著增加开销和日志量。
