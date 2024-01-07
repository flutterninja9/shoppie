import 'package:flutter/material.dart';
import 'package:dartz/dartz.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/auth/model/login_request.dart';
import 'package:shoppie_admin/features/auth/model/login_response.dart';
import 'package:shoppie_admin/features/auth/model/register_request.dart';
import 'package:shoppie_admin/features/auth/model/user_model.dart';
import 'package:shoppie_admin/features/auth/repository/auth_repository.dart';

class AuthNotifier with ChangeNotifier, UiStateMixin {
  final AuthRepository _authRepository;

  UserModel? _user;
  UserModel? get user => _user;

  bool _signedIn = false;
  bool get signedIn => _signedIn;

  AuthNotifier(this._authRepository);

  Future<void> register(RegisterRequest request) async {
    setLoading(true);
    final result = await _authRepository.register(request);
    result.fold(
      (unit) {
        _signedIn = true;
        _user = unit;
        resetState();
      },
      (failure) {
        setFailure(failureMessage);
      },
    );
  }

  Future<void> login(LoginRequest request) async {
    setLoading(true);
    Either<LoginResponse, Failure> result =
        await _authRepository.login(request);
    result.fold(
      (loginResponse) {
        _signedIn = true;
        resetState();
      },
      (failure) {
        setFailure(failureMessage);
      },
    );
  }

  Future<void> fetchSignedInUser() async {
    setLoading(true);
    Either<UserModel, Failure> result = await _authRepository.signedInUser();
    result.fold(
      (userModel) {
        _user = userModel;
        resetState();
      },
      (failure) {
        setFailure(failureMessage);
      },
    );
  }

  Future<void> checkIfSignedIn() async {
    setLoading(true);
    final loggedIn = await _authRepository.checkIfSignedIn();
    _signedIn = loggedIn;
    print(_signedIn);
    resetState();
  }
}
