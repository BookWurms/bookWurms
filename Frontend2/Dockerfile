# dev: Vite dev server with hot reload
FROM node:26-alpine3.24 AS dev

WORKDIR /app

COPY package*.json .npmrc ./
RUN npm ci

COPY . .

CMD ["npm", "run", "dev", "--", "--host"]


# build: create the static files
FROM dev AS build

RUN npm run build


# prod: Caddy serves the static files
FROM caddy:2-alpine AS prod

COPY --from=build /app/build /srv/frontend
