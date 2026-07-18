## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/README.md)
- [扩展 control 文件](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/inetrange.control)
- [1.0.2 版本 SQL 对象](https://github.com/pjungwir/inetrange/blob/9d62352afda4b84be01d2fb3375805a04a64995f/inetrange--1.0.2.sql)
- [当前 PGXN 发行版](https://pgxn.org/dist/inetrange/)

`inetrange` 增加以 `inet` 为子类型的 PostgreSQL 范围类型。它支持普通范围包含与重叠运算符、GiST 索引，以及用于禁止 IP 地址区间重叠的排斥约束。

```sql
CREATE EXTENSION inetrange;
SELECT '1.0.1.1'::inet <@ inetrange('1.0.0.0', '1.1.0.0', '[]');
CREATE TABLE network_block (
  addresses inetrange NOT NULL,
  owner_name text NOT NULL,
  EXCLUDE USING gist (addresses WITH &&)
);
```

边界遵循 PostgreSQL 范围规则，子类型则遵循 `inet` 对 IPv4 与 IPv6 值的排序。应明确选择包含或排除端点，并在把等值关系作为业务键前验证应用的规范表示。

1.0.2 有当前源码与 PGXN 发行版，但没有明确的现代 PostgreSQL 兼容矩阵。存储长期数据前，应在目标大版本上测试 GiST 索引构建、排斥约束、混合地址族、转储恢复与升级。
