import 'package:flutter/material.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';

class ViewProductsPage extends StatelessWidget {
  const ViewProductsPage({super.key});
  static const route = "${DashboardPage.route}/view-products";

  @override
  Widget build(BuildContext context) {
    return const Text("View Products");
  }
}
