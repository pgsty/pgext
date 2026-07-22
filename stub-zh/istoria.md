## 用法

来源：

- [扩展控制文件](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria.control)
- [版本 1.1 安装 SQL](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/istoria--1.1.sql)
- [官方触发器示例](https://github.com/KiriakosGeorgiou/istoria/blob/e1a21fc8e6d262c1990d4ad7fbd80fb1a75d981f/examples/sector.sql)

`istoria` 版本 `1.1` 记录行版本，并为应用表提供可分支的非线性撤销与重做。它是由触发器驱动的 SQL 框架，并不是 PostgreSQL 事务恢复工具。

### 核心流程

被跟踪的表必须有主键。添加 AFTER 行级触发器，并用 JSON 数组指定构成历史会话的列；空数组表示整张表共用一个会话。

```sql
CREATE EXTENSION istoria;

CREATE TRIGGER sector_history_tr
AFTER INSERT OR UPDATE OR DELETE ON sector
FOR EACH ROW
WHEN (pg_trigger_depth() < 1)
EXECUTE PROCEDURE history_trigger_func('["chart_id"]');

SELECT * FROM history_session_actions(1);
SELECT history_session_undo(1);
SELECT history_session_redo(1);
```

移动到较早动作后再修改数据，会创建另一条时间线。`history_session_set_active_action` 可直接移动到指定动作；撤销与重做则通过删除和插入重放保存的行镜像。

### 对象与注意事项

扩展会创建 `history_sessions`、`history_timelines` 和 `history_actions`，并在安装模式内创建一组全局 `history_*` 函数。保存的行使用 JSONB 快照。

重放会修改应用表，并可能触发表约束和其他触发器。上游有多个函数使用 `SECURITY DEFINER` 与动态 SQL，且向 PUBLIC 授权，因此使用前必须审核所有权、权限和搜索路径。请在真实数据副本上测试分支、并发写入、模式变更与重放失败；不要把这个老项目当作审计日志或备份。
