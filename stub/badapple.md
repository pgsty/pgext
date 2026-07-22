## Usage

Sources:

- [Official README](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/README.md)
- [Extension SQL](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple--1.0.0.sql)
- [C implementation](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.c)
- [Extension control file](https://github.com/higuoxing/pg_badapple/blob/92ddbeea6fee794bfaad70f1b1d9c4bfab068c91/badapple.control)

`badapple` is a novelty extension that renders the Bad Apple animation, or a basketball animation, as PostgreSQL `NOTICE` messages. It is suitable for an interactive terminal demonstration, not application or production workloads: the calling backend remains occupied while it reads and plays the server-side animation file.

### Core Workflow

Install the extension in a disposable database, enable notice output, and invoke one of its two entry points from a terminal that interprets ANSI/VT100 control sequences:

```sql
CREATE EXTENSION badapple;
SET client_min_messages = notice;

SELECT play_badapple();
-- Or:
SELECT play_basketball();
```

The C code opens `badapple.txt` or `basketball.txt` from PostgreSQL's shared extension directory on the database server. It reads 30 text lines per frame, emits the frame as a `NOTICE`, clears the terminal with an ANSI escape sequence, and sleeps between frames. Cancel the statement to stop playback.

### User-Facing Objects

- `play_badapple()`: play the bundled Bad Apple text frames.
- `play_basketball()`: play the bundled basketball text frames.

Both functions return `void`; their visible output is the stream of `NOTICE` messages.

### Operational Boundaries

The media files must be installed beside the extension's shared data files on the server. A graphical SQL client, suppressed notices, or a terminal without ANSI cursor handling will not display the animation as intended.

Playback generates many messages and deliberately keeps one PostgreSQL backend busy. Do not invoke it through connection pools, APIs, or production monitoring paths. The implementation copies source lines into a fixed-size frame buffer, so use only the upstream-supplied, trusted media files; replacing them with oversized or untrusted lines can corrupt backend memory. Test cancellation and logging behavior in an isolated database.
