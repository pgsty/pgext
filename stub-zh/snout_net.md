## 用法

来源：

- [README 0.1.0](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/README.md)
- [Control](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/snout_net.control)
- [SQL 0.1.0](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/sql/snout_net--0.1.0.sql)
- [Cargo](https://github.com/snoutdata/snout-net/blob/6c309d0075e123ef88a55c521becd308046a49b0/Cargo.toml)

`snout_net` 0.1.0 在 PostgreSQL 中排队保存 HTTP 请求，并在提交事务后由后台工作进程发送。它面向 PostgreSQL 17，运行时需要 libcurl 7.85 或更高版本。超级用户需要先预加载库、重启服务器，再在配置的数据库中安装扩展。

### 基本用法

```conf
shared_preload_libraries = 'snout_net'
snout_net.database_name = 'postgres'
```

```sql
CREATE EXTENSION snout_net;
SELECT net.check_worker_is_up();

BEGIN;
SELECT net.http_post(
  url := 'https://example.com/hook',
  body := '{"event":"signup"}'::jsonb
) AS request_id \gset
COMMIT;

SELECT status_code, content, error_msg
FROM net._http_response WHERE id = :request_id;
```

示例使用 psql 保存请求 ID。首次查询响应时可能还没有记录，需要等工作进程完成后再次查询。回滚的事务不会发送请求。不要在提交请求的同一个事务中同步等待响应。

### 对象与配置

- `net.http_get`、`net.http_post`、`net.http_delete`：返回请求 ID，接受查询参数、请求头与超时设置；POST 辅助函数接受 JSON 请求体。
- `net.http_request_queue` 和 `net._http_response`：保存待处理请求与响应。
- `net._http_collect_response`：获取响应；同步等待应放在后续事务中。
- `net.check_worker_is_up`、`net.wait_until_running`、`net.worker_restart`：工作进程健康检查与重启辅助函数。
- `snout_net.database_name` 指定唯一服务的数据库；`snout_net.username` 指定工作进程角色，默认为初始化集群的超级用户。
- `snout_net.ttl` 默认 6 小时，`snout_net.max_concurrent` 默认 200，`snout_net.max_timeout_ms` 默认 600000，`snout_net.max_response_bytes` 默认 64 MB。
- `snout_net.allowed_networks` 允许显式列出的内部 CIDR；否则会拒绝私网、回环、链路本地等非公网地址，并在 DNS 解析和重定向后再次检查。
- `snout_net.worker_type` 在服务器启动时设置，其他参数通过重新加载配置生效。升级共享库需要重启服务器。

### 投递与安全边界

请求队列和响应表均为 unlogged 表：数据库崩溃后不保证保留内容，也不会复制到物理备库。工作进程在发送途中重启可能重复发送请求，接收端需要容忍重复副作用。同一端点收到的请求顺序也不保证一致。如果触发器或约束导致响应无法写入，工作进程会记录警告，不会再次发送该请求。

安装脚本向 PUBLIC 授予 schema、表与序列的访问权限。开放 SQL 访问前，应检查这些权限、工作进程角色和出站网络策略。固定的 `net` 对象与 `pg_net` 重叠；这里不提供两者共存或自动迁移流程。
