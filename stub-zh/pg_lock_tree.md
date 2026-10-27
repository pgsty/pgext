## 用法

来源：

- [README](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/README.md)
- [Control](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/pg_lock_tree.control)
- [Install and grants](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/INSTALL.md)
- [Install SQL](https://github.com/apjapj/pg_lock_tree/blob/c5f6f8f3f2dbd94505f9f85a469946c8e9c932e3/pg_lock_tree--1.5.1.sql)

`pg_lock_tree` 1.5.1 是展示锁等待树和根阻塞者的纯 SQL 扩展，上游测试 PostgreSQL 14–19rc1。需要超级用户安装，不含动态库，不要求预加载或重启。此源码修订比仓库最新标签版本更新。

### 查看等待队列

```sql
CREATE EXTENSION pg_lock_tree;
SELECT * FROM lock_tree_text();
SELECT * FROM lock_tree_roots();
SELECT command FROM lock_tree_command();
```

命令辅助函数只返回待审阅 SQL，不执行它们，不应将完整输出交给自动执行命令。预备事务可能在没有活动后端的情况下阻塞会话，无法靠终止后端解决。wide 与 summary 读取器适用于宽队列，采样历史与 NOTIFY 告警为可选功能。

### 权限与维护

监控角色同时需要扩展的对象、函数授权和 pg_read_all_stats；单独 pg_monitor 不够。历史读取还需额外授权。采样、告警、保留和归档维护需要外部调度，扩展没有内置 worker。历史中的查询文本可能包含敏感字面量，更新已有安装前应审阅升级 SQL 和安装授权矩阵。
