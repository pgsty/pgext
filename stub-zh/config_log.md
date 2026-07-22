## 用法

来源：

- [官方 README](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/README.md)
- [官方扩展 SQL](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log--0.1.7.sql)
- [官方 C 实现](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log.c)
- [官方扩展控制文件](https://github.com/ibarwick/config_log/blob/db66842eec9398701bd045fad7e6a09f6b77eeb5/config_log.control)

`config_log` 是一个实验性后台工作进程，在服务器启动和配置重载后记录 PostgreSQL 配置文件设置。版本 `0.1.7` 每个集群只跟踪一个数据库；它不会持续监视任意文件。

### 核心流程

启动工作进程前先选择数据库和模式。默认值为 `postgres` 与 `public`：

```conf
config_log.database = 'postgres'
config_log.schema = 'public'
shared_preload_libraries = 'config_log'
```

重启 PostgreSQL，连接到该数据库，并在配置的模式中安装扩展：

```sql
CREATE EXTENSION config_log;
SELECT * FROM pg_settings_log_current ORDER BY name;
```

超级用户的搜索路径必须能够访问该模式。随后调用 `pg_reload_conf()` 会向工作进程发送重载信号；工作进程会调用 `pg_settings_logger()` 并追加配置差异。

### 对象与注意事项

`pg_settings_log` 按时间记录 INSERT、UPDATE 和 DELETE 历史。`pg_settings_log_current` 展示最新状态，`pg_settings_logger()` 执行差异比较。初始快照仅包含 `pg_settings` 中来源为 `configuration file` 的行。

扩展会撤销公众对其表、视图和记录函数的访问，但记录行仍包含 `sourcefile` 与 `sourceline`。这些路径和设置可能泄露部署细节，应谨慎授权。它只监控配置的数据库；更改 `config_log.database` 或 `config_log.schema` 时，必须协调扩展对象、预加载配置与服务器重启。
