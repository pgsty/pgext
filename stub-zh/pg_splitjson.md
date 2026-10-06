## 用法

来源：

- [README.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/README.md)
- [docs/api-reference.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/docs/api-reference.md)
- [sql/pg_splitjson--0.1.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/sql/pg_splitjson--0.1.0.sql)
- [CHANGELOG.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/CHANGELOG.md)
- [pg_splitjson.control](https://github.com/Haiwen-Yin/pg_splitjson/blob/b531bab029fad13bfea3a11ead2e020c68ca1ff3/pg_splitjson.control)

`pg_splitjson` 0.1.0 是面向 PostgreSQL 18 的初始实现，适用于频繁更新 JSON 字段。托管视图呈现完整文档，专用操作可更新已有热字段而不重写冷数据模板。

### 核心用法

```sql
CREATE EXTENSION pg_splitjson;
SELECT splitjson.create_table('public.events', '[["counter"],["state"]]'::jsonb);
INSERT INTO public.events(id, doc)
VALUES (1, '{"counter":1,"state":"new","payload":{"body":"cold"}}');
SELECT splitjson.set_field('public.events', 1, ARRAY['state'], '"ready"'::jsonb);
SELECT splitjson.increment_field('public.events', 1, ARRAY['counter'], 1);
SELECT * FROM public.events;
```

### 运行边界

创建、迁移和索引管理函数由扩展安装者执行。应用只获得业务视图及 API 模式权限，内部存储应保持隔离。`splitjson.set_fields` 批量更新字段，`splitjson.increment_field` 在行锁内递增，`splitjson.get_field` 与 `splitjson.find_ids` 读取字段，`splitjson.create_path_index` 创建 B-tree，`splitjson.drop_table` 删除托管对象。

声明 1–64 条互不重叠的热路径。缺失字段、结构调整或普通的整文档更新可能重新打包文档；索引维护、WAL、行锁和 vacuum 成本仍然存在。普通视图更新遇到并发变化可能返回 `40001`，应用需要重试。

可选的 SELECT 改写要求在规划前通过 `LOAD 'pg_splitjson'` 或会话预加载载入库，并由 `splitjson.enable_query_rewrite` 控制。核心 API 不要求全局预加载或重启。早期同版本原型需要重新安装并逻辑迁移。`splitjson.migrate_table` 将快照复制到新视图，不同步后续写入，也不复制全部约束与权限。
