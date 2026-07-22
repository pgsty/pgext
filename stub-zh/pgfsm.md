## 用法

来源：

- [Official README](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/README.md)
- [Extension SQL](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/pgfsm--0.0.1.sql)
- [Official worked example](https://github.com/michelp/pgfsm/blob/ee79a0e12ef73474a07b34fc36b3b2214e85b42e/sql/test.sql)

pgfsm 是一个紧凑的 SQL 示例，用于在 PostgreSQL 中存储有限状态机定义及其实例。它保证每个实例都处于已声明状态，并要求状态更新遵循允许的目标状态。它是教学代码，不是通用工作流引擎。

### 核心流程

把扩展安装到专用 `fsm` 模式，定义转换，创建实例，再按转换标签推进：

```sql
CREATE SCHEMA fsm;
CREATE EXTENSION pgfsm WITH SCHEMA fsm;
SET search_path = fsm, public;

INSERT INTO transition(name, from_state, transition, to_state) VALUES
  ('turnstile', 'locked',   'coin', 'unlocked'),
  ('turnstile', 'unlocked', 'push', 'locked');

INSERT INTO machine(name, state)
VALUES ('turnstile', 'locked');

SELECT * FROM do_transition(1, 'coin');
```

每个可用状态都必须作为 `from_state` 出现，因为 machine 表有一个指向该列的外键。

### 重要对象

- `transition` 存储状态机名称、当前状态、转换标签和目标状态。
- `machine` 存储状态机实例。
- `transitions_for` 返回实例当前状态可用的转换行。
- `states_for` 返回状态机定义中声明的源状态。
- `do_transition` 应用匹配标签并返回更新后的实例。
- `check_valid_state_update` 和 `check_valid_state_insert` 为校验触发器提供实现。

### 设计限制

`transition(name, from_state)` 上的主键只允许每个状态有一条出向转换，因此不能按标签建立分支。直接更新可以不提供标签而选择允许的目标；如果转换标签有意义，应撤销直接表更新权限。它没有历史记录、回调执行、乐观锁版本、所有权模型或工作流定义迁移支持。将该模式用于应用状态前，需要有意识地扩展其设计。
