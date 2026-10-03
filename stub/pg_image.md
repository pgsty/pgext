## Usage

Sources:

- [extensions/pg_image/pg_image.control](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/pg_image.control)
- [README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/README.md)
- [extensions/pg_image/Cargo.toml](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/Cargo.toml)
- [extensions/pg_image/src/lib.rs](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/src/lib.rs)
- [CHANGELOG.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/CHANGELOG.md)
- [LICENSE](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/LICENSE)
- [extensions/pg_image/README.md](https://github.com/matroidbe/pg_extensions-releases/blob/2617326c54d1ef3c51cf996b2e2ef1123aefae25/extensions/pg_image/README.md)

`pg_image` 0.3.0 processes image BYTEA values through the `pgimg` schema: dimensions, EXIF/GPS, thumbnails, format conversion and perceptual hashes.

### Core Workflow

```sql
CREATE EXTENSION pg_image;
CREATE TABLE image_sample (id bigint PRIMARY KEY, photo bytea);
SELECT id, pgimg.img_width(photo), pgimg.img_height(photo),
       pgimg.img_thumbnail(photo, 128, 128)
FROM image_sample;
```

### Operational Boundaries

The control does not require superuser-only installation, but object-creation privileges still apply. No preload is required. Basic image operations work independently of optional ONNX inference. Detection requires ONNX Runtime and matching model files; `pg_image.model_dir`, `pg_image.default_model` and confidence/IoU settings select them. GPS-to-geometry helpers require PostGIS separately. Bound image sizes and restrict access to model paths and expensive inference. Perceptual hash similarity is not proof of image identity. This is an unsupported proof of concept under Matroid Source Available License 1.0. APIs may change. Version 0.3.0 is the new upgrade baseline: earlier 0.2.0 installations require a rehearsed data migration/recreation, not ordinary ALTER EXTENSION UPDATE. Back up data and dependencies before following that destructive upstream path.
