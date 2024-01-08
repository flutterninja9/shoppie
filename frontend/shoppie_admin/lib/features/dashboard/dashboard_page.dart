import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/dashboard/controllers/dashboard_controller.dart';
import 'package:shoppie_admin/features/orders/pages/view_orders.dart';
import 'package:shoppie_admin/features/products/pages/add_prodcust.dart';
import 'package:shoppie_admin/features/products/pages/view_products.dart';

class DashboardPage extends StatefulWidget {
  const DashboardPage({super.key, required this.page});
  static const route = "/dashboard";
  final Widget page;
  @override
  State<DashboardPage> createState() => _DashboardPageState();
}

class _DashboardPageState extends State<DashboardPage> {
  final GlobalKey<ScaffoldState> _scaffoldKey = GlobalKey();

  bool activePage(String route) {
    return GoRouterState.of(context).uri.toString().contains(route);
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => sl<DashboardController>()..getSignedInUser(),
      child: Scaffold(
        key: _scaffoldKey,
        body: Row(
          children: [
            Drawer(
              child: Column(
                children: [
                  ListTile(
                    onTap: () => context.go(ViewProductsPage.route),
                    leading: activePage(ViewProductsPage.route)
                        ? const Icon(Icons.circle)
                        : const Icon(Icons.circle_outlined),
                    title: const Text("View Products"),
                  ),
                  ListTile(
                    onTap: () => context.go(AddProductPage.route),
                    leading: activePage(AddProductPage.route)
                        ? const Icon(Icons.circle)
                        : const Icon(Icons.circle_outlined),
                    title: const Text("Add product"),
                  ),
                  ListTile(
                    onTap: () => context.go(ViewOrdersPage.route),
                    leading: activePage(ViewOrdersPage.route)
                        ? const Icon(Icons.circle)
                        : const Icon(Icons.circle_outlined),
                    title: const Text("View Orders"),
                  ),
                  ListTile(
                    onTap: () {
                      context.read<AuthNotifier>().logout();
                    },
                    leading: const Icon(Icons.logout),
                    title: const Text("Logout"),
                  ),
                ],
              ),
            ),
            Expanded(
              child: Column(
                children: [
                  AppBar(
                    title: const Text("Dashboard"),
                    actions: [
                      Padding(
                        padding: const EdgeInsets.all(8.0),
                        child: Consumer<DashboardController>(
                            builder: (context, provider, child) {
                          if (provider.isLoading) {
                            return const Center(
                              child: CircularProgressIndicator(),
                            );
                          }

                          return CircleAvatar(
                            child: Text(
                              provider.user?.firstName.characters.first ?? "-",
                            ),
                          );
                        }),
                      )
                    ],
                  ),
                  Expanded(child: widget.page),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
