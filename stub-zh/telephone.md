## 用法

来源：

- [上游 telephone 类型文档](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/doc/telephone.md)
- [已复核 commit 的上游 README](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/README.md)
- [扩展 control 文件](https://github.com/pgstuff/telephone/blob/abdf084b6460f2dbfe747c0c19092a845ad67e02/telephone.control)

`telephone` 添加规范化电话号码类型。它会规范标点和空格，把大写字母按电话键盘映射为数字以比较号码身份，并支持暂停、确认、分机与部分号码指令。因此，唯一约束比较的是号码身份，而不是显示形式。

```sql
CREATE EXTENSION telephone;

CREATE TABLE contact_phone (
  phone telephone PRIMARY KEY,
  label text NOT NULL
);

INSERT INTO contact_phone VALUES ('1 (800) 555-01AZ', 'support');
SELECT * FROM contact_phone
WHERE phone = '18005550129'::telephone;
```

以 `+` 开头的号码会按扩展自带的拨号计划数据校验，并可使用国内格式、E.123、服务类型和地理区段辅助函数。拨号计划不受支持时仍可使用纯数字模式。

### 注意事项

0.0.1 版被标记为 unstable，国际拨号计划覆盖也不完整；不能把它当成完备的全球号码验证库。虽然 control 文件声明可重定位，但安装 SQL 将几个辅助函数硬编码在 `public` 中。应审查 schema 放置，并测试应用实际接受的准确拨号计划。
