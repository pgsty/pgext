## 用法

来源：

- [Official usage documentation](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/doc/pg_cache_tree.md)
- [Version 1.0.0 SQL implementation](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/sql/pg_cache_tree--1.0.0.sql)
- [Official regression workflow](https://github.com/ozum/pg_cache_tree/blob/95591ed10a1eebfaf2736a9863bd71fe550d7b1b/test/sql/base.sql)

`pg_cache_tree` 1.0.0 为邻接表树维护物化的祖先和后代 ID 数组。行触发器会在插入、删除以及节点 ID 或父 ID 变化后更新数组。它适合频繁读取完整祖先或后代的场景，代价是写放大和非规范化状态。

### 准备树表

文档定义且最安全的布局使用默认列名：

```sql
CREATE EXTENSION pg_cache_tree;

CREATE TABLE business_unit (
  id integer PRIMARY KEY,
  "parentId" integer REFERENCES business_unit(id)
    ON DELETE CASCADE ON UPDATE CASCADE,
  "parentsCache" integer[] NOT NULL DEFAULT '{}',
  "childrenCache" integer[] NOT NULL DEFAULT '{}'
);

SELECT ct_create_trigger('business_unit');
```

之后插入和移动节点会同步更新两个方向的缓存：

```sql
INSERT INTO business_unit (id, "parentId") VALUES
  (1, NULL), (2, 1), (20, 2), (21, 2);

SELECT id, "parentId", "parentsCache", "childrenCache"
FROM business_unit
ORDER BY id;
```

使用 `ct_drop_trigger('business_unit')` 删除生成的两个触发器。

### 主要对象

- `ct_create_trigger(regclass, id, parent, parents, children)`：创建 INSERT/DELETE 与 UPDATE 触发器。
- `ct_trigger_update_cache()`：维护祖先和后代数组的触发器函数。
- `ct_drop_trigger(regclass, parent)`：删除 `ct_create_trigger` 创建的触发器。
- `ct_insert_array` 与 `ct_subtract_array`：内部数组辅助函数；`-#` 操作符暴露减法。
- `ct_assign` 与 `ct_copy_value`：触发器使用的动态复合字段辅助函数。

### 完整性与版本边界

1.0.0 的可配置列接口背后存在实现缺陷。触发器仍会按字面量读取 `NEW."parentId"`，因此即使向 `ct_create_trigger` 传入自定义父列也可能失败。标量重载 `ct_insert_array(anyarray, anyelement, anyelement)` 还会调用无关的 `extra_modules.t_cpc_insert_array` 函数。应使用文档默认列、避免该辅助重载，并对任何本地修复运行回归测试。

扩展不会阻止环。自引用外键只能保证父行存在，不能保证图仍是一棵树。启用缓存维护前，应添加应用校验或约束触发器进行环检测。

移动或删除高层节点会在同一事务中更新大量祖先和后代行。应按真实深度和扇出评估写入、WAL 与锁开销。禁用触发器、失败的批量加载或直接修改缓存列都可能使数组过期，而上游没有提供重建命令；必须保留规范的 `id`/`parentId` 数据源和经过测试的对账流程。
