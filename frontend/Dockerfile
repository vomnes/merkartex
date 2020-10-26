FROM node:13

COPY . /app
WORKDIR /app

RUN npm install
RUN npm run build
RUN npm install -g serve

ENTRYPOINT serve -s dist -l 5001
