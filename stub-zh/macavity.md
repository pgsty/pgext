## 用法

来源：

- [README.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/README.md)
- [CHANGELOG.md](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/CHANGELOG.md)
- [macavity.control](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/macavity.control)
- [sql/macavity--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/sql/macavity--0.2.0.sql)
- [sql/macavity--0.1.0--0.2.0.sql](https://github.com/CrystallineCore/Macavity/blob/v0.2.1/sql/macavity--0.1.0--0.2.0.sql)

`macavity` 0.2.1 为可丢弃的 PostgreSQL 测试集群提供确定性故障注入。上游支持 Linux 上的 PostgreSQL 16–19，其他平台尚未验证，Windows 不支持崩溃注入。崩溃动作可能断开所有会话并触发崩溃恢复。

### 事件用法

```sql
CREATE EXTENSION macavity;
SELECT * FROM macavity_points();
SELECT macavity_arm('executor_start', 'error', 1);
SELECT 1; -- expected injected error
SELECT * FROM macavity_status();
SELECT macavity_disarm();
SELECT macavity_reset();
```

### 事件表与升级

每个会话保存多个事件，各有固定 ID、独立计数及 `armed`、`completed` 或 `disarmed` 状态。`macavity_arm` 返回整数事件 ID；仅接收 ID 的重载会重新启用已完成或已停用事件。`macavity_status` 返回零到多行事件。`macavity_disarm` 保留事件历史，`macavity_reset` 则清空历史并重置 ID。

注入点为 `executor_start`、`executor_end`、`before_commit` 与 `before_abort`，动作为 `error`、`delay` 与 `crash`。延迟固定一秒，回滚期间不允许注入错误。同一次命中按延迟、崩溃、错误的顺序处理，同类事件按 ID 排序。

创建需要超级用户，无需共享预加载。只有注入点枚举向 `PUBLIC` 开放。从 0.1.0 使用 `ALTER EXTENSION macavity UPDATE` 升级会以新签名重新创建函数：须恢复显式授权、修改调用方并重新连接旧会话。当前 Pigsty 打包字段仍对应 0.1.0。

### 版本 0.2.1

0.2.0 到 0.2.1 的 SQL 升级脚本为空操作；本版本新增 PostgreSQL 19 兼容性，SQL API 不变。安装匹配的库和 SQL 文件，执行 `ALTER EXTENSION macavity UPDATE TO '0.2.1';`，并在测试前重新连接旧会话。
