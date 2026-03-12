

## 用法

> [pg_get_functiondef: 获取函数定义](https://github.com/IvorySQL/IvorySQL/tree/master/contrib/pg_get_functiondef)

`pg_get_functiondef` 扩展提供获取 PostgreSQL 函数和过程完整定义（DDL）的功能，在 IvorySQL 的 Oracle 兼容上下文中特别有用。

### 启用

```sql
CREATE EXTENSION pg_get_functiondef;
```

### 获取函数定义

```sql
-- 通过 OID 获取函数的 DDL
SELECT pg_get_functiondef(oid) FROM pg_proc WHERE proname = 'my_function';

-- 通过名称获取函数定义
SELECT pg_get_functiondef('my_function'::regproc);
```

该扩展扩展了内置的 `pg_get_functiondef()`，以支持 IvorySQL 使用的 Oracle 兼容函数和过程语法，包括 PL/iSQL 过程体和 Oracle 风格参数声明。
