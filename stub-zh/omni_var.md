


## 用法

> [omni_var: 作用域变量](https://docs.omnigres.org/omni_var/variables/)

`omni_var` 扩展提供具有事务、会话和语句作用域的类型化变量。

### 设置变量

```sql
SELECT omni_var.set('my_var', true);              -- 事务作用域
SELECT omni_var.set_session('my_var', true);       -- 会话作用域
SELECT omni_var.set_statement('my_var', true);     -- 语句作用域
```

显式指定类型：

```sql
SELECT omni_var.set('text_var', 'value'::text);
```

### 获取变量

```sql
SELECT omni_var.get('my_var', false);              -- 事务作用域，默认值=false
SELECT omni_var.get_session('my_var', false);       -- 会话作用域
SELECT omni_var.get_statement('my_var', false);     -- 语句作用域
```

默认值有双重用途：变量未找到时返回默认值，同时其类型指定了期望的返回类型。

### 行为

- 如果变量被显式设置为 `null`，`get` 返回 `null`（而非默认值）
- 类型不匹配会引发错误：`ERROR: type mismatch DETAIL: expected integer, got boolean`
- 事务作用域变量绑定到所在事务的生命周期
- 会话变量在整个会话期间持久存在
