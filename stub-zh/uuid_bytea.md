## 用法

来源：

- [postgresql/uuid_bytea/README](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/README)
- [postgresql/uuid_bytea/uuid_bytea--1.0.sql](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea--1.0.sql)
- [postgresql/uuid_bytea/uuid_bytea.c](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea.c)
- [postgresql/uuid_bytea/uuid_bytea.control](https://github.com/ringerc/scrapcode/blob/66a8caa0dbbd821e9a7ebc6b59b30ad5f7d875ba/postgresql/uuid_bytea/uuid_bytea.control)

`uuid_bytea` 是在 uuid 与 bytea 之间显式转换的小型实验 C 扩展。上游作者将其定位为试验代码，曾在 PostgreSQL 9.1 上测试。

### 核心用法

```sql
CREATE EXTENSION uuid_bytea;
SELECT uuid_to_bytea('0fcc6350-118d-11e4-a559-7de5338eb025'::uuid);
SELECT bytea_to_uuid(decode('0fcc6350118d11e4a5597de5338eb025','hex'));
```

### 运行边界

`uuid_to_bytea` 返回 UUID 的 16 字节值，`bytea_to_uuid` 将 16 字节值转回 UUID，同时安装显式类型转换。不要传入任意长度的二进制值，也不要据 UUID 文本字节顺序推断其他外部协议。

安装要求超级用户，上游未要求预加载。历史测试环境不能证明当前兼容性，所核对的目录未声明许可证。
