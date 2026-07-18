## 用法

来源：

- [上游 README](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/README.md)
- [扩展 control 文件](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/postpic.control)
- [SQL 安装脚本](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/sql/postpic.sql)
- [C 实现](https://github.com/drotiro/postpic/blob/98307c665f4f83aa0cd89812e9640f737f1df0ea/server/src/postpic.c)

`postpic` `0.9.1` 版添加由 GraphicsMagick 支持的 `image` 类型。SQL 函数可以从 `bytea` 或大对象载入图像，读取宽高、日期及类似 EXIF 的属性，并生成缩略图、裁剪、旋转、绘图和联系表。

### 示例

```sql
CREATE EXTENSION postpic;
CREATE TABLE pictures (id bigint GENERATED ALWAYS AS IDENTITY, body image);
INSERT INTO pictures(body) VALUES (image_from_bytea(pg_read_binary_file('/srv/import/photo.jpg')));
SELECT id, width(body), height(body), colorspace(body) FROM pictures;
SELECT thumbnail(body, 320) FROM pictures WHERE id = 1;
```

示例会读取服务器端路径，因此需要特权文件访问；应用应改为通过受控导入路径传递获准字节。图像解码会在 PostgreSQL 后端内运行原生 GraphicsMagick 代码，畸形媒体可能消耗大量内存/CPU 或触发解码器漏洞。这是没有当前 PostgreSQL 或 GraphicsMagick 兼容性声明的旧代码。应限制执行权限、约束输入大小、及时修补编解码器，并只在隔离环境评估。
