## 用法

来源：

- [已复核 commit 的 go_fdw README](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/README.md)
- [已复核 commit 的 go_fdw 安装 SQL](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw--1.0.sql)
- [已复核 commit 的 go_fdw.control](https://github.com/dennwc/go-fdw/blob/93f15338153f6e07f1673e8733863cc17258a181/go_fdw.control)

`go_fdw` 是用于开发 PostgreSQL 外部数据包装器的实验性 Go/cgo 模板，并不是连接任意 Go 数据的通用连接器。附带示例支持表扫描、EXPLAIN 和表选项；开发者需要用自己的数据源逻辑替换示例中的 `SetTable` 实现。

### 运行附带示例

```sql
CREATE EXTENSION go_fdw;

CREATE FOREIGN TABLE public.gotest (
  id integer NOT NULL,
  name text NOT NULL
)
SERVER "go-fdw"
OPTIONS (foo 'bar');

SELECT * FROM public.gotest;
```

扩展安装脚本会自动创建 `go_fdw` 包装器和名为 `go-fdw` 的服务器。

### 注意事项

- 上游将项目称为实验性项目，并且只说明在 PostgreSQL 9.6 和 Go 1.8.1 上测试过。用于当前服务器前必须先验证并移植。
- 这是带基本可运行实现的开发模板。其表选项和返回行取决于编译进动态库的 Go 代码。
- README 说明，开发过程中替换已编译动态库后需要重启 PostgreSQL 才能重新测试；这并不是 `shared_preload_libraries` 要求。
