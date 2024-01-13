import 'package:json_annotation/json_annotation.dart';
import 'package:shoppie_admin/features/orders/model/order_model.dart';

part 'orders_model.g.dart';

@JsonSerializable()
class OrdersModel {
  @JsonKey(name: "data")
  final List<OrderModel> orders;

  OrdersModel({
    required this.orders,
  });

  factory OrdersModel.fromJson(Map<String, dynamic> json) =>
      _$OrdersModelFromJson(json);

  Map<String, dynamic> toJson() => _$OrdersModelToJson(this);
}
