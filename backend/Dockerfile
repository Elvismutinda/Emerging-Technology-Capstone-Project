
# Change this if necessary. I just wrote it to fill it up for now
# Delete this comment when you're done

FROM node:20

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

EXPOSE 4000

CMD ["node", "index.js"]