FROM node:18-alpine3.17 as build
WORKDIR /app
COPY . /app
RUN rm -rf node_modules
RUN npm install
RUN env VITE_BASE_URL="/backend" VITE_PLATFORM="kubernetes" npm run build

FROM nginx:1.21-alpine
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]