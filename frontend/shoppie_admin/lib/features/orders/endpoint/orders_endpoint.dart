import 'dart:async';

import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';
import 'package:shoppie_admin/features/orders/model/orders_model.dart';

part 'orders_endpoint.g.dart';

@RestApi()
abstract class OrdersEndpoint {
  factory OrdersEndpoint(Dio dio, {String baseUrl}) = _OrdersEndpoint;

  @GET('/orders/all')
  Future<OrdersModel> getOrders();

  @GET('/orders/{orderId}/cancel')
  Future<void> cancel(@Path() String orderId);
}
