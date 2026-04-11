
## 用法

> 语法：
>
> ```sql
> CREATE EXTENSION pg_liquid;
> SELECT liquid.query('Edge("a","b"). Edge("b","c"). Path(X,Y) :- Edge(X,Y). Path(X,Y) :- Edge(X,Z), Path(Z,Y). Path("a",Y)?');
> ```
>
> 来源：[README](https://github.com/michael-golfi/pg_liquid)，[文档站](https://michael-golfi.github.io/pg_liquid/)

`pg_liquid` 将 Liquid 博客语言及其数据模型映射到原生 PostgreSQL 存储与执行环境。该扩展提供 SQL 入口，用于运行 Liquid 风格程序、按主体查询，以及管理将关系行投影为 Liquid 复合体的行归一化器。

## 核心函数

上游 README 列出的主要函数包括：

- `liquid.query(program text)`
- `liquid.query_as(principal text, program text)`
- `liquid.read_as(principal text, program text)`

这些函数分别支持普通执行、按主体查询和 CLS 感知读取。

## 语言特性

当前 README 说明支持的程序特性包括：

- `%` 注释
- 以 `.` 结尾的事实与规则定义
- 一个终结 `?` 查询
- `Edge(...)`
- 形如 `Type@(cid=..., role=...)` 的命名复合体
- 查询局部递归规则

## 示例形态

程序以文本形式传入，可定义事实、规则和最终查询：

```sql
SELECT liquid.query($$
  Edge("a","b").
  Edge("b","c").
  Path(X,Y) :- Edge(X,Y).
  Path(X,Y) :- Edge(X,Z), Path(Z,Y).
  Path("a",Y)?
$$);
```

## 说明

项目 README 指向 VitePress 文档站作为主要文档入口，并说明运行和部署细节也记录在该站点中。该扩展当前以 PGXN 包版本 `0.1.1` 发布，并已在 PostgreSQL 14 到 18 上验证。
