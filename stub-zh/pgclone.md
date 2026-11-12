## 用法

来源：

- [pgclone v4.5.0 README](https://github.com/valehdba/pgclone/blob/v4.5.0/README.md)
- [pgclone v4.5.0 使用指南](https://github.com/valehdba/pgclone/blob/v4.5.0/docs/USAGE.md)
- [异步克隆指南](https://github.com/valehdba/pgclone/blob/v4.5.0/docs/ASYNC.md)
- [pgclone v4.5.0 发行说明](https://github.com/valehdba/pgclone/releases/tag/v4.5.0)

pgclone 通过 PostgreSQL 连接克隆表、模式、函数、角色或整个数据库。它还提供预检、结构差异、屏蔽、一致快照以及可选的后台作业。使用它进行受控数据库复制，而不是将其用作备份和恢复的无人值守替代品。

### 创建并运行一个克隆

    CREATE EXTENSION pgclone;
    SELECT pgclone.version();

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'customers',
      true
    );

模式和数据库入口点遵循相同的连接优先模式：

    SELECT pgclone.schema(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true
    );

    SELECT pgclone.database(
      'host=source.example dbname=app user=clone_user',
      true
    );

主要 API 包括 pgclone.table、pgclone.schema、pgclone.functions、pgclone.database 和 pgclone.database_create。_ex 变体暴露了对索引、约束和触发器的显式选择。

### 过滤和屏蔽数据

JSON 选项可以限制列和行：

    SELECT pgclone.table(
      'host=source.example dbname=app user=clone_user',
      'public',
      'users',
      true,
      'users_lite',
      '{"columns":["id","name","email"],"where":"active"}'
    );

4.4 版本增加了模式级和数据库级屏蔽、表包含模式以及 exclude_tables。屏蔽表达式在源端的 COPY 查询中运行，因此成功屏蔽的数据不会以未屏蔽的形式到达目标。

4.4.2 版本的屏蔽验证器会跳过不安全或不兼容的屏蔽：无法转换为列的常量值、NOT NULL 列中的 NULL 值、唯一或主键列上的非哈希屏蔽，以及外键列上的屏蔽。被跳过的屏蔽会使该列保持未屏蔽状态。将警告视为隐私门失败，并在分发克隆之前检查结果。

### 冲突处理与远端脱敏报告

从 4.4.3 起，同步表、模式和数据库克隆会执行显式提供的 `conflict` 选项。`replace` 先级联删除已有对象再重建；`skip`、`error` 和 `rename` 分别选择文档规定的其他处理方式。过去被忽略的选项现在可能触发破坏性操作，升级前应检查调用参数。

4.4.4 增加了银行账号与 IBAN 识别，以及远端报告重载：

```sql
SELECT * FROM pgclone.masking_report(
    'host=source.example dbname=app user=clone_user', 'public');
```

### 4.5.0 的 Oracle 数据源

`pgclone.table_from_oracle()` 需要启用了可选 Oracle OCI 支持的动态库，数据库主机也必须能加载 Oracle 客户端。普通构建虽然提供该函数，但会拒绝执行此操作。独立的 `pgclone.oracle_type_to_pg()` 类型映射函数不需要 Oracle 客户端。

```sql
SELECT pgclone.table_from_oracle(
    'dsn=ORACLE_SOURCE user=clone_user',
    'HR', 'EMPLOYEES', dry_run => true);
```

省略 Oracle 密码时，扩展从服务器进程的 `PGCLONE_ORACLE_PASSWORD` 环境变量读取密码。试运行只返回 DDL，不创建目标。实际克隆要求目标模式已存在、目标表不存在，且不会替换该表。DDL 与复制的数据在同一个 PostgreSQL 事务中完成，Oracle 端使用只读快照。

不支持的 Oracle 列类型会在发出警告后跳过，因此必须核对结果列集合。默认值、外键与检查约束、非唯一索引、触发器、分区、标识列属性以及脱敏和过滤逻辑均不复制。替换动态库后，还需升级扩展以安装新增 SQL 函数：

```sql
ALTER EXTENSION pgclone UPDATE TO '4.5.0';
```

### 预检、差异和一致性

    SELECT pgclone.preflight(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

    SELECT pgclone.diff(
      'host=source.example dbname=app user=clone_user',
      'public'
    )::jsonb;

预检检查连接性、版本、权限、容量、名称、角色、扩展和表空间。差异报告 DDL 差异而不应用更改。

模式和数据库克隆默认使用共享导出快照，因此相关表可以一致地复制。长时间的快照可能会延迟源真空清理和 WAL 回收。仅在明确接受跨表不一致性时才将 consistent 选项设置为 false。

### 异步作业

异步执行需要预加载和重启：

    shared_preload_libraries = 'pgclone'

    SELECT pgclone.schema_async(
      'host=source.example dbname=app user=clone_user',
      'sales',
      true,
      '{"parallel":4}'
    );

    SELECT * FROM pgclone.jobs_view;
    SELECT pgclone.progress(1);
    SELECT pgclone.cancel(1);

pgclone 还暴露了 progress_detail、resume 和 clear_jobs 用于作业管理。根据所需的并行度调整 max_worker_processes。

### 重要边界

- 上游使用指南要求超级用户权限来安装和使用 pgclone。
- 异步模式下的模式/数据库/并行路径在 v4.5.0 中不应用脱敏、表筛选或 exclude_tables。当这些控制是安全需求时，请使用文档中同步路径。
- 请勿将密码存储在 SQL 和日志中；优先考虑 libpq 服务文件、密钥文件或其他受控凭据机制。
- v4.4.2 版本改进了序列状态的复制，并保护 PostgreSQL 17 源会话免受 transaction_timeout 影响，但调用者仍需验证对象所有权、扩展、角色和克隆后应用程序行为。
