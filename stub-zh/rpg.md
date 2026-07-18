## 用法

来源：

- [上游 README](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/README.md)
- [扩展 SQL](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/rpg--0.1.sql)
- [扩展 control 文件](https://github.com/ryapric/rpg/blob/76c6a508ac04fca27ebb0ea3ffc675e7a3fb6442/rpg.control)

`rpg` 版本 `0.1` 是一个小型 SQL/PLpgSQL 原型，包含三个工具函数。`year_month` 用于格式化日期，`letters` 则返回小写或大写英文字母表：

```sql
CREATE EXTENSION rpg;

SELECT year_month(DATE '2026-07-16');
SELECT letters('upper');
```

### 原型限制

第三个函数 `bind_rows(output_table, table_a, table_b)` 尝试通过两个指定表的 `UNION ALL` 创建新表。其动态 SQL 直接拼接表名和列名，未做标识符引用，而且输入关系仍须满足 union 兼容性。绝不能将不可信名称传给 `bind_rows`；实际使用前应审查并修补实现。上游 README 没有提供 API、兼容性或支持文档。
