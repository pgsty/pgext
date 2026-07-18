## 用法

来源：

- [固定提交的扩展 control 文件](https://github.com/MasahikoSawada/debug_funcs/blob/8508ef191735ef48375a2e626be1986af4d35dd8/debug_funcs.control)
- [固定提交的 SQL 对象定义](https://github.com/MasahikoSawada/debug_funcs/blob/8508ef191735ef48375a2e626be1986af4d35dd8/debug_funcs--1.0.sql)

`debug_funcs` 暴露一组底层 C 函数，用于实验 PostgreSQL buffer cleanup lock、buffer lock mode、extension/relation lock 和锁计时。SQL 接口包括 `pg_LockBufferForCleanup(regclass,bigint,int)`、`pg_LockBuffer(regclass,bigint,text,int)`、`pg_lockforextension(regclass)`、`replock()`、`show_define_variables()`、`rel_lock(regclass,bool)`、`rel_unlock(regclass)`、`rel_lock_unlock(regclass,int)` 和 `extlock_bench(regclass,int)`。

```sql
CREATE EXTENSION debug_funcs;
SELECT show_define_variables();
```

锁函数接收休眠时长；`pg_LockBuffer` 的模式为 `share` 或 `exclusive`。它们会有意获取内部锁，可能阻塞并发会话，或在指定时间内持续占用资源。不要在应用流量、连接池健康检查或自动化生产诊断中调用。

已复核仓库没有用户指南或兼容性矩阵；版本 `1.0` 直接使用 PostgreSQL 内部 buffer 与 relation-locking API。只能针对准确的目标服务器头文件构建，在可丢弃测试数据上运行，设置 statement timeout，并保留一个可执行取消或终止后端的独立会话。
