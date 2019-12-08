var http = require("http");
var fs = require("fs");

http
  .createServer((request, response) => {
    console.log("require come", request.url);

    const html = fs.readFileSync("./index.html");
    response.writeHead(200, {
      "Content-Type": "text/html"
    });
    response.end(html);
  })
  .listen(8082);

console.log("listen at localhost:8082");
