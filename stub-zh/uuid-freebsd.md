## 用法

来源：

- [官方 README](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/README.md)
- [2.0 扩展 SQL](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd--2.0.sql)
- [FreeBSD UUID 实现](https://github.com/rhodiumtoad/uuid-freebsd/blob/fb25ce7e66f28670772cf3b97b6a54661ceae5ed/uuid-freebsd.c)

`uuid-freebsd` 2.0 为 FreeBSD 上的 PostgreSQL 提供 UUID 生成函数。它支持随机、基于时间和基于命名空间的 UUID，以及标准命名空间常量；这是针对 FreeBSD 设施构建的平台专用替代方案。

### 核心流程

```sql
CREATE EXTENSION "uuid-freebsd";

SELECT uuid_generate_v4();
SELECT uuid_generate_v5(uuid_ns_url(), 'https://example.com/orders/42');
```

版本 4 每次生成新的随机 UUID。版本 3 与 5 以确定性方式组合命名空间 UUID 和名称，适合让同一个逻辑名称始终映射到同一个标识符。

### 函数

- `uuid_generate_v1()`：使用主机节点标识创建基于时间的 UUID。
- `uuid_generate_v1mc()`：使用随机组播节点标识创建基于时间的 UUID。
- `uuid_generate_v3(namespace, name)`：创建 MD5 命名空间 UUID。
- `uuid_generate_v4()`：创建随机 UUID。
- `uuid_generate_v5(namespace, name)`：创建 SHA-1 命名空间 UUID。
- `uuid_nil()`、`uuid_ns_dns()`、`uuid_ns_url()`、`uuid_ns_oid()` 与 `uuid_ns_x500()`：返回标准常量。

该扩展可重定位，不需要预加载或重启。基于时间的版本 1 值可能暴露稳定的主机与时间特征；如果不需要确定性或时间标识，应优先使用版本 4；如果需要版本 1 的顺序特性但不希望使用真实节点标识，可使用 `uuid_generate_v1mc()`。项目已无人维护，因此必须针对实际 FreeBSD 与 PostgreSQL 版本验证软件包。
