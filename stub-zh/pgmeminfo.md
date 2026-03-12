


## 用法

> [pgmeminfo: PostgreSQL 内存上下文信息](https://github.com/okbob/pgmeminfo)

pgmeminfo 提供函数来检查 PostgreSQL 后端的内存使用情况和内存上下文层次结构。

### 函数

**内存信息概览：**

```sql
-- 显示整体内存信息
SELECT * FROM pgmeminfo();
```

**内存上下文层次结构：**

```sql
-- 显示累积的内存上下文大小
SELECT * FROM pgmeminfo_contexts();

-- 显示指定深度的内存上下文
SELECT * FROM pgmeminfo_contexts(deep => 1);

-- 显示所有上下文（不累积）
SELECT * FROM pgmeminfo_contexts(deep => -1, accum_mode => 'off');
```
