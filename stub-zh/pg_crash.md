

## 用法

> [pg_crash: 向随机进程发送随机信号](https://github.com/cybertec-postgresql/pg_crash)

pg_crash 是一个混沌工程扩展，定期向 PostgreSQL 后端进程发送 kill 信号，适用于高可用和故障转移测试。必须添加到 `shared_preload_libraries` 中。

### 配置

添加到 `postgresql.conf`：

```
shared_preload_libraries = 'pg_crash'

# 要发送的 POSIX 信号（空格分隔）
crash.signals = '1 2 3'

# 发送信号之间的延迟秒数
crash.delay = 30
```

### 信号参考

常用 POSIX 信号：`1`（SIGHUP）、`2`（SIGINT）、`3`（SIGQUIT）、`9`（SIGKILL）、`15`（SIGTERM）。

配置后重启服务器。后台工作进程将按照指定的间隔定期向随机后端进程发送配置的信号。
