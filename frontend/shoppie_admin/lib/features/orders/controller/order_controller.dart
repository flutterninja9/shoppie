import 'package:flutter/foundation.dart';
import 'package:shoppie_admin/core/constants/constants.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/core/mixins/ui_state_mixins.dart';
import 'package:shoppie_admin/features/orders/model/order_model.dart';
import 'package:shoppie_admin/features/orders/model/orders_model.dart';
import 'package:shoppie_admin/features/orders/repository/order_repository.dart';

class OrdersController extends ChangeNotifier with UiStateMixin {
  final OrderRepository _repository;
  OrdersModel? _orders;
  OrdersModel? get orders => _orders;

  OrdersController(this._repository);

  Future<void> fetchAllOrders() async {
    setLoading(true);
    final resultOrFailure = await _repository.getAllOrders();

    resultOrFailure.fold(
      (l) {
        _orders = l;

        if (_orders?.orders.isEmpty ?? false) {
          setEmpty();
        } else {
          resetState();
        }
      },
      (r) {
        if (r is UnknownFailure) {
          setFailure(r.error.toString());
        }
        setFailure(failureMessage);
      },
    );
  }

  Future<void> cancel(OrderModel order) async {
    setLoading(true);
    if (order.status == "cancelled") {
      resetState();
      throw UnsupportedError("Already cancelled");
    }
    final resultOrFailure = await _repository.cancel(order.id);

    resultOrFailure.fold(
      (l) {
        fetchAllOrders();
      },
      (r) {
        if (r is UnknownFailure) {
          setFailure(r.error.toString());
        }
        setFailure(failureMessage);
      },
    );
  }
}
