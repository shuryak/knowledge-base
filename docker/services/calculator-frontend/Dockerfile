FROM node:18-alpine as build
ARG API_HOST
ARG API_PORT

WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .

RUN REACT_APP_API_HOST=${API_HOST} \
    REACT_APP_API_PORT=${API_PORT} \
    npm run build

FROM nginx:stable-alpine
COPY --from=build /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
