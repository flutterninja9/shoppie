abstract class Failure {
  Failure(String string);
}

class UnknownFailure implements Failure {
  final Object? error;
  final StackTrace? stack;

  UnknownFailure({this.error, this.stack});
}
