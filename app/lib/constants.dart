import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class Constants {
  static Future<void> initialize() async {
    // Load .env files in order of precedence

    try {
      // Local development overrides
      await dotenv.load(fileName: '.env.local');
    } catch (_) {
      // Allow to fail if file doesn't exist
    }
    
    try {
      // Environment-specific values
      await dotenv.load(fileName: '.env.${kReleaseMode ? 'prod' : 'dev'}');
    } catch (_) {
      // Allow to fail if file doesn't exist
    }
    
    try {
      // Default values
      await dotenv.load(fileName: '.env');
    } catch (_) {
      // Allow to fail if file doesn't exist
    }
  }

  static String? get apiUrl {
    return dotenv.env['NGROK_API_URL'];
  }
}