## 用法

来源：

- [官方 README](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/README.md)
- [扩展 SQL](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw--1.0.sql)
- [内置 Go 示例](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/fdw.go)

`go_fdw` 1.0 是使用 Go 编写 PostgreSQL 外部数据包装器的实验模板。分发示例不会连接外部系统：它只注册一个硬编码生成器，返回十条合成行。开发者需要替换 `fdw.go` 并为自己的数据源重新构建共享库。

### 内置示例

```sql
CREATE EXTENSION go_fdw;

CREATE FOREIGN TABLE public.gotest (
  id integer NOT NULL,
  name text NOT NULL
)
SERVER "go-fdw"
OPTIONS (foo 'bar');

SELECT * FROM gotest;
```

安装会创建外部数据包装器 `go_fdw` 和服务器 `go-fdw`。在内置实现中，`id` 来自行计数器，`name` 变成行/列标签之类文本。可空列保持 NULL；演示只显式处理整数和文本类型。表选项会传给 Go 回调，但示例中的 `foo='bar'` 没有实际含义。

### 开发边界

框架实现表估算、扫描、重扫/关闭、表选项以及 `EXPLAIN` 属性。它没有通用远程协议、写入回调、事务集成、条件下推、连接下推、认证或重试行为。这些语义完全取决于传给 `SetTable` 的自定义 Go 代码。

上游只报告在 Ubuntu x64 上使用 PostgreSQL 9.6 与 Go 1.8.1 进行测试。模块跨越 PostgreSQL C 与 Go 运行时边界，并使用旧服务器 API；必须针对精确 PostgreSQL 和 Go 版本重新构建和压力测试，尤其关注内存、信号、并发与后端关闭。替换库后需要启动新的 PostgreSQL 后端进程，因为已经加载的共享对象不会原地重载。
