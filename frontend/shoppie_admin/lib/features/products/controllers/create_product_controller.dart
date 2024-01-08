import 'package:flutter/foundation.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/products/models/create_product_request.dart';
import 'package:shoppie_admin/features/products/repository/products_repository.dart';

class CreateProductController extends ChangeNotifier with UiStateMixin {
  final ProductsRepository _repository;
  CreateProductController(this._repository);

  Future<void> createProduct(
      CreateProductRequest req, Function() postSuccess) async {
    setLoading(true);
    final resultOrFailure = await _repository.createProduct(req);

    resultOrFailure.fold(
      (l) {
        resetState();
        postSuccess();
      },
      (r) {
        setFailure(failureMessage);
      },
    );
  }
}
