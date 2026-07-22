## 用法

来源：

- [TencentDB for PostgreSQL 内核版本更新动态](https://cloud.tencent.com/document/product/409/67017)
- [TencentDB 支持插件版本概览](https://cloud.tencent.com/document/product/409/75121)
- [TencentDB for PostgreSQL 产品页](https://cloud.tencent.com/product/postgres)

`rds_server_handler` 1.0 是 TencentDB for PostgreSQL 的服务器维护扩展。腾讯公开的内核更新动态将它描述为可通过命令清空全部服务器缓存，并在较新内核中支持进程级 syscache/relcache 淘汰。它是云厂商内部运维工具，不是可移植的社区 PostgreSQL 扩展。

### 核心流程

先确认目标 TencentDB 内核列出了 1.0 版，再用高权限账号创建并核验扩展：

```sql
CREATE EXTENSION rds_server_handler;

SELECT extname, extversion
FROM pg_extension
WHERE extname = 'rds_server_handler';
```

腾讯公开文档没有给出缓存淘汰 SQL 命令或函数签名。只能在目标托管实例上查看实际安装对象，并从腾讯云支持或该内核对应的控制台文档获取调用方式：

```sql
SELECT n.nspname AS schema_name,
       p.proname AS function_name,
       pg_get_function_identity_arguments(p.oid) AS arguments
FROM pg_proc AS p
JOIN pg_namespace AS n ON n.oid = p.pronamespace
JOIN pg_depend AS d ON d.objid = p.oid AND d.deptype = 'e'
JOIN pg_extension AS e ON e.oid = d.refobjid
WHERE e.extname = 'rds_server_handler';
```

该清单查询只用于诊断，并不意味着可以调用未公开的例程。

### 运维边界

缓存淘汰会突然损失热缓存性能，并可能影响其他会话。它应作为协调后的故障处置或维护操作，使用厂商支持的最小作用域，并在操作前后测量实例状态。不要猜测函数名，也不要把它作为常规应用查询执行。

扩展的可用性、所需权限、具体对象接口与行为均绑定 TencentDB 内核版本。腾讯扩展矩阵在受支持规格中列出 1.0 版，内核更新动态则记录了缓存淘汰能力和安全修复；应保持托管内核更新，并在升级后重新核对厂商矩阵。
