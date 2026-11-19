## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/pg_context.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/pg_context--1.1.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_context/meson.build)

`pg_context` `1.1` 提供实体消歧与上下文相关排序。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_ontology;
CREATE EXTENSION pg_context;
SELECT context_register('Apple', 'customer:42', 'customer', NULL, 100, 'sales', ARRAY['invoice']);
SELECT * FROM context_resolve('Apple', 'Apple invoice');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

本扩展依赖 AlohaDB 的 pg_ontology 模块。实体档案、别名、关键词和本体线索共同决定消歧评分，context_rerank 可调整候选排序。这是规则匹配，不是训练得到的语义模型。应审核词典写入和 public 模式权限；1.1 修改了实体档案键以支持多个词义。

此模块对应上游旧名 `alohadb_context`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
