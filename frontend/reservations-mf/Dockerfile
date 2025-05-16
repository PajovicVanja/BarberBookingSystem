# ---------- 1) BUILD stage ----------------------------------------------------
FROM node:18-alpine AS build

WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm ci            # install deps *inside the image*

COPY . .
RUN npm run build     # → creates /app/dist

# ---------- 2) RUNTIME stage --------------------------------------------------
FROM nginx:alpine

# CORS + SPA fallback
COPY nginx.conf /etc/nginx/conf.d/default.conf

# remove default html and copy our dist
RUN rm -rf /usr/share/nginx/html/*
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
