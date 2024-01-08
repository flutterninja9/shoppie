import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/dashboard/controllers/dashboard_controller.dart';

class DashboardPage extends StatelessWidget {
  const DashboardPage({super.key});
  static const route = "/dashboard";

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (context) => sl<DashboardController>()..getSignedInUser(),
      child: Scaffold(
        appBar: AppBar(
          title: const Text("Dashboard"),
          actions: [
            Consumer<DashboardController>(builder: (context, provider, child) {
              if (provider.isLoading) {
                return const Center(
                  child: CircularProgressIndicator(),
                );
              }

              return CircleAvatar(
                child: Text(provider.user?.firstName.characters.first ?? "-"),
              );
            })
          ],
        ),
        floatingActionButton: FloatingActionButton(
          child: const Icon(Icons.add),
          onPressed: () {},
        ),
      ),
    );
  }
}
