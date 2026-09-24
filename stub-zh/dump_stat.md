## 用法

来源：

- [Official documentation](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/dump_stat--1.0.sql)
- [Control file](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/anyarray_elemtype.control)
- [Version 1.0 SQL](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/dump_stat--1.0.sql)
- [Native helper](https://github.com/postgrespro/postgrespro/blob/9086b754eb42fc0ee3481a3016360f31a27f2271/contrib/dump_stat/anyarray_elemtype.c)

`dump_stat` 1.0 将历史 Postgres Pro 9.5 分支的优化器统计信息导出为 SQL。本条目只描述该分支的固定源码版本，不表示兼容当前 PostgreSQL 版本。

### 核心流程

安装匹配分支的扩展文件后，具有相应权限的管理员可以生成并检查语句：

```sql
CREATE EXTENSION dump_stat;
SELECT * FROM dump_statistic('public'::text);
```

结果是 SQL 文本，不会自动执行恢复。应先审查，再用于兼容的测试环境。生成的语句会更新或插入目标库的 `pg_catalog.pg_statistic`，并按名称解析表、类型、列和操作符标识。

### 函数与边界

`dump_statistic()` 导出系统模式之外可用的统计信息；`dump_statistic(schema_name text)` 限定模式；`dump_statistic(schema_name text, table_name text)` 和 `dump_statistic(relid oid)` 选择单个关系。其他辅助函数用于解析限定对象名，`anyarray_elemtype()` 由共享库实现。

安装包含 C 模块及 PL/pgSQL 例程，不需要预加载。读取原始目录统计以及执行生成的目录写入，都需要相应的管理权限；输出中可能含有敏感的采样值。目标对象和目录布局必须匹配，源码固定了历史分支的统计布局，因此不能将结果当作可跨版本迁移的通用转储，也不能替代正常备份与 ANALYZE。
