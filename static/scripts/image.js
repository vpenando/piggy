'use strict';

class Size {
    constructor(w, h) {
        this.width = w;
        this.height = h;
    }
}

class Image {
    constructor(src, cssClass, size) {
        this._src = src;
        this._cssClass = cssClass;
        this._size = size;
    }

    toHTML() {
        let result = '<img src="'+ this._src + '"';
        if (this._cssClass) {
            result += ' class="' + this._cssClass + '"';
        }
        if (this._size) {
            result += ' width="' + this._size.width + '"';
            result += ' height="' + this._size.height + '"';
        }
        result += ' />';
        return result;
    }
}