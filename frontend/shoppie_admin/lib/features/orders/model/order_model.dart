import 'package:json_annotation/json_annotation.dart';

part 'order_model.g.dart';

@JsonSerializable()
class OrderModel {
  final String id;
  final String status;
  @JsonKey(name: "total_amount")
  final double? totalAmt;
  @JsonKey(name: "order_items")
  final List<OrderItem> items;

  OrderModel({
    required this.id,
    required this.status,
    required this.totalAmt,
    required this.items,
  });

  factory OrderModel.fromJson(Map<String, dynamic> json) =>
      _$OrderModelFromJson(json);

  Map<String, dynamic> toJson() => _$OrderModelToJson(this);
}

@JsonSerializable()
class OrderItem {
  final String id;
  @JsonKey(name: "order_id")
  final String orderId;
  @JsonKey(name: "product_id")
  final String productId;
  final double? price;
  final int? quantity;

  OrderItem({
    required this.id,
    required this.orderId,
    required this.price,
    required this.productId,
    required this.quantity,
  });

  factory OrderItem.fromJson(Map<String, dynamic> json) =>
      _$OrderItemFromJson(json);

  Map<String, dynamic> toJson() => _$OrderItemToJson(this);
}
