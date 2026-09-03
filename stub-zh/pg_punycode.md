## 用法

来源：

- [官方文档](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/README.md)
- [扩展控制文件](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/crates/pg_punycode/pg_punycode.control)
- [构建清单](https://github.com/sindrip/pg_punycode/blob/67792c9a5e2fc17627532e4696c0861096f6efba/crates/pg_punycode/Cargo.toml)

`pg_punycode` 提供 Punycode 与 IDNA/UTS-46 域名转换函数。

### 启用

为目标服务器安装文件，然后在目标数据库中创建 `pg_punycode`：

```sql
CREATE EXTENSION pg_punycode;
```

### 核心流程

以下示例取自经审查的上游文档。使用前应调整对象名、路径、凭据与工作负载参数。

```sql
CREATE EXTENSION pg_punycode;

SELECT domain_to_ascii('Bücher.DE');                  -- xn--bcher-kva.de
SELECT domain_to_unicode('xn--fiqs8s');               -- 中国
SELECT punycode_encode('bücher');                     -- bcher-kva
SELECT punycode_decode('MajiKoi5-783gue6qz075azm5e'); -- MajiでKoiする5秒前
```

### 主要对象

经审查的安装接口包含以下主要对象：

| 对象 | 类型 | 运维作用 |
|---|---|---|
| `domain_to_ascii` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `punycode_encode` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `domain_to_unicode` | FUNCTION | 经审查安装接口中的可调用函数。 |
| `punycode_decode` | FUNCTION | 经审查安装接口中的可调用函数。 |

### 运维与边界

- 已验证的 PostgreSQL 主版本证据覆盖 13, 14, 15, 16, 17, 18, 19；不要推断未列出的主版本。
- 目录生命周期为 preview；生产使用前应测试升级、备份恢复与服务器兼容性。
