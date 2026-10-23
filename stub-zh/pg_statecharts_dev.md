## 用法

来源：

- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/README.md)
- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts_dev/pg_statecharts_dev.control)
- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)

`pg_statecharts_dev` 0.1.0 是 `pg_statecharts` 的开发工具扩展，用于导入 SCXML 定义与生成 Sqitch 迁移。生产状态机运行时不需要这套服务器文件访问工具。

### 核心流程

安装运行时及其依赖后，启用开发工具：

```sql
CREATE EXTENSION pg_statecharts_dev;
SELECT * FROM fsm.import_scxml_files(
  source_path => '/repo/statecharts',
  recursive => false,
  on_conflict_do_nothing => false
);
```

源目录位于数据库主机上。状态图回调必须已经存在，且签名符合要求。解析或验证失败会回滚导入事务，不会只接收状态图的一部分。

### 生成迁移

```sql
SELECT fsm.gen_statechart_sqitch_migrations(
  source_path => '/repo/statecharts',
  sqitch_plan_file_path => '/repo/sqitch/sqitch.plan',
  recursive => false
);
```

生成前应建立目标目录，并保证数据库操作系统用户可以访问。生成器写入部署、回退与验证脚本，更新计划文件，并可能覆盖已有的状态图生成文件。作为应用迁移使用前，必须检查这些改动。

### 权限与边界

读取服务器文件需要 `pg_read_server_files`，写入需要 `pg_write_server_files`。这些是广泛权限，不是仅限于所提供路径的沙箱。应将扩展及相关角色限制在受控的开发数据库中。它共用 `fsm` 模式，为纯 SQL 实现，无需预加载；文档测试覆盖 PostgreSQL 16–18。

生成迁移时使用运行时当前的语义版本辅助函数。旧迁移若仍转换到已移除的 semver 类型，应按上游迁移指南修订；重新生成全部历史状态图可能引入回调依赖顺序错误。生产端部署审查后的 SQL 即可，无需安装本开发工具。
