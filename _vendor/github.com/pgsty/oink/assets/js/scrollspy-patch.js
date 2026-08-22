/*
 * Oink compatibility patch based on Bootstrap PR #41726:
 * https://github.com/twbs/bootstrap/pull/41726
 */

(function () {
  'use strict';

  // Wait for Bootstrap to be available
  if (typeof bootstrap === 'undefined' || !bootstrap.ScrollSpy) {
    console.warn('[Oink] ScrollSpy patch: Bootstrap not found, patch skipped');
    return;
  }

  const ScrollSpy = bootstrap.ScrollSpy;

  // Check if _initializeTargetsAndObservables exists
  if (!ScrollSpy.prototype._initializeTargetsAndObservables) {
    console.warn(
      '[Oink] ScrollSpy patch: _initializeTargetsAndObservables method not found, patch skipped',
    );
    return;
  }

  // parseSelector function from Bootstrap PR #41726
  // Escapes CSS selectors, particularly IDs, to handle special characters
  function parseSelector(selector) {
    if (selector && window.CSS && window.CSS.escape) {
      // document.querySelector needs escaping to handle IDs (html5+) containing special characters
      // Match # followed by ID characters (excluding quotes and #)
      selector = selector.replace(
        /#([^\s"']+)/g,
        (match, id) => `#${CSS.escape(id)}`,
      );
    }
    return selector;
  }

  // SelectorEngine - mimics Bootstrap's SelectorEngine API using native DOM methods
  const SelectorEngine = {
    find(selector, element) {
      if (!element) {
        return [];
      }
      return Array.from(element.querySelectorAll(selector));
    },
    findOne(selector, element) {
      if (!element) {
        return null;
      }
      return element.querySelector(selector);
    },
  };

  const SELECTOR_TARGET_LINKS = '[href]';

  // Simple fallback implementations matching Bootstrap's behavior
  // These match the internal isDisabled and isVisible utilities
  function isDisabled(element) {
    return (
      element.hasAttribute('disabled') || element.classList.contains('disabled')
    );
  }

  function isVisible(element) {
    return (
      element &&
      element.offsetParent !== null &&
      window.getComputedStyle(element).visibility !== 'hidden'
    );
  }

  // Bootstrap's _initializeTargetsAndObservables with escaped selectors.
  // prettier-ignore-start
  function patchedInitializeTargetsAndObservables() {
    this._targetLinks = new Map()
    this._observableSections = new Map()

    const targetLinks = SelectorEngine.find(SELECTOR_TARGET_LINKS, this._config.target)

    for (const anchor of targetLinks) {
      // ensure that the anchor has an id and is not disabled
      if (!anchor.hash || isDisabled(anchor)) {
        continue
      }

      const decodedHash = decodeURI(anchor.hash)
      const escapedSelector = parseSelector(decodedHash)
      const observableSection = SelectorEngine.findOne(escapedSelector, this._element)

      // ensure that the observableSection exists & is visible
      if (isVisible(observableSection)) {
        this._targetLinks.set(decodedHash, anchor)
        this._observableSections.set(anchor.hash, observableSection)
      }
    }
  }
  // prettier-ignore-end

  // Apply the patched method
  ScrollSpy.prototype._initializeTargetsAndObservables =
    patchedInitializeTargetsAndObservables;
})();
