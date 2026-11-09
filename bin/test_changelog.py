#!/usr/bin/env python3
"""Regression checks for complete, consistent release-note ingestion."""
import contextlib
import csv
import importlib.util
import io
import sys
import tempfile
import unittest
from pathlib import Path
from unittest.mock import patch

sys.dont_write_bytecode = True
SPEC = importlib.util.spec_from_file_location(
    "changelog", Path(__file__).with_name("changelog.py")
)
changelog = importlib.util.module_from_spec(SPEC)
SPEC.loader.exec_module(changelog)


def table(rows, date="2026-04-19"):
    return (
        f"## {date}\n\n"
        "| Package | Old | New | Comment |\n"
        "|:--------|:----|:----|:--------|\n"
        + "\n".join(rows) + "\n"
    )


class ChangelogTest(unittest.TestCase):
    def test_invalid_sources_leave_previous_csv_intact(self):
        row = "| postgis | 3.6.2 | 3.6.3 | bump |"
        other = "| pgvector | - | 0.8.0 | new |"
        cases = {
            "Chinese-only row": (table([row]), table([row, other])),
            "English-only row": (table([row, other]), table([row])),
            "Chinese-only batch": (table([row]), table([row]) + table([other], "2026-04-20")),
            "English-only batch": (table([row]) + table([other], "2026-04-20"), table([row])),
            "version mismatch": (table([row]), table([row.replace("3.6.2", "3.6.1")])),
            "conflicting duplicates": (
                table([row, row.replace("3.6.3", "3.6.4")]),
                table([row, row.replace("3.6.3", "3.6.4")]),
            ),
            "missing column": (table(["| postgis | 3.6.2 | 3.6.3 |"]), table([row])),
            "empty name": (table([row.replace("postgis", "")]), table([row.replace("postgis", "")])),
        }
        for name, (en, zh) in cases.items():
            with self.subTest(name=name), tempfile.TemporaryDirectory() as temp:
                root = Path(temp)
                for side in ("rpm", "deb"):
                    (root / f"{side}.md").write_text(en, encoding="utf-8")
                    (root / f"{side}.zh.md").write_text(zh, encoding="utf-8")
                output = root / "changelog.csv"
                output.write_text("previous validated output\n", encoding="utf-8")
                with patch.object(changelog, "RELEASE", root), patch.object(changelog, "OUTPUT", output):
                    with self.assertRaises(ValueError):
                        changelog.main()
                self.assertEqual(output.read_text(), "previous validated output\n")

    def test_reordered_twins_preserve_versions_notes_and_side_differences(self):
        with tempfile.TemporaryDirectory() as temp:
            root = Path(temp)
            sources = {
                "rpm.md": table(["| postgis | 3.6.2 | 3.6.3-2 | RPM build |"]),
                "rpm.zh.md": table(["| postgis | 3.6.2 | 3.6.3-2 | RPM 构建 |"]),
                "deb.md": table([
                    "| postgis | 3.6.2 | 3.6.3-1 | DEB build |",
                    "| vector | - | [0.8.0](https://example.org/release) | new |",
                ]),
                "deb.zh.md": table([
                    "| `vector` | — | 0.8.0 | 新增 |",
                    "| postgis | 3.6.2 | 3.6.3-1 | DEB 构建 |",
                ]),
            }
            for name, text in sources.items():
                (root / name).write_text(text, encoding="utf-8")
            output = root / "changelog.csv"
            catalog = ({"postgis": "postgis", "vector": "pgvector"}, {"postgis", "pgvector"}, {})
            with (
                patch.object(changelog, "ROOT", root),
                patch.object(changelog, "RELEASE", root),
                patch.object(changelog, "OUTPUT", output),
                patch.object(changelog, "load_catalog", return_value=catalog),
                contextlib.redirect_stdout(io.StringIO()),
            ):
                self.assertEqual(changelog.main(), 0)
            with output.open(newline="", encoding="utf-8") as stream:
                rows = list(csv.DictReader(stream))
            self.assertEqual([r["pkg"] for r in rows], ["pgvector", "postgis"])
            self.assertEqual((rows[0]["old_ver"], rows[0]["new_ver"], rows[0]["note_zh"]), ("", "0.8.0", "新增"))
            self.assertEqual(rows[1]["new_ver"], "3.6.3-2")
            self.assertEqual(rows[1]["note_en"], "RPM build; DEB build; RPM 3.6.3-2, DEB 3.6.3-1")
            self.assertEqual(rows[1]["note_zh"], "RPM 构建；DEB 构建；RPM 3.6.3-2，DEB 3.6.3-1")


if __name__ == "__main__":
    unittest.main()
