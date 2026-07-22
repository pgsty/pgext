## 用法

来源：

- [Amazon RDS 传输概览与限制](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.html)
- [Amazon RDS 设置流程](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Setup.html)
- [传输流程](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Transporting.html)
- [函数参考](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.transport.import_from_server.html)
- [参数参考](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/PostgreSQL.TransportableDB.Parameters.html)

`pg_transport` 是 Amazon RDS for PostgreSQL 的扩展，可把一个数据库从源 RDS 实例物理拉取到目标 RDS 实例。它面向大型数据库迁移，相比转储恢复可减少处理量与停机时间，但不适用于自管 PostgreSQL、Amazon EC2 或非 RDS 目标。

### 核心流程

源与目标必须运行相同 PostgreSQL 版本。在两端使用自定义参数组，预加载扩展，在目标端配置足够的工作进程容量，重启后在两端安装 `pg_transport`。

```conf
shared_preload_libraries = 'pg_transport'
pg_transport.num_workers = 3
max_worker_processes = 18
```

```sql
CREATE EXTENSION pg_transport;

SELECT transport.import_from_server(
  'source.example.region.rds.amazonaws.com',
  5432,
  'transport_source',
  'source-temporary-password',
  'appdb',
  'destination-temporary-password',
  true
);
```

先在目标端调用 `transport.import_from_server` 并把 `dry_run` 设为 `true`。兼容性、空间、安全组、凭据与工作进程检查通过后，再把 `dry_run` 设为 `false` 重复调用。源用户与目标调用者都必须拥有 `rds_superuser`；应使用临时角色，或在传输完成后立即轮换传入的两个密码。

`pg_transport.num_workers` 取值 1–32，默认 3。AWS 定义目标端工作进程要求为 `(3 * pg_transport.num_workers) + 9`。`pg_transport.work_mem` 限制每个工作进程的内存，`pg_transport.timing` 控制进度计时消息。

### 迁移边界

开始传输会终止源数据库会话，并把该数据库置于特殊只读状态。传输过程中目标数据库不可访问。操作不具备事务性，也不会通过普通 WAL 记录以供时间点恢复；在传输及后续备份完成前，目标实例不可执行 PITR。失败后的清理可能让源数据库继续只读，需要显式恢复。

传输期间源数据库只能安装 `pg_transport`。源对象必须位于 `pg_default`，不支持 `reg` 类型，目标端不能已有同名数据库，而且只读副本及其父实例不能参与。角色、所有权与 ACL 不会迁移；目标对象由目标端本地用户拥有。必须单独盘点并重建安全配置，验证备份与应用切换，并遵循最新 AWS 文档，因为可用性取决于托管 RDS 引擎版本。
