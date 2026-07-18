## 用法

来源：

- [扩展 control 文件](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/saio.control)
- [上游警告与配置](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/doc/saio.md)
- [规划器钩子实现](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/src/saio.c)
- [PGXN 包元数据](https://github.com/wulczer/saio/blob/e6abbd9f0886e4067b8c42701a12b6e2589b641d/META.json)

`saio` `0.0.1` 版是 2011 年的实验性规划器模块，通过模拟退火搜索连接顺序。它没有扩展 DDL 文件：安装兼容共享库后，会话用 `LOAD 'saio'` 激活规划器钩子。加载后，其 `saio` 开关默认启用，而 `saio_seed` 及若干温度、平衡、移动和重算设置用于控制搜索。

### 只在隔离的旧版构建中检查

```sql
LOAD 'saio';
SET saio = off;
SET saio_seed = 0.5;
SET saio = on;

EXPLAIN
SELECT n.nspname, c.relname
FROM pg_namespace AS n
JOIN pg_class AS c ON c.relnamespace = n.oid;

SET saio = off;
```

上游文档开篇就明确警告不要使用该模块，并称其优化器质量很差。代码面向 2011 年的 PostgreSQL 内部接口，不能假定兼容现代服务器；加载后还会改变整个后端会话的规划行为。不要在生产环境部署，也不要把它加载到承载有价值工作的会话。若历史研究确有需要，应使用符合其年代的可丢弃服务器，在受控查询前保持 `saio` 禁用，记录执行计划与崩溃，并在完成后终止会话。
