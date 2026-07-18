## 用法

来源：

- [项目 README](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/README.md)
- [扩展 control 文件](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/yeti.control)
- [1.14.1 版升级 SQL](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/yeti--1.14.0--1.14.1.sql)
- [共享内存限速器](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/src/f_tbf_rate_check.c)
- [LNP 解析函数](https://github.com/yeti-switch/yeti-pg-ext/blob/a069acd6e32f26f32061b6a0852a3b83f3d9b7a8/src/f_lnp_resolve.c)

`yeti` 1.14.1 为 YETI softswitch 提供专用原生辅助函数。其表面包括 LNP endpoint 与解析调用、tag 操作、随机正则替换、模板处理、DNS SRV 排序和共享内存 token-bucket 限速检查。它不是通用电信抽象。

### 安装并启用限速器

只有 `tbf_rate_check()` 需要服务器预加载。依赖它之前应配置库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'yeti_pg_ext'
```

```sql
CREATE SCHEMA yeti_ext;
CREATE EXTENSION yeti WITH SCHEMA yeti_ext;

SELECT yeti_ext.tbf_rate_check(1, 1001, 5.0);
```

若未预加载该库，`tbf_rate_check()` 会发出警告并返回 `true`。这属于 fail-open 行为：配置缺失会禁用执行，而不是拒绝流量。服务请求前，应加入验证预加载状态的启动探针。

### 易失状态与网络边界

token bucket 只存在于单个 PostgreSQL 实例的共享内存中，使用固定 5,000 项哈希，并会在重启或故障转移时丢失。它不会在主库、备库或多个集群间协调。必须压力测试 hook 链接、争用、时钟行为、rate 语义和容量耗尽；不能把它作为唯一计费、滥用或安全限制。

LNP 函数把与电话号码相关的请求发送到配置的外部 endpoint，并可能在传输失败时返回 null 或 error 字段。应设置有限超时，限制 endpoint 配置，保护 SQL 和日志中的电话号码数据，并在扩展外定义重试与 circuit-breaker 行为。应一起固定 YETI 应用和扩展版本，并遵循完整 SQL 升级链；辅助函数签名在 1.14.1 中发生了变化。
