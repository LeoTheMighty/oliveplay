use dotenv::dotenv;

// Load environment variables from .env file
dotenv().ok();

// Define constants
pub const SQS_QUEUE_NAME: &str = env!("SQS_QUEUE_NAME");
pub const API_URL: &str = env!("API_URL");
pub const DB_HOST: &str = env!("DB_HOST");
pub const DB_NAME: &str = env!("DB_NAME");
pub const DB_USER: &str = env!("DB_USER");
pub const DB_PASS: &str = env!("DB_PASS");
