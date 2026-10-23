## 用法

来源：

- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/README.md)
- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/pg_statecharts.control)
- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/function/start_machine_with_latest_statechart.sql)
- [官方资料](https://github.com/kronor-io/statecharts/blob/472349e864440c39a30863a0db815050391804c7/pg_statecharts/sql/table/state_machine_event.sql)

`pg_statecharts` 0.1.0 在 PostgreSQL 事务内执行层次状态机及其事件。本版本为纯 SQL 实现，测试覆盖 PostgreSQL 16–18，无需预加载或重启。

### 核心流程

将 `ltree` 安装在 `public` 模式，再使用具备所需对象创建权限的角色安装扩展：

```sql
CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;
CREATE EXTENSION pg_statecharts;
```

启动机器前，应先通过 SQL 迁移或开发工具扩展部署状态图定义。对于已经部署且名为 order 的状态图，以下操作启动最新版本并返回状态机记录：

```sql
SELECT * FROM fsm.start_machine_with_latest_statechart(1, 'order', '{}');
```

向 `fsm.state_machine_event` 插入分片、状态机标识、事件名称和 JSON 数据即可提交事件。触发器会在插入事务内处理事件，并运行进入与退出回调。调用者必须具备对应表和函数的权限；状态图回调属于应用代码，需要单独审查。

### 目录与版本

定义与运行中的状态机位于 `fsm` 模式。`fsm.to_semver` 解析数字版本，`fsm.semver_text` 将其格式化；版本以整数数组存储，原生排序保持数值顺序。需要比较版本先后时，不应按格式化后的文本排序。

### 备份与旧版迁移

本次核对的 0.1.0 源码将状态表与序列登记为数据库转储对象。事件收件箱有意排除，以免恢复时重复执行回调。0.0.0 版本产生的备份不含状态机数据；升级后应重新备份并验证恢复。

较早安装的 0.1.0 可能尚未登记转储对象，因此应检查实际目录，不能只看版本字符串。不要为了刷新登记而删除含有数据的扩展。旧 Rust 或 Sqitch 部署的迁移需要超级用户操作，转换期间还依赖旧 semver 扩展；应遵循迁移指南并在恢复副本上演练。上游已不再单独测试旧 0.0.0 迁移路径。
