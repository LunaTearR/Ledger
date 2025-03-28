FROM node:20.18.3-alpine

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY . /app/

RUN rm -rf package-lock.json

RUN npm install

RUN npm install -g @angular/cli@18.1.0

CMD ["ng", "serve", "--host", "0.0.0.0"]