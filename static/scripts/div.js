'use strict';

class Div {
    constructor(content, cssClass) {
        if (content) {
            this._content = [...content];
        } else {
            this._content = [];
        }
        this._class = cssClass;
    }

    toHTML() {
        let result = '<div';
        if (this._class) {
            result += ' class="' + this._class + '"';
        }
        result += '>';
        this._content.forEach(c => result += c);
        result += '</div>';
        return result;
    }
}