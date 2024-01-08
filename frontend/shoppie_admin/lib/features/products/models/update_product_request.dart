import 'package:json_annotation/json_annotation.dart';

part "update_product_request.g.dart";

@JsonSerializable()
class UpdateProductRequest {
  final String id;
  final String name;
  final String description;
  final double price;
  final int quantity;

  UpdateProductRequest({
    required this.description,
    required this.id,
    required this.name,
    required this.price,
    required this.quantity,
  });

  factory UpdateProductRequest.fromJson(Map<String, dynamic> json) =>
      _$UpdateProductRequestFromJson(json);

  Map<String, dynamic> toJson() => _$UpdateProductRequestToJson(this);
}
