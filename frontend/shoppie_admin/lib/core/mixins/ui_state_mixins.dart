import 'package:flutter/widgets.dart';

mixin UiStateMixin on ChangeNotifier {
  bool _isLoading = false;
  bool _isEmpty = false;
  String? _errorMessage;

  bool get isLoading => _isLoading;
  bool get isLoaded => !_isLoading && _errorMessage == null && !_isEmpty;
  bool get isEmpty => _isEmpty;
  String? get errorMessage => _errorMessage;

  void setLoading(bool loading) {
    _isLoading = loading;
    _isEmpty = false;
    _errorMessage = null;
    notifyListeners();
  }

  void setEmpty() {
    _isLoading = false;
    _isEmpty = true;
    _errorMessage = null;
    notifyListeners();
  }

  void setFailure(String message) {
    _isLoading = false;
    _isEmpty = false;
    _errorMessage = message;
    notifyListeners();
  }

  void resetState() {
    _isLoading = false;
    _isEmpty = false;
    _errorMessage = null;
    notifyListeners();
  }
}
