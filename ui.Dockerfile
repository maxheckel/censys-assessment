FROM node:10

WORKDIR /usr/src/app/censys-ui

RUN npm install

EXPOSE 4200

CMD ["npm", "run", "serve", "--", "--port", "4200"]