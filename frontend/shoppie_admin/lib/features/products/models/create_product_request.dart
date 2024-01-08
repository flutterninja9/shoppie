import 'package:json_annotation/json_annotation.dart';

part "create_product_request.g.dart";

@JsonSerializable()
class CreateProductRequest {
  final String name;
  final String description;
  final double price;
  final int quantity;

  CreateProductRequest({
    required this.description,
    required this.name,
    required this.price,
    required this.quantity,
  });

  factory CreateProductRequest.fromJson(Map<String, dynamic> json) =>
      _$CreateProductRequestFromJson(json);

  Map<String, dynamic> toJson() => _$CreateProductRequestToJson(this);
}
