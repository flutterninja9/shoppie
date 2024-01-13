import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';
import 'package:shoppie_admin/features/orders/controller/order_controller.dart';

class ViewOrdersPage extends StatelessWidget {
  const ViewOrdersPage({super.key});
  static const route = "${DashboardPage.route}/view-orders";

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => sl<OrdersController>()..fetchAllOrders(),
      child: Scaffold(
        body: Consumer<OrdersController>(builder: (context, provider, child) {
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
              final order = provider.orders!.orders[index];

              return ListTile(
                title: Text(order.id),
                subtitle: Text(
                  "Status: ${order.status} Items: ${order.items.length} Total: \$${order.totalAmt}",
                ),
                trailing: SizedBox(
                  width: 150,
                  child: Row(
                    children: [
                      TextButton.icon(
                        icon: const Icon(
                          Icons.cancel,
                          color: Colors.red,
                        ),
                        label: const Text(
                          "Cancel",
                          style: TextStyle(color: Colors.red),
                        ),
                        onPressed: () async {
                          try {
                            await context
                                .read<OrdersController>()
                                .cancel(order);
                          } on UnsupportedError catch (e) {
                            if (!context.mounted) return;
                            ScaffoldMessenger.of(context).showSnackBar(
                                SnackBar(content: Text(e.message!)));
                          }
                        },
                      ),
                    ],
                  ),
                ),
              );
            },
            itemCount: provider.orders?.orders.length,
          );
        }),
      ),
    );
  }
}
