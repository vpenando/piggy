'use strict';

class ComboBoxOption {
    constructor(value, selected=false) {
        this._value = value;
        this._selected = selected;
    }

    toHTML() {
        if (this._selected) {
            return '<option selected>' + this._value + '</option>';
        } else {
            return '<option>' + this._value + '</option>';
        }
    }
}

class ComboBox {
    constructor(id, c, choices, selectedIndex) {
        this._id = id;
        this._class = c;
        this._choices = choices;
        this._selectedIndex = selectedIndex;
    }

    toHTML() {
        let options = '';
        for (let i = 0; i < this._choices.length; i++) {
            const selected = i == this._selectedIndex;
            options += '\t' + new ComboBoxOption(this._choices[i], selected).toHTML() + '\n';
        }

        let result = '<select';
        if (this._id) {
            result += ' id="' + this._id + '"';
        }
        if (this._class) {
            result += ' class="' + this._class + '"';
        }
        result += '>\n';
        result += options + '</select>\n';
        return result;
    }
}