## 用法

来源：

- [项目 README 与兼容矩阵](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/README.md)
- [扩展 control 文件](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/pg_pathman.control)
- [Hook 实现](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/src/hooks.c)
- [分区函数](https://github.com/postgrespro/pg_pathman/blob/09c0afe25154a034f5ae9cf5b15f977ce8a7a1e9/src/pl_funcs.c)

`pg_pathman` 1.5 提供哈希与范围分区管理、行路由、运行时分区选择和数据迁移 worker。上游声明项目已停止开发，并建议使用 PostgreSQL 原生分区。其文档支持上限为 PostgreSQL 15，且多个版本需要打补丁的内核构建；所有新设计都应使用原生分区。

### 加载现有部署

必须预加载该库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'pg_pathman'
```

把它安装到不允许非受信角色创建对象的专用模式：

```sql
CREATE SCHEMA pathman;
REVOKE CREATE ON SCHEMA pathman FROM PUBLIC;
GRANT USAGE ON SCHEMA pathman TO PUBLIC;
CREATE EXTENSION pg_pathman WITH SCHEMA pathman;

SELECT pathman.pathman_version();
```

使用干净模式很重要，因为未精确匹配函数签名的调用可能受到恶意重载攻击。应检查扩展模式权限和调用者的 `search_path` 设置。

### 迁移与退役风险

`create_hash_partitions()` 和 `create_range_partitions()` 等函数会改变表结构；启用 `partition_data` 时还会复制行，并持锁直至提交。`partition_table_concurrently()` 用短事务分批移动数据，但仍会修改在线数据且需要运行监控。执行任何转换前，应在恢复副本上演练约束、索引、触发器、序列、外键、复制、回退和应用行为。

该模块使用规划器、执行器、utility 和共享内存 hook。上游特别警告 `ProcessUtility_hook` 顺序可能与 `pg_stat_statements` 等扩展相互干扰；预加载顺序不能替代兼容性测试。不要在未列出的 PostgreSQL 大版本上加载此构建。对现有部署，应在操作系统或 PostgreSQL 升级淘汰兼容构建环境之前，规划可度量的声明式分区迁移。
