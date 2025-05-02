# API Setup

1. **Install sqlc**  
   Follow instructions at https://docs.sqlc.dev/en/stable/ (e.g. `brew install sqlc` or use Docker).

2. **Install Nx**  
   `npm install -g nx` (or use it via `npx nx`).

3. **Build**  
   ```sh
   # Make sure cate_utils is built first (generates Go code via sqlc)
   nx build cate_utils
   # Then build your api
   nx build api
   ```

4. **Migrations**  
   - To run migrations locally:
     ```sh
     nx build migrator
     nx run migrator:run
     ```
   - Or via Docker:
     ```sh
     docker build -t migrator -f Migrator/Dockerfile .
     docker run -e DATABASE_URL=$DATABASE_URL migrator
     ```

5. **Run**  
   ```sh
   # From the api folder
   go run main.go
   ``` 