import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';
import 'package:shoppie_admin/features/products/controllers/products_controller.dart';
import 'package:shoppie_admin/features/products/pages/update_product.dart';

class ViewProductsPage extends StatelessWidget {
  const ViewProductsPage({super.key});
  static const route = "${DashboardPage.route}/view-products";

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => sl<ProductsController>()..fetchAllProducts(),
      child: Scaffold(
        body: Consumer<ProductsController>(builder: (context, provider, child) {
          if (provider.isLoading) {
            return const Center(
              child: CircularProgressIndicator(),
            );
          }

          if (provider.isEmpty) {
            return const Center(
              child: Text("No data found!"),
            );
          }

          return ListView.builder(
            itemBuilder: (context, index) {
              final product = provider.products!.data[index];

              return ListTile(
                title: Text(product.name),
                subtitle: Text(
                  "Price: ${product.price} Qty: ${product.quantity} ${product.description}",
                ),
                trailing: SizedBox(
                  width: 100,
                  child: Row(
                    children: [
                      IconButton(
                        onPressed: () =>
                            context.go(UpdateProductPage.makeRoute(product.id)),
                        icon: const Icon(Icons.edit),
                      ),
                      IconButton(
                        onPressed: () => context
                            .read<ProductsController>()
                            .deleteProduct(product.id),
                        icon: const Icon(Icons.delete),
                        color: Colors.red,
                      ),
                    ],
                  ),
                ),
                onTap: () {},
              );
            },
            itemCount: provider.products?.data.length,
          );
        }),
      ),
    );
  }
}
