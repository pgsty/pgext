## 用法

来源：

- [pg_rational v0.0.3 README](https://github.com/begriffs/pg_rational/blob/v0.0.3/README.md)
- [pg_rational v0.0.3 控制文件](https://github.com/begriffs/pg_rational/blob/v0.0.3/pg_rational.control)
- [截至 v0.0.3 的变更](https://github.com/begriffs/pg_rational/compare/v0.0.2...v0.0.3)

`pg_rational` 通过固定的 64 位 PostgreSQL 类型提供精确分数运算。对于必须保持精确的数值，以及需要在现有位置之间插入新位置而无需重新编号整张表的用户自定义行排序，可以使用 `rational`。

### 精确运算

```sql
CREATE EXTENSION pg_rational;

SELECT 1::rational / 3 * 3 = 1;
SELECT '1/3'::rational + '2/7'::rational;
SELECT rational_simplify('36/12');
```

扩展会检测算术溢出，而不是静默回绕。`ratt` 是用于元组强制转换的辅助类型：

```sql
SELECT 1 + (i, i + 1)::ratt
FROM generate_series(1, 5) AS i;
```

整数值、浮点值与有理数之间可以相互转换。浮点数转换会寻找有理数近似值；将有理数转换成浮点数则会失去精确性。

### 稳定的用户自定义排序

```sql
CREATE SEQUENCE todos_seq AS integer;

CREATE TABLE todos (
  prio rational UNIQUE DEFAULT nextval('todos_seq')::integer,
  what text NOT NULL
);

INSERT INTO todos (what)
VALUES ('install extension'), ('read about it'), ('try it');

UPDATE todos
SET prio = rational_intermediate(1, 2)
WHERE what = 'try it';

SELECT * FROM todos ORDER BY prio;
```

请使用 `integer` 序列，并显式转换 `nextval()`。该扩展有意不提供从 `bigint` 到 `rational` 的隐式转换，因为其分子受 PostgreSQL `integer` 范围限制。

### 索引、聚合与注意事项

- `rational` 支持 btree 和 hash 操作符类，因此可用于排序索引和等值索引。
- 除算术及比较操作符外，该扩展还提供 `min(rational)`、`max(rational)` 和 `sum(rational)` 聚合函数。
- `rational_intermediate(lower, upper)` 沿 Stern-Brocot 树查找两个参数之间的分数。范围极窄时耗时会更长，而 v0.0.3 没有最大深度参数；在没有语句超时保护的情况下，不要向攻击者开放由其控制的病态边界。
- 只有当算术运算保持在该类型的分子和分母限制内时，值才是精确的。应处理溢出错误，不要静默回退为浮点数。
- 版本 0.0.3 主要是构建兼容性和文档发行版；面向用户的有理数运算接口保持稳定。
