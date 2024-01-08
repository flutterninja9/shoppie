import 'package:json_annotation/json_annotation.dart';
import 'package:shoppie_admin/features/products/models/get_product_by_id_response.dart';

part "products_result.g.dart";

@JsonSerializable()
class ProductsResults {
  final List<ProductEntity> data;

  ProductsResults({
    required this.data,
  });

  factory ProductsResults.fromJson(Map<String, dynamic> json) =>
      _$ProductsResultsFromJson(json);

  Map<String, dynamic> toJson() => _$ProductsResultsToJson(this);
}
