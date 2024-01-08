import 'package:go_router/go_router.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/auth/view/auth_view.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';
import 'package:shoppie_admin/features/orders/pages/view_orders.dart';
import 'package:shoppie_admin/features/products/pages/add_prodcust.dart';
import 'package:shoppie_admin/features/products/pages/view_products.dart';

class AppRouter {
  AppRouter._();

  static GoRouter router = GoRouter(
    initialLocation: ViewProductsPage.route,
    refreshListenable: sl<AuthNotifier>(),
    redirect: (context, state) {
      if (sl<AuthNotifier>().signedIn) {
        return null;
      }

      return AuthView.route;
    },
    routes: [
      GoRoute(
        path: AuthView.route,
        builder: (context, state) {
          return const AuthView();
        },
      ),
      ShellRoute(
        builder: (context, state, page) {
          return DashboardPage(page: page);
        },
        routes: [
          GoRoute(
            path: ViewProductsPage.route,
            builder: (context, state) {
              return const ViewProductsPage();
            },
          ),
          GoRoute(
            path: AddProductPage.route,
            builder: (context, state) {
              return const AddProductPage();
            },
          ),
          GoRoute(
            path: ViewOrdersPage.route,
            builder: (context, state) {
              return const ViewOrdersPage();
            },
          ),
        ],
      ),
    ],
  );
}
