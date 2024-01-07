import 'package:dio/dio.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AuthInterceptor extends Interceptor {
  final SharedPreferences sp;

  AuthInterceptor(this.sp);
  @override
  void onRequest(
      RequestOptions options, RequestInterceptorHandler handler) async {
    final token = sp.getString("token");
    if (token != null && token.isNotEmpty) {
      options.headers.update('Authorization', (_) => "Bearer $token",
          ifAbsent: () => "Bearer $token");
    }

    super.onRequest(options, handler);
  }
}
