## 用法

来源：

- [扩展初始化与配置](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw.c)
- [外部表实现](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/fdw_driver.c)
- [扩展 SQL](https://github.com/troels/hbase_fdw/blob/a4c0c5d981113a0f37186cb85374601f9d69eb60/hbase_fdw--1.0.sql)

`hbase_fdw` 版本 `1.0` 是实验性的只读外部数据包装器，它启动嵌入式 JVM，并通过 HBase Java 客户端执行查询。仓库没有用户 README，因此以下接口仅限扩展源码中可确认的行为。

### 服务器设置

模块只在 PostgreSQL 处理 `shared_preload_libraries` 时注册共享内存与后台工作进程。重启服务器前先配置 JVM：

```conf
shared_preload_libraries = 'hbase_fdw'
hbase_fdw.java_home = '/path/to/java-home'
hbase_fdw.classpath = '/path/to/hbase-client.jar:/path/to/dependencies/*'
```

实现会在 `hbase_fdw.java_home` 下查找 amd64 `libjvm.so`；应确认所选 Java 布局与 HBase 客户端 JAR 集合符合这些源码假设。

### 核心流程

```sql
CREATE EXTENSION hbase_fdw;
CREATE SERVER hbase_srv FOREIGN DATA WRAPPER hbase_fdw;

CREATE FOREIGN TABLE events (
  row_key text OPTIONS (hbase_type 'row_key'),
  payload text OPTIONS (
    hbase_type 'column', family 'f', qualifier 'payload'
  )
) SERVER hbase_srv OPTIONS (hbase_table 'events');

SELECT * FROM events WHERE row_key = 'event-42';
```

`hbase_table` 选择远端表，默认使用外部表名称。列选项 `hbase_type` 接受 `row_key`、`family` 或 `column`；普通列还可用 `family` 与 `qualifier` 标识 HBase 单元格。

### 限制

该 FDW 只暴露扫描回调，且只下推行键列上的等值条件。它没有定义服务器或用户映射连接选项；HBase 连接来自嵌入式 JVM 可见的 Java 客户端配置。后台工作进程失败后不会自动重启。应把这份年代较久且文档极少的代码仅用于评估，并在使用前测试服务器启动、JVM 加载、HBase 认证、类型转换与故障恢复。
