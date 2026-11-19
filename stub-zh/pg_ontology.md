## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/pg_ontology.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/pg_ontology--1.1.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_ontology/meson.build)

`pg_ontology` `1.1` 提供本体三元组、层次遍历与历史。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_ontology;
SELECT ontology_assert('Honolulu', 'capital_of', 'Hawaii');
SELECT * FROM ontology_query('Honolulu', NULL, NULL);
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

ontology_subclass 管理经过循环检查的类型层次，ontology_related 提供有界图遍历，ontology_history 记录断言和撤回。这是关系式三元组存储，不是通用 OWL 推理引擎。应控制本体表写入并检查历史保留策略。1.1 提供匹配的安装、升级脚本，增加双向遍历和变更跟踪。

此模块对应上游旧名 `alohadb_ontology`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
