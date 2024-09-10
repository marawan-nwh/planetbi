import { hasValue } from "$lib/js/common.js";

let BaseURI = "https://planet.bi";

// https://stackoverflow.com/a/71601719/17931895
if (import.meta.env.DEV) {
  BaseURI = "http://localhost:8080";
}

// http://stackoverflow.com/a/14525299/1420619
let paramterize = function (data) {
  return Object.keys(data)
    .map(function (k) {
      return encodeURIComponent(k) + "=" + encodeURIComponent(data[k]);
    })
    .join("&");
};

let parseJSON = function (str) {
  try {
    return JSON.parse(str);
  } catch (e) {
    return {};
  }
};

// HTTP
let act = function (xhr, success, error) {
  error = error || function () {};
  return function () {
    if (xhr.readyState == 4) {
      let response =
        xhr.response == "" ? xhr.response : parseJSON(xhr.response);
      if (xhr.status == 200) {
        success(response, xhr.status);
      } else {
        error(response, xhr.status);
      }
    }
  };
};

export let http = {
  get: function (url, success, error, headers, timeout, ontimeout) {
    let xhr = new XMLHttpRequest();
    xhr.onreadystatechange = act(xhr, success, error);
    xhr.open("GET", BaseURI + url, true);
    if (!!headers && Object.keys(headers).length > 0) {
      let hs = Object.keys(headers);
      for (let i = 0; i < hs.length; i++) {
        const element = hs[i];
        xhr.setRequestHeader(element, headers[element]);
      }
    }
    if (hasValue(timeout)) {
      xhr.timeout = timeout;
    }
    if (hasValue(ontimeout)) {
      xhr.ontimeout = ontimeout;
    }
    xhr.send();
  },
  json: function (url, params, success, error, headers, timeout, ontimeout) {
    let xhr = new XMLHttpRequest();
    xhr.onreadystatechange = act(xhr, success, error);
    xhr.open("POST", BaseURI + url, true);
    xhr.setRequestHeader("Content-type", "application/json");
    if (!!headers && Object.keys(headers).length > 0) {
      let hs = Object.keys(headers);
      for (let i = 0; i < hs.length; i++) {
        const element = hs[i];
        xhr.setRequestHeader(element, headers[element]);
      }
    }
    if (hasValue(timeout)) {
      xhr.timeout = timeout;
    }
    if (hasValue(ontimeout)) {
      xhr.ontimeout = ontimeout;
    }
    xhr.send(JSON.stringify(params));
  },
  post: function (url, params, success, error, headers, timeout, ontimeout) {
    let xhr = new XMLHttpRequest();

    // TODO: Check if we can/should drop this in production
    xhr.withCredentials = true;

    xhr.onreadystatechange = act(xhr, success, error);
    xhr.open("POST", BaseURI + url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
    if (!!headers && Object.keys(headers).length > 0) {
      let hs = Object.keys(headers);
      for (let i = 0; i < hs.length; i++) {
        const element = hs[i];
        xhr.setRequestHeader(element, headers[element]);
      }
    }
    if (hasValue(timeout)) {
      xhr.timeout = timeout;
    }
    if (hasValue(ontimeout)) {
      xhr.ontimeout = ontimeout;
    }
    xhr.send(paramterize(params));
  },
  upload: function (url, params, success, error, headers, timeout, ontimeout) {
    let xhr = new XMLHttpRequest();
    xhr.onreadystatechange = act(xhr, success, error);
    xhr.open("POST", BaseURI + url, true);
    if (!!headers && Object.keys(headers).length > 0) {
      let hs = Object.keys(headers);
      for (let i = 0; i < hs.length; i++) {
        const element = hs[i];
        xhr.setRequestHeader(element, headers[element]);
      }
    }
    if (hasValue(timeout)) {
      xhr.timeout = timeout;
    }
    if (hasValue(ontimeout)) {
      xhr.ontimeout = ontimeout;
    }
    let formData = new FormData();
    for (i in params) {
      formData.append(i, params[i]);
    }
    xhr.send(formData);
  },
};
