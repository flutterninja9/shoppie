import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:shoppie_admin/features/auth/view/login_page.dart';
import 'package:shoppie_admin/features/auth/view/register_page.dart';
import 'package:shoppie_admin/features/dashboard/dashboard_page.dart';

class AuthView extends StatefulWidget {
  const AuthView({super.key});

  static const route = "/auth";

  @override
  State<AuthView> createState() => _AuthViewState();
}

class _AuthViewState extends State<AuthView> {
  bool authModeLogin = true;

  @override
  Widget build(BuildContext context) {
    if (authModeLogin) {
      return LoginPage(
        postLoginSucess: () {
          context.go(DashboardPage.route);
        },
        onClickRegister: () {
          setState(() {
            authModeLogin = false;
          });
        },
      );
    }

    return RegisterPage(
      onClickLogin: () {
        setState(() {
          authModeLogin = true;
        });
      },
      onRegisterSuccess: () {
        ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Success!, Login to continue')));
        setState(() {
          authModeLogin = true;
        });
      },
    );
  }
}
