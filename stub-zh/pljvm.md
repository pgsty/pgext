## 用法

来源：

- [项目 README](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/README.md)
- [扩展 control 文件](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/pljvm.control)
- [1.0.0 版安装 SQL](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/pljvm--1.0.0.sql)
- [PostgreSQL 调用处理器](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/src/pljvm.c)
- [Java RPC 服务](https://github.com/davecramer/pljvm/blob/39696d3c0a157d4f9b295d9180620cabd6dc6242/plj-java/src/main/java/org/postgresql/plj/PLJVM.java)

`pljvm` 1.0.0 定义了一种 PostgreSQL 过程语言：它通过 RPC 把函数参数发送给独立 JVM 进程中的代码，再反序列化返回值。函数体指定 Java 类和方法；Java 端服务加载该类、创建实例并调用公开方法。

### 启动服务并定义函数

```shell
cd plj-java
mvn exec:java
```

```sql
CREATE EXTENSION pljvm;

CREATE FUNCTION pljvm_add(i int2, j int2)
RETURNS int2
AS 'org.postgresql.plj.test.Int.add'
LANGUAGE pljvm;

SELECT pljvm_add(2::int2, 3::int2);
```

Java 类必须位于服务 classpath 中。已审查源码的两端都使用固定 TCP 端点 `localhost:8000`；扩展不会启动或监管 JVM 服务。

### 安全与实现限制

虽然 README 将其描述为可信执行引擎，但安装 SQL 使用的 `CREATE LANGUAGE pljvm` 没有 `TRUSTED` 关键字。因此 PostgreSQL 会把它视为非可信语言，创建函数需要相应的高权限。已审查 RPC 路径没有可配置端点、认证或传输加密；应隔离该服务，且不得对外暴露 8000 端口。

这个源码快照属于早期实现：C 处理器注明不支持触发器和多列返回值，Java 服务也使用较早的 Netty/Maven 依赖。JVM 进程生命周期、classpath 部署、故障处理、资源限制以及具体 PostgreSQL/JDK 兼容性都需要由运维方负责，并在生产使用前完成验证。
