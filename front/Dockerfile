FROM node:14.9-alpine


WORKDIR /usr/src/app

COPY ["package.json", "package-lock.json", "/usr/src/app/"]

RUN npm i --silent

CMD npm start