import 'dart:io';

class Config {
  final isProduction = Platform.environment['IS_PRODUCTION'] == '1';
  static const String apiUrl = 'http://0.0.0.0:3000';
}