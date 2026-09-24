## 用法

来源：

- [README](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/README.md)
- [Control file / 控制文件](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/nabla.control)
- [Cargo.toml](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/Cargo.toml)
- [src/lib.rs](https://github.com/tomaquet18/nabla/blob/3c35e44e60f29b9f18bc8412a69120d1bda373f9/src/lib.rs)

`nabla` 通过逻辑 WAL 异步维护视图，并向订阅者提供有序增量。审核的 0.1.0 源码面向 PostgreSQL 17，仍属于实验性实现；视图内容可能落后于已提交写入。

### 启用

配置 `wal_level = logical`，将 `nabla` 加入 `shared_preload_libraries`，将 `nabla.database` 设为已存在的数据库，并预留复制槽。重启后以超级用户安装。工作进程只维护配置指定的一个数据库。

```sql
CREATE EXTENSION nabla;
CREATE TABLE orders (id bigserial PRIMARY KEY, k int, amount numeric, status text);
ALTER TABLE orders REPLICA IDENTITY FULL;
SELECT nabla.create_view('orders_by_k',
  'SELECT k, count(*) AS n, sum(amount) AS total FROM orders WHERE status = ''paid'' GROUP BY k');
INSERT INTO orders(k, amount, status) VALUES (1, 10, 'paid');
SELECT nabla.await_ready('orders_by_k');
SELECT nabla.wait_for('orders_by_k', pg_current_wal_lsn());
SELECT * FROM orders_by_k;
```

### 订阅与恢复

`nabla.current_seq` 返回增量游标，`nabla.changes` 按世代读取指定序号之后的变更。通知只是唤醒信号，不是持久投递机制。订阅者必须遵循快照与游标协议；世代改变或历史已被清理时，需重新读取快照。`nabla.refresh` 重建视图，`nabla.drop_view` 删除注册。

只应使用固定源码接受的查询形式，并监控工作进程错误和视图状态。上述简单过滤与分组示例避开了更复杂的连接限制。复制槽可能保留大量 WAL；超过配置的保留边界后，视图可能失效并需要重建。基础表的复制标识必须提供增量维护所需的旧值。
