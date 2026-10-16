## 用法

来源：

- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.0.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.0.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.0--1.1.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.0--1.1.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.1--1.2.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.1--1.2.sql)
- [contrib/pgcrypto_pqc/pgcrypto_pqc--1.2--1.3.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc--1.2--1.3.sql)
- [contrib/pgcrypto_pqc/sql/pgcrypto_pqc.sql](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/sql/pgcrypto_pqc.sql)
- [contrib/pgcrypto_pqc/meson.build](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/meson.build)
- [README.md](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/README.md)
- [contrib/pgcrypto_pqc/pgcrypto_pqc.control](https://github.com/dyber-pqc/FortressQL/blob/474609628f61767e766865ee6621b297a72928af/contrib/pgcrypto_pqc/pgcrypto_pqc.control)

`pgcrypto_pqc` 是基于 PostgreSQL 17 的 FortressQL 分叉随附的 SQL 密码扩展。control 版本为 1.3，由基础 SQL 与升级脚本组成。它提供 KEM、签名、加密及算法查询函数；启用它本身不会自动启用该内核的 TLS、TDE 或 WAL 功能。

### 核心用法

```sql
CREATE EXTENSION pgcrypto_pqc;
SELECT name, type, nist_level FROM pqc_algorithms() ORDER BY name;
SELECT length(public_key) AS public_key_bytes,
       length(secret_key) AS secret_key_bytes
FROM pqc_kem_keygen('ML-KEM-768');
```

### 运行边界

需要启用后量子支持并具备 liboqs/OpenSSL 的 FortressQL 构建。本条目不宣称兼容未经修改的 PostgreSQL。安装需要超级用户；control 标记为可迁移模式，也没有要求扩展专用预加载。

`pqc_kem_keygen`、`pqc_kem_encapsulate` 和 `pqc_kem_decapsulate` 实现密钥封装；`pqc_sig_keygen`、`pqc_sign` 和 `pqc_verify` 提供签名接口。`pqc_encrypt` 与 `pqc_decrypt` 提供文档所述的混合加密接口，`pqc_algorithms` 列出可用算法。后续 SQL 修订还加入依赖 FortressQL 的审计、列加密、托管元数据与报告对象。

私钥会返回给 SQL 调用者，并在服务器端处理。应保护其存储、传输和日志，限制执行权限，配对操作必须使用相同算法。提供某种算法或报告函数不等于通过独立安全或合规认证。使用更广泛的加密和审计功能前，应检查该内核的运行文档。
