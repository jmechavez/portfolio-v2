// static/js/main.js
window.loadStylesheets = function () {
  const stylesheets = [
    "/static/css/styles.css",
    "/static/css/components/header.css",
    "/static/css/components/welcome.css",
    "/static/css/components/about.css",
    "/static/css/components/footer.css",
    "/static/css/utils.css",
  ];
  stylesheets.forEach((url) => {
    if (!document.querySelector(`link[href="${url}"]`)) {
      const link = document.createElement("link");
      link.rel = "stylesheet";
      link.href = url;
      document.head.appendChild(link);
      console.log(`Loaded stylesheet: ${url}`);
    }
  });
};

console.log("main.js loaded");
