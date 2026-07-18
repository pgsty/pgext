## 用法

来源：

- [扩展控制文件](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree.control)
- [0.0.1 版安装 SQL](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree--0.0.1.sql)
- [访问方法实现](https://github.com/zhongjn/mtree/blob/40773bc60f4516015feda337dc0082d1e4466657/mtree.c)

mtree 0.0.1 安装一个名为 mtree 的索引访问方法处理器。它不安装任何数据类型、操作符或操作符类，因此自身无法创建可用的用户索引。

### 安装检查

只能在为目标 PostgreSQL 主版本构建的可丢弃服务器上使用：

```sql
CREATE EXTENSION mtree;
SELECT amname, amtype FROM pg_am WHERE amname = 'mtree';
```

该查询只确认访问方法已注册，并不是端到端索引示例。

### 注意事项

- 扩展没有提供兼容操作符类。普通 CREATE INDEX 命令没有受支持的类型与操作符语义可用。
- 实现是 2020 年对 PostgreSQL GiST 内部代码的分支，调用私有后端、缓冲区、锁与预写日志接口。这类代码与 PostgreSQL 主版本紧密绑定。
- 上游没有使用文档、回归测试、崩溃恢复证据、兼容矩阵或许可证。
- 应将 mtree 视为未完成的索引基础设施。未经独立代码审计，以及并发、崩溃、恢复和升级测试，不得用于承载生产数据。
