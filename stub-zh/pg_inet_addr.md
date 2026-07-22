## 用法

来源：

- [官方 README](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/README.md)
- [1.0 版扩展 SQL](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/pg_inet_addr--1.0.sql)
- [C 实现](https://github.com/eulerto/pg_inet_addr/blob/ac8bee911bd3ac3055f13a871bb2af34df254e29/pg_inet_addr.c)

`pg_inet_addr` 列出 PostgreSQL 后端可见的操作系统接口所分配的 IPv4 与 IPv6 地址。对于 PostgreSQL 的 `inet_server_addr()` 可能返回 NULL 的本地或 Unix socket 连接，它尤其有用。

### 核心流程

扩展把函数安装到 `pg_catalog`，并撤销 PUBLIC 的 EXECUTE 权限。超级用户或所有者必须显式授权。

```sql
CREATE EXTENSION pg_inet_addr;

GRANT EXECUTE ON FUNCTION pg_catalog.pg_inet_addr()
TO monitoring_role;
```

可以查询全部接口，也可以按地址族和接口名过滤。

```sql
SELECT interface, address, netmask, family
FROM pg_catalog.pg_inet_addr()
ORDER BY interface, family, address;

SELECT interface, address::inet
FROM pg_catalog.pg_inet_addr()
WHERE family = 'IPv4'
  AND interface <> 'lo';
```

### 返回列

- `interface`：操作系统接口名。
- `address`：以文本格式化的接口地址。
- `netmask`：以文本格式化的接口掩码。
- `family`：`IPv4` 或 `IPv6`。

该函数声明为 VOLATILE，每次调用都返回新快照。地址与掩码都是文本，因此应先过滤或校验再转换。链路本地 IPv6 地址可能包含 `%eth0` 之类的作用域后缀。

### 运维说明

结果描述的是 PostgreSQL 服务器进程所在的网络命名空间，而不是客户端网络、DNS 名、负载均衡地址或公网端点。在容器中，它通常代表容器或 Pod 的命名空间。两次调用之间接口可能增减，函数也不会判断哪个地址实际接受 PostgreSQL 流量。

实现依赖主机的 `getifaddrs` 接口，面向 Unix 类系统。由于输出会暴露主机网络细节，除非监控角色确实需要，否则应保留默认权限边界。
