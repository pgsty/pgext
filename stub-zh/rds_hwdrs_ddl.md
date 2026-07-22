## 用法

来源：

- [华为云 RDS for PostgreSQL 文档](https://support.huaweicloud.com/intl/en-us/usermanual-rds-pg/rds_09_0067.html)

`rds_hwdrs_ddl` 是华为云 RDS for PostgreSQL 在增量 DDL 同步期间使用的服务商扩展。当数据库用户无法直接创建所需对象时，它会创建华为同步服务需要的表、函数、事件触发器和授权。它属于托管服务能力，并非可移植到自建 PostgreSQL 的开源扩展。

### 可用性与启用

首先检查目标 RDS 实例的实时扩展清单；支持版本取决于服务实例，文档没有给出通用 PostgreSQL 兼容矩阵：

```sql
SELECT *
FROM pg_available_extension_versions
WHERE name = 'rds_hwdrs_ddl';
```

华为文档中的 SQL 安装路径使用服务商函数，而不是直接执行 `CREATE EXTENSION`：

```sql
SELECT control_extension('create', 'rds_hwdrs_ddl');
```

启用后，在开始同步任务前确认服务商创建的对象：

```sql
SELECT relname, relowner::regrole, relacl
FROM pg_class
WHERE relname = 'hwdrs_ddl_info';

SELECT proname, proowner::regrole
FROM pg_proc
WHERE proname = 'hwdrs_ddl_function';

SELECT evtname, evtevent
FROM pg_event_trigger
WHERE evtname = 'hwdrs_ddl_event';
```

### 同步生命周期

文档所述流程明确是临时性的：

1. 按华为同步流程要求，在源端或目标 RDS for PostgreSQL 实例启用 `rds_hwdrs_ddl`；
2. 创建并运行 PostgreSQL 到 RDS 的同步任务；
3. 同步完成后，通过服务商支持的控制台或 SQL 管理路径移除扩展，从而删除其跟踪表、函数与事件触发器。

具体拓扑与移除命令应遵循当前控制台说明。服务商页面显示的移除标识符与该页其他位置使用的扩展名称不同，因此应针对实例的可用扩展核实命令，或使用 RDS 控制台，不要盲目复制。

### 运维边界

`hwdrs_ddl_event` 事件触发器在 `ddl_command_end` 触发，`hwdrs_ddl_info` 保存同步所用信息。因此安装扩展会改变集群侧 DDL 行为，并授予服务所需访问权限。只应在批准的迁移期间启用，监控同步任务，并在结束后确认移除。

不要用手工创建的同名表、函数或触发器替代它：对象定义和权限由华为云控制，并可能随服务变化。该文档页面更新于 2025-11-10；实际生产任务开始前，应重新检查实时服务商文档与实例清单。
