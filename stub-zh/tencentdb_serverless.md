## 用法

来源：

- [腾讯云官方数据库资源隔离指南](https://cloud.tencent.com/document/product/409/105554)
- [腾讯云 PostgreSQL 官方插件版本矩阵](https://cloud.tencent.com/document/product/409/75121)

`tencentdb_serverless` 是腾讯云 PostgreSQL 的服务商扩展，用于在多租户实例中按数据库设置 CPU 资源限制。它由腾讯云支持团队作为托管服务的一部分安装与配置，不是可下载、可移植的 PostgreSQL 扩展。

### 启用与核心流程

官方功能指南要求提交工单，以启用数据库资源隔离、安装 `tencentdb_serverless` 并设置服务商管理参数。指南记录了符合要求的 PostgreSQL 14、15 与 16 内核版本；其他版本应查询当前服务矩阵。该功能目前适用于主实例。

腾讯云确认启用后，检查托管边界并为每个非模板数据库分配限制：

```sql
SHOW tencentdb_serverless.min_cpu_cores;
SHOW tencentdb_serverless.max_cpu_cores;

SELECT tencentdb_serverless.set_database_cpu_limit(
    'tenant_001', 2.0, 2.5
);

SELECT *
FROM tencentdb_serverless.resource_limit_view;
```

`set_database_cpu_limit(database_name, min_cpu_cores, max_cpu_cores)` 设置相对保证的最小值与绝对最大值。最大值为 `-1` 表示不设置上限。由服务商管理的 GUC 描述实例级允许范围，只应查看，不应由用户直接修改。

可使用以下函数重置单个数据库或整个实例的配置：

```sql
SELECT tencentdb_serverless.reset_database_limit('tenant_001');
SELECT tencentdb_serverless.reset_all_database_limit();
```

### 运维说明

限制实时生效，但指南要求配置每一个数据库，隔离才能按预期工作。资源不足时，最小 CPU 值是分配权重，并非永久独占核心。`resource_limit_view` 包含内存列，指南将其标为预留且当前未使用。

删除数据库会自动删除其配置。关闭功能同样需要提交腾讯云工单，由服务商重置参数、清除限制并删除扩展。应在资源竞争下验证租户 SLO，并遵循服务商当前的内核、权限、故障切换、计费与支持规则。
