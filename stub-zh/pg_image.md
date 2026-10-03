## 用法

来源：

- [extensions/pg_image/pg_image.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/pg_image.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_image/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/Cargo.toml)
- [extensions/pg_image/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_image/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/README.md)

`pg_image` 0.3.0 通过 `pgimg` 模式处理图像 BYTEA 值，提供尺寸、EXIF／GPS、缩略图、格式转换与感知哈希。

### 核心用法

```sql
CREATE EXTENSION pg_image;
CREATE TABLE image_sample (id bigint PRIMARY KEY, photo bytea);
SELECT id, pgimg.img_width(photo), pgimg.img_height(photo),
       pgimg.img_thumbnail(photo, 128, 128)
FROM image_sample;
```

### 运行边界

控制文件不限定超级用户安装，但仍需具备创建相应对象的权限。无需预加载。基础图像操作不依赖可选的 ONNX 推理；目标检测需要 ONNX Runtime 与匹配模型文件，通过 `pg_image.model_dir`、`pg_image.default_model` 及置信度／IoU 设置选择。GPS 转几何对象的辅助函数另需 PostGIS。应限制图像尺寸，并控制模型路径与高开销推理的访问。感知哈希接近不代表图像完全相同。 这是采用 Matroid Source Available License 1.0 的无支持概念验证项目，API 可能变化。0.3.0 是新的升级起点：旧 0.2.0 安装需要预演数据迁移／重建，不能直接执行普通 ALTER EXTENSION UPDATE。遵循上游这一破坏性路径前，必须备份数据并检查依赖。
