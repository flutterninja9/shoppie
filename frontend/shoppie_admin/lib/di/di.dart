import 'package:dio/dio.dart';
import 'package:get_it/get_it.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:shoppie_admin/app/app_config.dart';
import 'package:shoppie_admin/core/interceptors/auth_interceptor.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/auth/endpoint/auth_api_client.dart';
import 'package:shoppie_admin/features/auth/repository/auth_repository.dart';
import 'package:shoppie_admin/features/auth/repository/auth_repository_impl.dart';
import 'package:shoppie_admin/features/dashboard/controllers/dashboard_controller.dart';
import 'package:shoppie_admin/features/orders/controller/order_controller.dart';
import 'package:shoppie_admin/features/orders/endpoint/orders_endpoint.dart';
import 'package:shoppie_admin/features/orders/repository/order_repository.dart';
import 'package:shoppie_admin/features/orders/repository/order_repository_impl.dart';
import 'package:shoppie_admin/features/products/controllers/create_product_controller.dart';
import 'package:shoppie_admin/features/products/controllers/products_controller.dart';
import 'package:shoppie_admin/features/products/controllers/update_product_controller.dart';
import 'package:shoppie_admin/features/products/endpoint/products_endpoint.dart';
import 'package:shoppie_admin/features/products/repository/products_repository.dart';
import 'package:shoppie_admin/features/products/repository/products_repository_impl.dart';

final sl = GetIt.I;

Future<void> setupSl() async {
  // core

  // controllers
  sl.registerLazySingleton<AuthNotifier>(() => AuthNotifier(sl()));
  sl.registerFactory<DashboardController>(() => DashboardController(sl()));
  sl.registerFactory<ProductsController>(() => ProductsController(sl()));
  sl.registerFactory<CreateProductController>(
    () => CreateProductController(sl()),
  );
  sl.registerFactory<UpdateProductController>(
    () => UpdateProductController(sl()),
  );
  sl.registerFactory<OrdersController>(
    () => OrdersController(sl()),
  );

  // repo
  sl.registerLazySingleton<AuthRepository>(
    () => AuthRepositoryImpl(sl(), sl()),
  );
  sl.registerLazySingleton<ProductsRepository>(
    () => ProductsRepositoryImpl(sl()),
  );
  sl.registerLazySingleton<OrderRepository>(
    () => OrderRepositoryImpl(sl()),
  );

  // endpoints
  sl.registerLazySingleton<AuthApiClient>(() => AuthApiClient(sl()));
  sl.registerLazySingleton<ProductspiClient>(() => ProductspiClient(sl()));
  sl.registerLazySingleton<OrdersEndpoint>(() => OrdersEndpoint(sl()));

  // external
  final sp = await SharedPreferences.getInstance();
  sl.registerLazySingleton<SharedPreferences>(() => sp);

  final interceptor = AuthInterceptor(sp);
  sl.registerLazySingleton<AuthInterceptor>(() => interceptor);

  final dio = Dio(
    BaseOptions(
        baseUrl: AppConfig.baseUrl,
        connectTimeout: const Duration(seconds: 15),
        receiveTimeout: const Duration(seconds: 15)),
  );
  dio.interceptors.addAll([LogInterceptor(), interceptor]);
  sl.registerLazySingleton<Dio>(() => dio);
}
