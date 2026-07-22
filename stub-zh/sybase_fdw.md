## 用法

来源：

- [官方上游 README](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/README)
- [官方扩展 SQL](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw--1.0.sql)
- [官方实现](https://github.com/raitech/sybase_fdw/blob/a230646830bb363fca5d659b1a68408f780140f5/sybase_fdw.c)

`sybase_fdw` 是不完整的外部数据包装器骨架，而不是可用的 Sybase 连接器。1.0 版会登记处理器和验证器，但所有外部扫描回调都是 NULL，验证器也没有实现选项校验。它无法规划或读取外部表。

### 安装创建的内容

在一次性开发数据库中，安装只能用于查看登记后的包装器对象：

```sql
CREATE EXTENSION sybase_fdw;

SELECT fdwname, fdwhandler::regproc, fdwvalidator::regproc
FROM pg_foreign_data_wrapper
WHERE fdwname = 'sybase_fdw';
```

扩展会创建 `sybase_fdw_handler()`、`sybase_fdw_validator(text[], oid)` 和 `sybase_fdw` 外部数据包装器。它没有记录连接、认证、服务器、用户映射或外部表选项。

### 不可用边界

处理器返回的 `FdwRoutine` 中，规划、解释、开始、迭代、重新扫描和结束回调全都是 NULL。因此，不要发布 `CREATE SERVER` 或 `CREATE FOREIGN TABLE` 流程并暗示查询可以成功。空验证器也不能防止拼错或不安全的选项。

只能把该项目当作历史 FDW 骨架，或者在实现并测试完整 PostgreSQL FDW 生命周期、Sybase 客户端连接层、类型转换、事务、凭据、错误、取消和清理之后，将其作为开发起点。实际访问 Sybase 时，应选择有兼容性与安全行为文档的受维护连接器；1.0 标签并不表示已经适合生产。
