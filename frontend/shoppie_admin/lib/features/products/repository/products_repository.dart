import 'package:dartz/dartz.dart';
import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/products/models/create_product_request.dart';
import 'package:shoppie_admin/features/products/models/get_product_by_id_response.dart';
import 'package:shoppie_admin/features/products/models/products_result.dart';
import 'package:shoppie_admin/features/products/models/update_product_request.dart';

abstract class ProductsRepository {
  Future<Either<ProductsResults, Failure>> getAllProducts();

  Future<Either<Unit, Failure>> createProduct(CreateProductRequest request);

  Future<Either<ProductEntity, Failure>> getProductDetails(String productId);

  Future<Either<Unit, Failure>> updateProduct(
      String productId, UpdateProductRequest request);

  Future<Either<Unit, Failure>> deleteProduct(String productId);
}
