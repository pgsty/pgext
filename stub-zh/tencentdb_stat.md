## 用法

来源：

- [腾讯云 PostgreSQL 扩展支持矩阵](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_stat` 是 TencentDB for PostgreSQL 云厂商扩展。官方支持矩阵在 PostgreSQL 10 至 14 中列出 1.0 版，而 PostgreSQL 15 至 18 未列出该扩展。它不是可移植的社区软件包，公开的云厂商页面也没有记录 SQL 对象或函数签名。

### 启用与验证

在支持矩阵列出该扩展的 TencentDB 实例和引擎版本上，使用服务提供的扩展管理权限创建扩展，并确认安装版本：

```sql
CREATE EXTENSION tencentdb_stat;

SELECT extversion
FROM pg_extension
WHERE extname = 'tencentdb_stat';
```

如果创建被拒绝或找不到扩展，应在控制台核对准确的 TencentDB PostgreSQL 大版本与内核小版本，并咨询当前云厂商支持渠道。不要将其他服务或自托管 PostgreSQL 主机上的二进制文件复制到 TencentDB。

### 云厂商边界

可用性、权限、升级和移除均由腾讯云控制。升级引擎大版本前应重新核对官方矩阵：在 PostgreSQL 10 至 14 上使用 `tencentdb_stat` 的数据库，升级到 PostgreSQL 15 或更高版本后可能无法获得同一扩展。

公开官方页面只确认名称、版本和支持的引擎大版本，因此不要根据目录描述虚构 QPS 视图、函数、GUC 或权限。创建后，应在非生产实例上检查实际安装对象，并在授权访问或接入监控前取得适用的腾讯云文档或支持说明。应按照云厂商的真实服务行为测试升级、故障转移、只读节点、性能开销和扩展移除。
