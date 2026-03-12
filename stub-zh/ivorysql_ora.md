

## 用法

> [ivorysql_ora: PostgreSQL 数据库上的 Oracle 兼容扩展](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/ivorysql_ora)

`ivorysql_ora` 扩展作为 IvorySQL 项目的一部分，为 PostgreSQL 提供 Oracle 兼容特性。它添加了 Oracle 兼容的数据类型、函数和 PL/SQL 行为。

### 启用

```sql
CREATE EXTENSION ivorysql_ora;
```

### Oracle 兼容数据类型

该扩展添加了 Oracle 风格的数据类型，包括：

- `NUMBER` / `NUMBER(p,s)` - Oracle 兼容的数值类型
- `VARCHAR2(n)` - Oracle 兼容的变长字符串
- `DATE` - 带时间部分的 Oracle 风格 DATE
- `BINARY_FLOAT` / `BINARY_DOUBLE` - IEEE 浮点类型

### Oracle 兼容函数

提供 Oracle 风格的字符串操作、日期运算、数值操作和类型转换内置函数，行为与 Oracle 语义一致。

### 兼容模式

IvorySQL 支持 Oracle 兼容模式，可更改解析器行为：

```sql
SET compatible_mode TO oracle;  -- 启用 Oracle 兼容
SET compatible_mode TO pg;      -- 恢复为标准 PostgreSQL
```

在 Oracle 模式下，SQL 解析器接受 Oracle 风格语法，包括：

- Oracle 风格外连接（`(+)` 语法）
- `CONNECT BY` 层次查询
- Oracle 风格序列（`sequence.NEXTVAL`）
- 包风格对象引用
