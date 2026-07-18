## 用法

来源：

- [Alpha 版本文档](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/README.md)
- [预加载与后台工作进程注册代码](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/src/extension/entrypoint.c)
- [配置示例](https://github.com/nextgres/nextgres-idcp/blob/b90bbf19a158f937951a52edb6f888f8ffa0b30e/nextgres_idcp.conf)

`nextgres_idcp` 是 alpha 阶段的数据库内连接池。它在共享预加载期间注册后台工作进程，把客户端会话代理到较小的后端连接池。已复核版本面向 Ubuntu 22.04 上的 PostgreSQL 16。

```conf
shared_preload_libraries = 'nextgres_idcp'
nextgres_idcp.thread_count = 1
nextgres_idcp.session_pool_size = 10
nextgres_idcp.port = 6543
```

启用预加载后需重启集群，再在目标数据库创建扩展对象。上游明确说明此 alpha 版本不可用于生产：查询取消、清理和配置仍不完整，也不提供分片或读写负载均衡。代理只应绑定到受控接口，应用与 PostgreSQL 同等级的认证和 TLS 控制，并在任何非实验室用途前压测故障恢复。
