## 用法

来源：

- [database.dev package](https://database.dev/dventimi/pg_partition_magician)
- [Version 0.6.0 release SQL](https://github.com/dventimisupabase/pg_partition_magician/releases/download/v0.6.0/pg_partition_magician--0.6.0.sql)
- [Version 0.6.0 control file](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/pgpm_core/extension.control)
- [Version 0.6.0 README](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/README.md)
- [Version 0.6.0 guide](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/docs/guide.md)
- [Version 0.6.0 API reference](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/docs/reference.md)
- [Version 0.6.0 changelog](https://github.com/dventimisupabase/pg_partition_magician/blob/v0.6.0/CHANGELOG.md)

`dventimi@pg_partition_magician` 使用 SQL 和 PL/pgSQL 管理原生 RANGE 分区，支持时间、数值、UUIDv7 以及符合编码要求的文本键。转换时会将原表保留为有明确边界的历史分区；把这部分历史数据拆成更小的分区是另一项操作。

### 启用与核心工作流

本条目对应 database.dev 包。虽然 GitHub 所有者名为 `dventimisupabase`，registry 中的注册名称仍为 `dventimi@pg_partition_magician`。其 `0.6.0` 版本的 registry 安装内容与官方发布 SQL 一致。在已配置 database.dev 安装器和 `pg_tle` 的环境中，先注册包，再使用带双引号的扩展名称启用。必须先在当前数据库中启用依赖的 `pg_cron` 扩展。

```sql
SELECT dbdev.install('dventimi@pg_partition_magician');
CREATE EXTENSION "dventimi@pg_partition_magician" VERSION '0.6.0';
SELECT pgpm.version();
```

下面假设已有 `public.events` 表，其中 `created_at` 列非空且随写入单调增长。主键或唯一约束必须包含该列；没有主键或唯一约束的表也受支持。转换必须在自动提交模式下作为顶层语句执行：`pgpm.transmute` 会在不同阶段之间提交事务，不能放入外层事务。

```sql
CALL pgpm.transmute(
  p_parent   => 'public.events',
  p_control  => 'created_at',
  p_interval => interval '1 month',
  p_obtain   => 7,
  p_retain   => NULL
);
SELECT pgpm.schedule();
SELECT * FROM pgpm.status();
SELECT pgpm.resume('public.events');
```

转换后，表的定时维护默认处于暂停状态。恢复维护前应检查分区边界。此例无限期保留历史分区。`pgpm.schedule()` 会创建两个独立调度的维护任务，默认均每分钟运行一次。

### 重要对象

| 对象 | 用途 |
| --- | --- |
| `pgpm.transmute` | 转换普通表并登记分区策略。 |
| `pgpm.obtain`, `pgpm.extend_to` | 创建未来分区；在已知键值将大幅跳跃之前主动扩展覆盖范围。 |
| `pgpm.set_regrain`, `pgpm.regrain_history` | 启用分批历史拆分，或手动驱动历史拆分。 |
| `pgpm.set_retain`, `pgpm.retain` | 设置或执行保留策略；符合条件的旧分区会被删除。 |
| `pgpm.pause`, `pgpm.resume` | 暂停或恢复某张表的定时维护。 |
| `pgpm.schedule`, `pgpm.unschedule` | 在安装调度器的数据库中管理维护任务。 |
| `pgpm.status()`, `pgpm.config`, `pgpm.log` | 检查分区覆盖范围、策略、进度与失败记录。 |
| `pgpm.version()`, `pgpm.installed` | 查看已安装的实现版本和 SQL 安装历史。 |

### 要求与运行边界

- 上游测试覆盖 PostgreSQL 15–18。对象位于固定的 `pgpm` 模式。扩展使用 SQL 和 PL/pgSQL，自身不含共享库；定时运行依赖服务器既有的 `pg_cron` 配置。
- 安装和表转换需要相应的模式、DDL 权限与表所有权。调度和维护应由有权操作目标表并使用调度器的角色执行；这些 SQL 例程不提供提权封装。
- 转换会扫描原表以验证边界，并短暂持有目录操作所需的锁。转换期间，超出这些边界的写入会失败。应为并发插入预留足够的边界余量，并检查引用目标表的外键：默认拒绝此类外键，除非选择文档规定的保留模式。
- 不设 DEFAULT 分区。过晚或超前、落在已有范围之外的键值会导致写入失败；应监控覆盖范围，并在主动跳跃键值之前调用 `pgpm.extend_to`。任意回填更早的键值不符合单调键设计。
- 历史拆分会复制行、产生 WAL，并临时占用额外磁盘空间。可以长期保留原始的大分区，但其范围内的细粒度分区裁剪和保留策略要等拆分后才能生效。保留策略会删除数据，应明确配置并做好相应备份。归档是独立的可选模块。
- 从 0.5.0 之前的版本升级时，须重新运行 `pgpm.schedule()`。未来分区的创建已移到独立任务；仅安装新 SQL 不会注册该任务。独立 SQL 安装方式支持重新执行安装文件；database.dev 方式则需要使用已注册包的安装或更新路径。

0.6.0 使用关系 OID 约束分区操作，并在安装时回填 `pgpm.part.child_oid`。归档、写阻断、保留策略及超表切换遇到名称指向错误对象时会停止。`fail_archive_identity`、`fail_write_block_identity` 和 `fail_retain_identity` 事件会计入 `status().retain_drop_failures`，且不会自动清除。应调查记录中的对象身份不匹配；从 0.5.0 升级无需额外手工迁移步骤。
