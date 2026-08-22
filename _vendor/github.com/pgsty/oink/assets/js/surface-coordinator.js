/**
 * surface-coordinator.js — coordinate mutually exclusive transient UI.
 *
 * A surface registers a close callback and may request that the coordinator
 * close every other registered surface before it opens. The callbacks own
 * their own focus-restoration policy; coordinated closes deliberately do not
 * steal focus from the surface being opened.
 */
(function () {
  'use strict';

  var surfaces = new Map();

  function register(name, close) {
    if (!name || typeof close !== 'function') return function () {};
    var callbacks = surfaces.get(name) || [];
    callbacks.push(close);
    surfaces.set(name, callbacks);
    return function () {
      var current = surfaces.get(name) || [];
      current = current.filter(function (callback) {
        return callback !== close;
      });
      if (current.length) {
        surfaces.set(name, current);
      } else {
        surfaces.delete(name);
      }
    };
  }

  function closeOthers(activeName, keepNames) {
    var keep = new Set(keepNames || []);
    surfaces.forEach(function (callbacks, name) {
      if (name !== activeName && !keep.has(name))
        callbacks.slice().forEach(function (close) {
          close(false);
        });
    });
  }

  function closeAll(restoreFocus) {
    surfaces.forEach(function (callbacks) {
      callbacks.slice().forEach(function (close) {
        close(restoreFocus === true);
      });
    });
  }

  window.OinkSurfaceCoordinator = {
    register: register,
    closeOthers: closeOthers,
    closeAll: closeAll,
  };
})();
