## 用法

来源：

- [上游 README](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/README.md)
- [扩展 control 文件](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-properties/custom_properties.control)
- [可信语言扩展 SQL](https://github.com/GaryAustin1/custom-properties/blob/41f8434fab75171b0fe4080c0d150d85bcc16b18/tle-custom-properties/custom_properties--0.0.3.sql)
- [dbdev 软件包页面](https://database.dev/garyaustin/custom_properties)

`custom_properties` 是面向 Supabase 的可信语言扩展，用于保存用户的命名属性，并提供适合行级安全策略调用的辅助函数。已发布的 `dbdev` 软件包和现有 SQL 版本为 `0.0.3`；当前 control 文件写的是 `0.0.4`，但复核提交中没有对应的 `0.0.4` 安装脚本。

### 软件包工作流

先生成迁移，复核内容并选定目标模式，再通过常规迁移流程应用：

```sh
dbdev add -o ./migrations -s extensions -v 0.0.3 package -n "garyaustin@custom_properties"
```

安装后，策略代码可以查询这些辅助函数（若选择了其他模式，请替换 `user_roles`）：

```sql
SELECT user_roles.user_has_property('Teacher');
SELECT user_roles.get_user_properties();
```

SQL 假定存在 `auth.users`、`auth.uid()`、`authenticated` 和 `service_role` 等 Supabase 对象；它会创建表、RLS 策略、授权和 `SECURITY DEFINER` 函数，并定义一个默认禁用、可改写 `auth.users.raw_app_meta_data` 的触发器。部署前应审计生成的迁移、函数搜索路径、授权和可选触发器。在上游补齐相应 SQL 前，不要认定 control 文件中的版本可以安装。
