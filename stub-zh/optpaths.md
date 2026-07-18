## 用法

来源：

- [上游 README](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/README.md)
- [0.1 版安装 SQL](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/optpaths--0.1.sql)
- [规划器钩子实现](https://github.com/danolivo/anotherSelfJoin/blob/b4461d76e07fb83e1a18b7e6cee8953b6d0c27c4/optpaths.c)

optpaths 0.1 是一个实验性规划器钩子，试图在唯一性条件允许时消除一小类冗余自连接。

### 隔离评估

在服务器启动时加载动态库，重启后仅在可丢弃数据库中测试：

```conf
shared_preload_libraries = 'optpaths'
```

```sql
CREATE EXTENSION optpaths;
EXPLAIN SELECT p.* FROM test AS p
JOIN (SELECT * FROM test WHERE b = 0) AS q ON p.a = q.a;
SET remove_self_joins = off;
EXPLAIN SELECT p.* FROM test AS p
JOIN (SELECT * FROM test WHERE b = 0) AS q ON p.a = q.a;
```

在考虑实际工作负载前，应比较钩子启用与禁用时的计划和结果。

### 注意事项

- 安装 SQL 会创建一个通用名称的 test 表并插入两行。安装可能与已有表冲突，也会修改选定的扩展模式。
- CREATE EXTENSION 不会加载规划器动态库。要可靠激活，必须预加载并重启服务器。
- 钩子会修改内部规划器结构，且只覆盖有限的连接形态。没有当前主版本兼容矩阵或证明结果等价的回归套件。
- remove_self_joins 与 log_level 都可由用户设置。详细日志可能淹没客户端或服务器日志。
- 必须隔离测试精确查询语义、预备计划、并行计划、分区、行级安全与升级。未经审查，不得在集群范围启用该原型。
