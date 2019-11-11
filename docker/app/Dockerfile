FROM node:13.1-alpine AS builder

WORKDIR /app
COPY ./app/package.json /app/package.json
RUN npm install
COPY ./app /app
RUN npm run build

FROM nginx:1.17-alpine
COPY --from=builder /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]
