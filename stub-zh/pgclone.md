## 用法

来源：

- [pgclone v4.4.2 README](https://github.com/valehdba/pgclone/blob/v4.4.2/README.md)
- [pgclone v4.4.2 使用指南](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/USAGE.md)
- [异步克隆指南](https://github.com/valehdba/pgclone/blob/v4.4.2/docs/ASYNC.md)
- [pgclone v4.4.2 发行说明](https://github.com/valehdba/pgclone/releases/tag/v4.4.2)

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
- 异步模式下的模式/数据库/并行路径在 v4.4.2 中不尊重屏蔽、表或 exclude_tables。当这些控制是安全需求时，请使用文档中同步路径。
- 请勿将密码存储在 SQL 和日志中；优先考虑 libpq 服务文件、密钥文件或其他受控凭据机制。
- v4.4.2 版本改进了序列状态的复制，并保护 PostgreSQL 17 源会话免受 transaction_timeout 影响，但调用者仍需验证对象所有权、扩展、角色和克隆后应用程序行为。
