FROM node:alpine AS build
COPY ./ .
RUN npm install && \
npm run build

FROM nginx:alpine
COPY --from=build ./build /usr/share/nginx/html
EXPOSE 80
