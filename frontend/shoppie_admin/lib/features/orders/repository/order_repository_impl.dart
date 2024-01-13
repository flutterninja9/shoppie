import 'package:dartz/dartz.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/orders/endpoint/orders_endpoint.dart';
import 'package:shoppie_admin/features/orders/model/orders_model.dart';
import 'package:shoppie_admin/features/orders/repository/order_repository.dart';

class OrderRepositoryImpl implements OrderRepository {
  final OrdersEndpoint _apiClient;

  OrderRepositoryImpl(this._apiClient);

  @override
  Future<Either<OrdersModel, Failure>> getAllOrders() async {
    try {
      final result = await _apiClient.getOrders();
      return Left(result);
    } catch (e, s) {
      return Right(UnknownFailure(error: e, stack: s));
    }
  }

  @override
  Future<Either<Unit, Failure>> cancel(String orderId) async {
    try {
      await _apiClient.cancel(orderId);
      return const Left(unit);
    } catch (e, s) {
      return Right(UnknownFailure(error: e, stack: s));
    }
  }
}
