## Usage

Sources:

- [Upstream README](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/README.md)
- [Extension control file](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/pg_ffmpeg.control)
- [Rust package manifest](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/Cargo.toml)
- [Filter safety implementation](https://github.com/sweatybridge/pg_ffmpeg/blob/4a07afed5e04fb23937cf6ff8cd220c41667b2f9/src/filter_safety.rs)

`pg_ffmpeg` version `0.3.4` exposes FFmpeg 8 media inspection and transformation functions in the fixed `ffmpeg` schema for PostgreSQL 16–18. Inputs and outputs are commonly whole media objects carried as `bytea`; the extension also provides transcoding, audio extraction, trimming, frame extraction, HLS, GIF, waveform, subtitle, overlay, filtering, concatenation, and encoding operations.

### Inspect a stored upload

```sql
CREATE EXTENSION pg_ffmpeg;

SELECT ffmpeg.media_info(body)
FROM media_uploads
WHERE id = 1;

SELECT ffmpeg.thumbnail(body, 1.0, 'jpeg')
FROM media_uploads
WHERE id = 1;
```

FFmpeg parses attacker-controlled binary formats inside the PostgreSQL backend, and media jobs can consume substantial CPU, memory, and temporary storage. Large `bytea` values and `concat_agg` can amplify backend memory use. The `hls` interface can fetch a URL and stores playlists and segments in extension tables, so unrestricted execution can create outbound-request and storage-abuse paths. The upstream filter allowlist blocks several file-reading filters, but that is not a complete trust boundary. Grant execution only to a dedicated workload role, apply statement/resource/network limits, validate media outside PostgreSQL when possible, and never expose server-side file operations to untrusted SQL users.
