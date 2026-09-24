## 用法

来源：

- [Logical apply visibility guide](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/lsn-wait.md)
- [Control file](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_wait/pgwrh_wait.control)
- [Extension SQL](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/pgwrh_wait/pgwrh_wait--1.0.0-alpha1.sql)
- [Alpha release boundaries](https://github.com/mkleczek/pgwrh/blob/v1.0.0-alpha1/docs/releases/1.0.0-alpha1.md)

`pgwrh_wait` 1.0.0-alpha1 使读取在获取快照前，等待逻辑复制订阅端应用到发布端的指定水位。它面向 PostgreSQL 18 和内置 `pgoutput` 插件，可独立于其他 pgwrh 扩展运行。此 alpha 版本只提供全新安装脚本。

### 在订阅端启用

将库追加到 `shared_preload_libraries`，保留其他现有条目，然后重启 PostgreSQL。随后以管理员身份在每个需要提供受保护读取的数据库中安装 SQL API：

```ini
shared_preload_libraries = 'pgwrh_wait'
```

```sql
CREATE EXTENSION pgwrh_wait;
SELECT pgwrh.applied_lsn('pgwrh_replica_subscription');
```

依赖事务设置前，先在独立的健康检查事务中执行诊断调用。未预加载时可以创建扩展，但调用 API 会报告前置条件错误。自定义设置赋值成功不能证明库已启用。监控初始化前，`pgwrh.applied_lsn` 可能返回 NULL，该函数本身也不证明表的初始同步已经完成。

### 在获取快照前等待

必须使用实际执行读取的订阅端连接。将示例 LSN 替换为相关写入提交后从正确发布端取得的水位，并使用实际订阅名和表名：

```sql
BEGIN ISOLATION LEVEL REPEATABLE READ READ ONLY;
SET LOCAL pgwrh.read_after_subscription = 'pgwrh_replica_subscription';
SET LOCAL pgwrh.wait_timeout_ms = '5s';
SET LOCAL pgwrh.read_after_lsn = '0/12345678';
SELECT * FROM my_table;
COMMIT;
```

必须在任何查询或快照之前发送水位，包括 `SELECT 1` 和 `set_config`。设置必须直接应用于显式顶层事务，不能放在函数或保存点中。支持读已提交、可重复读和可串行化隔离级别。超时或取消会使操作失败，此时应回滚，不能继续读取旧数据。

`pgwrh.wait_for_lsn(text, pg_lsn, integer)` 提供另一种 SQL 函数接口，默认超时为 10,000 毫秒。它只能保护随后单独执行的读已提交语句；同一语句内的等待与读取仍使用旧快照。该函数拒绝可重复读和可串行化隔离级别。普通用户可以调用两种等待接口，但仍受常规表权限约束。

### 复制与恢复要求

订阅中的所有表都必须完成初始同步。不支持两阶段提交订阅，也不接受待执行的事务跳过操作。LSN 属于特定发布端的复制历史，路由必须保持这一对应关系。进度依据已应用的提交推进，而非已收到的 WAL 或保活消息。提交后采样的水位可能超过最后一次已发布提交，因此空闲或经过过滤的订阅可能需要后续被复制的心跳事务。

`pgwrh.max_tracked_subscriptions` 默认为 256，修改后需要重启。已删除的订阅仍占用条目，直到重启才释放。使用外部表时，每个实际远程事务都必须在获取快照前应用屏障。该保证只是可见性的下界，并非集群级快照或故障转移后的持久性保证。跳过变更、出现数据分歧或执行恢复后，需要重新建立有效的数据基线。
