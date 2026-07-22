## 用法

来源：

- [官方 README](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/README.md)
- [1.0 版本 SQL 函数](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck--1.0.sql)
- [解析器实现](https://github.com/jconway/pgsynck/blob/8836bdd4c70e9417dd94a11487ee50e42ba7a969/pgsynck.c)

`pgsynck` 将文本拆分成 SQL 语句，在不执行的情况下把每条语句送入 PostgreSQL 原始解析器，并为每条语句返回一行诊断。它适合做早期语法筛查，但不能证明 SQL 可以执行。

### 检查脚本

```sql
CREATE EXTENSION pgsynck;

SELECT sql, cursorpos, sqlerrcode, message, hint
FROM pgsynck($sql$
  SELECT 1;
  SELECT 1 FROM;
  INSERT INTO target_table(id) VALUES (1);
$sql$);
```

原始解析器接受语句时，`cursorpos` 与 `sqlerrcode` 为零，`message` 与 `hint` 为空字符串。语法错误会返回解析器光标位置、PostgreSQL 内部编码的整数错误码、主要消息与可选提示。`sqlerrcode` 不是常见的五字符 SQLSTATE 文本。

### 校验边界

解析器不会解析表名或列名、检查权限、推断参数类型、执行重写/规划，也不会运行函数与约束。因此，不存在的关系或类型错误仍可能通过。需要语义保证时，应以正确模式与角色在事务中执行后续验证。

1.0 版本实现了自己的语句分割器。它可以处理普通单/双引号、美元引用字符串与注释，但早于许多当前 PostgreSQL 版本，也不能替代完整服务器词法分析器。接受不可信脚本前，应测试嵌套注释、转义字符串设置、带标签美元引用、过程体、空语句与光标位置。还应限制输入大小与语句数，避免把数据库后端变成无限制解析服务。

控制文件注释称扩展会执行 shell 命令，但已复核的 C 与 SQL 实现只调用 PostgreSQL 原始解析器。这个不一致也说明：使用该已废弃版本前必须本地审阅源码，并针对准确版本测试。
