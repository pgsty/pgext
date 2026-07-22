## 用法

来源：

- [官方扩展 SQL](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw--1.0.sql)
- [官方 C 实现](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/hello_fdw.c)
- [官方 README](https://github.com/wikrsh/hello_fdw/blob/925ade901f9badcc6ed7e01a1c99b6c0f9e3fab1/README.md)

`hello_fdw` 是演示 PostgreSQL 外部数据包装器 API 的最小教学示例。每次扫描只返回一行，并把字符串 `Hello,World` 传给所有声明的列；它不会连接外部系统，也不是通用数据源。

### 核心流程

```sql
CREATE EXTENSION hello_fdw;

CREATE SERVER hello_server
    FOREIGN DATA WRAPPER hello_fdw;

CREATE FOREIGN TABLE hello_demo (
    first_value text,
    second_value text
) SERVER hello_server;

SELECT * FROM hello_demo;
```

结果包含一行，两列都是 `Hello,World`。重新扫描会重置这个单行迭代器。

### 已实现接口

- `hello_fdw_handler` 注册规划、扫描、重扫、解释与分析回调。
- `hello_fdw_validator` 不执行任何操作，因此不会校验包装器、服务器、用户映射或表选项。
- `EXPLAIN` 会增加一个固定的 `Hello` 属性。
- 实现固定估算一行，没有插入、更新、删除、导入模式或远程连接回调。

### 限制

扩展通过各列类型的文本输入函数传值。非文本列可能以不同方式接受 `Hello,World`，也可能拒绝它，因此演示应使用文本列。即使空校验器接受选项，选项也不会改变返回数据。

这是示例代码，不是运维型 FDW。它只有少量较老的提交，不校验选项、不执行远程 I/O，且实现的是历史 FDW API。它只适合在一次性环境中学习回调；若要作为模板，必须先针对目标 PostgreSQL 版本更新内存、规划器、API 兼容性、错误处理与写入路径假设。
