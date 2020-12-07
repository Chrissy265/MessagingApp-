var http = require("http");
var url = require("url");
var fs = require("fs");
const { Router } = require("express");
var server = http.createServer(function (request, response) {
  css(request, response);
  var path = url.parse(request.url).pathname;
  switch (path) {
    case "/login.html":
      fs.readFile(__dirname + path, function (error, data) {
        if (error) {
          response.writeHead(404);
          response.write(error);
          response.end();
        } else {
          response.writeHead(200, {
            "Content-Type": "text/html",
          });
          response.write(data);
          response.end();
        }
      });
      break;
    case "/conversations.html":
      fs.readFile(__dirname + path, function (error, data) {
        if (error) {
          response.writeHead(404);
          response.write(error);
          response.end();
        } else {
          response.writeHead(200, {
            "Content-Type": "text/html",
          });
          response.write(data);
          response.end();
        }
      });
      break;
    case "/contacts.html":
      fs.readFile(__dirname + path, function (error, data) {
        if (error) {
          response.writeHead(404);
          response.write(error);
          response.end();
        } else {
          response.writeHead(200, {
            "Content-Type": "text/html",
          });
          response.write(data);
          response.end();
        }
      });
      break;
    default:
      response.writeHead(404);
      response.write("opps this doesn't exist - 404");
      response.end();
      break;
  }
});
server.listen(8081);

function css(request, response) {
  console.log(request.url);
  if (request.url === "/style.css") {
    response.writeHead(200, { "Content-type": "text/css" });
    var fileContents = fs.readFileSync("style.css", {
      encoding: "utf8",
    });
    console.log(fileContents);
    response.write(fileContents);
  }
}
