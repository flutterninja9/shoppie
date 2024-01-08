// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'products_result.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ProductsResults _$ProductsResultsFromJson(Map<String, dynamic> json) =>
    ProductsResults(
      data: (json['data'] as List<dynamic>)
          .map((e) => ProductEntity.fromJson(e as Map<String, dynamic>))
          .toList(),
    );

Map<String, dynamic> _$ProductsResultsToJson(ProductsResults instance) =>
    <String, dynamic>{
      'data': instance.data,
    };
