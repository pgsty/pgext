

## 用法

> [E-Maj: 记录并回滚表内容变更](https://github.com/dalibo/emaj)
>
> [**文档**](https://emaj.readthedocs.io/en/latest/) | [Emaj_web 图形界面](https://github.com/dalibo/emaj_web)

E-Maj 记录表上执行的变更操作（Insert、Update、Delete、Truncate），涵盖一个或多个表集合，并可在需要时高效地撤销这些变更，将表集合恢复到预定义的稳定状态。

在**开发环境**中，它可以轻松回滚程序执行产生的所有更新，方便反复重放测试流程。

在**生产环境**中，它提供：

- 表变更历史记录，供检查与审计
- 表组的批间保存点
- 无需停机即可将表组"恢复"到稳定状态
- 批处理窗口内的多个保存点，每个都可用作恢复点


## 核心概念

### 表组（Tables Group）

表组是一组生命节奏相同的应用表——它们的内容可以作为整体恢复。一个组可以包含不同 schema 中的表和序列。每个组处于两种状态之一：**LOGGING**（记录中）或 **IDLE**（空闲），并可指定为：

- **ROLLBACKABLE**（标准）——支持记录和回滚
- **AUDIT_ONLY**——只记录变更，不支持回滚，适用于没有主键的表或 UNLOGGED 表

### 标记（Mark）

标记代表表组生命周期中的快照时刻，捕获所有组成员的稳定状态。标记名称在组内唯一。

### 回滚（Rollback）

回滚操作将表和序列恢复到建立特定标记时的状态：

- **非日志回滚**（Unlogged rollback）——撤销的更新不留痕迹
- **日志回滚**（Logged rollback）——撤销操作也会被记录，允许后续再次反转

注意：E-Maj 的回滚与 PostgreSQL 原生的事务回滚有本质区别。


## 主要函数

### 启动表组

```sql
SELECT emaj.emaj_start_group('my_group', 'mark_1');
```

激活变更记录。表组必须处于 IDLE 状态。会自动创建初始标记。

### 设置中间标记

```sql
SELECT emaj.emaj_set_mark_group('my_group', 'mark_2');
```

记录应用状态的时间点快照。表组必须处于 LOGGING 状态。

### 回滚表组

非日志回滚（恢复表，不留撤销痕迹）：

```sql
SELECT * FROM emaj.emaj_rollback_group('my_group', 'mark_1');
```

日志回滚（允许撤销本次回滚）：

```sql
SELECT * FROM emaj.emaj_logged_rollback_group('my_group', 'mark_1');
```

两者都支持 `'_EMAJ_LAST_MARK_'` 关键字来指定最近的标记。

### 停止表组

```sql
SELECT emaj.emaj_stop_group('my_group');
```

停用记录功能，将表组状态从 LOGGING 切换为 IDLE。


## 多组操作

支持同时对多个表组进行批量操作：

```sql
SELECT emaj.emaj_start_groups('{"group1","group2"}', 'multi_mark');
SELECT emaj.emaj_set_mark_groups('{"group1","group2"}', 'common_mark');
SELECT * FROM emaj.emaj_rollback_groups('{"group1","group2"}', 'common_mark');
SELECT emaj.emaj_stop_groups('{"group1","group2"}');
```


## 变更检查

E-Maj 提供函数用于统计和检查标记之间的数据变更内容，并可生成重放已记录变更的 SQL 脚本。这对审计和调试非常有用。


## Emaj_web

**Emaj_web** 是一个基于 PHP 的 Web 图形管理工具，提供友好的 E-Maj 管理界面。可在 [GitHub](https://github.com/dalibo/emaj_web) 获取，详见[文档](https://emaj.readthedocs.io/en/latest/webOverview.html)。
