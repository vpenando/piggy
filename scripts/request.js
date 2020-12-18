'use strict';

const HttpRequestStatus = {
    UNSENT: 0,
    OPENED: 1,
    HEADERS_RECEIVED: 2,
    LOADING: 3,
    DONE: 4
};

class HttpRequest {
    constructor(method, url) {
        this._request = new XMLHttpRequest();
        this._request.open(method, url);
        this._request.setRequestHeader("Content-type", "application/json");
        this._method = method;
        this._url = url;
        this._onSuccess = function() {};
        this._onError = function() {};
        this._onDone = function() {};
    }

    onSuccess(callback) {
        this._onSuccess = callback;
        return this;
    }

    onError(callback) {
        this._onError = callback;
        return this;
    }

    send(body) {
        let success = this._onSuccess;
        let error = this._onError;
        let request = this._request;
        request.onreadystatechange = () => {
            if (request.readyState == HttpRequestStatus.DONE) {
                if (request.status == 200) {
                    success(request);
                } else {
                    error(request);
                }
            }
        }
        request.send(body);
    }
}
