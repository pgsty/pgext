/**
 * search-engine.js — deterministic page search for the Command Palette.
 *
 * The index builder is injected so this module can be unit-tested without a
 * DOM. Both Latin/Lunr and CJK substring paths apply keywords and the same
 * final per-document boost multiplier.
 */
(function () {
  'use strict';

  var CJK = /[぀-ヿ㐀-䶿一-鿿豈-﫿]/;

  function number(value, fallback) {
    var parsed = Number(value);
    return Number.isFinite(parsed) && parsed > 0 ? parsed : fallback;
  }

  function lexical(a, b) {
    a = String(a || '');
    b = String(b || '');
    return a < b ? -1 : a > b ? 1 : 0;
  }

  function compare(a, b) {
    return (
      b.score - a.score ||
      lexical(a.doc.title, b.doc.title) ||
      lexical(a.doc.ref, b.doc.ref)
    );
  }

  function group(results) {
    var groups = [];
    var groupByKey = new Map();
    (results || []).forEach(function (result) {
      var rawKey = result.doc.root || result.doc.type || 'pages';
      var key = String(rawKey).trim().toLowerCase() || 'pages';
      var current = groupByKey.get(key);
      if (!current) {
        current = {
          key: key,
          label:
            (result.doc.breadcrumb && result.doc.breadcrumb[0]) ||
            result.doc.root ||
            result.doc.type ||
            'Pages',
          results: [],
        };
        groupByKey.set(key, current);
        groups.push(current);
      }
      current.results.push(result);
    });
    return groups;
  }

  function create(documents, lunrApi, maxResults) {
    var docs = documents || [];
    var limit = number(maxResults, 10);
    var docByRef = new Map();
    docs.forEach(function (doc) {
      docByRef.set(doc.ref, doc);
    });

    // Case-fold every searchable field once. queryCjk scans the whole corpus on
    // every keystroke, so folding here avoids re-allocating a lowercase copy of
    // each document per character typed.
    var folded = docs.map(function (doc) {
      var keywordText = Array.isArray(doc.keywords)
        ? doc.keywords.join(' ')
        : doc.keywords || '';
      return {
        title: (doc.title || '').toLowerCase(),
        keywordText: keywordText,
        keywords: keywordText.toLowerCase(),
        headings: (doc.headings || '').toLowerCase(),
        description: (doc.description || '').toLowerCase(),
        body: (doc.body || '').toLowerCase(),
      };
    });

    var index = lunrApi(function () {
      this.ref('ref');
      this.field('title', { boost: 5 });
      this.field('categories', { boost: 3 });
      this.field('tags', { boost: 3 });
      this.field('keywords', { boost: 4 });
      this.field('headings', { boost: 3 });
      this.field('description', { boost: 2 });
      this.field('body');
      docs.forEach(function (doc) {
        this.add(doc);
      }, this);
    });

    function queryCjk(query) {
      var needle = query.toLowerCase();
      var hits = [];
      docs.forEach(function (doc, docIndex) {
        var fields = folded[docIndex];
        var titleAt = fields.title.indexOf(needle);
        var keywordText = fields.keywordText;
        var keywordAt = fields.keywords.indexOf(needle);
        var headingAt = fields.headings.indexOf(needle);
        var descAt = fields.description.indexOf(needle);
        var bodyAt = fields.body.indexOf(needle);
        var textScore =
          (titleAt >= 0 ? 100 : 0) +
          (keywordAt >= 0 ? 80 : 0) +
          (headingAt >= 0 ? 50 : 0) +
          (descAt >= 0 ? 30 : 0) +
          (bodyAt >= 0 ? 10 : 0);
        if (!textScore) return;
        var excerpt = doc.excerpt || '';
        if (bodyAt >= 0) {
          var start = Math.max(0, bodyAt - 24);
          excerpt =
            (start > 0 ? '…' : '') +
            doc.body.slice(start, bodyAt + 56) +
            '…';
        } else if (descAt >= 0) {
          excerpt = doc.description;
        } else if (keywordAt >= 0) {
          excerpt = keywordText;
        }
        hits.push({
          doc: doc,
          excerpt: excerpt,
          score: textScore * number(doc.boost, 1),
          textScore: textScore,
        });
      });
      return hits.sort(compare).slice(0, limit);
    }

    function queryLatin(query) {
      var found = index.query(function (builder) {
        lunrApi.tokenizer(query.toLowerCase()).forEach(function (token) {
          var term = token.toString();
          builder.term(term, { boost: 100 });
          builder.term(term, {
            wildcard:
              lunrApi.Query.wildcard.LEADING |
              lunrApi.Query.wildcard.TRAILING,
            boost: 10,
          });
          builder.term(term, { editDistance: 2 });
        });
      });
      return found
        .map(function (result) {
          var doc = docByRef.get(result.ref);
          return doc
            ? {
                doc: doc,
                excerpt: doc.excerpt || doc.description || '',
                score: result.score * number(doc.boost, 1),
                textScore: result.score,
              }
            : null;
        })
        .filter(Boolean)
        .sort(compare)
        .slice(0, limit);
    }

    return {
      query: function (query) {
        return CJK.test(query) ? queryCjk(query) : queryLatin(query);
      },
      queryCjk: queryCjk,
      queryLatin: queryLatin,
    };
  }

  window.OinkSearchEngine = {
    CJK: CJK,
    compare: compare,
    group: group,
    lexical: lexical,
    create: create,
  };
  if (typeof module === 'object' && module.exports)
    module.exports = window.OinkSearchEngine;
})();
