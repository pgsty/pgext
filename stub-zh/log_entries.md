## 用法

来源：

- [1.0 版扩展 SQL](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries--1.0.sql)
- [触发器实现](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/log_entries.c)
- [上游回归示例](https://github.com/Houdini/log_entries/blob/fd6e860c91abf9990c4a5e2a2d28fe536a7a2845/sql/log_entries.sql)

`log_entries` 是一个历史 C 触发器，在行插入和更新期间把当前时间与 PostgreSQL 角色写入两个指定列。1.0 版要求第一个目标列使用已淘汰的 `abstime` 类型，第二个使用 `text`，因此仅适用于旧版 PostgreSQL。

### 历史工作流

在仍提供 `abstime` 的 PostgreSQL 版本上，以恰好两个列名参数挂接函数：

```sql
CREATE EXTENSION log_entries;

CREATE TABLE legacy_records (
  record_id integer PRIMARY KEY,
  payload text,
  updated_at abstime,
  updated_by text
);

CREATE TRIGGER legacy_records_log_entries
BEFORE INSERT OR UPDATE OR DELETE ON legacy_records
FOR EACH ROW
EXECUTE FUNCTION log_entries(updated_at, updated_by);
```

对于插入与更新事件，触发器写入当前绝对时间和 `current_user`。删除事件则原样返回旧元组，不修改审计列。直接调用与语句级触发器会被拒绝，而且实现要求恰好两个参数及严格匹配的旧类型。

### 兼容性边界

现代 PostgreSQL 已移除 `abstime`，因此上游 SQL 与 C 代码无法原样用于当前受支持版本。源码只有一次 2016 年提交，没有 README、升级脚本或兼容矩阵。其 C 实现还使用旧触发器与元组 API，不能假定在新版服务器头文件下仍安全工作。

不能只把表中的 `abstime` 改为 `timestamptz`：二进制函数会检查旧类型 OID 并拒绝新类型。迁移需要经审查的源码移植与新回归测试；更合适的做法通常是编写使用 `timestamptz` 并明确操作者策略的小型 PL/pgSQL 触发器。

### 审计注意事项

该触发器保存的是有效数据库角色，不是应用用户、原始会话角色、客户端身份、语句文本或不可变历史。有权更新审计列或禁用触发器的用户可以改变结果。删除行也不会留下独立审计记录。

1.0 版可重定位，且没有声明预加载或扩展依赖。应把它视为旧数据库的源码考古。当前审计需求应使用受支持类型，保护审计数据与触发器所有权，定义操作者语义，并明确测试更新、删除、复制、批量加载与恢复行为。
