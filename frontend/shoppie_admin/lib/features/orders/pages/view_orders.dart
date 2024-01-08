import 'package:flutter/material.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';

class ViewOrdersPage extends StatelessWidget {
  const ViewOrdersPage({super.key});
  static const route = "${DashboardPage.route}/view-orders";

  @override
  Widget build(BuildContext context) {
    return const Text("View Orders");
  }
}
