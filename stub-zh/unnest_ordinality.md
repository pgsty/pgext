## 用法

来源：

- [已复核 commit 的官方文档](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/doc/unnest_ordinality.md)
- [SQL 声明](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/sql/unnest_ordinality.sql)
- [集合返回实现](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/src/unnest_ordinality.c)
- [扩展 control 文件](https://github.com/pgexperts/unnest_ordinality/blob/35d67aee91c14ed57b66d1bc57a22c0a0827fec8/unnest_ordinality.control)
- [官方 PGXN 发行页](https://pgxn.org/dist/unnest_ordinality/)

`unnest_ordinality` 是旧式集合返回函数，把任意 PostgreSQL 数组展开为 `(element_number, element)` 行。它早于 PostgreSQL 内置的 `WITH ORDINALITY` 语法，在受维护的服务器版本上已经弃用。

### 核心流程

```sql
CREATE EXTENSION unnest_ordinality;

SELECT *
FROM unnest_ordinality(ARRAY['a', NULL, 'c']);

-- Preferred on PostgreSQL 9.4 and later:
SELECT ordinality::integer AS element_number, element
FROM unnest(ARRAY['a', NULL, 'c']) WITH ORDINALITY AS u(element, ordinality);
```

扩展函数返回从 1 开始的 `element_number integer` 与 `element anyelement`。它保留 NULL 元素，并按存储顺序扁平化多维数组。编号是连续序号，不是源数组声明的下标；自定义下界与维度坐标不会保留。空数组不返回任何行。

### 弃用与性能边界

C 函数会解构整个源数组，并在返回前把全部输出行物化到 tuplestore。大型或 toasted 数组所需的内存与临时文件容量因此与展开结果成正比。它使用旧式后端元组和数组 API，主要只在 PostgreSQL 9.2 上测试过。

PostgreSQL 9.4 或更高版本应使用内置 `unnest(...) WITH ORDINALITY`。只有固定的旧版兼容需求才应保留 `unnest_ordinality`，并在升级或删除旧二进制文件前迁移 SQL。
