import 'package:flutter/material.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/auth/model/user_model.dart';
import 'package:shoppie_admin/features/auth/repository/auth_repository.dart';

class DashboardController extends ChangeNotifier with UiStateMixin {
  final AuthRepository _repository;

  DashboardController(this._repository);
  UserModel? _user;
  UserModel? get user => _user;

  Future<void> getSignedInUser() async {
    setLoading(true);

    final resultOrFailure = await _repository.signedInUser();

    resultOrFailure.fold(
      (l) {
        _user = l;
        resetState();
      },
      (r) {
        setFailure(failureMessage);
      },
    );
  }
}
