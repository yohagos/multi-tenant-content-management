FROM node:26-alpine3.24 as builder

WORKDIR /app

COPY package.json package.json
RUN npm install --legacy-peer-deps

COPY . .

RUN npm run build:shell
RUN npm run build:tenant-admin
RUN npm run build:content-editor
RUN npm run build:public-view

FROM nginx:1.31.4-alpine3.24

COPY --from=builder /app/dist/shell/browser /usr/share/nginx/html/shell
COPY --from=builder /app/dist/tenant-admin/browser /usr/share/nginx/html/tenant-admin
COPY --from=builder /app/dist/content-editor/browser /usr/share/nginx/html/content-editor
COPY --from=builder /app/dist/public-view/browser /usr/share/nginx/html/public-view

COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD [ "nginx", "-g", "daemon off;" ]
