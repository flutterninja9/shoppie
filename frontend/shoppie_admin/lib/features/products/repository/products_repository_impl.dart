import 'package:shoppie_admin/core/exceptions/failure.dart';
import 'package:shoppie_admin/features/products/endpoint/products_endpoint.dart';
import 'package:shoppie_admin/features/products/models/create_product_request.dart';
import 'package:shoppie_admin/features/products/models/get_product_by_id_response.dart';
import 'package:shoppie_admin/features/products/models/products_result.dart';
import 'package:shoppie_admin/features/products/models/update_product_request.dart';
import 'package:shoppie_admin/features/products/repository/products_repository.dart';
import 'package:dartz/dartz.dart';

class ProductsRepositoryImpl implements ProductsRepository {
  final ProductspiClient _apiClient;

  ProductsRepositoryImpl(this._apiClient);

  @override
  Future<Either<ProductsResults, Failure>> getAllProducts() async {
    try {
      final products = await _apiClient.getAllProducts();
      return Left(products);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<Unit, Failure>> createProduct(
      CreateProductRequest request) async {
    try {
      await _apiClient.createProduct(request);
      return const Left(unit);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<ProductEntity, Failure>> getProductDetails(
      String productId) async {
    try {
      final result = await _apiClient.getProductDetails(productId);
      return Left(result.data);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<Unit, Failure>> updateProduct(
      String productId, UpdateProductRequest request) async {
    try {
      await _apiClient.updateProduct(productId, request);
      return const Left(unit);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }

  @override
  Future<Either<Unit, Failure>> deleteProduct(String productId) async {
    try {
      await _apiClient.deleteProduct(productId);
      return const Left(unit);
    } catch (e) {
      return Right(UnknownFailure());
    }
  }
}
