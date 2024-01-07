import 'package:flutter/material.dart';
import 'package:shoppie_admin/router/router.dart';

class ShoppieAdmin extends StatelessWidget {
  const ShoppieAdmin({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'Shoppie Admin',
      routerConfig: AppRouter.router,
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
    );
  }
}
