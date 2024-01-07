import 'package:flutter_dotenv/flutter_dotenv.dart';

class AppConfig {
  static String userServiceBaseUrl = dotenv.env["USER_SERVICE_BASE_URL"] ?? "";

  static String productServiceBaseUrl =
      dotenv.env["PRODUCT_SERVICE_BASE_URL"] ?? "";

  static String cartServiceBaseUrl = dotenv.env["CART_SERVICE_BASE_URL"] ?? "";

  static String orderServiceBaseUrl =
      dotenv.env["ORDER_SERVICE_BASE_URL"] ?? "";

  static String jwtSecret = dotenv.env["JWT_SECRET"] ?? "";

  static String baseUrl = dotenv.env["BASE_URL"] ?? "";
}
