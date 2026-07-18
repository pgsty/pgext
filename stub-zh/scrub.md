## 用法

来源：

- [项目 README 与检查范围](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/README.md)
- [扩展 control 文件](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub.control)
- [1.0 版 SQL API](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub--1.0.sql)
- [后台 worker 实现](https://github.com/tvondra/scrub/blob/8d39b4a2c7419715c9b7726c5d7b478ee5e0a9c6/scrub.c)

`scrub` 1.0 启动后台 worker，读取数据库页面并检查损坏。启用数据校验和时，它可以验证页面校验和，还会检查通用页头边界、heap 页面和元组（包括 TOAST 值）以及 B-tree 页面和元组。它报告发现，但不会修复损坏。

### 预加载并启动限速检查

加入该库并重启 PostgreSQL：

```conf
shared_preload_libraries = 'scrub'
```

安装扩展，启动一个数据库的检查并监控计数器：

```sql
CREATE EXTENSION scrub;

SELECT scrub_start(
  dbname => 'appdb',
  cost_delay => 10,
  cost_limit => 1000,
  reset_status => true
);

SELECT * FROM scrub_status();
SELECT scrub_is_running();
```

成本设置以类似 vacuum 成本控制的方式限速。整个集群同一时间只能运行一个 scrub。`scrub_stop()` 请求终止，`scrub_reset()` 清除计数器；否则状态会跨多次运行累积。

### 检测边界与事件处理

worker 会读取所有选中页面且不会主动写入，但仍可能产生大量存储 I/O、占用缓存，并与备份、vacuum、备库和前台查询竞争。应确定 I/O 预算，监控延迟和复制滞后，并把首次运行安排在非高峰时段。

检测失败的详情写入 PostgreSQL 服务端日志。应保留这些日志和存储证据，停止可避免的写入，并遵循经过测试的损坏响应流程；不能假设一个失败计数器已经识别唯一受损对象。检查器不会覆盖所有访问方法、不会把每个 heap 行与每个索引交叉核对，也不能证明不存在潜在损坏。其底层页面代码没有当前 PostgreSQL 兼容矩阵，因此生产使用前必须测试准确大版本、校验和设置、存储栈和备份流程。
