## 用法

来源：

- [固定提交的上游 README](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/README.md)
- [1.0 版安装 SQL](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/pg_foobar--1.0.sql)
- [固定提交的 parallel worker 实现](https://github.com/MasahikoSawada/pg_foobar/blob/1cf8a2ea6c701726b14241a16d95ece30e643ad8/pg_foobar.c)

pg_foobar 1.0 版是一个教学用途的 dynamic parallel worker 示例。它提供一个函数来启动指定数量的 worker；每个 worker 都会向 PostgreSQL server log 写入一串 foo 和 bar 消息。函数本身始终返回 null text。

### 仅用于测试的示例

只在一次性 PostgreSQL 实例上使用很小的数值，并查看 server log，而不是查询结果：

```sql
CREATE EXTENSION pg_foobar;

SELECT pg_foobar(2, 3, 4);
```

这会请求两个 worker，每个分别记录三条 foo 和四条 bar 消息。可用 worker slot、服务器日志配置和其他 background worker 的并发使用情况决定它能否启动。

### API 缺陷与兼容性

安装 SQL 把第一个参数声明为 regclass，但 C 函数却把它读取为普通 32 位 worker count；README 仍然传入 integer。这个 catalog/API 不一致会造成容易误解的类型转换、显示和函数解析行为；应检查实际安装的 signature，绝不能把关系名或源于 OID 的值当作 worker count 传入。

该扩展可以占用大量 dynamic worker slot 并产生大量日志，因此不要向普通用户开放函数。应撤销 public execute，并在任何演示 wrapper 中限制全部三个参数。

上游只记录了 PostgreSQL 9.6 和 10 beta 1。源码依赖旧版 parallel worker 内部接口，自 2017 年后没有更新，也没有当前兼容矩阵，尽管仓库并未归档。它是示例代码，不是生产日志、调度或监控设施。
