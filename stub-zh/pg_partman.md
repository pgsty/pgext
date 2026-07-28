## 用法

来源：

- [pg_partman 5.5.0 README](https://github.com/pgpartman/pg_partman/blob/v5.5.0/README.md)
- [pg_partman 5.5.0 变更日志](https://github.com/pgpartman/pg_partman/blob/v5.5.0/CHANGELOG.md)
- [pg_partman 使用指南](https://github.com/pgpartman/pg_partman/blob/v5.5.0/doc/pg_partman_howto.md)
- [pg_partman 参考文档](https://github.com/pgpartman/pg_partman/blob/v5.5.0/doc/pg_partman.md)
- [pg_partman 5.5.0 控制文件](https://github.com/pgpartman/pg_partman/blob/v5.5.0/pg_partman.control)

`pg_partman` 按时间或整数 ID 自动管理 PostgreSQL 声明式分区集。它可以创建未来分区、执行保留策略、移动已有数据，并通过 SQL 调度或可选后台工作进程运行维护。底层表仍是普通 PostgreSQL 原生分区表。

### 核心流程

```sql
CREATE SCHEMA partman;
CREATE EXTENSION pg_partman SCHEMA partman;

CREATE TABLE public.measurements (
    id bigint GENERATED ALWAYS AS IDENTITY,
    created_at timestamptz NOT NULL,
    value numeric
) PARTITION BY RANGE (created_at);

SELECT partman.create_partition(
    p_parent_table := 'public.measurements',
    p_control := 'created_at',
    p_interval := '1 day'
);

CALL partman.run_maintenance_proc();
SELECT * FROM partman.show_partitions('public.measurements');
```

`create_partition()` 是创建受管分区集的当前名称。旧的 `create_parent()` 在 5.x 系列中仍为向后兼容保留。模板表用于携带 PostgreSQL 不会自动传播的属性；子分区已经存在后再修改模板，只会影响未来子分区，除非另外调整旧分区。

### 保留与数据移动

```sql
UPDATE partman.part_config
SET retention = '30 days',
    retention_keep_table = false
WHERE parent_table = 'public.measurements';

CALL partman.partition_data_proc('public.measurements');
CALL partman.undo_partition_proc('public.measurements');
```

当配置为删除子表时，保留策略具有破坏性。如果其他表通过外键引用该分区集，只有在确认引用行不再阻止分离或删除后，才应设置 `detach_before_drop`。使用 `retention_schema` 时，5.5 要求目标 schema 与每个被移动子表具有同一所有者。

### 后台工作进程

在服务启动前加入工作进程库：

```conf
shared_preload_libraries = 'pg_partman_bgw'
pg_partman_bgw.interval = 3600
pg_partman_bgw.dbname = 'mydb'
pg_partman_bgw.role = 'partman_maintainer'
```

修改 `shared_preload_libraries` 需要重启；其余工作进程设置可以重载。工作进程角色需要完整访问 pg_partman schema 及全部受管分区集。应使用专用非超级用户角色，并把拥有这些表的角色授予它：

```sql
CREATE ROLE partman_maintainer WITH LOGIN;
GRANT table_owner TO partman_maintainer;
```

5.5 将 `pg_partman_bgw.role` 的默认值改为 `partman_maintainer`。因此升级后，之前依赖隐式配置的工作进程会停止成功运行，直到该角色存在并获得所需权限。

### 5.5 版本升级

```sql
ALTER EXTENSION pg_partman UPDATE TO '5.5.0';
```

5.5 修复多条 SQL 注入与权限提升路径，增加用于对配置行实施 RLS 策略的 `maintenance_role` 列，并允许某一分区集失败后继续维护其他分区集。失败分区集会记录 warning，并把最后运行标记设为空，因此监控必须同时检查 PostgreSQL 日志与配置状态。

该版本还增加 `detach_before_drop`、继承列级统计目标，并改变保留 schema 的所有权规则。部分升级脚本会重建扩展函数或过程，因此扩展升级后要复查 PUBLIC 授权。

### 运维边界

- 要求 PostgreSQL 14 或更高版本；版本 5 只使用原生声明式分区。
- `pg_jobmon` 是可选依赖。安装后会增加任务监控，也会增加一层权限边界。
- 按文档配置所有者、schema、表、过程、函数、临时表及可选 RLS 权限后，无需超级用户也可以安装和运行 pg_partman。
- 日常维护应只有一个调度器负责。除非经过明确协调，不要同时使用后台工作进程与外部调度器。
- 大型维护可能获取许多锁并移动大量数据。应在代表性数据上测试保留与迁移，监控默认分区，并使用独立于分区保留策略的备份。
