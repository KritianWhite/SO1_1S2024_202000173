FROM node:20.11.0-alpine as BUILDER_IMAGE

WORKDIR /app/frontend

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:1.25.3-alpine as SERVER_IMAGE

# Remove the default Nginx configuration
RUN rm /etc/nginx/conf.d/default.conf

# Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/nginx.conf


COPY --from=BUILDER_IMAGE /app/frontend/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]