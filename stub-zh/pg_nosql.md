## 用法

来源：

- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/README.md)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/LICENSE)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/pg_nosql.control)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/pg_nosql--1.0.sql)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/meson.build)
- [官方资料](https://github.com/Vern-AllWorks-LLC/alohadb/blob/v21.1.0/contrib/pg_nosql/nosql_query.c)

`pg_nosql` `1.0` 提供JSON 文档集合与查询函数。本条目限定于 AlohaDB 21.1.0 的 PostgreSQL 18 分支，是预览模块；不表示可直接安装在标准 PostgreSQL 上。

### 工作流程

```sql
CREATE EXTENSION pg_nosql;
SELECT doc_create_collection('demo');
SELECT doc_insert('demo', '{"city":"Hilo","age":31}');
SELECT * FROM doc_search('demo', '{"city":"Hilo"}');
```

安装需要超级用户；运行函数还需相应对象权限。

### 运行边界

doc_put、doc_patch 和 doc_remove 修改文档，doc_search 支持文档列出的过滤运算及 sort/field/limit 选项。集合元数据和底层表仍是 PostgreSQL 对象。本 API 不实现 MongoDB 线协议。应审核生成表的权限、模式验证规则和动态过滤条件，不应向应用角色开放无限制的集合管理能力。

此模块对应上游旧名 `alohadb_nosql`，应遵循 AlohaDB 的集群及对象改名流程，不能将其他项目的同名扩展当作替代品。升级前应演练依赖对象、权限和备份恢复。
