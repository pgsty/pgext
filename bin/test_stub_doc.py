#!/usr/bin/env python3
import importlib.util
import inspect
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
    class FakeConn:
        def __init__(self, names):
            self.names = names
            self.executed = []
            self.synced = []

        def execute(self, sql):
            self.executed.append(sql)

        def query_csv(self, query):
            self.assert_universe_query(query)
            return [{"ext": name} for name in self.names]

        def assert_universe_query(self, query):
            if "pgext.universe4" not in query:
                raise AssertionError(f"unexpected catalog query: {query}")

        def sync_docs_from_csv(self, rows, expected_count, commit=True):
            self.synced.append((list(rows), expected_count, commit))

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

    def test_load_stub_docs_requires_complete_bilingual_universe(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        conn = self.FakeConn(["a", "b"])
        with self.assertRaisesRegex(ValueError, "missing_zh=1"):
            stub_doc.load_stub_docs(
                conn,
                {"a": "English A", "b": "English B"},
                {"a": "中文 A"},
            )

        self.assertEqual(conn.synced, [])

    def test_load_stub_docs_syncs_one_filtered_atomic_stage(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        conn = self.FakeConn(["a", "b"])
        stats = stub_doc.load_stub_docs(
            conn,
            {"a": "English A", "b": "English B", "notes": "not a stub"},
            {"a": "中文 A", "b": "中文 B", "task": "not a stub"},
        )

        self.assertEqual(
            conn.synced,
            [([("a", "English A", "中文 A"), ("b", "English B", "中文 B")], 2, True)],
        )
        self.assertEqual(stats["base_rows"], 2)
        self.assertEqual(stats["extra_en"], ["notes"])
        self.assertEqual(stats["extra_zh"], ["task"])

    def test_load_stub_docs_dry_run_rolls_back_staged_sync(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        conn = self.FakeConn(["a"])
        stats = stub_doc.load_stub_docs(
            conn,
            {"a": "English A"},
            {"a": "中文 A"},
            dry_run=True,
        )

        self.assertEqual(conn.synced, [([("a", "English A", "中文 A")], 1, False)])
        self.assertTrue(stats["dry_run"])

    def test_doc_ddl_is_not_foreign_keyed_to_rebuildable_catalogs(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        self.assertNotIn("REFERENCES pgext.extension", stub_doc.DOC_TABLE_SQL)
        self.assertNotIn("REFERENCES pgext.universe4", stub_doc.DOC_TABLE_SQL)
        self.assertIn("cargo_url", stub_doc.DOC_TABLE_SQL)

    def test_live_doc_comments_are_updated_inside_atomic_sync(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        source = inspect.getsource(stub_doc.PsqlClient.sync_docs_from_csv)
        comment_pos = source.index("COMMENT ON TABLE pgext.doc")
        transaction_end_pos = source.index("{transaction_end}")
        self.assertLess(comment_pos, transaction_end_pos)
        self.assertIn("against pgext.universe4", source)
        self.assertIn("COMMENT ON COLUMN pgext.doc.cargo_url", source)
        self.assertIn("pgext.doc metadata differs from pgext.universe4", source)
        self.assertIn("pgext.doc still has a foreign key", source)
        self.assertIn("COMMENT ON TABLE {DOC_BACKUP_TABLE}", source)
        self.assertIn("Historical immutable pgext.extension", source)

    def test_atomic_sync_replaces_rows_when_universe_ids_move(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        source = inspect.getsource(stub_doc.PsqlClient.sync_docs_from_csv)
        delete_pos = source.index("DELETE FROM pgext.doc;")
        insert_pos = source.index("INSERT INTO pgext.doc (")
        self.assertLess(delete_pos, insert_pos)
        self.assertNotIn("ON CONFLICT (id)", source)
        self.assertNotIn("id/name mapping conflicts", source)

    def test_initial_backup_is_built_by_name_from_the_staged_docs(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        source = inspect.getsource(stub_doc.PsqlClient.sync_docs_from_csv)
        self.assertIn("JOIN stub_doc_load AS s ON s.ext = e.name", source)
        self.assertIn("pgext.extension is not covered", source)
        self.assertIn("documentation baseline backup is empty or incomplete", source)

    def test_historical_backup_does_not_freeze_live_markdown(self):
        if stub_doc is None:
            self.fail("bin/stub_doc.py is missing")

        source = inspect.getsource(stub_doc.PsqlClient.sync_docs_from_csv)
        self.assertNotIn("baseline Markdown bodies differ", source)
        self.assertNotIn("baseline Markdown bodies changed", source)
        self.assertNotIn("documentation baseline backup is not the exact extension set", source)


if __name__ == "__main__":
    unittest.main()
