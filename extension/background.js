var username = "paulil";
var password = "Smile@1314";
var retry = 3;

chrome.webRequest.onAuthRequired.addListener(
  function handler(details) {
    if (--retry < 0)
      return {cancel: true};
    return {authCredentials: {username: username, password: password}};
  },
  {urls: ["<all_urls>"]},
  ['blocking']
);
