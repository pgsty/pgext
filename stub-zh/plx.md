## 用法

来源：

- [PGXN plx 3.0.0 README](https://pgxn.org/dist/plx/3.0.0/README.html)
- [plx 3.0.0 用户指南](https://api.pgxn.org/src/plx/plx-3.0.0/doc/USERGUIDE.md)
- [plx 3.0.0 兼容性说明](https://api.pgxn.org/src/plx/plx-3.0.0/doc/COMPATIBILITY.md)
- [plx 3.0.0 变更日志](https://api.pgxn.org/src/plx/plx-3.0.0/CHANGELOG.md)
- [plx 控制文件](https://api.pgxn.org/src/plx/plx-3.0.0/plx.control)

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

### 3.0.0 升级

安装新库并运行 `ALTER EXTENSION plx UPDATE TO '3.0.0'` 后，用原始方言 DDL 重新创建受影响函数。扩展升级不会重写已有的生成 PL/pgSQL。迁移前保留原始源码，重点检查依赖插值、NULL 处理或异常退出结果的函数。

扩展不再可重定位。安装时选定模式，并在恢复时保持相同模式，否则需重新生成嵌入辅助对象名称的函数。创建方言函数或 DO 块要求对 plpgsql 具有 USAGE 权限。启用函数体检查时，创建过程会验证生成的 PL/pgSQL，拒绝不支持的参数或结果类型以及格式错误的类型声明。

SQL 插值中的值现在绑定为参数；名称使用 `ident()`，有意拼接的 `raw()` SQL 片段须限于可信输入。向字符串追加 NULL 现在传播 NULL；若要视为空字符串，应使用 `coalesce(x, '')`。多数方言现在会在 return、break、continue 时执行 finally/ensure，但查询取消和断言失败仍不在该处理范围内。重新部署函数前应阅读各方言的破坏性变更。
