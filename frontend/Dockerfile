FROM node:18-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Use Node.js to serve the Next.js app dynamically
FROM node:18-alpine
WORKDIR /app
COPY --from=build /app ./

# curl for health checks
RUN apk add --no-cache curl

EXPOSE 3000
CMD ["npm", "run", "start"]