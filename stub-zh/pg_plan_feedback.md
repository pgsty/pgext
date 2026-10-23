## 用法

来源：

- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/README.md)
- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/pg_plan_feedback.control)
- [Official source](https://github.com/pganalyze/pg_plan_feedback/blob/26769778708c9e49cef1ea3f4c3e48cdc40b6f3f/pg_plan_feedback--1.0.sql)
- [Official source](https://github.com/lfittl/postgres/pull/82)

`pg_plan_feedback` 从执行结果学习各查询、各关系的行数修正系数，并可将其用于后续规划。它依赖尚需额外应用的 PostgreSQL 内核补丁，不能视为兼容标准发行版。上游明确将其标记为实验性质，不适合生产。

### 先观察再应用

在具有所需内核钩子且已安装库文件的服务器上：

```sql
LOAD 'pg_plan_feedback';
CREATE EXTENSION pg_plan_feedback;
SET pg_plan_feedback.mode = 'observe';
SELECT count(*) FROM pg_class;
SELECT pg_plan_feedback_flush();
SELECT * FROM pg_plan_feedback;
```

查看结果的角色需要 `pg_read_all_stats` 或显式授予的等效权限。`observe` 只记录反馈，不改变估计；`apply` 使用学习到的修正；`off` 关闭功能。切换模式前应运行有代表性的查询并检查反馈。

### 对象与限制

`pg_plan_feedback_reset(queryid)` 清除修正，省略参数时全部清除。`pg_plan_feedback_set(queryid, relation, correction)` 设置人工修正系数。两者默认限制权限，因为会影响其他用户的计划。`pg_plan_feedback_stats()` 返回数量统计，视图暴露跨数据库的估计值、实际行数和修正系数。

库可以通过共享预加载、会话预加载或显式 `LOAD` 加载；动态共享内存本身不要求共享预加载。SQL 扩展可以迁移模式，安装需要超级用户。它学习基础关系的行数，不强制连接顺序或完整计划。未完成的扫描不能提供可靠反馈；缓存计划需要重新规划才会应用新估计。不要假定实验性修订之间的存储格式和接口稳定。
