#!/usr/bin/env python3
import importlib.util
import pathlib
import sys
import tempfile
import unittest


sys.dont_write_bytecode = True

MODULE_PATH = pathlib.Path(__file__).with_name("stub_doc.py")
if MODULE_PATH.exists():
    spec = importlib.util.spec_from_file_location("stub_doc", MODULE_PATH)
    stub_doc = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(stub_doc)
else:
    stub_doc = None


class StubDocFileSyncTest(unittest.TestCase):
    def test_collect_stub_docs_reads_markdown_by_extension_name(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        with tempfile.TemporaryDirectory() as tmp:
            stub_dir = pathlib.Path(tmp)
            (stub_dir / "vector.md").write_text("# vector\n\nEnglish docs\n", encoding="utf-8")
            (stub_dir / "README.txt").write_text("ignored", encoding="utf-8")

            docs = stub_doc.collect_stub_docs(stub_dir)

        self.assertEqual(docs, {"vector": "# vector\n\nEnglish docs\n"})

    def test_safe_stub_name_rejects_path_components(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        self.assertEqual(stub_doc.safe_stub_name("uuid-ossp"), "uuid-ossp")
        with self.assertRaises(ValueError):
            stub_doc.safe_stub_name("../bad")
        with self.assertRaises(ValueError):
            stub_doc.safe_stub_name("bad/name")

    def test_dump_stub_docs_writes_non_null_docs(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        rows = [
            {"ext": "vector", "en_doc": "# vector\n", "zh_doc": "# 向量\n"},
            {"ext": "missing_zh", "en_doc": "# English only\n", "zh_doc": None},
        ]

        with tempfile.TemporaryDirectory() as tmp:
            root = pathlib.Path(tmp)
            en_count, zh_count = stub_doc.dump_stub_docs(rows, root / "stub", root / "stub-zh")

            self.assertEqual((en_count, zh_count), (2, 1))
            self.assertEqual((root / "stub" / "vector.md").read_text(encoding="utf-8"), "# vector\n")
            self.assertEqual((root / "stub-zh" / "vector.md").read_text(encoding="utf-8"), "# 向量\n")
            self.assertFalse((root / "stub-zh" / "missing_zh.md").exists())

    def test_connect_uses_psql_client_without_python_database_dependency(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        try:
            client = stub_doc.connect("postgres:///data")
        except ModuleNotFoundError as exc:
            self.fail(f"connect should not require a Python database driver: {exc}")

        self.assertEqual(client.pgurl, "postgres:///data")
        self.assertTrue(hasattr(client, "execute"))


if __name__ == "__main__":
    unittest.main()
