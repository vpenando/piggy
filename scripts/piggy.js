'use strict';

const searchBar = document.getElementById('search-bar');

searchBar.addEventListener('input', function() {
    table.search(searchBar.value || '');
    refresh();
});

function formatDateSelectorContent(d) {
    return d.getFullYear() + '-'
        + ('0' + (d.getMonth() + 1)).slice(-2) + '-'
        + ('0' + d.getDate()).slice(-2);
}

function refresh() {
    index = 0;
    mainTableBody.innerHTML = '';
    table.items.forEach(addRow);
    if (typeof onOperationsLoaded === 'function') {
        onOperationsLoaded();
    }
}

function updateTotal() {
    total = 0;
    table.items.forEach(it => total += it.operation.amount);
}

async function loadOperations() {
    const url = '/operations?year=' + date.getFullYear() + '&month=' + (date.getMonth()+1);
    new HttpRequest('GET', url)
        .onSuccess(r => {
            const operations = JSON.parse(r.responseText) || [];
            operations.forEach(op => {
                op.date = new Date(op.date);
                op.creation_date = new Date(op.creation_date);
                table.addOperation(op);
            });
            refresh();
        })
        .onError(r => {
            throw new Error('ERROR: ' + r);
        })
        .send();
}

async function loadPage() {
    table.reset();
    new HttpRequest('GET', '/categories')
        .onSuccess(r => {
            const cats = JSON.parse(r.responseText) || [];
            cats.forEach(c => table.addCategory(c));
            loadOperations();
        })
        .onError(r => {
            throw new Error('ERROR: ' + r);
        })
        .send();
    new HttpRequest('GET', '/months')
        .onSuccess(r => {
            const months = JSON.parse(r.responseText) || [];
            if (months != []) {
                monthLabel.innerHTML = months[date.getMonth()] + ' ' + date.getFullYear();
            }
        })
        .onError(r => {
            throw new Error('ERROR: ' + r);
        })
        .send();
}