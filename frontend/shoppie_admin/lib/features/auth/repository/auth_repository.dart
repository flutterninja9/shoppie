import 'package:dartz/dartz.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/auth/model/login_request.dart';
import 'package:shoppie_admin/features/auth/model/login_response.dart';
import 'package:shoppie_admin/features/auth/model/register_request.dart';
import 'package:shoppie_admin/features/auth/model/user_model.dart';

abstract class AuthRepository {
  Future<Either<UserModel, Failure>> register(RegisterRequest request);

  Future<Either<LoginResponse, Failure>> login(LoginRequest request);

  Future<Either<UserModel, Failure>> signedInUser();

  Future<bool> checkIfSignedIn();
}
