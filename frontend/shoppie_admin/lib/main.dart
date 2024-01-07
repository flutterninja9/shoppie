import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/app/shoppie_admin.dart';
import 'package:shoppie_admin/di/di.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';

Future<void> main() async {
  await dotenv.load();
  await setupSl();
  runApp(
    ChangeNotifierProvider.value(
      value: sl<AuthNotifier>()..checkIfSignedIn(),
      child: const ShoppieAdmin(),
    ),
  );
}
