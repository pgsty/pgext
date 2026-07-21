## 用法

来源：

- [plx 1.3.1 README](https://github.com/commandprompt/plx/blob/v1.3.1/README.md)
- [plx 文档](https://commandprompt.github.io/plx/)
- [plx 用户指南](https://github.com/commandprompt/plx/blob/v1.3.1/doc/USERGUIDE.md)
- [plx 限制](https://github.com/commandprompt/plx/blob/v1.3.1/doc/LIMITATIONS.md)
- [plx 1.3.1 发行版](https://github.com/commandprompt/plx/releases/tag/v1.3.1)

`plx` 提供了熟悉的程序语言方言，当 `CREATE FUNCTION` 执行时会编译为普通的 PL/pgSQL。PostgreSQL 存储并执行生成的 PL/pgSQL，使用其内置的信任处理程序；无需加载 Ruby、PHP、JavaScript、Python、Go、COBOL、Oracle 或 SQL Server 运行时到后端。

```sql
CREATE EXTENSION plx;
```

### 可用方言

| 语言 | 表面语法 |
|---|---|
| `plxruby` | Ruby |
| `plxphp` | PHP |
| `plxjs` | JavaScript |
| `plxts` | TypeScript 注解在 JavaScript 方言之上 |
| `plxpython3` | Python 3 |
| `plxgo` | Go |
| `plxcobol` | ISO COBOL |
| `plxplsql` | Oracle PL/SQL |
| `plxtsql` | Transact-SQL |

所有方言都针对相同的 PL/pgSQL 语句表面，包括赋值、条件判断、循环、查询迭代、动态 SQL、游标、异常处理、触发器和集合返回函数。

### 创建一个函数

在 `LANGUAGE` 子句中选择一种方言，并保持函数签名使用 PostgreSQL 类型：

```sql
CREATE FUNCTION grade(score integer)
RETURNS text
LANGUAGE plxruby
AS $$
  grade #:: text
  if score >= 90
    grade = "A"
  elsif score >= 80
    grade = "B"
  else
    grade = "F"
  end
  return grade
$$;

SELECT grade(85);
```

翻译只会在创建函数时发生一次。存储在 `pg_proc.prosrc` 中的可执行体是普通的 PL/pgSQL，因此可以导出、审查和运行而无需单独的解释器。

### 检查和调试生成代码

```sql
SELECT pg_get_functiondef('grade(integer)'::regprocedure);
SELECT prosrc
FROM pg_proc
WHERE oid = 'grade(integer)'::regprocedure;

SELECT plx_source('grade(integer)'::regprocedure);
```

运行时错误行号指向生成的 PL/pgSQL。使用 `plx_source()` 获取原始嵌入方言体；在与源码相关联时，将其与 `pg_get_functiondef()` 一起使用。

### SQL 和字符串构建

表达式保留 PostgreSQL 的 SQL 语义，而不是模拟完整的源语言运行时环境。对于 SQL 使用每个方言的查询/执行形式，并且非字面量表达式使用显式的 PostgreSQL 类型。`plx_strbuild` 扩展对象辅助函数在 PostgreSQL 18 中加速了重复字符串追加：

```sql
CREATE FUNCTION labels(n integer)
RETURNS text
LANGUAGE plxjs
AS $$
  let out: text = "";
  for (let i = 1; i <= n; i++) {
    out += `item-${i},`;
  }
  return out;
$$;
```

该构建器在 PostgreSQL 13-17 中仍然正确，但其就地优化需要 PostgreSQL 18。

### 边界和注意事项

- plx 实现的是语法表面，而不是源语言的运行时环境：没有 gems、Python 模块、JavaScript 导入、Go 协程、PHP 类或 SQL Server 事务命令。
- 函数在 PL/pgSQL 的信任沙箱中运行，没有任何直接文件系统访问、网络访问、任意本地代码访问或事务控制权限。
- 参数和返回类型必须是 PostgreSQL 类型。局部变量的类型推断有限；对于调用和复合表达式需要显式声明类型。
- SQL 使用三值逻辑和 PostgreSQL 的数值/字符串语义。源语言中的真假性和使用 `+` 进行字符串连接没有被复制。
- 局部变量被提升到一个 PL/pgSQL 的 `DECLARE` 块中，因此块局部作用域和具有不同类型的重新声明不可用。
- 版本 1.3.1 是一个仅包含代码的安全发布：它增加了词法分析/字符串构建容量保护、堆栈深度检查、有限缩进处理以及对原始字符串、PHP 插值和非十进制整数字面量解析的修复。安装二进制文件后，请运行 `ALTER EXTENSION plx UPDATE TO '1.3.1'`。
