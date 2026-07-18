## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/worden341/pgchronos/blob/16731fb0cd98a915ef8732523dffddbef09f49ce/README.md)
- [1.0.2 版本 SQL 对象](https://github.com/worden341/pgchronos/blob/16731fb0cd98a915ef8732523dffddbef09f49ce/sql/pgchronos--1.0.2.sql)
- [PGXN 发行版](https://pgxn.org/dist/pgchronos/)

`pgchronos` 是在 `daterange` 与 `tstzrange` 数组上执行集合操作的纯 SQL 演示。它提供并集、交集、差集、归并、包含函数，以及对应的 `+`、`*`、`-`、`<@` 和 `@>` 操作符。

```sql
CREATE EXTENSION pgchronos;
SELECT ARRAY[daterange('2026-01-01', '2026-02-01')]
       - ARRAY[daterange('2026-01-10', '2026-01-15')];
```

结果是互不重叠的剩余区间数组。并集和 `reduce` 合并重叠或相邻区间，交集则保留公共片段。应针对精确范围子类型验证边界包含性、空与无限边界、空值、顺序、重复、夏令时行为和规范化。

上游明确纠正了早期“stable”标签：整个 1.x 系列都是 alpha 质量概念演示，不适合生产，仓库也已放弃。应优先使用受支持 PostgreSQL 版本中持续维护的 multirange 功能，或经过彻底审计的应用 SQL。保留旧数据时，应为代数恒等式建立属性测试，并谨慎迁移依赖操作符的表达式，因为自定义操作符会全局改变查询含义。
