# Stage 1: Build the Flutter app
FROM instrumentisto/flutter:latest as build

WORKDIR /app
COPY . .

# Build the Flutter web app
RUN flutter build web

# Stage 2: Serve the Flutter app using nginx
FROM nginx:alpine

# Copy the build output to nginx's web directory
COPY --from=build /app/build/web /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]