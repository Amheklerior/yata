FROM node:22-alpine AS builder

WORKDIR /app

COPY . .

RUN npm install -g pnpm
RUN CI=1 pnpm install
RUN pnpm build 

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]