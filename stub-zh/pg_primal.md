## 用法

来源：

- [官方扩展源码](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/src/lib.rs)
- [官方扩展清单](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/pg_primal/Cargo.toml)
- [官方 Primal Server README](https://github.com/PrimalHQ/primal-server/blob/4b6f384c00fb68d7df10a38bdc33e0e951d2d446/README.md)

`pg_primal` 是 Nostr 缓存服务 Primal Server 使用的小型 PostgreSQL 辅助扩展。它提供 Nostr 标识符解析和 Primal 媒体缓存 URL 构造；不会安装服务端庞大的应用模式，也不会自行运行 Nostr 服务。

### 核心流程

```sql
CREATE EXTENSION pg_primal;

SELECT octet_length(nostr_parse(repeat('00', 32)));
-- 32

SELECT cdn_url(
  'https://example.org/media/photo.jpg',
  'large',
  false
);
```

### 函数索引

- `nostr_parse(text) RETURNS bytea` 先让 Nostr SDK 将输入解析为公钥，再尝试解析为事件 ID。可识别的标识符会转换为 32 字节表示；无法识别的值返回 `NULL`。
- `cdn_url(text, text, boolean) RETURNS text` 在 Primal 的 `https://primal.b-cdn.net/media-cache` 端点下构造 URL。尺寸参数的首字符成为 `s` 参数；动画标志转换为 `a=1` 或 `a=0`；源 URL 经百分号编码后放入 `u`。

### 边界

扩展清单提供 PostgreSQL `11` 到 `16` 的构建特性，版本为 `0.0.0`，且不要求预加载。两个函数都是应用专用辅助函数。`cdn_url` 会直接解包三个参数并读取尺寸参数首字符，因此必须传入非 NULL 值和非空尺寸字符串。Nostr 缓存数据库模式与 websocket API 由 Primal Server 的独立安装流程初始化和运行。
