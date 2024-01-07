import 'dart:async';

import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';
import 'package:shoppie_admin/features/auth/model/login_request.dart';
import 'package:shoppie_admin/features/auth/model/login_response.dart';
import 'package:shoppie_admin/features/auth/model/register_request.dart';
import 'package:shoppie_admin/features/auth/model/user_model.dart';

part 'auth_api_client.g.dart';

@RestApi()
abstract class AuthApiClient {
  factory AuthApiClient(Dio dio, {String baseUrl}) = _AuthApiClient;

  @POST('/login')
  Future<LoginResponse> login(@Body() LoginRequest request);

  @POST('/register')
  Future<UserModel> regsiter(@Body() RegisterRequest request);

  @GET('/user')
  Future<UserModel> getUser();
}
