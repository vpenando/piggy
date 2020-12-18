'use strict';

class Input {
    constructor(inputType) {
        this._inputType = inputType;
        this.properties = new Map();
    }

    addProperty(value, key) {
        this.properties.set(value, key);
    }

    toHTML() {
        let result ='<input type="' + this._inputType + '"';
        this.properties.forEach(function(value, key, _) {
            result += ' ' + key + '="' + value + '"';
        });
        result += '></input>';
        return result;
    }
}

class DateInput extends Input {
    constructor(cssClass, id) {
        super('date');
        this.addProperty('class', cssClass);
        this.addProperty('id', id);
    }
}

class ImageInput extends Input {
    constructor(src, onclick, title) {
        super('image');
        this.addProperty('src', src);
        this.addProperty('onclick', onclick);
        this.addProperty('title', title);
    }
}

class NumberInput extends Input {
    constructor(id) {
        super('number');
        this.addProperty('id', id);
    }
}

class TextInput extends Input {
    constructor(cssClass, id, value) {
        super('text');
        this.addProperty('class', cssClass);
        this.addProperty('id', id);
        this.addProperty('value', value);
    }
}