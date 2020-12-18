'use strict';

const searchBar = document.getElementById('search-bar');

searchBar.addEventListener('input', function() {
    table.search(searchBar.value || '');
    fillTable();
});

function fillTable() {
    index = 0;
    mainTableBody.innerHTML = '';
    table.items.forEach(addRow);
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
            fillTable();
            if (typeof onOperationsLoaded === 'function') {
                onOperationsLoaded();
            }
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