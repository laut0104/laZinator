# FROM node:14-alpine As builder
FROM node

WORKDIR /projects/frontend

COPY frontend/package.json ./
COPY frontend/package-lock.json ./

RUN npm install -g @angular/cli
RUN npm install

EXPOSE 4200

# RUN npm install -g @angular/cli @angular-devkit/build-angular && npm install

# COPY . .

# ARG BUILD_ENV

# RUN npm run build -- --configuration ${BUILD_ENV}

# EXPOSE 4200
# CMD ["npm", "start"]

# FROM nginx:1.17.1-alpine

# COPY --from=builder /usr/src/app/dist/angular-app/ /usr/share/nginx/html
# COPY --from=builder /usr/src/app/nginx/default.conf /etc/nginx/conf.d/default.conf
