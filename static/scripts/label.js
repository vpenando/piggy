'use strict';

class Label {
    constructor(content, cssClass) {
        this._content = this._escapeHTML(content);
        this._class = cssClass;
    }

    _escapeHTML(str) {
        return str.replace(/&/g,'&amp;')
            .replace(/</g,'&lt;')
            .replace(/>/g,'&gt;');
    }

    toHTML() {
        if (this._class) {
            return '<label class="' + this._class + '">' + this._content + '</label>';
        }
        return '<label>' + this._content + '</label>';
    }
}