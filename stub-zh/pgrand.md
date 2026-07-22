## 用法

来源：

- [官方 README](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/readme.md)
- [随机 FDW 实现](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/src/fdw/mod.rs)
- [GUC 实现](https://github.com/lij55/pgrand/blob/607b698255007094ca5c44ab4e2bda1ad8d03558/src/utils/guc.rs)

`pgrand` 通过 PostgreSQL 外部数据包装器生成合成行。它适合快速填充测试表或为向量测试生成数组；生成值属于测试数据，不提供统计均匀性或密码学安全随机性。

### 核心流程

扩展会安装 `random_fdw_handler`，但必须显式创建包装器与服务器。

```sql
CREATE EXTENSION pgrand;
CREATE FOREIGN DATA WRAPPER random
HANDLER random_fdw_handler;
CREATE SERVER random_server
FOREIGN DATA WRAPPER random;

SET random.seed = 42;
SET random.array_length = 8;
SET random.min_text_length = 8;
SET random.max_text_length = 20;

CREATE FOREIGN TABLE generated_orders (
    order_id integer,
    amount numeric(10,2),
    tags text[],
    payload jsonb,
    source_ip inet
)
SERVER random_server
OPTIONS (total '1000');

SELECT * FROM generated_orders LIMIT 10;

INSERT INTO test_orders
SELECT * FROM generated_orders;
```

外部表选项 `total` 控制一次扫描生成的行数，默认值为 `1000`。非零 `random.seed` 会为每次扫描初始化可重复的 ChaCha8 流；默认值 `0` 使用熵源。

### GUC

- `random.min_int` 与 `random.max_int` 限制整数生成范围。
- `random.min_text_length` 与 `random.max_text_length` 限制文本长度。
- `random.array_length` 设置数组长度，范围为 `1` 到 `16384`。
- `random.float_scale` 控制生成浮点数的小数尺度。
- `random.seed` 选择确定性或熵播种输出。

实现可以生成多种常用的 PostgreSQL 数值、文本、数组、JSON、网络与几何类型。不支持的类型会产生 NULL，而不是一个值。

### 注意事项

上游明确说明生成数字并非均匀分布，且不支持用户定义类型。表访问方法仍标记为开发中；文档中可用的接口是 FDW。扫描会合成全部请求行，不会把条件下推到外部源，因此应保守设置 `total`，并只物化所需数据。已复核的 `0.1.0` Cargo 特性矩阵面向 PostgreSQL `15` 与 `16`；扩展不可重定位，且需要超级用户安装。不需要预加载或重启。
