import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/auth/endpoint/auth_api_client.dart';
import 'package:shoppie_admin/features/auth/model/login_request.dart';
import 'package:shoppie_admin/features/auth/model/login_response.dart';
import 'package:shoppie_admin/features/auth/model/register_request.dart';
import 'package:shoppie_admin/features/auth/model/user_model.dart';
import 'package:shoppie_admin/features/auth/repository/auth_repository.dart';

class AuthRepositoryImpl implements AuthRepository {
  final AuthApiClient _apiClient;
  final SharedPreferences _preferences;

  AuthRepositoryImpl(this._apiClient, this._preferences);

  @override
  Future<Either<UserModel, Failure>> register(RegisterRequest request) async {
    try {
      final result = await _apiClient.regsiter(request);
      return Left(result);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<LoginResponse, Failure>> login(LoginRequest request) async {
    try {
      final response = await _apiClient.login(request);
      await _preferences.setString("token", response.token);
      return Left(response);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<UserModel, Failure>> signedInUser() async {
    try {
      final user = await _apiClient.getUser();
      return Left(user);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<bool> checkIfSignedIn() async {
    final token = _preferences.getString("token");

    if (token == null || token.isEmpty) {
      return false;
    }

    return true;
  }

  @override
  Future<Unit> logout() async {
    await _preferences.remove("token");

    return unit;
  }
}
