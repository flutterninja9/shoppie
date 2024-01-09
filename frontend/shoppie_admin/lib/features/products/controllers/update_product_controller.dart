import 'package:flutter/foundation.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/products/models/get_product_by_id_response.dart';
import 'package:shoppie_admin/features/products/models/update_product_request.dart';
import 'package:shoppie_admin/features/products/repository/products_repository.dart';

class UpdateProductController extends ChangeNotifier with UiStateMixin {
  final ProductsRepository _repository;
  UpdateProductController(this._repository);

  ProductEntity? _productEntity;
  ProductEntity? get productEntity => _productEntity;

  Future<void> getProductById(String id) async {
    setLoading(true);
    final resultOrFailure = await _repository.getProductDetails(id);

    resultOrFailure.fold(
      (l) {
        _productEntity = l;
        resetState();
      },
      (r) {
        setFailure(failureMessage);
      },
    );
  }

  Future<void> updateProduct(
    String id,
    UpdateProductRequest req,
    Function() postSuccess,
  ) async {
    setLoading(true);
    final resultOrFailure = await _repository.updateProduct(id, req);

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
