## 用法

来源：

- [官方 README](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/README.md)
- [扩展控制文件](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres.control)
- [扩展 SQL](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres--0.1.sql)
- [必需的 PostgreSQL 核心补丁](https://github.com/danolivo/pargres/blob/d62ebdff2f9758f1b444639dc5cf65fd1edc7e4f/pargres_core_changes.diff)

`pargres` 是一个 2018 年的研究原型，用于在无共享 PostgreSQL 节点间执行并行查询。它不是即装即用的扩展：所核对的源码必须配合定制 PostgreSQL 核心补丁，而且明确既不提供全局快照，也不提供分布式写入提交。

### 实验设置

应在一次性 PostgreSQL 源码树上应用随附核心补丁，在所有节点安装完全匹配的二进制，预加载模块，并配置顺序一致的主机和端口列表。源码公开的设置如下：

```conf
shared_preload_libraries = 'pargres'
pargres.node = 1
pargres.nnodes = 2
pargres.hosts = 'node1,node2'
pargres.ports = '5432,5432'
pargres.eports = '1000,1000'
```

```sql
CREATE EXTENSION pargres;
SELECT * FROM relsfrag;
```

`relsfrag` 记录关系名、分片属性编号与分片函数标识。SQL 还暴露 `isLocalValue(text,integer)` 和内部协调辅助函数 `set_query_id(integer,integer)`。规划和执行钩子会加入自定义交换节点，并同时使用普通 PostgreSQL 连接和独立交换套接字。

### 原型限制

设计声明的范围只有只读查询。没有集群级快照时，各节点可能观察到不同数据库状态；没有分布式提交时，写入无法保持原子，只能在该原型之外顺序执行。主机连接仅由主机名和端口拼接，交换流量使用原始套接字，因此认证、加密、防火墙与网络故障处理都由部署方负责。

所需补丁会改变哈希连接与嵌套循环执行器行为，并针对历史服务器源码树。未经完整代码审计，不要把它应用到受支持的生产 PostgreSQL 构建。项目没有维护中的升级、故障转移、恢复、安全或兼容保证。`pargres` 只应放在可复现实验环境中，并在每次实验中验证补丁服务器、计划正确性、节点顺序、端口连通性与结果一致性。
