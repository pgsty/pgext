

## 用法

> [session_variable: 会话变量和常量的注册与操作](https://github.com/splendiddata/session_variable)

### 创建变量和常量

```sql
CREATE EXTENSION session_variable;

-- 创建带初始值的变量
SELECT session_variable.create_variable('my_var', 'text'::regtype, 'initial text'::text);

-- 创建初始值为 NULL 的变量
SELECT session_variable.create_variable('my_date_var', 'date'::regtype);

-- 创建常量（不能通过 set() 更改）
SELECT session_variable.create_constant('my_env', 'text'::regtype, 'Production'::text);
```

### 获取和设置值

```sql
-- 获取变量值（第二个参数是类型提示）
SELECT session_variable.get('my_var', null::text);

-- 设置变量值（返回之前的值）
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
-- 修改初始/常量值（影响新会话）
SELECT session_variable.alter_value('my_env', 'Development'::text);

-- 从数据库定义重新加载所有变量
SELECT session_variable.init();

-- 删除变量或常量
SELECT session_variable.drop('my_var');

-- 检查变量是否存在
SELECT session_variable.exists('my_var');

-- 获取变量类型
SELECT session_variable.type_of('my_var');
```

### 关键行为

- 变量在数据库级别定义；每个会话获取本地副本
- `set()` 仅更改会话本地副本；其他会话不受影响
- `alter_value()` 更改存储的值；新会话将看到它，现有会话需要 `init()` 来刷新
- 常量不能通过 `set()` 更改，只能通过 `alter_value()`
- 变量和常量名称在两种类型之间必须唯一
