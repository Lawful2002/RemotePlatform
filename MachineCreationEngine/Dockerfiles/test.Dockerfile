FROM node:alpine
WORKDIR /app

RUN touch test.html
RUN echo '<h1>Test</h1>' > test.html

RUN npm i -g serve

CMD [ "serve", "index.html" ]