## 用法

来源：

- [官方 README](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/README.md)
- [Rust 函数实现](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/src/lib.rs)
- [Rust 包清单](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/Cargo.toml)
- [扩展控制文件](https://github.com/stella3d/pgzstd_rs/blob/d9b56279f41f1375e9db6a3d397c89346dade5f8/pgzstd_rs.control)

`pgzstd_rs` 是一个 pgrx 原型，用 Zstandard 手工压缩与解压 `bytea` 值。上游明确说明它尚未就绪、不可依赖；只能使用一次性数据做评估。

### 压缩与恢复

```sql
CREATE EXTENSION pgzstd_rs;

WITH compressed AS (
  SELECT to_zstd(convert_to('repeated repeated repeated', 'UTF8'), 3) AS value
)
SELECT convert_from(from_zstd(value), 'UTF8')
FROM compressed;
```

`to_zstd(bytea, integer)` 按指定 Zstandard 级别压缩值。`from_zstd(bytea)` 返回解压后的字节，对无效输入抛出通用解压错误。`from_maybe_zstd(bytea)` 则在解压失败时返回原始字节：

```sql
SELECT from_maybe_zstd('\x6e6f742d7a737464'::bytea);
```

Rust 源码还声明了批量压缩函数 `to_zstd_parallel`。其 SQL 签名由固定的 pgrx 工具链生成；使用前应在已安装构建中检查 `\df *zstd*`，并测试数组与空值行为。

### 安全与存储

`from_maybe_zstd` 无法区分有意未压缩数据与损坏/截断的压缩数据，因此可能静默掩盖损坏。每个载荷旁都应保存明确的格式/版本标记与校验和；已知值被压缩时应使用 `from_zstd`。接受不可信帧前，应限制解压后大小与资源，因为很小的输入可能急剧膨胀。

`0.0.0` 版本声明只有超级用户能安装，并提供 PostgreSQL 11 至 16 的功能标志，但上游没有稳定的兼容性或升级契约。API 只进行逐值手工压缩，不会改变 PostgreSQL TOAST 行为。应基准测试端到端空间、CPU、WAL、备份、复制与查询开销，并在往返与灾难恢复测试通过前保留原始数据。
