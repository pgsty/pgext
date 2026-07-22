## 用法

来源：

- [官方上游文档](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/README.md)
- [官方扩展 SQL](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/time_for_keys--0.0.1.sql)
- [官方覆盖聚合实现](https://github.com/pjungwir/time_for_keys/blob/d63e5d2884bb91151bcd28cf36c57c199086544b/completely_covers.c)

`time_for_keys` 0.0.1 在有效时间表之间实现基于触发器的外键。相同整数键的父表行只要合起来覆盖子表的完整 `tstzrange`，子引用就有效，即使没有单个父表行能独自覆盖它。它是面向状态时间历史的早期原型，不是系统时间审计或通用 SQL 时态约束实现。

### 定义时态关系

父表历史通常需要不重叠的排他约束。创建表后，安装扩展并创建四个约束触发器：

```sql
CREATE EXTENSION btree_gist;
CREATE EXTENSION time_for_keys;

CREATE TABLE houses (
  id integer NOT NULL,
  valid_at tstzrange NOT NULL,
  EXCLUDE USING gist (id WITH =, valid_at WITH &&)
);

CREATE TABLE rooms (
  id integer NOT NULL,
  house_id integer,
  valid_at tstzrange NOT NULL
);

SELECT create_temporal_foreign_key(
  'rooms_have_a_house',
  'rooms',  'house_id', 'valid_at',
  'houses', 'id',       'valid_at'
);
```

函数先验证现有子表行，再为子表 INSERT 或 UPDATE 和父表 UPDATE 或 DELETE 创建延迟约束触发器。违规使用外键 SQLSTATE `23503`。

### 重要对象

- `completely_covers(tstzrange, tstzrange)` 是聚合函数，用于测试输入时段是否合起来完整覆盖固定目标范围。
- `create_temporal_foreign_key(text, text, text, text, text, text, text)` 创建并验证时态关系。
- `drop_temporal_foreign_key(text, text, text)` 删除它生成的四个触发器。

### 原型限制

实现只支持整数键、一个键列和 `tstzrange`。表名不能包含模式。它没有实现 CASCADE、SET NULL、SET DEFAULT、MATCH 模式、复合键、其他范围类型、视图或跟踪已创建关系的目录。触发器名称包含所给约束名，并受 PostgreSQL 标识符长度上限约束。

检查使用动态 SQL，不具备内置外键成熟的锁定、计划缓存、DDL 依赖管理或并发分析。应测试并发父子变更、重叠与相邻范围、无限或空边界、延迟约束行为、表重命名、转储与恢复以及触发器清理。只有在应用事务隔离级别下验证准确语义后才能使用。
