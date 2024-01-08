import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';
import 'package:shoppie_admin/features/products/controllers/create_product_controller.dart';
import 'package:shoppie_admin/features/products/models/create_product_request.dart';
import 'package:shoppie_admin/features/products/pages/view_products.dart';

class AddProductPage extends StatefulWidget {
  const AddProductPage({super.key});
  static const route = "${DashboardPage.route}/add-product";

  @override
  State<AddProductPage> createState() => _AddProductPageState();
}

class _AddProductPageState extends State<AddProductPage> {
  final _formKey = GlobalKey<FormState>();
  late final CreateProductController controller;
  late final TextEditingController nameController;
  late final TextEditingController descController;
  late final TextEditingController priceController;
  late final TextEditingController qtyController;

  @override
  void initState() {
    super.initState();
    controller = sl<CreateProductController>();
    nameController = TextEditingController();
    descController = TextEditingController();
    priceController = TextEditingController();
    qtyController = TextEditingController();
  }

  @override
  void dispose() {
    nameController.dispose();
    descController.dispose();
    priceController.dispose();
    qtyController.dispose();
    super.dispose();
  }

  void _submitProduct() {
    if (_formKey.currentState!.validate()) {
      final req = CreateProductRequest(
        description: descController.text,
        name: nameController.text,
        price: num.parse(priceController.text).toDouble(),
        quantity: num.parse(qtyController.text).toInt(),
      );

      controller.createProduct(req, () {
        context.go(ViewProductsPage.route);
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => controller,
      child: Scaffold(
        body: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Form(
            key: _formKey,
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.stretch,
              children: <Widget>[
                TextFormField(
                  controller: nameController,
                  decoration: const InputDecoration(labelText: 'Product Name'),
                  validator: (value) =>
                      value!.isEmpty ? 'Please enter product name' : null,
                ),
                const SizedBox(height: 15),
                TextFormField(
                  controller: descController,
                  decoration:
                      const InputDecoration(labelText: 'Product Description'),
                  validator: (value) => value!.isEmpty
                      ? 'Please enter product description'
                      : null,
                ),
                const SizedBox(height: 15),
                TextFormField(
                  controller: priceController,
                  decoration: const InputDecoration(labelText: 'Price'),
                  validator: (value) =>
                      value!.isEmpty ? 'Please enter product price' : null,
                ),
                const SizedBox(height: 15),
                TextFormField(
                  controller: qtyController,
                  decoration:
                      const InputDecoration(labelText: 'Product Quantity'),
                  validator: (value) =>
                      value!.isEmpty ? 'Please enter product quantity' : null,
                ),
                const SizedBox(height: 20),
                Consumer<CreateProductController>(
                    builder: (context, provider, child) {
                  if (provider.isLoading) {
                    return const Center(
                      child: CircularProgressIndicator(),
                    );
                  }

                  return ElevatedButton(
                    onPressed: _submitProduct,
                    child: const Text('Add Product'),
                  );
                }),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
