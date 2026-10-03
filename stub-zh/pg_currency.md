## 用法

来源：

- [extensions/pg_currency/pg_currency.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/pg_currency.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_currency/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/Cargo.toml)
- [extensions/pg_currency/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_currency/src/seed.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/seed.rs)
- [extensions/pg_currency/src/rate.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/rate.rs)
- [extensions/pg_currency/src/convert.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_currency/src/convert.rs)

`pg_currency` 0.3.0 在 `pgcurrency` 中保存币种与分日期汇率，并以基准币种进行交叉换算。

### 核心用法

```sql
CREATE EXTENSION pg_currency;
SELECT pgcurrency.seed_iso();
SELECT pgcurrency.set_rate('EUR', 0.9, DATE '2026-10-01');
SELECT pgcurrency.convert(100, 'USD', 'EUR', DATE '2026-10-01');
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。汇率由使用方录入，不包含实时行情源。`pgcurrency.base_currency` 默认 USD，基准汇率为一。`set_rate` 保存汇率、日期、类型与来源，`get_rate`、`convert` 和 `round_to_currency` 提供历史查找及币种精度处理。缺少所需汇率时会报错。它与源码包名称相近的既有币种类型扩展是不同项目。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
