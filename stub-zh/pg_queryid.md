## 用法

来源：

- [已复核 commit 的 README](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/README.md)
- [扩展 control 文件](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid.control)
- [SQL 函数定义](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid--0.0.1.sql)
- [指纹实现](https://github.com/rjuju/pg_queryid/blob/7ea79d8a79c51c049b48d19364ce7b5502cdb7d9/pg_queryid.c)

`pg_queryid` 是一个面向 PostgreSQL 14 及以上版本的外部查询指纹概念验证模块。预加载后，它会替代核心 query ID 计算，并可按名称而不是 OID 对引用对象生成指纹，还可选择忽略 schema 或涉及临时表的查询。创建扩展还会提供单语句函数 `pg_queryid(text) returns bigint`。

### 配置与查询

该模块必须位于 `shared_preload_libraries` 的最后一项；上游还要求在替代 hook 生效时禁用核心 `compute_query_id`。

```conf
shared_preload_libraries = 'pg_stat_statements,pg_queryid'
compute_query_id = off
pg_queryid.use_object_names = on
pg_queryid.ignore_schema = off
pg_queryid.ignore_temp_tables = off
```

重启 PostgreSQL 后：

```sql
CREATE EXTENSION pg_queryid;

SELECT pg_queryid('SELECT * FROM public.orders WHERE id = 42');
```

### 注意事项

- 上游把版本 `0.0.1` 标为尚未达到生产就绪水平的概念验证。已复核源码最后更新于 2022 年，除 PostgreSQL 14 及以上版本外，没有当前兼容性声明。
- 按名称生成指纹可能带来显著开销。启用 `pg_queryid.ignore_schema = on` 后，不同 schema 中的同名对象会有意共享一个指纹。
- 修改任意指纹选项都会改变消费者看到的 query ID。修改后应重置 `pg_stat_statements` 以及其他按 query ID 索引的状态。
- 启用 `pg_queryid.ignore_temp_tables = on` 后，引用临时关系的语句会被故意排除在该模块的指纹之外，并可能从依赖 query ID 的监控中消失。
- 自动替代需要预加载并重启。仅执行 `CREATE EXTENSION` 只会安装 SQL 辅助函数，不会安装 parser hook。
