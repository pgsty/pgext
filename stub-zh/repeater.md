## 用法

来源：

- [已复核 commit 的上游 README](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/README.md)
- [PostgreSQL 12 开发补丁](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/pg12_devel.patch)
- [Repeater 实现](https://github.com/danolivo/execplan/blob/a61ae072f8d10d0836ba491a549cb198100c8c49/repeater.c)

`repeater` 是把序列化 PostgreSQL 原始查询计划发送到远程节点并执行相同计划的演示模块。它属于 `execplan` 研究项目，依赖 `postgres_fdw`，还要求侵入式修改 PostgreSQL 12 核心和 libpq。

```conf
shared_preload_libraries = 'repeater'
```

加载库后，由修补后的服务器与客户端机制管理可移植对象标识符、计划序列化、传输、远程会话和事务，而不是标准 PostgreSQL SQL API。不要加载到未修补服务器，也不要混用不同源码树构建的二进制。

序列化计划依赖 PostgreSQL 内部节点布局、目录对象、类型、操作符、模式、统计假设和二进制版本。执行外部原始计划会绕过正常规划安全，版本或目录偏差可能导致崩溃或语义损坏。这是 PostgreSQL 12 时代、仅有简化演示测试的研究代码，不是生产分布式执行器。任何研究都必须使用隔离可丢弃集群，并验证认证、计划来源、对象映射、故障、取消、远端事务和升级。
