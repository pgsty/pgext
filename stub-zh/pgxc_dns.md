## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/README)
- [1.0 版本 SQL 对象](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns--1.0.sql)
- [C 实现](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns.c)
- [扩展 control 文件](https://github.com/storm-db/pgxc_dns/blob/441b94223abd693620b5125ae473dfd40a36e790/pgxc_dns.control)

`pgxc_dns` 是历史 Postgres-XC 1.0 代码，根据协调器相对使用率生成 BIND SDB 区域记录。它原本通过 `pgxc_dns_zone` 视图暴露协调器 A 记录，以实现基于 DNS 的负载分配。

```sql
CREATE EXTENSION pgxc_dns;
SELECT name, ttl, rdtype, rdata
FROM pgxc_dns_zone
ORDER BY name, rdtype, rdata;
```

配置项包括 `pgxc_dns.zone`、`pgxc_dns.name`、`pgxc_dns.host`、`pgxc_dns.zone_ttl` 和 `pgxc_dns.ttl`。每个协调器需报告自身地址，同时共享区域设置。

这不是面向当前 PostgreSQL、Postgres-XL 或现代分布式 PostgreSQL 产品的扩展。上游只在 CentOS 6.3 与基于 PostgreSQL 9.2 的 Postgres-XC 上测试，文档中的 BIND 配置还嵌入数据库凭据。只能把它视为归档源码；实际部署应改用受支持的服务发现、健康检查与带认证 DNS 集成。
