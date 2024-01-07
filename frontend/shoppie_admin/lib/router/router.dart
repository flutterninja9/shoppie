import 'package:go_router/go_router.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/auth/view/auth_view.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';

class AppRouter {
  AppRouter._();

  static GoRouter router = GoRouter(
    initialLocation: DashboardPage.route,
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
      GoRoute(
        path: DashboardPage.route,
        builder: (context, state) {
          return const DashboardPage();
        },
      ),
    ],
  );
}
