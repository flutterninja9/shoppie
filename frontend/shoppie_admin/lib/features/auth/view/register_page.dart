import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:shoppie_admin/features/auth/controller/auth_notifier.dart';
import 'package:shoppie_admin/features/auth/model/register_request.dart';

class RegisterPage extends StatelessWidget {
  const RegisterPage({
    super.key,
    required this.onClickLogin,
    required this.onRegisterSuccess,
  });

  final Function() onClickLogin;
  final Function() onRegisterSuccess;
  static const route = "/register";

  @override
  Widget build(BuildContext context) {
    final usernameController = TextEditingController();
    final firstNameController = TextEditingController();
    final lastNameController = TextEditingController();
    final emailController = TextEditingController();
    final passwordController = TextEditingController();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Register'),
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          crossAxisAlignment: CrossAxisAlignment.stretch,
          children: [
            TextField(
              controller: usernameController,
              decoration: const InputDecoration(
                labelText: 'Username',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 10),
            TextField(
              controller: firstNameController,
              decoration: const InputDecoration(
                labelText: 'First name',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 10),
            TextField(
              controller: lastNameController,
              decoration: const InputDecoration(
                labelText: 'Last name',
                border: OutlineInputBorder(),
              ),
            ),
            const SizedBox(height: 10),
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
                onPressed: () => _register(
                  context,
                  usernameController.text,
                  emailController.text,
                  passwordController.text,
                  firstNameController.text,
                  lastNameController.text,
                ),
                child: const Text('Register'),
              );
            }),
            const SizedBox(height: 20),
            ElevatedButton(
              onPressed: onClickLogin,
              child: const Text('Already have an account? Click here'),
            ),
          ],
        ),
      ),
    );
  }

  Future<void> _register(
    BuildContext context,
    String username,
    String email,
    String password,
    String firstName,
    String lastName,
  ) async {
    if (username.isNotEmpty &&
        email.isNotEmpty &&
        password.isNotEmpty &&
        firstName.isNotEmpty &&
        lastName.isNotEmpty) {
      final registerRequest = RegisterRequest(
        username: username,
        email: email,
        password: password,
        firstName: firstName,
        lastName: lastName,
      );

      await context
          .read<AuthNotifier>()
          .register(registerRequest)
          .then((_) => onRegisterSuccess());
    } else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Please fill in all fields')),
      );
    }
  }
}
