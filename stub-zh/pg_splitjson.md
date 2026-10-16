## 用法

来源：

- [README.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/README.md)
- [docs/api-reference.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/api-reference.md)
- [docs/support-matrix.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/support-matrix.md)
- [docs/production-runbook.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/docs/production-runbook.md)
- [sql/pg_splitjson--0.3.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/sql/pg_splitjson--0.3.0.sql)
- [sql/pg_splitjson--0.2.0--0.3.0.sql](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/sql/pg_splitjson--0.2.0--0.3.0.sql)
- [pg_splitjson.control](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/pg_splitjson.control)
- [CHANGELOG.md](https://github.com/Haiwen-Yin/pg_splitjson/blob/9982e76f8ebf6bc7e08a7efc10d240f21cc76b5a/CHANGELOG.md)

`pg_splitjson` 0.3.0 适用于频繁更新 JSON 字段；上游验证环境为 PostgreSQL 18.6、Linux x86_64，未承诺其他小版本和架构。托管视图呈现完整文档，专用操作可更新已有热字段而不重写冷数据模板。

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

可选的 SELECT 改写要求在规划前通过 `LOAD 'pg_splitjson'` 或会话预加载载入库，并由 `splitjson.enable_query_rewrite` 控制。核心 API 不要求全局预加载或重启。正式版本支持 0.1.0 → 0.2.0 → 0.3.0 升级；无法识别的早期原型仍需逻辑迁移。`splitjson.migrate_table` 将快照复制到新视图，不同步后续写入，也不复制全部约束与权限。

### 条件更新与访问控制

`splitjson.row_revision` 返回 `epoch:id:revision` 令牌；`splitjson.set_fields_if_revision` 可拒绝过期写入并返回 `40001`。`splitjson.apply_patch` 支持 RFC 6902（不支持删除根文档），`splitjson.merge_patch` 支持 RFC 7396，`splitjson.upsert_document` 用于整文档写入。固定非负数组下标可用于热路径；普通视图不支持 ON CONFLICT。

`splitjson.enable_rls` 与 `splitjson.set_policy` 在私有存储上按调用者执行策略，应用继续通过业务视图访问。策略表达式不能引用冷数据、完整文档、子查询或用户自定义函数。`splitjson.create_partitioned_table` 仅支持按标识列进行 HASH/RANGE 分区。

### 升级、备份与恢复

先备份并在各节点安装匹配的共享库，随后执行 `ALTER EXTENSION pg_splitjson UPDATE TO '0.3.0'`。已载入旧库的会话需要重建。应备份整个数据库；只导出业务视图不会包含私有存储。

恢复备份、PITR 或接管后，写入前调用 `splitjson.reinitialize_table`，按实际事件选择 `restore`、`recovery` 或 `takeover`，再运行 `splitjson.check_all(true)`。这会轮换版本令牌的纪元。在线迁移使用独立的 WAL 流程；原生发布/订阅只支持相同布局，接管需要隔离旧写入端，不支持双主写入。
