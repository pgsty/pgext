

## 用法

> [plisql: PL/iSQL 过程语言](https://github.com/IvorySQL/IvorySQL/tree/master/src/pl/plisql)

PL/iSQL 是 IvorySQL 项目提供的 PostgreSQL Oracle 兼容过程语言。它使用 Oracle PL/SQL 语法和语义扩展了 PL/pgSQL。

### 启用

```sql
CREATE EXTENSION plisql;
```

### 创建函数

```sql
CREATE OR REPLACE FUNCTION hello_world
RETURN VARCHAR2
AS
BEGIN
    RETURN 'Hello, World!';
END;
/
```

### Oracle 风格过程

```sql
CREATE OR REPLACE PROCEDURE update_salary(
    p_emp_id IN NUMBER,
    p_amount IN NUMBER
)
AS
BEGIN
    UPDATE employees SET salary = salary + p_amount WHERE emp_id = p_emp_id;
END;
/

CALL update_salary(100, 5000);
```

### 关键特性

- Oracle 风格 `BEGIN...END` 块
- `IN`、`OUT`、`IN OUT` 参数模式
- Oracle 风格异常处理，带命名异常
- `%TYPE` 和 `%ROWTYPE` 属性引用
- Oracle 兼容游标语法（`CURSOR...IS`、`OPEN`、`FETCH`、`CLOSE`）
- 函数声明中使用 `RETURN` 替代 `RETURNS`
- 类包变量作用域

### 与 PL/pgSQL 的区别

- 使用 `AS` 关键字替代 `$$` 分隔符
- 支持 Oracle 风格 `/` 作为语句终结符
- 原生支持 `VARCHAR2`、`NUMBER` 和其他 Oracle 类型
- Oracle 兼容 `DBMS_OUTPUT` 集成
