/* eslint-disable */

// Service Worker
const cacheName = "Logger.Fitness";
const filesToCache = [
  "/",
  "index.html",
  "./js/index.js",
  "./styles/styles.css"
];

self.addEventListener("install", event => {
  console.log("[ServiceWorker**] Install");
  event.waitUntil(
    caches.open(cacheName).then(cache => {
      console.log("[ServiceWorker**] Caching app shell");
      return cache.addAll(filesToCache);
    })
  );
});

self.addEventListener("activate", event => {
  caches.keys().then(keyList => {
    return Promise.all(
      keyList.map(key => {
        if (key !== cacheName) {
          console.log("[ServiceWorker] - Removing old cache", key);
          return caches.delete(key);
        }
      })
    );
  });
});

self.addEventListener("fetch", event => {
  event.respondWith(
    caches.match(event.request, { ignoreSearch: true }).then(response => {
      return response || fetch(event.request);
    })
  );
});
