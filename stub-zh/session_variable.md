## 用法

来源：

- [3.6 README](https://github.com/splendiddata/session_variable/blob/3.6/README.md)
- [3.6 SQL API](https://github.com/splendiddata/session_variable/blob/3.6/session_variable--3.6.sql)
- [3.6 control file](https://github.com/splendiddata/session_variable/blob/3.6/session_variable.control)
- [3.6 initialization code](https://github.com/splendiddata/session_variable/blob/3.6/session_variable.c)

`session_variable` 定义数据库级变量和常量，并为每个会话保存独立的值。修改会话本地值不会影响其他连接。

### 创建变量和常量

```sql
CREATE EXTENSION session_variable;

-- Create a variable with initial value
SELECT session_variable.create_variable('my_var', 'text'::regtype, 'initial text'::text);

-- Create a variable with NULL initial value
SELECT session_variable.create_variable('my_date_var', 'date'::regtype);

-- Create a constant (cannot be changed via set())
SELECT session_variable.create_constant('my_env', 'text'::regtype, 'Production'::text);
```

### 获取和设置值

```sql
-- Get variable value (second arg is type hint)
SELECT session_variable.get('my_var', null::text);

-- Set variable value (returns true on success)
SELECT session_variable.set('my_var', 'new text'::text);
```

### 在 PL/pgSQL 中使用

```sql
DO $$
DECLARE
    my_field text;
BEGIN
    my_field := session_variable.get('my_var', my_field);
    RAISE NOTICE 'Value: %', my_field;
END
$$ LANGUAGE plpgsql;
```

### 管理函数

```sql
-- Alter the initial/constant value (affects new sessions)
SELECT session_variable.alter_value('my_env', 'Development'::text);

-- Reload all variables from database definitions
SELECT session_variable.init();

-- Drop a variable or constant
SELECT session_variable.drop('my_var');

-- Check if a variable exists
SELECT session_variable.exists('my_var');

-- Get the type of a variable
SELECT session_variable.type_of('my_var');
```

### 关键行为

- 变量在数据库级别定义；每个会话获取本地副本
- `set()` 仅更改会话本地副本；其他会话不受影响
- `alter_value()` 更改存储的值；新会话将看到它，现有会话需要 `init()` 来刷新
- 常量不能通过 `set()` 更改，只能通过 `alter_value()`
- 变量和常量名称在两种类型之间必须唯一

### 权限、读取与版本边界

由超级用户在 `session_variable` 模式中安装这个不可重定位的扩展；上述 SQL 工作流无需共享预加载或重启。管理定义应授予 `session_variable_administrator_role`，普通访问应授予 `session_variable_user_role`。管理员角色包含用户角色。

`session_variable.set` 和 `session_variable.alter_value` 成功时返回布尔值 true，不返回原值。成功修改存储定义后，调用者立即看到变化，提交后启动的会话也能看到变化；已有会话保留本地状态，直到 `session_variable.init()` 重置所有值。定义保存在 `session_variable.variables` 中，并纳入逻辑备份。

`session_variable.get_stable` 可能在语句执行期间缓存结果；如果触发器在同一语句中修改该值，应使用普通读取函数。`session_variable.get_constant` 标记为 IMMUTABLE，在管理员修改常量后可能返回缓存结果；修改期间应使用普通读取函数。

3.6 移除过时的版本 1 初始化支持。此前的 3.5 禁止为 `collection` 和 `icollection` 设置非空初始值，因为从文本初始化可能导致后端崩溃。使用这些类型时，应创建初始值为空的变量，并按源码中的初始化钩子 `session_variable.variable_initialisation()` 设置值；README 中另一个钩子名称与实现不一致。上游列出 PostgreSQL 14-18 支持，以及暂定的 PostgreSQL 19 支持。
