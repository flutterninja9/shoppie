import 'package:json_annotation/json_annotation.dart';

part "get_product_by_id_response.g.dart";

@JsonSerializable()
class GetProductByIdResponse {
  final ProductEntity data;

  GetProductByIdResponse({
    required this.data,
  });

  factory GetProductByIdResponse.fromJson(Map<String, dynamic> json) =>
      _$GetProductByIdResponseFromJson(json);

  Map<String, dynamic> toJson() => _$GetProductByIdResponseToJson(this);
}

@JsonSerializable()
class ProductEntity {
  final String id;
  final String name;
  final String description;
  final double price;
  final int quantity;

  ProductEntity({
    required this.description,
    required this.id,
    required this.name,
    required this.price,
    required this.quantity,
  });

  factory ProductEntity.fromJson(Map<String, dynamic> json) =>
      _$ProductEntityFromJson(json);

  Map<String, dynamic> toJson() => _$ProductEntityToJson(this);
}
