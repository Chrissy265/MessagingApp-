FROM node:12

WORKDIR /usr/src/app

COPY . .

RUN npm install -g http-server

EXPOSE 8000

CMD  http-server -a localhost -p 8000