import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/auth/model/login_request.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({
    super.key,
    required this.onClickRegister,
    required this.postLoginSucess,
  });

  final Function() onClickRegister;
  final Function() postLoginSucess;
  static const route = "/login";

  @override
  Widget build(BuildContext context) {
    final emailController = TextEditingController();
    final passwordController = TextEditingController();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Login'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            TextField(
              controller: emailController,
              decoration: const InputDecoration(
                labelText: 'Email',
                border: OutlineInputBorder(),
              ),
              keyboardType: TextInputType.emailAddress,
            ),
            const SizedBox(height: 10),
            TextField(
              controller: passwordController,
              decoration: const InputDecoration(
                labelText: 'Password',
                border: OutlineInputBorder(),
              ),
              obscureText: true,
            ),
            const SizedBox(height: 20),
            Consumer<AuthNotifier>(builder: (context, provider, child) {
              if (provider.isLoading) {
                return const Center(
                  child: CircularProgressIndicator(),
                );
              }

              return ElevatedButton(
                onPressed: () => _login(
                  context,
                  emailController.text,
                  passwordController.text,
                ),
                child: const Text('Login'),
              );
            }),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: () {
                onClickRegister();
              },
              child: const Text('Don\'t have an account? Click here'),
            ),
          ],
        ),
      ),
    );
  }

  void _login(BuildContext context, String email, String password) {
    if (email.isNotEmpty && password.isNotEmpty) {
      final loginRequest = LoginRequest(email: email, password: password);
      context.read<AuthNotifier>().login(loginRequest).then((_) {
        postLoginSucess();
      });
    } else {
      ScaffoldMessenger.of(context).showSnackBar(const SnackBar(
        content: Text('Please enter both email and password'),
      ));
    }
  }
}
