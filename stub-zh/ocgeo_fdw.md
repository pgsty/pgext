## 用法

来源：

- [上游固定版本 README](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/README.md)
- [1.0 版本安装 SQL](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/sql/ocgeo_fdw--1.0.sql)
- [固定版本 FDW 实现](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/ocgeo_fdw.c)
- [固定版本 API 客户端实现](https://github.com/sgsfak/ocgeo_fdw/blob/7e9c66697b7b4fa3ec58eb47d29544f6339c15a1/ocgeo_api.c)

ocgeo_fdw 1.0 把 OpenCage 的 JSON 地理编码 API 映射为只读外部表。对输入列 q 使用等值条件可执行正向或反向地理编码，最低置信度条件也可下推为 API 参数。返回字段可包含原始 JSON、PostgreSQL point 与 box、格式化地址、置信度和地址组件。

### 服务器、用户映射与查询

```sql
CREATE EXTENSION ocgeo_fdw;

CREATE SERVER ocdata_server
FOREIGN DATA WRAPPER ocgeo_fdw
OPTIONS (uri 'https://api.opencagedata.com/geocode/v1/json');

CREATE USER MAPPING FOR current_user
SERVER ocdata_server
OPTIONS (api_key '<api-key>', max_reqs_sec '1');

CREATE FOREIGN TABLE geocode_api (
    json_response jsonb,
    location point,
    confidence integer,
    formatted text,
    q text
) SERVER ocdata_server;

SELECT confidence, location, formatted
FROM geocode_api
WHERE q = 'Eiffel Tower, France' AND confidence >= 5;
```

普通扫描如果没有可用的 q 等值限制就不会发出 API 请求；嵌套循环连接则可能对外表侧的每一行都发一次请求。

### 网络、配额与可观测性限制

每次 API 请求都在 PostgreSQL 后端内同步执行；DNS、TLS、限速等待和远端处理期间，会一直占用事务、连接以及可能持有的锁。内置限速器只限制单个后端每秒请求数，不会跨会话全局执行每日配额。连接与重扫描会放大调用量，应对稳定结果进行物化或缓存。

API 密钥保存在用户映射中。调试输出可能包含完整请求 URL；README 还明确演示了如何启用客户端调试查看它，因此在密钥可能进入客户端或服务器日志的环境中不要开启。应限制服务器使用权、设置出站白名单、采用低权限密钥和语句超时，并以服务商权威计数器监控配额。

ocgeo_stats 只报告进程内存中的本地计数，可能包含该后端切换使用过的多个角色条目，而且函数在数值会随请求变化时被错误声明为 IMMUTABLE。不要把它当作审计或计费来源。所审源码只声明支持 PostgreSQL 9.6 至 14，自 2021 年后没有更新；libcurl、规划器 API、JSON 解析、TLS 策略、服务响应字段及当前 PostgreSQL 大版本都必须重新做集成测试。
