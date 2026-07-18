## 用法

来源：

- [固定提交的上游 README](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/README.md)
- [0.0.1 版安装 SQL](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/micromanage--0.0.1.sql)
- [固定提交的 planner 实现](https://github.com/lithp/pg_micromanage/blob/1ed36659995a82a6346305f276762c8eed65d32b/micromanage.c)

micromanage 是一个面向 PostgreSQL 9.6 的实验性 planner hook 演示。它接收 base64 编码的 protobuf plan，拦截形状非常固定的调用，并构造支持顺序扫描、简单表达式和过滤、nested-loop join 以及排序的 PostgreSQL plan。

### 示例

构建扩展需要 protobuf-c 和 protoc。下面的上游示例运行一个预先编码的 plan，扫描只有一列的表：

```sql
CREATE EXTENSION micromanage;

CREATE TABLE a (a integer);
INSERT INTO a VALUES (1);

SELECT * FROM run_select('Cg0KBwoFCAESAWEaAggBEgMKAWE=');
```

结果为数值 1。安装的 API 还包括用于查看 plan 的 dump_query，以及把仓库 queries.proto 所描述的文本形式转换为编码数据的 encode_protobuf。

### 实验性限制

该 hook 只识别外层形状符合实现预期、参数为常量的 run_select SELECT。它只在 public schema 中解析关系名；如果已有其他 planner hook，它会拒绝加载而不是串联调用。

源码注释明确指出 cost 信息不完整、事务处理不确定、存在内存泄漏，而且可能缺少安全检查。代码依赖旧版 planner 内部接口，上游只记录了 PostgreSQL 9.6。仓库并未归档，但自 2018 年后没有更新。micromanage 只应作为隔离数据库中的教学或研究代码；不要接受不可信 protobuf plan，也不要用于生产查询执行。
