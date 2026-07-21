## 用法

来源：

- [pgmonitor-extension v2.2.0 README](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/README.md)
- [pgmonitor v2.2.0 control file](https://github.com/CrunchyData/pgmonitor-extension/blob/v2.2.0/pgmonitor.control)
- [pgmonitor-extension v2.2.0 release notes](https://github.com/CrunchyData/pgmonitor-extension/releases/tag/v2.2.0)

pgmonitor 通过一组精心挑选的视图、物化视图和表向外部收集器暴露 PostgreSQL 监控指标。其 SQL 指标无需后台工作进程即可使用；可选的 pgmonitor_bgw 后台工作进程会定期刷新物化数据。

### 安装扩展

创建一个专用模式并在此安装 pgmonitor：

    CREATE SCHEMA pgmonitor_ext;
    CREATE EXTENSION pgmonitor SCHEMA pgmonitor_ext;

仅授予收集器所需的访问权限，以访问指标对象。某些底层的 PostgreSQL 统计信息仍受内置角色和行可见性规则的约束。

### 收集指标

外部代理可以选择由扩展配置表描述的活动对象：

    SELECT *
    FROM pgmonitor_ext.metric_views
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_matviews
    WHERE active;

    SELECT *
    FROM pgmonitor_ext.metric_tables
    WHERE active;

这些表描述了指标名称、激活状态、作用范围和刷新间隔。查询安装的定义，而不是假设每个 PostgreSQL 版本都启用了所有指标。

指标表面包括活动情况、数据库和表统计信息、锁、复制、WAL 和归档状态、真空进度、配置设置、检查点以及当其依赖项可用时的扩展特定视图。

### 手动刷新物化指标

无需后台工作进程，可以调用为指定模式和指标配置的刷新过程：

    CALL pgmonitor_ext.refresh_metrics(
      'pgmonitor_ext',
      'pg_stat_statements'
    );

使用 metric_matviews 返回的名称；不要假设示例指标已安装或处于活动状态。该扩展保留了一个兼容性旧版刷新函数，但新的集成应使用文档中的过程。

### 可选后台工作进程

要在 PostgreSQL 内部安排刷新：

    shared_preload_libraries = 'pgmonitor_bgw'
    pgmonitor_bgw.dbname = 'postgres,app'
    pgmonitor_bgw.role = 'postgres'
    pgmonitor_bgw.interval = 30

更改 shared_preload_libraries 后重启 PostgreSQL。pgmonitor_bgw.dbname 是必需的，列出了要维护的数据库。上游 v2.2 当前要求工作进程角色为超级用户；使用最窄控制的角色，并保护其凭据和设置。

### 对象索引

- metric_views: 直接查询的指标视图及其收集元数据。
- metric_matviews: 物化指标及刷新间隔。
- metric_tables: 表支持的指标及维护元数据。
- refresh_metrics(schema, name): 为一个配置的指标调用刷新过程。
- pgmonitor_bgw.dbname: 由可选工作进程处理的数据库。
- pgmonitor_bgw.role: 用于刷新工作的角色。
- pgmonitor_bgw.interval: 工作循环间隔（秒）。

### 版本 2.2 及注意事项

版本 2.2 移除了设置校验指标，修复了 PostgreSQL 13 上的旧版刷新路径，并减少了常规日志噪音。

- 指标查询会增加共享统计信息、目录和扩展对象的负载。根据测量的成本设置收集间隔。
- 健康的收集器连接并不能证明物化视图是新鲜的；监控它们的时间戳和工作进程日志。
- 该扩展提供数据库指标，而不是主机、文件系统或进程指标。
- 安装 pgmonitor 不会自动配置 Prometheus、导出程序、仪表板或警报规则。
