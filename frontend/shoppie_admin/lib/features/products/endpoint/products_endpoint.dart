import 'dart:async';

import 'package:dio/dio.dart';
import 'package:retrofit/retrofit.dart';
import 'package:shoppie_admin/features/products/models/create_product_request.dart';
import 'package:shoppie_admin/features/products/models/products_result.dart';
import 'package:shoppie_admin/features/products/models/update_product_request.dart';

part 'products_endpoint.g.dart';

@RestApi()
abstract class ProductspiClient {
  factory ProductspiClient(Dio dio, {String baseUrl}) = _ProductspiClient;

  @GET('/products/')
  Future<ProductsResults> getAllProducts();

  @POST('/products/')
  Future<void> createProduct(@Body() CreateProductRequest request);

  @GET('/products/{pId}')
  Future<void> getProductDetails(
    @Path() String pId,
  );

  @PATCH('/products/{pId}')
  Future<void> updateProduct(
    @Path() String pId,
    @Body() UpdateProductRequest request,
  );

  @DELETE('/products/{pId}')
  Future<void> deleteProduct(
    @Path() String pId,
  );
}
