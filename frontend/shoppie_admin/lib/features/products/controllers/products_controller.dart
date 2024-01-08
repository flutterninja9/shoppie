import 'package:flutter/foundation.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/products/models/products_result.dart';
import 'package:shoppie_admin/features/products/repository/products_repository.dart';

class ProductsController extends ChangeNotifier with UiStateMixin {
  final ProductsRepository _repository;
  ProductsResults? _products;
  ProductsResults? get products => _products;

  ProductsController(this._repository);

  Future<void> fetchAllProducts() async {
    setLoading(true);
    final resultOrFailure = await _repository.getAllProducts();

    resultOrFailure.fold(
      (l) {
        _products = l;

        if (_products?.data.isEmpty ?? false) {
          setEmpty();
        } else {
          resetState();
        }
      },
      (r) {
        setFailure(failureMessage);
      },
    );
  }

  Future<void> deleteProduct(String id) async {
    setLoading(true);
    final resultOrFailure = await _repository.deleteProduct(id);

    await resultOrFailure.fold(
      (l) async {
        await fetchAllProducts();
      },
      (r) {
        setFailure(failureMessage);
      },
    );
  }
}
