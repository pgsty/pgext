## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/README.md)
- [0.0.1 版本 SQL 对象](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby--0.0.1.sql)
- [调用处理器实现](https://github.com/godfat/plruby/blob/548db739064a02dad4418376f46fc25e8e842f29/plruby.c)

`plruby` 是把 Ruby 嵌入 PostgreSQL 后端的极早期过程语言原型。0.0.1 版本安装带调用、内联与校验处理器的不可信语言 `plruby`；上游只称它在一台电脑上“勉强可用”。

```sql
CREATE EXTENSION plruby;
CREATE FUNCTION ruby_add(integer, integer)
RETURNS integer
LANGUAGE plruby
AS $$ args[0] + args[1] $$;
SELECT ruby_add(2, 3);
```

只有在审计处理器实际参数与返回约定后，该示例才适合隔离构建测试。不可信过程语言能够以数据库服务器进程的操作系统权限执行，必须保持仅超级用户可用。

不要把此原型部署到生产，也不要让租户提交代码。它大量复制了早期 PL/V8 模式，没有当前 Ruby/PostgreSQL 兼容矩阵，也没有沙箱、内存、异常、信号、线程、垃圾回收或事务安全保证。应改用持续维护的过程语言，或在数据库外执行 Ruby。任何归档测试都应位于没有秘密和网络信任的可丢弃主机。
