// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'get_product_by_id_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

GetProductByIdResponse _$GetProductByIdResponseFromJson(
        Map<String, dynamic> json) =>
    GetProductByIdResponse(
      data: ProductEntity.fromJson(json['data'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$GetProductByIdResponseToJson(
        GetProductByIdResponse instance) =>
    <String, dynamic>{
      'data': instance.data,
    };

ProductEntity _$ProductEntityFromJson(Map<String, dynamic> json) =>
    ProductEntity(
      description: json['description'] as String,
      id: json['id'] as String,
      name: json['name'] as String,
      price: (json['price'] as num).toDouble(),
      quantity: json['quantity'] as int,
    );

Map<String, dynamic> _$ProductEntityToJson(ProductEntity instance) =>
    <String, dynamic>{
      'id': instance.id,
      'name': instance.name,
      'description': instance.description,
      'price': instance.price,
      'quantity': instance.quantity,
    };
