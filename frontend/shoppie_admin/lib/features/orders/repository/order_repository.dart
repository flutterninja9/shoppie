import 'package:dartz/dartz.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/orders/model/orders_model.dart';

abstract class OrderRepository {
  Future<Either<OrdersModel, Failure>> getAllOrders();
  Future<Either<Unit, Failure>> cancel(String orderId);
}
