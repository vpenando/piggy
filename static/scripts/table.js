'use strict';

const TableItemStatus = {
    DISABLED: -1,
    NONE: 0,
    NEW: 1,
    UPDATED: 2,
    DELETED: 3
};

class TableItem {
    constructor(operation, status=TableItemStatus.NONE) {
        this._operation = operation;
        this.status = status;
    }

    get operation() {
        return this._operation;
    }

    markAsUpdated() {
        if (this.operation.id != 0) {
            this.status = TableItemStatus.UPDATED;
        }
    }

    markAsDeleted() {
        this.status = TableItemStatus.DELETED;
    }
}

const selectors = [
    it => (it.operation.category != 0) ? table.categories[it.operation.category].name.toLowerCase() : '',
    it => it.operation.date,
    it => it.operation.description.toLowerCase(),
    it => it.operation.amount,
    it => it.operation.creation_date
];

const Column = {
    CATEGORY: 0,
    DATE: 1,
    DESCRIPTION: 2,
    AMOUNT: 3,
    CREATION_DATE: 4
};

const sortAscending = (lhs, rhs) => lhs > rhs ? 1 : rhs > lhs ? -1 : 0;
const sortDescending = (lhs, rhs) => lhs < rhs ? 1 : rhs < lhs ? -1 : 0;

class Table {
    constructor() {
        this._searchText = '';
        this._index = 0;
        this._date = new Date();
        this._ascending = true;
        this._lastSortColumn = -1;
        this.reset();
    }

    get categories() {
        return this._categories;
    }

    get items() {
        return this.search(this._searchText);
    }

    get newItems() {
        return (this._itemsSource.filter(it => it.status == TableItemStatus.NEW));
    }

    get updatedItems() {
        return (this._itemsSource.filter(it => it.status == TableItemStatus.UPDATED));
    }

    get deletedItems() {
        return (this._itemsSource.filter(it => it.operation.id != 0 && it.status == TableItemStatus.DELETED));
    }

    get itemsSource() {
        return this._itemsSource.filter(it => it.status != TableItemStatus.DELETED);
    }

    sortBy(column) {
        if (column == this._lastSortColumn) {
            this._ascending = !this._ascending;
        } else {
            this._ascending = true;
        }
        const ascending = this._ascending;
        const sortMethod = function(item1, item2) {
            const selector = selectors[column];
            const lhs = selector(item1);
            const rhs = selector(item2);
            if (ascending) {
                return sortAscending(lhs, rhs);
            } else {
                return sortDescending(lhs, rhs);
            }
        }
        this._itemsSource.sort(sortMethod);
        this._lastSortColumn = column;
    }

    search(text) {
        this._searchText = text;
        if (this._searchText == '') {
            return this.itemsSource;
        }
        return this.itemsSource.filter(it => {
            const op = it.operation;
            const searchText = this._searchText.toLowerCase();
            return op.description.toLowerCase().includes(searchText)
                || (op.category != 0 && this._categories[op.category].name.toLowerCase().includes(searchText));
        });
    }

    reset() {
        this._categories = [{id:0, name:''}];
        this._itemsSource = [];
    }

    itemAt = (index) => this.items[index];
    addCategory = (category) => this._categories.push(category);
    addItem = (item) => this._itemsSource.push(item);
    addOperation = (operation) => this.addItem(new TableItem(operation));
    deleteItem = (index) => this.items[index].markAsDeleted();
}