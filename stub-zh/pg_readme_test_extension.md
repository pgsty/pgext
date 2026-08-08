## 用法

来源：

- [pg_readme 0.7.1 README](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/README.md)
- [测试扩展控制文件](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension.control)
- [测试扩展 SQL 固件](https://api.pgxn.org/src/pg_readme/pg_readme-0.7.1/pg_readme_test_extension/pg_readme_test_extension--forever.sql)
- [Pigsty 软件包矩阵](https://pgext.cloud/ext/pg_readme_test_extension)

`pg_readme_test_extension` 是 `pg_readme` 随附的集成测试固件。它会安装带注释的域、类型、表、视图、例程、触发器和处理指令，以便上游验证 `pg_extension_readme()`。它不是应用功能，也不是生产依赖。

### 使用测试固件

```sql
CREATE EXTENSION pg_readme CASCADE;
CREATE EXTENSION pg_readme_test_extension;

SELECT pg_extension_readme('pg_readme_test_extension'::name);
```

使用输出测试或演示生成器，然后从一次性数据库中移除该固件：

```sql
DROP EXTENSION pg_readme_test_extension;
```

### 边界与注意事项

- 上游发行版本为 0.7.1，但该固件的控制版本刻意使用字面量 `forever`。
- 该固件随 `pg_readme` 一起提供；当前 Pigsty DEB 软件包为 0.7.1，而 RPM 软件包仍为 0.7.0。两边的扩展版本都保持 `forever`。
- 它可重定位，本身不要求 `hstore`，并会创建使用通用名称的示例对象。仅应安装在这些对象不会与真实应用模式冲突的环境中。
- 它的 SQL 接口用于覆盖生成器行为，并可能随测试演进而改变。不要让应用代码依赖这些固件对象。
