## 用法

来源：

- [官方 README](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/README.md)
- [扩展控制文件](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.control)
- [扩展 SQL](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding--0.0.1.sql)
- [C 实现](https://github.com/tongxin/mypg_sharding/blob/567526b4292bc4b613b7d7c18ad0b423da1985d6/mypg_sharding.c)

`mypg_sharding` 是一个从 Postgres Pro 的 pg_shardman 工作派生而来的废弃原型。固定的 0.0.1 源码不能组成可安装且内部一致的扩展，因此没有可以安全执行或推荐给最终用户的分片流程。

### 预期流程

源码的设计意图是预加载动态库，在固定的 `mypg` 模式中创建元数据，定义节点组，通过 `postgres_fdw` 注册节点，创建哈希分区并进行再平衡。可见配置如下：

```conf
shared_preload_libraries = 'mypg_sharding'
mypg.is_master = on
mypg.node_name = 'node1'
```

计划中的 SQL 接口包括 `create_nodegroup`、`register_node`、`rebalance_nodegroup` 与 `create_hash_partitions`，以及内部连接、广播、复制和分区替换辅助函数。这些名称只说明未完成的设计意图，不是受支持 API。

### 源码阻塞缺陷

安装脚本包含无效 PL/pgSQL 语法，创建了 `nodegroup`、`node`、`tableinfo`，后面却引用不存在的 `mypg.nodeinfo` 等对象，还残留无法解析的 `shardman` 名称。脚本还调用类似 pg_pathman 的分区辅助函数，却没有声明相应依赖，而控制文件只列出 `postgres_fdw`。因此若不重新设计和修复，固定源码上的 `CREATE EXTENSION` 无法完成。

C 代码依赖 2018 年前后的服务器内部接口，也没有维护中的兼容与升级路径。不要在生产集群中预加载该动态库或暴露其 SQL。未来若要恢复开发，应先统一模式，声明所有依赖，加入确定性的安装和升级测试，定义认证与失败语义，并针对每个受支持 PostgreSQL 大版本测试路由、再平衡、备份、故障转移与恢复。
