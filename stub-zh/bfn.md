## 用法

来源：

- [官方上游文档](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/readme.md)
- [官方 Rust 实现](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/src/lib.rs)
- [官方 Rust 软件包清单](https://github.com/bpsbits-org/bfn/blob/f5185ee875e4b88eb56b9c233100fb493283dde5/Cargo.toml)

`bfn` 2.0.1 是基于 pgrx 的通用 PostgreSQL 辅助函数集合，涵盖 UUIDv7、日期、文本清理、JSON 地址、解析、哈希和数值默认值。它不是单一数据模型：应审查并只授权应用真正需要的函数。该发布声明了 PostgreSQL 16 至 18 构建特性，并要求超级用户安装。

### 安装到专用模式

扩展不可重定位，上游将其安装到专用的 `bfn` 模式：

```sql
CREATE SCHEMA IF NOT EXISTS bfn;
CREATE EXTENSION bfn SCHEMA bfn;

SELECT bfn.bfn_version();
SELECT bfn.new_uuid();
SELECT bfn.uuid_to_ts(bfn.new_uuid());
SELECT bfn.date_range(DATE '2026-01-01', DATE '2026-01-03');
```

应使用模式限定调用，不要把宽泛的工具模式放在不可信 `search_path` 的前部。

### 重要函数组

- `new_uuid()` 创建 UUIDv7 值，`uuid_to_ts(uuid)` 提取其中的时间戳；非 v7 或无效值返回 NULL。
- `all_dates_from(date, date)`、`date_range(date, date)`、`first_day_of_month(date)` 和 `last_day_of_month(date)` 提供日期辅助功能。
- `to_address(...)` 构造带可选坐标的 JSONB 地址。
- `isi(...)`、`is_empty(...)`、`is_null(...)`、`not_null(...)`、`trim(...)`、`san_trim(...)`、`strip_tags(...)` 和 `upper_first(text)` 提供常用谓词与文本规范化。
- `random_base64()`、`md5_as_base64(text)`、`md5_as_uuid(text)` 及其验证辅助函数生成 token 或确定性的 MD5 派生值。
- 解析、名称连接、公制缩放和零值默认辅助函数覆盖应用专用的便利用例。

### 安全与兼容性

MD5 派生辅助函数不适合密码存储、签名或抗碰撞身份标识。对 `random_base64()` 输出的判断应限于固定 Rust `rand` 依赖提供的保证，不能把它当成已有文档的认证协议。文本清理器只是便利变换，不是 HTML 或 SQL 安全边界。应固定扩展及 Rust 依赖图，因为输出格式和解析行为可能影响持久值与索引。在约束、生成列或持久键中使用辅助函数前，应在准确的 PostgreSQL 大版本上测试 UUID 排序、时间提取、Unicode 大小写、数值精度和 NULL 或默认值语义。
