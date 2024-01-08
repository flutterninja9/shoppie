import 'package:flutter/material.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';

class AddProductPage extends StatelessWidget {
  const AddProductPage({super.key});
  static const route = "${DashboardPage.route}/add-product";

  @override
  Widget build(BuildContext context) {
    return const Text("Add Product");
  }
}
