## 用法

来源：

- [PGXN 0.1.0](https://pgxn.org/dist/macavity/0.1.0/)

`macavity` 为 PostgreSQL 16–18 提供确定性的会话内故障注入。此测试版本仅用于可丢弃的测试集群。创建扩展需要超级用户，无需共享预加载。

### 注入故障

```sql
CREATE EXTENSION macavity;
SELECT * FROM macavity_points();
SELECT macavity_arm('executor_start', 'error', 1);
SELECT 1; -- expected injected error
SELECT * FROM macavity_status();
SELECT macavity_disarm();
```

故障在指定次数的执行点命中时触发，随后解除；注入错误后仍可读取计数器。每个会话最多设置一个故障，新会话初始没有故障配置。

### 动作与使用边界

执行点包括 `executor_start`、`executor_end`、`before_commit` 和 `before_abort`。动作包括 `error`、`delay` 和 `crash`；延迟固定为一秒，不允许在事务中止处理过程中注入错误。

崩溃动作会杀死调用它的后端，PostgreSQL 随后断开其他会话并执行崩溃恢复。不可在保存重要数据的集群上启用。计数器属于会话，后端退出后即丢失。
