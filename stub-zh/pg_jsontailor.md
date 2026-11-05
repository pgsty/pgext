## 用法

来源：

- [上游来源 1](https://github.com/baremonconsulting/pg_jsontailor/blob/2e884e1cac72cbd89aab0fa0293db0d77885768c/README.md)
- [上游来源 2](https://github.com/baremonconsulting/pg_jsontailor/blob/2e884e1cac72cbd89aab0fa0293db0d77885768c/pg_jsontailor.control)

`pg_jsontailor` 开发版本 0.2.0 在关系表之上生成带乐观并发检查的可写 JSON 视图。

### 核心流程

`jsontailor.duality_view_ddl` 返回 DDL，便于执行前审查。支持的形态是一个根表、一个可选的只读多对一对象，以及一个可选的可写一对多数组。更新会完整替换文档，必须携带先前读取的 ETag；受保护的删除同样检查该标记。

```sql
CREATE EXTENSION pg_jsontailor;
CREATE SCHEMA shop;
CREATE SCHEMA api;
CREATE TABLE shop.orders (order_id bigint PRIMARY KEY, status text);
SELECT jsontailor.create_duality_view(
  'api.orders_dv',
  '{"root":{"table":"shop.orders","key":"order_id",
    "fields":{"status":"status"},"insert":true,"update":true,"delete":true}}'::jsonb
);
SELECT data FROM api.orders_dv;
```

### 运行边界

这是面向 PostgreSQL 17–19 的未发布评估及受控试用版本，以 SQL 和 PL/pgSQL 实现，无需预加载。生成的访问路径遵循调用者权限和行级安全，创建视图不会赋予更广泛的表访问权限。生成 DDL 不依赖扩展运行时，且不构成完整的 Oracle 双重性视图或 ORDS 兼容层。
