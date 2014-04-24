require.config({
  baseUrl: '/public/vendor'
});

require({
  paths: {
    scribe: 'scribe/scribe',
    'scribe-plugin-blockquote-command': 'scribe-plugin-blockquote-command/scribe-plugin-blockquote-command',
    'scribe-plugin-curly-quotes': 'scribe-plugin-curly-quotes/scribe-plugin-curly-quotes',
    'scribe-plugin-formatter-plain-text-convert-new-lines-to-html': 'scribe-plugin-formatter-plain-text-convert-new-lines-to-html/scribe-plugin-formatter-plain-text-convert-new-lines-to-html',
    'scribe-plugin-heading-command': 'scribe-plugin-heading-command/scribe-plugin-heading-command',
    'scribe-plugin-intelligent-unlink-command': 'scribe-plugin-intelligent-unlink-command/scribe-plugin-intelligent-unlink-command',
    'scribe-plugin-keyboard-shortcuts': 'scribe-plugin-keyboard-shortcuts/scribe-plugin-keyboard-shortcuts',
    'scribe-plugin-link-prompt-command': 'scribe-plugin-link-prompt-command/scribe-plugin-link-prompt-command',
    'scribe-plugin-sanitizer': 'scribe-plugin-sanitizer/scribe-plugin-sanitizer',
    'scribe-plugin-smart-lists': 'scribe-plugin-smart-lists/scribe-plugin-smart-lists',
    'scribe-plugin-toolbar': 'scribe-plugin-toolbar/scribe-plugin-toolbar',
    'upload-light': 'upload-light/lib/upload-light'
    }
  }, [
    'scribe',
    'scribe-plugin-blockquote-command',
    'scribe-plugin-curly-quotes',
    'scribe-plugin-formatter-plain-text-convert-new-lines-to-html',
    'scribe-plugin-heading-command',
    'scribe-plugin-intelligent-unlink-command',
    'scribe-plugin-keyboard-shortcuts',
    'scribe-plugin-link-prompt-command',
    'scribe-plugin-sanitizer',
    'scribe-plugin-smart-lists',
    'scribe-plugin-toolbar',
    'upload-light'
  ], function(
    Scribe,
    scribePluginBlockquoteCommand,
    scribePluginCurlyQuotes,
    scribePluginFormatterPlainTextConvertNewLinesToHtml,
    scribePluginHeadingCommand,
    scribePluginIntelligentUnlinkCommand,
    scribePluginKeyboardShortcuts,
    scribePluginLinkPromptCommand,
    scribePluginSanitizer,
    scribePluginSmartLists,
    scribePluginToolbar,
    UploadLight
  ) {
    // Configure Scribe
    var scribe = new Scribe(document.querySelector('.scribe'), {allowBlockElements: true})

    scribe.on('content-changed', function() {
      document.querySelector('.scribe-html').textContent = scribe.getHTML()
    });

    // Keyboard shortcuts
    var ctrlKey = function(event) { event.metaKey || event.ctrlKey}

    var commandsToKeyboardShortcutsMap = Object.freeze({
      bold: function(event) { event.metaKey && event.keyCode === 66 }, // b
      italic: function(event) { event.metaKey && event.keyCode === 73 }, // i
      strikeThrough: function(event) { event.altKey && event.shiftKey && event.keyCode === 83 }, // s
      removeFormat: function(event) { event.altKey && event.shiftKey && event.keyCode === 65 }, // a
      linkPrompt: function(event) { event.metaKey && !event.shiftKey && event.keyCode === 75 }, // k
      unlink: function(event) { event.metaKey && event.shiftKey && event.keyCode === 75  }, //k,
      insertUnorderedList: function(event) { event.altKey && event.shiftKey && event.keyCode === 66 }, // b
      insertOrderedList: function(event) { event.altKey && event.shiftKey && event.keyCode === 78 }, // n
      blockquote: function(event) { event.altKey && event.shiftKey && event.keyCode === 87 }, // w
      h2: function(event) { ctrlKey(event) && event.keyCode === 50 } // 2
    });

    // Plugin
    scribe.use(scribePluginBlockquoteCommand());
    scribe.use(scribePluginHeadingCommand(2));
    scribe.use(scribePluginIntelligentUnlinkCommand());
    scribe.use(scribePluginLinkPromptCommand());
    scribe.use(scribePluginToolbar(document.querySelector('.scribe-toolbar')));
    scribe.use(scribePluginSmartLists());
    scribe.use(scribePluginCurlyQuotes());
    scribe.use(scribePluginKeyboardShortcuts(commandsToKeyboardShortcutsMap));

    // Formatters
    scribe.use(scribePluginSanitizer({
      tags: {
        p: {},
        br: {},
        b: {},
        strong: {},
        i: {},
        s: {},
        blockquote: {},
        ol: {},
        ul: {},
        li: {},
        a: {href: true},
        h2: {},
        img: {src: true},
      }
    }
    ));

    scribe.use(scribePluginFormatterPlainTextConvertNewLinesToHtml())

    // Configure upload-light
    new UploadLight(
      document.querySelector('.upload-light'),
      document.querySelector('#select-image'),
      {
        max: 1
    }).on('load', function(message) {
      // Stop spinning
      var parsed = JSON.parse(message);
      if (parsed.url !== "undefined") {
        scribe.insertHTML("<img src=\"#{url}\" class=\"img-responsive\">");
      }
    }).on('error', function(message) {
      // Stop spinning
      console.error(JSON.parse(message));
    });
});
