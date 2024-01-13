// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'order_model.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

OrderModel _$OrderModelFromJson(Map<String, dynamic> json) => OrderModel(
      id: json['id'] as String,
      status: json['status'] as String,
      totalAmt: (json['total_amount'] as num?)?.toDouble(),
      items: (json['order_items'] as List<dynamic>)
          .map((e) => OrderItem.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$OrderModelToJson(OrderModel instance) =>
    <String, dynamic>{
      'id': instance.id,
      'status': instance.status,
      'total_amount': instance.totalAmt,
      'order_items': instance.items,
    };

OrderItem _$OrderItemFromJson(Map<String, dynamic> json) => OrderItem(
      id: json['id'] as String,
      orderId: json['order_id'] as String,
      price: (json['price'] as num?)?.toDouble(),
      productId: json['product_id'] as String,
      quantity: json['quantity'] as int?,
    );

Map<String, dynamic> _$OrderItemToJson(OrderItem instance) => <String, dynamic>{
      'id': instance.id,
      'order_id': instance.orderId,
      'product_id': instance.productId,
      'price': instance.price,
      'quantity': instance.quantity,
    };
