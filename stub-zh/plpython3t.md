## 用法

来源：

- [官方扩展 SQL](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t--1.0.sql)
- [官方扩展控制文件](https://github.com/pykello/plpythont/blob/3b05e9ffbe405d7768bec8fc3a0764642ceff125/plpython3t.control)

`plpython3t` 把 PostgreSQL 现有的 PL/Python 3 处理器注册为可信过程语言。它没有增加沙箱或受限 Python 运行时；只是改变了同一个服务端 Python 解释器周围的数据库权限边界。

### 核心流程

创建包装语言前，必须已经安装 PL/Python 3 运行时和处理器函数，通常通过 `plpython3u` 完成。

```sql
CREATE EXTENSION plpython3u;
CREATE EXTENSION plpython3t;

CREATE FUNCTION py_add(a integer, b integer)
RETURNS integer
LANGUAGE plpython3t
AS $$
return a + b
$$;

SELECT py_add(2, 3);
```

### 安装对象

扩展使用 `plpython3_call_handler`、`plpython3_inline_handler` 和 `plpython3_validator` 执行 `CREATE TRUSTED PROCEDURAL LANGUAGE plpython3t`。control 文件指向正常的 `plpython3` 共享库。它没有额外 SQL 函数、GUC 或预加载要求。

### 安全边界

标准 PL/Python 被 PostgreSQL 标为不可信语言，是因为 Python 代码能以数据库服务端权限访问服务器文件、进程、库和操作系统环境。`plpython3t` 只是把同一组处理器标记为可信，从而允许具备相应数据库和模式权限的普通角色创建这类函数。

因此，安装这个扩展授予的能力远大于其短小 SQL 文件所暗示的范围。只能在严格受控的数据库使用，并确保每个能够创建 `plpython3t` 函数的角色都在操作系统层面可信。不能把其他无关扩展的 mruby 风格沙箱能力套用到这里。
