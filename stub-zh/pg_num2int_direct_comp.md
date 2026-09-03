## 用法

来源：

- [官方文档](https://github.com/datastone-inc/pg_num2int_direct_comp/blob/7dba5a74b9a7e4515664cbd601847e9abbd687fa/README.md)
- [扩展控制文件](https://github.com/datastone-inc/pg_num2int_direct_comp/blob/7dba5a74b9a7e4515664cbd601847e9abbd687fa/pg_num2int_direct_comp.control)
- [官方仓库](https://github.com/datastone-inc/pg_num2int_direct_comp)

`pg_num2int_direct_comp` 保持索引可用性的 numeric/float 与整数精确比较操作符。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_num2int_direct_comp`：

```sql
CREATE EXTENSION pg_num2int_direct_comp;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
-- Example: API passes product ID as numeric parameter
CREATE TABLE products(id int4 PRIMARY KEY, parent int4 REFERENCES products, name text);
PREPARE find_product(numeric) AS SELECT * FROM products WHERE id = $1;

-- Without this extension: sequential scan (casts indexed column)
EXPLAIN (COSTS OFF) EXECUTE find_product(42);
--  Seq Scan on products
--    Filter: ((id)::numeric = '42'::numeric)   ← full table scan!

-- With this extension: index scan (transforms to integer comparison)
EXPLAIN (COSTS OFF) EXECUTE find_product(42);
--  Index Scan using products_pkey on products
--    Index Cond: (id = 42)                     ← uses primary key index
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `float4_cmp_int2` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_cmp_int4` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_cmp_int8` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_eq_int2` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_eq_int4` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_eq_int8` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_ge_int2` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `float4_ge_int4` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 12, 13, 14, 15, 16；不要推断未列出的主版本。
